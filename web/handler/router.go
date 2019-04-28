package handler

import (
	"github.com/gin-gonic/gin"

	"web/option"
)

func NewRouter(options *option.Options) (*gin.Engine, error) {
	router := gin.Default()
	router.RedirectTrailingSlash = true

	linkHandler := NewLinkHandler(options)

	link := router.Group("/")
	{
		link.GET("/:hash", linkHandler.Visit)
	}

	return router, nil
}
