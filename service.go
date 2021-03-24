package main

import (
	"sort"

	log "github.com/sirupsen/logrus"
)

const (
	NO_VALUE  string = ""
	TOP_ITEMS int    = 5
)

type MarketDataQuery struct {
	VolumeQuoteAsset         string
	NumberOfTradesQuoteAsset string
}

type MarketData struct {
	TopVolume         Stats
	TopNumberOfTrades Stats
}

type Stats []*TickerChangeStatics

func (s Stats) Len() int      { return len(s) }
func (s Stats) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// ByVolume implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Stats value.
type ByVolume struct{ Stats }

func (s ByVolume) Less(i, j int) bool { return s.Stats[i].Volume > s.Stats[j].Volume }

// ByVolume implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Stats value.
type ByNumberOfTrades struct{ Stats }

func (s ByNumberOfTrades) Less(i, j int) bool { return s.Stats[i].Tradecount > s.Stats[j].Tradecount }

type MarketDataService interface {
	GetMarketData(query *MarketDataQuery) (*MarketData, error)
}

type service struct {
	client ApiClient
}

func NewMarketDataService(c *ApiClient) MarketDataService {
	return &service{
		client: *c,
	}
}

func (s *service) GetMarketData(q *MarketDataQuery) (*MarketData, error) {
	info, err := s.client.GetExchangeInfo()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Infof("Get %d symbols from exchange info", len(info.Symbols))
	metadata := make(map[string]Symbol, len(info.Symbols))
	for i := range info.Symbols {
		s := info.Symbols[i]
		metadata[s.Symbol] = s
	}

	tickerStats, err := s.client.GetTickerChangeStatistics(NO_VALUE)
	if err != nil {
		log.Error("Error occurred while getting ticker change statistics")
		return nil, err
	}

	var volumes []*TickerChangeStatics
	var numberOfTrades []*TickerChangeStatics
	for _, t := range tickerStats {
		s := metadata[t.Symbol]
		if s.Quoteasset == q.VolumeQuoteAsset {
			volumes = append(volumes, t)
		}
		if s.Quoteasset == q.NumberOfTradesQuoteAsset {
			numberOfTrades = append(numberOfTrades, t)
		}
	}
	log.WithField("quoteAsset", q.VolumeQuoteAsset).Infof("Found %d symbols to sort by volume", len(volumes))
	log.WithField("quoteAsset", q.NumberOfTradesQuoteAsset).Infof("Found %d symbols to sort by number of trades", len(numberOfTrades))

	sort.Sort(ByVolume{volumes})
	sort.Sort(ByNumberOfTrades{numberOfTrades})

	return &MarketData{
		TopVolume:         volumes[:TOP_ITEMS],
		TopNumberOfTrades: numberOfTrades[:TOP_ITEMS],
	}, nil
}
