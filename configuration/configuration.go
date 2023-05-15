package configuration

import (
	"github.com/caarlos0/env/v6"

	"github.com/Allan-Nava/go-bard/gobard"
)

type Configuration struct {
	IsDebug bool `env:"IS_DEBUG"`
	BaseUrl string
	BardApiKey string `env:"_BARD_API_KEY"`
	//RestClient *resty.Client
}

func GetConfiguration() *Configuration {
	configuration := Configuration{}
	err := env.Parse(&configuration)
	if err != nil {
		panic("failed to read configuration")
	}
	//
	configuration.BaseUrl = gobard.BASE_URL
	//
	return &configuration
}
