package helper

import (
	"github.com/gin-gonic/gin"

	accountProto "account/proto/account"
)

func CurrentUser(ctx *gin.Context) *accountProto.UserRequest {
	user, found := ctx.Get("user")
	if found {
		return user.(*accountProto.UserRequest)
	}
	return nil
}
