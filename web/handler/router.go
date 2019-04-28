package router

import (
	"github.com/gin-gonic/gin"

	"web/option"
)

func NewRouter(options *option.Options) *gin.Engine {
	router := gin.Default()
	router.RedirectTrailingSlash = true

	router.GET("/:hash")
}
