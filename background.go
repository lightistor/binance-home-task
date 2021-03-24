package main

import (
	"sort"

	"github.com/jasonlvhit/gocron"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
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
	// get top number of trades
	byTradeCountSort := func(symbols []*SymbolData) {
		sort.Sort(ByTradeCount{symbols: symbols})
	}
	topNumberOfTrades, _ := b.service.GetTopSymbols(
		"USDT", TOP_LIMIT, byTradeCountSort)

	// get spreds
	var spreadTargets []string
	for _, v := range topNumberOfTrades {
		spreadTargets = append(spreadTargets, v.Symbol)
	}
	spreads, _ := b.service.GetSpreads(spreadTargets)

	// calc delta, print, update current
	var delta decimal.Decimal
	var deltaSign string
	for _, s := range spreads {
		if old, found := b.state[s.Symbol]; found {
			delta = s.Value.Add(old.Neg())
			if delta.IsPositive() {
				deltaSign = "+"
			} else if delta.IsNegative() {
				// delta has '-' already
				deltaSign = ""
			} else {
				// no change
				deltaSign = "="
			}

			log.Infof("%s: %s (%s%s)", s.Symbol, s.Value, deltaSign, delta)
		} else {
			log.Infof("%s: %s (n/a)", s.Symbol, s.Value)
		}

		b.state[s.Symbol] = s.Value
	}
}
