package middleware

import (
	//"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strings"
	"time"
)

var jwtSecret = []byte("Og6zf&J#OJTkw4blmpeQ_(hx~!1p%r%fCq%Stv&^fL%6@4kL0i#l$O7(4ZddI71s)_&+KuX")

// AuthMiddleware 校验用户的token是否有效
func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			context.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			context.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			context.Abort()
			return
		}

		context.Set("username", claims["username"])
		context.Set("userRole", claims["role"])
		context.Next()
	}
}

// RoleMiddleware 第二段中间件 用于判断是管理员还是用户
func RoleMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		userRole, ok := context.Get("userRole")
		if !ok {
			context.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
			context.Abort()
			return
		}
		if roleStr, ok := userRole.(string); ok {
			if roleStr == "admin" {
				context.Next()
			} else if roleStr == "user" {
				if isUserPath(context.Request.URL.Path) {
					context.Next()
				} else {
					context.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
					context.Abort()
				}
			} else {
				context.JSON(http.StatusForbidden, gin.H{"error": "Unknown role"})
				context.Abort()
			}
		} else {
			context.JSON(http.StatusForbidden, gin.H{"error": "Unknown role"})
			context.Abort()
		}
	}
}

// GenerateToken 生成Token
func GenerateToken(username string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 6).Unix(), // 6 小时有效期
	})
	return token.SignedString(jwtSecret)
}

// isUserPath 请求的路径是用户的吗
func isUserPath(path string) bool {
	// 根据实际情况判断用户路径
	log.Println("请求路径", path)
	return strings.HasPrefix(path, "/api/user")
}
