package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lining4069/ops-go/backend/app/common/response"
	"github.com/lining4069/ops-go/backend/app/services"
	"github.com/lining4069/ops-go/backend/global"
)

// JWTAuth Token鉴权中间件
func JWTAuth(GuardName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == "" {
			response.TokenFail(c)
			c.Abort()
			return
		}

		tokenStr = tokenStr[len(services.TokenType)+1:]
		// Token解析
		token, err := jwt.ParseWithClaims(tokenStr, &services.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(global.App.Config.Jwt.Secret), nil
		})
		if err != nil {
			response.TokenFail(c)
			c.Abort()
			return
		}

		claims := token.Claims.(*services.CustomClaims)
		// 发布者校验
		if claims.Issuer != GuardName {
			response.TokenFail(c)
			c.Abort()
			return
		}

		c.Set("token", token)
		c.Set("id", claims.Id)
	}
}
