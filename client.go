package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type ApiClient interface {
	GetExchangeInfo() (*ExchangeInfoResponse, error)
	GetTickerChangeStatistics(symbol string) ([]*TickerChangeStatics, error)
}

type client struct{}

func NewApiClient() ApiClient {
	return &client{}
}

func (c *client) GetExchangeInfo() (*ExchangeInfoResponse, error) {
	response, err := http.Get(apiBaseUrl + "/api/v3/exchangeInfo")

	if err != nil {
		log.Fatal(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var exchangeInfo ExchangeInfoResponse
	err = json.Unmarshal(responseData, &exchangeInfo)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Completed request to get exchange info")
	return &exchangeInfo, nil
}

func (c *client) GetTickerChangeStatistics(symbol string) ([]*TickerChangeStatics, error) {
	var url string
	var isArray bool
	if symbol != "" {
		log.WithField("symbol", symbol).Infof("Get ticker change statistics for %s", symbol)
		url = apiBaseUrl + "/api/v3/ticker/24hr?symbol=" + symbol
		isArray = false
	} else {
		log.Info("Get ticker change statistics for all symbols")
		url = apiBaseUrl + "/api/v3/ticker/24hr"
		isArray = true
	}

	response, err := http.Get(url)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	var tickerChangeStatics []TickerChangeStatics
	if isArray {
		err = json.Unmarshal(responseData, &tickerChangeStatics)
	} else {
		var item TickerChangeStatics
		err = json.Unmarshal(responseData, &item)
		tickerChangeStatics = []TickerChangeStatics{item}
	}
	if err != nil {
		log.Error(err)
		return nil, err
	}

	log.WithField("count", len(tickerChangeStatics)).Info("Completed request to get ticker change statistics")

	var result = []*TickerChangeStatics{}
	for i := range tickerChangeStatics {
		result = append(result, &tickerChangeStatics[i])
	}
	return result, nil
}
