// internal/http/middleware/auth.go
package middleware

import (
	"net/http"
	"strings"

	apperror "moviediary/pkg/apperror"
	utils "moviediary/pkg/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			utils.Fail(c, http.StatusUnauthorized, apperror.ErrAuthorizationHeaderMissing.Message)
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := utils.ParseJWT(tokenStr, secret)
		if err != nil {
			utils.Fail(c, http.StatusUnauthorized, apperror.ErrInvalidToken.Message)
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)

		c.Next()
	}
}
