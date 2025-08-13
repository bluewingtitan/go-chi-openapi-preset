package service

import (
	"github.com/bluewingtitan/go-chi-openapi-preset/config"
	"github.com/bluewingtitan/go-chi-openapi-preset/logging"
	"github.com/rs/zerolog"
)

type ExampleService struct {
	Logger *zerolog.Logger
	Config *config.Config
}

func NewExampleService(conf *config.Config) *ExampleService {
	return &ExampleService{
		Logger: logging.GetLoggerFor("example-service"),
		Config: conf,
	}
}

func (s *ExampleService) GetExample() string {
	s.Logger.Info().Msg("GetExample")

	return "hello, world!"
}
