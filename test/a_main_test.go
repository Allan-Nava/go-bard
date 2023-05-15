package test

import (
	"os"
	"testing"

	"github.com/Allan-Nava/go-bard/configuration"
	"github.com/Allan-Nava/go-bard/env"
	"github.com/Allan-Nava/go-bard/gobard"
)

func TestMain(m *testing.M) {
	if os.Getenv("APP_ENV") == "" {
		err := os.Setenv("APP_ENV", "test")
		if err != nil {
			panic("could not set test env")
		}
	}
	env.Load()
	m.Run()
}

func GetGoBard() gobard.IGoBard {
	//
	configuration := configuration.GetConfiguration()
	g := gobard.NewGoBoard(
		configuration,
	)
	//
	return g
}
