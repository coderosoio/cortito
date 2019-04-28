package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"web/option"
)

type HomeHandler struct{}

func NewHomeHandler(options *option.Options) *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) Index(ctx *gin.Context) {
	user := CurrentUser(ctx)
	ctx.HTML(http.StatusOK, "home/index", gin.H{
		"user":      user,
		"logged_in": LoggedIn(ctx),
	})
}
