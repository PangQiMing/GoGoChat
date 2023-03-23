package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"
)

type MyClaims struct {
	GoGoID uint64 `json:"go_go_id"`
	jwt.StandardClaims
}

// GenerateToken  生成Token
func GenerateToken(goGoID uint64) (string, error) {
	claims := &MyClaims{
		GoGoID: goGoID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, err
}

// ParseToken 解析token
func ParseToken(tokenString string) (uint64, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims.GoGoID, nil
	} else {
		log.Println(err)
		return 0, err
	}
}

func VerificationToken(ctx *gin.Context) uint64 {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "没有找到Token",
		})
		return 0
	}

	ID, err := ParseToken(authHeader)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return 0
	}
	return ID
}
