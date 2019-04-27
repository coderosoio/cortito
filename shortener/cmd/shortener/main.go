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

	"common/config"
	"common/connection"
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
	_ = service.Server().Init(
		server.Wait(true),
	)

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

	options := option.NewOptions(
		option.WithLinkRepository(linkRepository),
	)

	linkHandler := handler.NewLinkHandler(options)

	if err := proto.RegisterLinkHandler(service.Server(), linkHandler); err != nil {
		log.Fatal(err)
	}

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
