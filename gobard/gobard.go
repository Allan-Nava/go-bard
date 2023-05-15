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

func NewGoBoard(configuration *configuration.Configuration) IGoBard {
	g := &gobard{
		configuration: configuration,
	}
	g.restClient = resty.New()
	//g.restClient.SetBaseURL(routes.BASE_URL)
	/*
		g.restClient.SetHeader("Referer", "https://bard.google.com/")
		//"Host": "bard.google.com",
		//"X-Same-Domain": "1",
		//"User-Agent": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36",
		//"Content-Type": "application/x-www-form-urlencoded;charset=UTF-8",
		//"Origin": "https://bard.google.com",
		//"Referer": "https://bard.google.com/",
	*/
	g.restClient.SetHeader("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	g.restClient.SetHeader("Host", "bard.google.com")
	g.restClient.SetHeader("X-Same-Domain", "1")
	g.restClient.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36")
	g.restClient.SetHeader("Origin", "https://bard.google.com")
	if configuration.IsDebug {
		g.restClient.SetDebug(true)
	}
	return g
}

// Resty Methods

func (o *gobard) restyPost(url string, body interface{}) (*resty.Response, error) {
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(body).
		Post(url)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *gobard) restyGet(url string, queryParams map[string]string) (*resty.Response, error) {
	resp, err := o.restClient.R().
		SetQueryParams(queryParams).
		Get(url)
	//
	if err != nil {
		return nil, err
	}
	return resp, nil
}
