package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	accountProto "account/proto/account"
)

func CurrentUser(ctx *gin.Context) *accountProto.UserResponse {
	user, exists := ctx.Get("user")
	if exists {
		user := user.(*accountProto.UserResponse)
		return user
	}
	return nil
}

func LoggedIn(ctx *gin.Context) bool {
	session := sessions.Default(ctx)
	token := session.Get("token")
	return token != nil
}
