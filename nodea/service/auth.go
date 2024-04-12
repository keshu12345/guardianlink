package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/keshu12345/guardianlink/model"
	"go.uber.org/fx"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gorp.v2"
)

var JwtSecretKey = "blocks"

// mockery --exported --name=AuthService --case underscore --output ../mocks/auth
// SqlExecutor
//mockery --exported --name=SqlExecutor --case underscore --output /Users/keshav/go/src/guardianlink/nodeb/mocks/db
type AuthService interface {
	Singup(c *gin.Context) (string, error)
	Singin(c *gin.Context) (string, error)
	Validate(c *gin.Context)
}

type authServiceImpl struct {
	fx.In
	Client *gorp.DbMap
}

func NewAuthService(as authServiceImpl) AuthService {
	return authServiceImpl{
		Client: as.Client,
	}
}

func (as authServiceImpl) Singin(c *gin.Context) (string, error) {

	var loginParams struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return "", fmt.Errorf("unable to bind user: %v", err)
	}

	var user model.User
	if err := as.Client.SelectOne(&user, "SELECT * FROM users WHERE username=?", loginParams.Username); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return "", fmt.Errorf("error:User not found: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginParams.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return "", fmt.Errorf("error:Invalid credentials: %w", err)
	}

	token, err := issueToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not issue token"})
		return "", fmt.Errorf("could not issue token: %w", err)

	}

	user.Token = token
	if _, err = as.Client.Update(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update user token"})
		return "", fmt.Errorf("could not update user token: %w", err)
	}
	return token, nil
}

func (as authServiceImpl) Singup(c *gin.Context) (string, error) {
	var newUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return "Failed to register user", fmt.Errorf("invalid request: %w", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return "Failed to hash password", fmt.Errorf("failed to hash password: %w", err)
	}

	var dbuser = &model.User{
		Username: newUser.Username,
		Password: string(hashedPassword),
	}

	if err := as.Client.Insert(dbuser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return "Failed to register user", fmt.Errorf("failed to register user: %w", err)
	}

	return "User registered successfully", nil
}

func (as authServiceImpl) Validate(c *gin.Context) {

	tokenString := c.GetHeader("Authorization")

	token, err := jwt.ParseWithClaims(tokenString, &model.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtSecretKey), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	if claims, ok := token.Claims.(*model.CustomClaims); ok && token.Valid {

		var user model.User

		if err := as.Client.SelectOne(&user, "SELECT * FROM users WHERE username=?", claims.Username); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return

		}
		if user.Token != tokenString {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token mismatch"})
			return
		}

		c.Set("username", claims.Username)
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	c.Next()

}

func issueToken(username string) (string, error) {
	claims := model.CustomClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JwtSecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
