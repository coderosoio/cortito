package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	shortenerProto "shortener/proto/shortener"

	"web/option"
)

type linkHandler struct {
	linkService shortenerProto.LinkService
}

func NewLinkHandler(options *option.Options) *linkHandler {
	return &linkHandler{
		linkService: options.LinkService,
	}
}

func (h *linkHandler) Visit(ctx *gin.Context) {
	hash := ctx.Param("hash")
	req := &shortenerProto.LinkRequest{
		Hash: hash,
	}
	res, err := h.linkService.FindLink(ctx, req)
	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}
	if _, err = h.linkService.IncreaseVisit(ctx, req); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.Header("Cache-Control", "no-cache")
	ctx.Redirect(http.StatusMovedPermanently, res.Url)
}
