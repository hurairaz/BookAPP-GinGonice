package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"strings"
	"time"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type claims struct {
	UserID uint `json:"user-id"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &claims{UserID: userID, RegisteredClaims: jwt.RegisteredClaims{
		Issuer:    "hurairaz",
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func AuthenticateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or malformed"})
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
