package handler

import (
	"common/keyvalue"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	shortenerProto "shortener/proto/shortener"

	"web/option"
)

type linkHandler struct {
	linkService     shortenerProto.LinkService
	keyValueStorage keyvalue.Storage
}

func NewLinkHandler(options *option.Options) *linkHandler {
	return &linkHandler{
		linkService:     options.LinkService,
		keyValueStorage: options.KeyValueStorage,
	}
}

func (h *linkHandler) Visit(ctx *gin.Context) {
	hash := ctx.Param("hash")
	req := &shortenerProto.LinkRequest{
		Hash: hash,
	}
	url, err := h.keyValueStorage.Get(hash, "")
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	if len(url) == 0 {
		res, err := h.linkService.FindLink(ctx, req)
		if err != nil {
			ctx.String(http.StatusNotFound, err.Error())
			return
		}
		url = res.Url
	}
	ctx.Header("Cache-Control", "no-cache")
	go func() {
		if _, err = h.linkService.IncreaseVisit(ctx, req); err != nil {
			log.Printf("error increasing link visit: %v", err)
		}
	}()

	ctx.Redirect(http.StatusMovedPermanently, url)
}
