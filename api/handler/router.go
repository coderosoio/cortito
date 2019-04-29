package handler

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"api/middleware"
	"api/option"
)

func NewRouter(options *option.Options) *gin.Engine {
	router := gin.Default()
	router.RedirectTrailingSlash = true

	log.Printf("AllowedHeaders: %+v", options.AllowedHeaders)
	log.Printf("AllowedMethods: %+v", options.AllowedMethods)
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    options.AllowedHeaders,
		AllowMethods:    options.AllowedMethods,
		ExposeHeaders: []string{
			"Content-Type",
			"Content-Length",
			"Access-Control-Allow-Headers",
			"Access-Control-Allow-Methods",
		},
		AllowCredentials: true,
	}))

	requireToken := middleware.RequireToken(options.AuthService)

	userHandler := NewUserHandler(options)
	authHandler := NewAuthHandler(options)
	linkHandler := NewLinkHandler(options)

	api := router.Group("/api")
	{
		account := api.Group("/account")
		{
			user := account.Group("/users")
			{
				user.POST("/", userHandler.Create)
				user.PUT("/:id", requireToken, userHandler.Update)
				user.GET("/:id", requireToken, userHandler.Find)
			}
			auth := account.Group("/auth")
			{
				auth.POST("/", authHandler.CreateToken)
				auth.DELETE("/", requireToken, authHandler.RevokeToken)
			}
		}
		shortener := api.Group("/shortener")
		{
			link := shortener.Group("/links")
			{
				link.GET("/", requireToken, linkHandler.Index)
				link.POST("/", requireToken, linkHandler.Create)
			}
		}
	}

	return router
}
