package middleware

import (
	"fmt"
	"strings"
	response "sukvij/galenfers/Response"
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

var jwtSecret = []byte("super-secret-key")

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

func Login(ctx *gin.Context, username string, role string) {
	// var credentials Credentials

	// credentials.UserName = username
	// dummy auth
	// if credentials.UserName != "sukvij" || credentials.Password != "12345" {
	// 	// ctx.JSON(401, gin.H{"error": "invalid username or password"})
	// 	response.SendResponse(ctx, nil, errors.New("invalid username or password"))
	// 	return
	// }

	token, err := GenerateToken(username, role)
	if err != nil {
		response.SendResponse(ctx, nil, err)
		return
	}

	response.SendResponse(ctx, gin.H{
		"access_token": token,
		"token_type":   "Bearer",
		"expires_in":   86400,
	}, err)
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			// ctx.AbortWithStatusJSON(401, gin.H{"error": "authorization header missing"})
			response.AbortWithStatus(ctx, "authorization header missing", 401)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			// ctx.AbortWithStatusJSON(401, gin.H{"error": "invalid authorization format"})
			response.AbortWithStatus(ctx, "invalid authorization format", 401)
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
			// ctx.AbortWithStatusJSON(401, gin.H{"error": "invalid or expired token"})
			response.AbortWithStatus(ctx, "invalid or expired token", 401)
			return
		}

		claims, ok := token.Claims.(*Claims)
		if !ok {
			// ctx.AbortWithStatusJSON(401, gin.H{"error": "invalid token claims"})
			response.AbortWithStatus(ctx, "invalid token claims", 401)
			return
		}

		ctx.Set("user_name", claims.UserName)
		ctx.Set("role", claims.Role)

		ctx.Next()
	}
}
