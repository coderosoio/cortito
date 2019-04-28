package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/micro/cli"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-web"

	accountProto "account/proto/account"
	commonConfig "common/config"
	shortenerProto "shortener/proto/shortener"

	"api/handler"
	"api/option"
)

const (
	ServiceName = "io.coderoso.cortito.api"
	Version     = "0.0.1"
)

func main() {
	var (
		configurationFile string

		errc = make(chan error)
	)
	service := web.NewService(
		web.Name(ServiceName),
		web.Version(Version),
		web.Flags(
			cli.StringFlag{
				Name:        "config, c",
				Value:       "config.yml",
				Usage:       "Path to the configuration file to use. Defaults to config.yml",
				EnvVar:      "CONFIG_FILE",
				Destination: &configurationFile,
			},
		),
		web.RegisterTTL(30*time.Second),
		web.RegisterInterval(10*time.Second),
	)
	_ = service.Init()

	config, err := commonConfig.GetConfig()
	if err != nil {
		log.Fatalf("error getting configuration: %v", err)
	}

	httpConfig := config.HTTP
	if httpConfig == nil {
		log.Fatal("no http configuration for web_tmp")
	}

	accountService := config.Service["account"]
	shortenerService := config.Service["shortener"]

	userService := accountProto.NewUserService(accountService.URL(), client.DefaultClient)
	authService := accountProto.NewAuthService(accountService.URL(), client.DefaultClient)
	linkService := shortenerProto.NewLinkService(shortenerService.URL(), client.DefaultClient)

	options := option.NewOptions(
		option.WithUserService(userService),
		option.WithAuthService(authService),
		option.WithLinkService(linkService),
	)

	router := handler.NewRouter(options)

	service.Handle("/", router)

	go func() {
		if err := service.Run(); err != nil {
			errc <- err
		}
	}()

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
		errc <- fmt.Errorf("%v", <-c)
	}()

	if err := <-errc; err != nil {
		log.Fatal(err)
	}
}
