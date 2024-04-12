package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"errors"

	"github.com/gin-gonic/gin"
	"github.com/keshu12345/guardianlink/model"
	mocks "github.com/keshu12345/guardianlink/nodeb/mocks/auth"
	mocksnode "github.com/keshu12345/guardianlink/nodeb/mocks/nodea"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testcase struct {
	name     string
	testcase func(*testing.T)
}

func Test_RegisterEndPoint(t *testing.T) {

	tests := []testcase{

		{
			name: "UPDATE BLOCK SUCCESS CASE",
			testcase: func(t *testing.T) {
				g := gin.New()
				auth := mocks.NewAuthService(t)
				nodea := mocksnode.NewNodeAService(t)
				auth.On("Validate", mock.Anything).Return()
				nodea.On("Update", mock.Anything).Return(model.Block{}, nil)
				RegisterEndpoint(g, auth, nodea)

				expected := http.StatusCreated
				req, _ := http.NewRequest("PUT", "/api/blocks/2", nil)
				resp := httptest.NewRecorder()
				g.ServeHTTP(resp, req)
				actual := resp.Code
				assert.Equal(t, expected, actual)
			},
		},

		{
			name: "UPDATE BLOCK ERROR CASE",
			testcase: func(t *testing.T) {
				g := gin.New()
				auth := mocks.NewAuthService(t)
				nodea := mocksnode.NewNodeAService(t)
				auth.On("Validate", mock.Anything).Return()
				nodea.On("Update", mock.Anything).Return(model.Block{}, errors.New("unable to update block"))
				RegisterEndpoint(g, auth, nodea)

				expected := http.StatusInternalServerError
				req, _ := http.NewRequest("PUT", "/api/blocks/2", nil)
				resp := httptest.NewRecorder()
				g.ServeHTTP(resp, req)
				actual := resp.Code
				assert.Equal(t, expected, actual)
			},
		},

		{
			name: " CREATE BLOCK SUCCESS CASE",
			testcase: func(t *testing.T) {
				g := gin.New()
				auth := mocks.NewAuthService(t)
				nodea := mocksnode.NewNodeAService(t)
				auth.On("Validate", mock.Anything).Return()
				nodea.On("Create", mock.Anything).Return(model.Block{}, nil)
				RegisterEndpoint(g, auth, nodea)

				expected := http.StatusCreated
				req, _ := http.NewRequest("POST", "/api/blocks/", nil)
				resp := httptest.NewRecorder()
				g.ServeHTTP(resp, req)
				actual := resp.Code
				assert.Equal(t, expected, actual)

			},
		},

		{
			name: "CREATE BLOCK ERROR CASE",
			testcase: func(t *testing.T) {
				g := gin.New()
				auth := mocks.NewAuthService(t)
				nodea := mocksnode.NewNodeAService(t)
				auth.On("Validate", mock.Anything).Return()
				nodea.On("Create", mock.Anything).Return(model.Block{}, errors.New("unable to create block"))
				RegisterEndpoint(g, auth, nodea)

				expected := http.StatusInternalServerError
				req, _ := http.NewRequest("POST", "/api/blocks/", nil)
				resp := httptest.NewRecorder()
				g.ServeHTTP(resp, req)
				actual := resp.Code
				assert.Equal(t, expected, actual)
			},
		},

		{
			name: " SIGNUP SUCCESS CASE",
			testcase: func(t *testing.T) {
				g := gin.New()
				auth := mocks.NewAuthService(t)
				nodea := mocksnode.NewNodeAService(t)
				var register string
				auth.On("Singup", mock.Anything).Return(register, nil)
				RegisterEndpoint(g, auth, nodea)

				expected := http.StatusCreated
				req, _ := http.NewRequest("POST", "/api/signup", nil)
				resp := httptest.NewRecorder()
				g.ServeHTTP(resp, req)
				actual := resp.Code
				assert.Equal(t, expected, actual)

			},
		},

		{
			name: " SIGNUP ERROR CASE",
			testcase: func(t *testing.T) {
				g := gin.New()
				auth := mocks.NewAuthService(t)
				nodea := mocksnode.NewNodeAService(t)
				auth.On("Singup", mock.Anything).Return(mock.Anything, errors.New("unable to create user"))
				RegisterEndpoint(g, auth, nodea)

				expected := http.StatusInternalServerError
				req, _ := http.NewRequest("POST", "/api/signup", nil)
				resp := httptest.NewRecorder()
				g.ServeHTTP(resp, req)
				actual := resp.Code
				assert.Equal(t, expected, actual)

			},
		},

		{
			name: " SIGNIN SUCCESS CASE",
			testcase: func(t *testing.T) {
				g := gin.New()
				auth := mocks.NewAuthService(t)
				nodea := mocksnode.NewNodeAService(t)
				var register string
				auth.On("Singin", mock.Anything).Return(register, nil)
				RegisterEndpoint(g, auth, nodea)

				expected := http.StatusOK
				req, _ := http.NewRequest("POST", "/api/signin", nil)
				resp := httptest.NewRecorder()
				g.ServeHTTP(resp, req)
				actual := resp.Code
				assert.Equal(t, expected, actual)

			},
		},

		{
			name: " SINGIN ERROR CASE",
			testcase: func(t *testing.T) {
				g := gin.New()
				auth := mocks.NewAuthService(t)
				nodea := mocksnode.NewNodeAService(t)
				auth.On("Singin", mock.Anything).Return(mock.Anything, errors.New("unable to create user"))
				RegisterEndpoint(g, auth, nodea)

				expected := http.StatusUnauthorized
				req, _ := http.NewRequest("POST", "/api/signin", nil)
				resp := httptest.NewRecorder()
				g.ServeHTTP(resp, req)
				actual := resp.Code
				assert.Equal(t, expected, actual)

			},
		},

		{
			name: " FETCH SUCESS CASE",
			testcase: func(t *testing.T) {
				g := gin.New()
				auth := mocks.NewAuthService(t)
				nodea := mocksnode.NewNodeAService(t)
				auth.On("Validate", mock.Anything).Return()
				nodea.On("Fetch", mock.Anything).Return([]model.Block{}, nil)
				RegisterEndpoint(g, auth, nodea)

				expected := http.StatusOK
				req, _ := http.NewRequest("GET", "/api/blocks/3", nil)
				resp := httptest.NewRecorder()
				actual := resp.Code
				g.ServeHTTP(resp, req)
				assert.Equal(t, expected, actual)

			},
		},

		{
			name: "FETCHING BLOCK ERROR",
			testcase: func(t *testing.T) {
				g := gin.New()
				auth := mocks.NewAuthService(t)
				nodea := mocksnode.NewNodeAService(t)
				auth.On("Validate", mock.Anything).Return()
				nodea.On("Fetch", mock.Anything).Return(nil, errors.New("unable to fetch blocks"))
				RegisterEndpoint(g, auth, nodea)
				expected := http.StatusInternalServerError
				req, _ := http.NewRequest("GET", "/api/blocks/3", nil)
				resp := httptest.NewRecorder()
				g.ServeHTTP(resp, req)
				actual := resp.Code
				assert.Equal(t, expected, actual)
			},
		},

		{
			name: "REQUIRED-AUTH NEGATIVE CASE",
			testcase: func(t *testing.T) {
				g := gin.New()
				auth := mocks.NewAuthService(t)
				nodea := mocksnode.NewNodeAService(t)
				nodea.On("Fetch", mock.Anything).Return(nil, errors.New("unable to fetch blocks"))
				RegisterEndpoint(g, auth, nodea)

				expected := http.StatusInternalServerError
				req, _ := http.NewRequest("GET", "/api/blocks/3", nil)
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Require-Auth", "false")
				resp := httptest.NewRecorder()
				g.ServeHTTP(resp, req)
				actual := resp.Code
				assert.Equal(t, expected, actual)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.testcase)
	}
}
