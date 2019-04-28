package router

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"common/config"

	"web/handler"
	"web/middleware"
	"web/option"
)

func NewRouter(options *option.Options) (*gin.Engine, error) {
	router := gin.Default()
	if config.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	renderer, err := setupRenderer()
	if err != nil {
		return nil, err
	}

	store := cookie.NewStore([]byte(options.SessionSecret))

	router.Use(sessions.Sessions(options.SessionName, store))

	// router.RedirectTrailingSlash = true
	router.HTMLRender = renderer
	router.StaticFS("/assets", http.Dir("static"))

	router.Use(middleware.CurrentUser())

	homeHandler := handler.NewHomeHandler(options)
	userHandler := handler.NewUserHandler(options)
	authHandler := handler.NewAuthHandler(options)
	linkHandler := handler.NewLinkHandler(options)

	home := router.Group("/")
	{
		home.GET("/", homeHandler.Index)
	}

	user := router.Group("/users")
	{
		user.GET("/new", userHandler.New)
		user.POST("/users", userHandler.Create)
	}

	auth := router.Group("/sessions")
	{
		auth.GET("/new", authHandler.New)
		auth.POST("/", authHandler.Create)
		auth.GET("/", authHandler.Delete)
	}

	link := router.Group("/links")
	{
		link.GET("/new", linkHandler.New)
		link.POST("/", linkHandler.Create)
	}

	visit := router.Group("/l")
	{
		visit.GET("/:hash", linkHandler.Visit)
	}

	return router, nil
}

func setupRenderer() (*gintemplate.TemplateEngine, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	templatePath, err := filepath.Abs(fmt.Sprintf("%s/template", wd))
	if err != nil {
		return nil, err
	}

	renderer := gintemplate.New(gintemplate.TemplateConfig{
		Root:      templatePath,
		Extension: ".html",
		Master:    "layout/master",
	})
	return renderer, nil
}
