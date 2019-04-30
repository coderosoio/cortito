package main

import (
	"common/keyvalue"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/micro/go-micro"

	commonConfig "common/config"
	shortenerProto "shortener/proto/shortener"

	"web/handler"
	"web/option"
)

func main() {

	var (
		configurationFile = flag.String("config,c", "config.yml", "Path to configuration file to use.")
		errc              = make(chan error)
	)
	flag.Parse()

	service := micro.NewService()
	service.Init()

	if err := commonConfig.SetConfigurationFile(*configurationFile); err != nil {
		log.Fatalf("error setting configuration file: %v", err)
	}

	config, err := commonConfig.GetConfig()
	if err != nil {
		log.Fatalf("error getting configuration: %v", err)
	}

	httpConfig := config.HTTP
	if httpConfig == nil {
		log.Fatal("no http configuration for web_tmp")
	}

	shortenerService := config.Service["shortener"]

	linkService := shortenerProto.NewLinkService(shortenerService.URL(), service.Client())

	keyValueStorage, err := keyvalue.NewKeyValueStorage("shortener")
	if err != nil {
		log.Fatalf("error getting key value storage: %v", err)
	}
	keyValueStorage = keyvalue.NamespaceMiddleware("shortener")(keyValueStorage)

	options := option.NewOptions(
		option.WithLinkService(linkService),
		option.WithKeyValueStorage(keyValueStorage),
	)

	router, err := handler.NewRouter(options)
	if err != nil {
		log.Fatalf("error getting handler: %v", err)
	}

	go func() {
		if httpConfig.IsSecure {
			if err := router.RunTLS(httpConfig.Address(false), httpConfig.SecureHTTP.CertFilepath, httpConfig.SecureHTTP.KeyFilepath); err != nil {
				errc <- fmt.Errorf("error running web_tmp server: %v", err)
			}
		} else {
			if err := router.Run(httpConfig.Address(false)); err != nil {
				errc <- fmt.Errorf("error running web_tmp server: %v", err)
			}
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
