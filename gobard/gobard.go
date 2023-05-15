package gobard

import (
	"github.com/Allan-Nava/go-bard/configuration"
	"github.com/go-resty/resty/v2"
)

type IGoBard interface {
}

type gobard struct {
	configuration *configuration.Configuration
	restClient    *resty.Client
}

func NewFakeYou(configuration *configuration.Configuration) IGoBard {
	fk := &gobard{
		configuration: configuration,
	}
	fk.restClient = resty.New()
	//fk.restClient.SetBaseURL(routes.BASE_URL)
	/*
		fk.restClient.SetHeader("Referer", "https://bard.google.com/")
		//"Host": "bard.google.com",
		//"X-Same-Domain": "1",
		//"User-Agent": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36",
		//"Content-Type": "application/x-www-form-urlencoded;charset=UTF-8",
		//"Origin": "https://bard.google.com",
		//"Referer": "https://bard.google.com/",
	*/
	fk.restClient.SetHeader("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	fk.restClient.SetHeader("Host", "bard.google.com")
	fk.restClient.SetHeader("X-Same-Domain", "1")
	fk.restClient.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36")
	fk.restClient.SetHeader("Origin", "https://bard.google.com")
	if configuration.IsDebug {
		fk.restClient.SetDebug(true)
	}
	return fk
}
