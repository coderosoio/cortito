package helper

import "github.com/gin-gonic/gin"

func CurrentToken(ctx *gin.Context) (string, bool) {
	token, found := ctx.Get("token")
	if found {
		return token.(string), true
	}
	return "", false
}
