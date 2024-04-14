package ratelimitor

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"gopkg.in/gorp.v2"
)

type Ratelimitor interface {
	Instance(c *gin.Context)
}

type Service struct {
	fx.In
	Client *gorp.DbMap
}

// Instance implements Ratelimitor.
func (s Service) Instance(c *gin.Context) {
	panic("unimplemented")
}

func NewService(as Service) Ratelimitor {
	return Service{
		Client: as.Client,
	}
}
