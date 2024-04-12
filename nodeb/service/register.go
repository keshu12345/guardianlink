package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// var jwtKey = []byte("speer")
var authServiceInterface AuthService
var nodeAServiceInterface NodeAService

func RegisterEndpoint(g *gin.Engine, as AuthService, nas NodeAService) {

	authServiceInterface = as
	nodeAServiceInterface = nas

	g.POST("/api/signup", signup)
	g.POST("/api/signin", signin)
	blocks := g.Group("/api/blocks")
	{
		blocks.Use(authMiddleware)
		blocks.POST("", create)
		blocks.PUT("/:height", update)
		blocks.GET("/:height", fetch)

	}

}

func create(c *gin.Context) {

	block, err := nodeAServiceInterface.Create(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to create block"})
	}

	c.JSON(http.StatusCreated, gin.H{"block": block})
}

func update(c *gin.Context) {

	block, err := nodeAServiceInterface.Update(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to update block"})
	}

	c.JSON(http.StatusCreated, gin.H{"block": block})

}

func fetch(c *gin.Context) {

	blocks, err := nodeAServiceInterface.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch blocks"})
	}

	c.JSON(http.StatusCreated, gin.H{"block": blocks})

}

func authMiddleware(c *gin.Context) {

	isAuthRequired := true
	if c.GetHeader("Require-Auth") == "false" {
		isAuthRequired = false
	}

	if isAuthRequired {
		authServiceInterface.Validate(c)
		if c.IsAborted() {
			return
		}
	}
	c.Next()
}

func signup(c *gin.Context) {

	successfully, err := authServiceInterface.Singup(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to create user"})
	}
	c.JSON(http.StatusCreated, gin.H{"message": successfully})
}

func signin(c *gin.Context) {

	token, err := authServiceInterface.Singin(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized access",
		})
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
