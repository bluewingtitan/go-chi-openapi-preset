package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type LoggingConfig struct {
	Level         string `yaml:"level"`
	EnableConsole bool   `yaml:"enable_console"`
	EnableFile    bool   `yaml:"enable_file"`
	Directory     string `yaml:"directory"`
	Filename      string `yaml:"filename"`
	// age in days
	MaxAge int `yaml:"max_age"`
	// size in mb
	MaxSize   int `yaml:"max_size"`
	FileCount int `yaml:"file_count"`
}

type TimeoutConfig struct {
	Handling   int `yaml:"handling"`
	Read       int `yaml:"read"`
	ReadHeader int `yaml:"read_header"`
	Write      int `yaml:"write"`
	Idle       int `yaml:"idle"`
}

type Config struct {
	Address        string        `yaml:"address"`
	AllowedOrigins []string      `yaml:"allowed_origins"`
	Timeout        TimeoutConfig `yaml:"timeout"`
	Logging        LoggingConfig `yaml:"logging"`
}

func GetDefaultConfig() Config {
	return Config{
		Address: ":8080",
		AllowedOrigins: []string{
			"http://localhost:8080",
		},
		Timeout: TimeoutConfig{
			Handling:   10,
			Read:       1,
			ReadHeader: 2,
			Write:      5,
			Idle:       30,
		},
		Logging: LoggingConfig{
			Level:         "info",
			Directory:     "logs",
			Filename:      "service.log",
			MaxAge:        7,
			MaxSize:       512,
			FileCount:     4,
			EnableFile:    true,
			EnableConsole: true,
		},
	}
}

func LoadConfig(filename string) (Config, error) {
	config := GetDefaultConfig()

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		contents, _ := yaml.Marshal(config)

		if err != nil {
			return config, err
		}

		err := os.WriteFile(filename, contents, 0644)

		if err != nil {
			return config, err
		}

		return config, nil
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
