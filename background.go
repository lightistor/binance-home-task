package main

import (
	"github.com/jasonlvhit/gocron"
	"github.com/shopspring/decimal"
)

type background struct {
	service MarketDataService
	state   map[string]decimal.Decimal
}

type BackgroundService interface {
	Start()
}

func NewBackgroundService(s *MarketDataService) BackgroundService {
	return &background{
		service: *s,
		state:   make(map[string]decimal.Decimal),
	}
}

func (b *background) Start() {
	gocron.Every(10).Second().Do(b.backgroundTask)
	<-gocron.Start()
}

func (b *background) backgroundTask() {

}
