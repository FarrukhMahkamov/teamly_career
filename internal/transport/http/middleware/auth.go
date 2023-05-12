package middleware

import (
	"strings"

	"github.com/FarrukhMahkamov/teamly_career/pkg"
	"github.com/gin-gonic/gin"
)

const (
	AuthorizationHeder = "Authorization"
	UserIdKey          = "user_id"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		Header := c.GetHeader(AuthorizationHeder)
		if Header == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"Message": "Unauthorized",
			})
			return
		}

		HeaderParts := strings.Split(Header, " ")
		if len(HeaderParts) != 2 {
			c.AbortWithStatusJSON(401, gin.H{
				"Message": "Invalid authorization header",
			})
			return
		}

		UserId, err := pkg.ParseToken(HeaderParts[1])
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"Message": "Unauthorized",
			})
			return
		}

		c.Set(UserIdKey, UserId)
		c.Next()
	}
}
