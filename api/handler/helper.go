package handler

import (
	"log"

	"github.com/gin-gonic/gin"

	accountProto "account/proto/account"
)

func CurrentUser(ctx *gin.Context) *accountProto.UserResponse {
	user, found := ctx.Get("user")
	log.Printf("%t: %+v", found, user)
	if found {
		return user.(*accountProto.UserResponse)
	}
	return nil
}
