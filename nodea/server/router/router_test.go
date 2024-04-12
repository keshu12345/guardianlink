package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/keshu12345/guardianlink/nodea/config"
	"github.com/stretchr/testify/assert"
)

type testcase struct {
	name     string
	testcase func(*testing.T)
}

func Test_NewGinRouter(t *testing.T) {

	tests := []testcase{
		{
			name: "CASE SUCCESS",

			testcase: func(t *testing.T) {
				cfg := &config.Configuration{
					EnvironmentName: "test",
					Server: config.Server{
						Port:         8080,
						ReadTimeout:  10,
						WriteTimeout: 10,
						IdleTimeout:  10,
					},
					Swagger: config.Swagger{},
					SQLite:  config.DB{},
				}
				g, err := NewGinRouter(cfg)
				assert.NoError(t, err, "NewGinRouter should not return an error")

				assert.NotNil(t, g, "NewGinRouter should return a non-nil *gin.Engine instance")
				g.GET("/test", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "test"})
				})
				w := performRequest(g, "GET", "/test")
				assert.Equal(t, 200, w.Code)
				assert.Contains(t, w.Body.String(), "test", "Expected response body to contain 'test'")

			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.testcase)
	}
}

func performRequest(r *gin.Engine, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
