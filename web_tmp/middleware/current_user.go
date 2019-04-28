package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	accountProto "account/proto/account"
)

func CurrentUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		user, ok := session.Get("user").(*accountProto.UserResponse)
		if ok {
			ctx.Set("user", user)
		}
		ctx.Next()
	}
}
