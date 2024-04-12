package server

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/keshu12345/guardianlink/nodeb/config"
	"go.uber.org/fx/fxtest"
)

type testCase struct {
	name      string
	testCases func(*testing.T)
}

func Test_InitServer(t *testing.T) {

	tests := []testCase{
		{
			name: "CASE SUCCESS",

			testCases: func(t *testing.T) {
				router := gin.Default()
				cfg := &config.Configuration{
					Server: config.Server{
						Port:         8080,
						ReadTimeout:  10,
						WriteTimeout: 10,
						IdleTimeout:  10,
					},
				}
				lifecycle := fxtest.NewLifecycle(t)

				Initialize(router, cfg, lifecycle)
				lifecycle.RequireStart()
				lifecycle.RequireStop()
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.testCases)
	}
}
