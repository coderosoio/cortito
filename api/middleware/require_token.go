package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	accountProto "account/proto/account"
)

func RequireToken(authService accountProto.AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.GetHeader("Authorization")
		parts := strings.Split(authorization, " ")
		if len(parts) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  "invalid authorization header",
			})
			return
		}
		authType := strings.TrimSpace(parts[0])
		switch authType {
		case "Bearer":
			fallthrough
		case "JWT":
			token := strings.TrimSpace(parts[1])
			if len(token) == 0 {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"status": http.StatusUnauthorized,
					"error":  "invalid token",
				})
				return
			}
			ctx.Set("token", token)
			ctx.Keys["token"] = token
			req := &accountProto.VerifyTokenRequest{
				Token: token,
			}
			res, err := authService.VerifyToken(ctx, req)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"status": http.StatusUnauthorized,
					"error":  err,
				})
				return
			}
			ctx.Set("user", res.User)

		default:
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  "invalid authorization type",
			})
			return
		}
		ctx.Next()
	}
}
