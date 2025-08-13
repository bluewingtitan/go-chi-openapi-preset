package logging

import (
	"errors"
	"fmt"
	"github.com/bluewingtitan/go-chi-openapi-preset/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	stdlog "log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func newRollingFile(config *config.Config) io.Writer {
	absPath, err := filepath.Abs(config.Logging.Directory)

	if err != nil {
		log.Fatal().Err(err).Str("directory", absPath).
			Str("filename", config.Logging.Filename).
			Msg("Unable to resolve absolute path")
		return nil
	}

	if err := os.MkdirAll(absPath, 0744); err != nil {
		log.Error().Err(err).Str("path", absPath).Msg("can't create log directory")
		return nil
	}

	log.Info().Str("path", absPath).Str("filename", config.Logging.Filename).Msg("Created new log file")

	return &lumberjack.Logger{
		Filename:   path.Join(absPath, config.Logging.Filename),
		MaxBackups: config.Logging.FileCount, // files
		MaxSize:    config.Logging.MaxSize,   // megabytes
		MaxAge:     config.Logging.MaxAge,    // days
	}
}

var _logger *zerolog.Logger = nil

func InitializeLogger(config *config.Config) (*zerolog.Logger, error) {
	if config == nil {
		return nil, errors.New("config is nil")
	}

	logLevel, err := zerolog.ParseLevel(strings.ToLower(config.Logging.Level))
	if err != nil {
		logLevel = zerolog.InfoLevel
	}

	var writers []io.Writer

	if config.Logging.EnableConsole {
		cw := zerolog.ConsoleWriter{Out: os.Stderr}
		cw.FormatTimestamp = func(i interface{}) string {
			return fmt.Sprintf("%v", i)
		}
		writers = append(writers, cw)
	}
	if config.Logging.EnableFile {
		writers = append(writers, newRollingFile(config))
	}

	mw := io.MultiWriter(writers...)

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	logger := zerolog.New(mw).With().Timestamp().Logger()

	logger = logger.Level(logLevel)

	log.Logger = logger
	_logger = &logger
	stdlog.SetFlags(0)
	stdlog.SetOutput(logger)

	return &logger, nil
}

func GetLoggerFor(moduleName string) *zerolog.Logger {
	if _logger == nil {
		logger := log.Logger.With().Str("module", moduleName).Logger()
		return &logger
	}

	logger := _logger.With().Str("module", moduleName).Logger()
	return &logger
}
