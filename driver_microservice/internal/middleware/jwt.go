package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

/*
Product Levelda kullanılacak generic JWT Bearer Token örneği budur. Test amaçlı token const string yapılacaktır.

var JwtSecret = []byte("bitaksi-top-secret-key")

	func JwtAuthMiddleware() gin.HandlerFunc {
		return func(c *gin.Context) {

			//Gelen istekteki Authorization Header'ı token kontrol amaçlı alınır
			authHeader := c.GetHeader("Authorization")
			if authHeader == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
				c.Abort()
				return
			}

			//JWT için Bearer Token kullanıldı bu sebepten ötürü gelen tokenın cinsi kontrol edilir
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization format"})
				c.Abort()
				return
			}

			tokenStr := parts[1]

			//Token içindeki bilgileri parselama
			claims := jwt.MapClaims{}
			token, err := jwt.ParseWithClaims(
				tokenStr,
				claims,
				func(t *jwt.Token) (interface{}, error) {

					//Signing kontrol edilir. Gelen Token HS256 signing olmalı!!
					if t.Method != jwt.SigningMethodHS256 {
						return nil, errors.New("unexpected signing method")
					}

					return JwtSecret, nil
				},
			)

			//Tokenda bir hata meydana gelirse sonraki sayfaya geçmesin diye hata gönderilir
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				c.Abort()
				return
			}

			//Token Valid değilse hata döner
			if !token.Valid {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
				c.Abort()
				return
			}

			//İleride herhangi bir user eklenmesi durumunda user bilgileri token aracılığı ile sayfadan sayfaya taşınır!
			c.Set("user", claims)

			c.Next()
		}
	}
*/
const StaticAPIToken = "BITAKSI-TEST-TOKEN-12345"

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing Authorization header"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid Authorization format"})
			c.Abort()
			return
		}

		token := parts[1]

		if token != StaticAPIToken {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
