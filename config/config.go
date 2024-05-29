package config

import (
	"github.com/fatih/color"
	"github.com/spf13/viper"
)

const (
	DefaultServerRoot     = "localhost:9090"
	DefaultPort           = 9090
	DefaultConfigFilePath = "./config.json"
	DefaultIsDebug        = true
	DefaultDevMode        = true
)

type Configuration struct {
	ServerRoot string `json:"serverRoot" mapstructure:"serverRoot"`
	Port       int    `json:"port" mapstructure:"port"`
	IsDebug    bool   `json:"isDebug" mapstructure:"isDebug"`
	DevMode    bool   `json:"dev" mapstructure:"dev"`
}

// ReadConfigFile reads the configuration file and returns the configuration.
func ReadConfigFile(configFilePath string) (*Configuration, error) {
	if configFilePath == "" {
		viper.SetConfigFile(DefaultConfigFilePath)
	} else {
		viper.SetConfigFile(configFilePath)
	}

	viper.SetDefault("serverRoot", DefaultServerRoot)
	viper.SetDefault("port", DefaultPort)
	viper.SetDefault("isDebug", DefaultIsDebug)
	viper.SetDefault("dev", DefaultDevMode)

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return nil, err
	}

	configuration := Configuration{}

	err = viper.Unmarshal(&configuration)
	if err != nil {
		return nil, err
	}

	regular := color.New()
	regular.Printf(" ➜ ReadConfigFile: %+v\n", removeSecurityData(configuration))

	return &configuration, nil
}

func removeSecurityData(config Configuration) Configuration {
	clean := config
	return clean
}
