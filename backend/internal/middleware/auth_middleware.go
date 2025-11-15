package middleware

import (
	"hih-yadebi-backend/internal/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var cfgJWT = config.LoadJWT()

func JWTAuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// cookie, err := c.Cookie("auth_token")
		// if err != nil {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		// 	c.Abort()
		// 	return
		// }
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			c.Abort()
			return
		}
		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		tokenstr := headerParts[1]
		// tokenstr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenstr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenMalformed
			}
			return []byte(cfgJWT.JWTSecret), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("user_id", int(claims["user_id"].(float64)))
			c.Set("username", claims["username"].(string))
		}
		c.Next()
	}
}
