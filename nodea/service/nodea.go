package service

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/keshu12345/guardianlink/model"
	"github.com/keshu12345/guardianlink/nodea/constant"

	logger "github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gopkg.in/gorp.v2"
)

type NodeAService interface {
	Create(c *gin.Context) (model.Block, error)
	Update(c *gin.Context) (model.Block, error)
	Fetch(c *gin.Context) ([]model.Block, error)
}

func ProvideMutex() *sync.Mutex {
	return &sync.Mutex{}
}

type nodeAServiceImpl struct {
	fx.In
	DBLock *sync.Mutex
	Client *gorp.DbMap
}

func NewNodeAService(as nodeAServiceImpl) NodeAService {
	return &nodeAServiceImpl{
		Client: as.Client,
		DBLock: &sync.Mutex{},
	}
}

func (nas nodeAServiceImpl) Create(c *gin.Context) (model.Block, error) {

	var block model.Block
	if err := c.ShouldBindJSON(&block); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return model.Block{}, fmt.Errorf("error binding JSON: %w", err)
	}

	block.Hash = generateRandomHash()
	block.Status = "pending"

	nas.DBLock.Lock()
	defer nas.DBLock.Unlock()

	trans, err := nas.Client.Begin()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
		return model.Block{}, fmt.Errorf("failed to start transaction: %w", err)
	}

	if err := trans.Insert(&block); err != nil {
		trans.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert block"})
		return model.Block{}, fmt.Errorf("failed to insert block: %w", err)
	}

	if err := trans.Commit(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
		return model.Block{}, fmt.Errorf("failed to commit transaction: %w", err)
	}

	go func(block model.Block) {
		if err := notifyNodeB(block, constant.NodeBURL.ToString()); err != nil {
			logger.Errorf("Failed to notify NodeB: %v", err)
			compensateBlockFailure(block, nas)
		} else {
			markBlockAsComplete(block, nas)
		}
	}(block)

	return block, nil

}

func (nas nodeAServiceImpl) Fetch(c *gin.Context) ([]model.Block, error) {

	height, err := strconv.Atoi(c.Param("height"))
	if err != nil {
		return []model.Block{}, fmt.Errorf("invalid height parameter: %w", err)
	}

	var blocks []model.Block
	_, err = nas.Client.Select(&blocks, "SELECT * FROM blocks WHERE height >= ?", height)
	if err != nil {
		return []model.Block{}, fmt.Errorf("failed to fetch blocks: %w", err)
	}

	return blocks, nil
}

func (nas nodeAServiceImpl) Update(c *gin.Context) (model.Block, error) {

	var block model.Block
	if err := c.ShouldBindJSON(&block); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return model.Block{}, fmt.Errorf("error binding JSON: %w", err)
	}

	block.Hash = generateRandomHash()
	block.Status = "pending"

	nas.DBLock.Lock()
	defer nas.DBLock.Unlock()

	trans, err := nas.Client.Begin()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
		return model.Block{}, fmt.Errorf("failed to start transaction: %w", err)
	}

	if _, err := trans.Update(&block); err != nil {
		trans.Rollback()
		return model.Block{}, fmt.Errorf("failed to update block: %w", err)
	}

	if err := trans.Commit(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
		return model.Block{}, fmt.Errorf("failed to commit transaction: %w", err)
	}

	go func(block model.Block) {
		if err := notifyNodeBUpdate(block, constant.NodeBURL.ToString()); err != nil {
			logger.Errorf("Failed to notify NodeB: %v", err)
			compensateBlockFailure(block, nas)
		} else {
			markBlockAsComplete(block, nas)
		}
	}(block)

	return block, nil
}

func notifyNodeB(block model.Block, nodebURL string) error {

	jsonValue, err := json.Marshal(block)
	if err != nil {
		logger.Errorf("Failed to marshal block: %v", err)
		return fmt.Errorf("failed to marshal block: %v", err)
	}

	req, err := http.NewRequest("POST", nodebURL, bytes.NewBuffer(jsonValue))
	if err != nil {
		logger.Errorf("Failed to create request: %v", err)
		return fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Require-Auth", "false")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Errorf("Failed to synchronize block with Node B: %v", err)
		return fmt.Errorf("failed to synchronize block with Node B: %v", err)
	}
	defer resp.Body.Close()

	return nil
}

func notifyNodeBUpdate(block model.Block, nodebURL string) error {

	jsonValue, err := json.Marshal(block)
	if err != nil {
		logger.Errorf("Failed to marshal block: %v", err)
		return fmt.Errorf("failed to marshal block: %v", err)
	}

	url := fmt.Sprintf("%s/%v", nodebURL, block.Height)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
	if err != nil {
		logger.Errorf("Failed to create request: %v", err)
		return fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Require-Auth", "false")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Errorf("Failed to synchronize block with Node B: %v", err)
		return fmt.Errorf("failed to synchronize block with Node B: %v", err)
	}
	defer resp.Body.Close()
	return nil
}

func compensateBlockFailure(block model.Block, nas nodeAServiceImpl) {

	nas.DBLock.Lock()
	defer nas.DBLock.Unlock()

	trans, err := nas.Client.Begin()
	if err != nil {
		logger.Errorf("Failed to start compensating transaction: %v", err)
		return
	}

	block.Status = "failed"
	if _, err := trans.Update(&block); err != nil {
		logger.Errorf("Failed to update block status to 'failed': %v", err)
		trans.Rollback()
		return
	}

	if err := trans.Commit(); err != nil {
		logger.Errorf("Failed to commit compensating transaction: %v", err)
		return
	}

	logger.Infof("Compensating transaction completed: Block %v marked as 'failed'", block.Height)
}

func markBlockAsComplete(block model.Block, nas nodeAServiceImpl) {

	nas.DBLock.Lock()
	defer nas.DBLock.Unlock()

	trans, err := nas.Client.Begin()
	if err != nil {
		logger.Errorf("Failed to start transaction for marking block as complete: %v", err)
		return
	}

	block.Status = "completed"
	if _, err := trans.Update(&block); err != nil {
		logger.Errorf("Failed to mark block as 'complete': %v", err)
		trans.Rollback()
		return
	}

	if err := trans.Commit(); err != nil {
		logger.Errorf("Failed to commit transaction for block completion: %v", err)
	}

}

func generateRandomHash() string {
	rand.Seed(time.Now().UnixNano())
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	hash := sha256.Sum256(randBytes)

	return hex.EncodeToString(hash[:])
}
