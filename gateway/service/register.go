package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keshu12345/guardianlink/gateway/service/auth"
	"github.com/keshu12345/guardianlink/gateway/service/ratelimitor"
)

var authencator auth.Authencator
var ratelimitorInterface ratelimitor.Ratelimitor

func RegisterEndpoint(g *gin.Engine, a auth.Authencator, rl ratelimitor.Ratelimitor) {

	authencator = a
	ratelimitorInterface=rl
	gateway := g.Group("/api")
	{   gateway.GET("",rateLimit)
		gateway.POST("/signup", signup)
		gateway.POST("/signin", signin)
		gateway.Use(authMiddleware)

	}

}


func rateLimit(c*gin.Context){
	ratelimitorInterface.Instance(c)

}
func authMiddleware(c *gin.Context) {

	isAuthRequired := true
	if c.GetHeader("Require-Auth") == "false" {
		isAuthRequired = false
	}

	if isAuthRequired {
		authencator.Validate(c)
		if c.IsAborted() {
			return
		}
	}
	c.Next()
}

func signup(c *gin.Context) {

	successfully, err := authencator.Singup(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to create user"})
	}

	c.JSON(http.StatusCreated, gin.H{"message": successfully})

}

func signin(c *gin.Context) {

	token, err := authencator.Singin(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized access",
		})
	}

	c.JSON(http.StatusCreated, gin.H{"token": token})

}
