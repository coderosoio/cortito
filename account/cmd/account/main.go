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
	"common/hashing"

	"account/handler"
	"account/model"
	"account/option"
	proto "account/proto/account"
	"account/repository"
)

const (
	// ServiceName is the service name.
	ServiceName = "io.coderoso.cortito.account"
	// Version is the current version of this service.
	Version = "0.0.1"
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

	db, err := connection.GetDatabaseConnection("account")
	if err != nil {
		log.Fatalf("error getting database connection: %v", err)
	}
	if err := model.Migrate(db); err != nil {
		log.Fatalf("error migrating database: %v", err)
	}

	// Repositories
	userRepository := repository.NewUserRepository(db)
	userTokenRepository := repository.NewUserTokenRepository(db)

	// Hashing strategy
	hashingStrategy := hashing.NewHashingStrategy(hashing.Argon2HashingStrategy, nil)

	// Auth strategy
	authStrategy, err := auth.NewAuthStrategy()
	if err != nil {
		log.Fatalf("error getting auth strategy: %v", err)
	}

	options := option.NewOptions(
		option.WithUserRepository(userRepository),
		option.WithUserTokenRepository(userTokenRepository),
		option.WithHashingStrategy(hashingStrategy),
		option.WithAuthStrategy(authStrategy),
	)

	userHandler := handler.NewUserHandler(options)
	if err := proto.RegisterUserHandler(service.Server(), userHandler); err != nil {
		log.Fatalf("error registering user handler: %v", err)
	}

	authHandler := handler.NewAuthHandler(options)
	if err := proto.RegisterAuthHandler(service.Server(), authHandler); err != nil {
		log.Fatalf("error registering auth handler: %v", err)
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
