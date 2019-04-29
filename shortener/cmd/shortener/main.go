package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"

	"common/auth"
	"common/config"
	"common/connection"
	commonWrapper "common/wrapper"
	proto "shortener/proto/shortener"

	"shortener/handler"
	"shortener/model"
	"shortener/option"
	"shortener/repository"
)

const (
	ServiceName = "io.coderoso.cortito.shortener"
	Version     = "0.0.1"
)

func main() {
	var (
		configurationFile string

		errc = make(chan error)
	)
	service := micro.NewService(
		micro.Name(ServiceName),
		micro.Version(Version),
		micro.Flags(
			cli.StringFlag{
				Name:        "config, c",
				Value:       "config.yml",
				Usage:       "Path to the configuration file to use. Defaults to config.yml",
				EnvVar:      "CONFIG_FILE",
				Destination: &configurationFile,
			},
		),
		micro.RegisterTTL(30*time.Second),
		micro.RegisterInterval(10*time.Second),
	)
	service.Init()

	if err := config.SetConfigurationFile(configurationFile); err != nil {
		log.Fatalf("error setting configuration file: %v", err)
	}

	db, err := connection.GetDatabaseConnection("shortener")
	if err != nil {
		log.Fatalf("error getting database connection: %v", err)
	}
	if err := model.Migrate(db); err != nil {
		log.Fatalf("error migrating database: %v", err)
	}

	linkRepository := repository.NewLinkRepository(db)

	// Auth strategy
	authStrategy, err := auth.NewAuthStrategy()
	if err != nil {
		log.Fatalf("error getting auth strategy: %v", err)
	}

	options := option.NewOptions(
		option.WithLinkRepository(linkRepository),
	)

	linkHandler := handler.NewLinkHandler(options)

	if err := proto.RegisterLinkHandler(service.Server(), linkHandler); err != nil {
		log.Fatal(err)
	}

	_ = service.Server().Init(
		server.WrapHandler(
			commonWrapper.NewAuthHandlerWrapper(authStrategy),
		),
		server.Wait(true),
	)

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
