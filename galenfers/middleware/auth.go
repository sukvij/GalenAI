package middleware

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Credentials struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type Claims struct {
	UserName string `json:"user_name"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte("super-secret-key") // use env in prod

func GenerateToken(userName string, role string) (string, error) {
	claims := Claims{
		UserName: userName,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "employee-service",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func Login(ctx *gin.Context) {
	var credentials Credentials

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(400, gin.H{"error": "invalid request body"})
		return
	}

	// dummy auth
	if credentials.UserName != "sukvij" || credentials.Password != "12345" {
		ctx.JSON(401, gin.H{"error": "invalid username or password"})
		return
	}

	token, err := GenerateToken(credentials.UserName, "admin")
	if err != nil {
		ctx.JSON(500, gin.H{"error": "failed to generate token"})
		return
	}

	ctx.JSON(200, gin.H{
		"access_token": token,
		"token_type":   "Bearer",
		"expires_in":   86400,
	})
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "authorization header missing"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "invalid authorization format"})
			return
		}

		tokenStr := parts[1]

		token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "invalid or expired token"})
			return
		}

		claims, ok := token.Claims.(*Claims)
		if !ok {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "invalid token claims"})
			return
		}

		ctx.Set("user_name", claims.UserName)
		ctx.Set("role", claims.Role)

		ctx.Next()
	}
}
