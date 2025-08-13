package main

import (
	"fmt"
	"github.com/bluewingtitan/go-chi-openapi-preset/config"
	"github.com/bluewingtitan/go-chi-openapi-preset/core"
	"github.com/bluewingtitan/go-chi-openapi-preset/incoming"
	"github.com/bluewingtitan/go-chi-openapi-preset/logging"
	"github.com/bluewingtitan/go-chi-openapi-preset/middleware"
	"github.com/bluewingtitan/go-chi-openapi-preset/service"
	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func main() {
	conf, err := config.LoadConfig("config.yaml")

	if err != nil {
		fmt.Printf("Could not load config: %v\n", err)
		return
	}

	logger, err := logging.InitializeLogger(&conf)

	if err != nil {
		fmt.Printf("Could not create logger: %v\n", err)
		return
	}

	logger.Info().Msg("starting server")

	srv := service.NewExampleService(&conf)
	startServer(srv, &conf) // todo: configure values

}

func startServer(service core.Service, conf *config.Config) {
	handler := incoming.NewReqHandler(service)

	strictHandler := incoming.NewStrictHandler(handler, []nethttp.StrictHTTPMiddlewareFunc{})

	router := chi.NewRouter()

	h := incoming.HandlerFromMux(strictHandler, router)
	h = middleware.CorsMiddleware(conf.AllowedOrigins)(h)

	server := &http.Server{
		Addr:                         conf.Address,
		Handler:                      http.TimeoutHandler(h, time.Duration(conf.Timeout.Handling)*time.Second, "request timed out"),
		DisableGeneralOptionsHandler: false,
		TLSConfig:                    nil,
		ReadTimeout:                  time.Duration(conf.Timeout.Read) * time.Second,
		ReadHeaderTimeout:            time.Duration(conf.Timeout.ReadHeader) * time.Second,
		WriteTimeout:                 time.Duration(conf.Timeout.Write) * time.Second,
		IdleTimeout:                  time.Duration(conf.Timeout.Idle) * time.Second,
	}

	log.Info().Str("address", conf.Address).Msg("starting server")

	err := server.ListenAndServe()

	if err != nil {
		log.Err(err)
	}
}
