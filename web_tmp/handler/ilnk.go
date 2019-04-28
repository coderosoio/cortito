package handler

import (
	"github.com/gin-gonic/gin"

	shortenerProto "shortener/proto/shortener"

	"web/option"
)

type LinkHandler struct {
	linkService shortenerProto.LinkService
}

func NewLinkHandler(options *option.Options) *LinkHandler {
	return &LinkHandler{
		linkService: options.LinkService,
	}
}

func (h *LinkHandler) Visit(ctx *gin.Context) {

}

func (h *LinkHandler) New(ctx *gin.Context) {

}

func (h *LinkHandler) Create(ctx *gin.Context) {

}
