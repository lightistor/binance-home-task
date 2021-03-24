package main

import (
	"net/http"
	"text/template"
)

type TopStatsSection struct {
	Title string
	Stats Stats
}

type PageData struct {
	PageTitle         string
	TopVolumes        TopStatsSection
	TopNumberOfTrades TopStatsSection
}

func (c *controller) index(w http.ResponseWriter, req *http.Request) {
	client := NewApiClient()
	service := NewMarketDataService(&client)

	marketData, _ := service.GetMarketData(
		&MarketDataQuery{
			VolumeQuoteAsset:         "BTC",
			NumberOfTradesQuoteAsset: "USDT",
		})

	tmpl := template.Must(template.ParseFiles("index.html"))

	data := PageData{
		PageTitle: "Binance Market Data",
		TopVolumes: TopStatsSection{
			Title: "TOP 5 highest volume over the last 24h for quote asset BTC",
			Stats: marketData.TopVolume,
		},
		TopNumberOfTrades: TopStatsSection{
			Title: "TOP 5 highest number of trades over the last 24h for quote asset USDT",
			Stats: marketData.TopNumberOfTrades,
		},
	}
	tmpl.Execute(w, data)
}
