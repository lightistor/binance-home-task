package main

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

func (e *ApiError) Error() string {
	return e.Message
}

type ExchangeInfoResponse struct {
	Timezone        string      `json:"timezone"`
	ServerTime      int64       `json:"timestamp"`
	RateLimits      []RateLimit `json:"rateLimits"`
	ExchangeFilters []struct{}  `json:"exchangeFilters"`
	Symbols         []Symbol    `json:"symbols"`
}

type RateLimit struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
}

type Symbol struct {
	Symbol                     string   `json:"symbol"`
	Status                     string   `json:"status"`
	Baseasset                  string   `json:"baseAsset"`
	Baseassetprecision         int      `json:"baseAssetPrecision"`
	Quoteasset                 string   `json:"quoteAsset"`
	Quoteprecision             int      `json:"quotePrecision"`
	Quoteassetprecision        int      `json:"quoteAssetPrecision"`
	Basecommissionprecision    int      `json:"baseCommissionPrecision"`
	Quotecommissionprecision   int      `json:"quoteCommissionPrecision"`
	Ordertypes                 []string `json:"orderTypes"`
	Icebergallowed             bool     `json:"icebergAllowed"`
	Ocoallowed                 bool     `json:"ocoAllowed"`
	Quoteorderqtymarketallowed bool     `json:"quoteOrderQtyMarketAllowed"`
	Isspottradingallowed       bool     `json:"isSpotTradingAllowed"`
	Ismargintradingallowed     bool     `json:"isMarginTradingAllowed"`
	Filters                    []Filter `json:"filters"`
	Permissions                []string `json:"permissions"`
}

type Filter struct {
	Filtertype       string `json:"filterType"`
	Minprice         string `json:"minPrice,omitempty"`
	Maxprice         string `json:"maxPrice,omitempty"`
	Ticksize         string `json:"tickSize,omitempty"`
	Multiplierup     string `json:"multiplierUp,omitempty"`
	Multiplierdown   string `json:"multiplierDown,omitempty"`
	Avgpricemins     int    `json:"avgPriceMins,omitempty"`
	Minqty           string `json:"minQty,omitempty"`
	Maxqty           string `json:"maxQty,omitempty"`
	Stepsize         string `json:"stepSize,omitempty"`
	Minnotional      string `json:"minNotional,omitempty"`
	Applytomarket    bool   `json:"applyToMarket,omitempty"`
	Limit            int    `json:"limit,omitempty"`
	Maxnumorders     int    `json:"maxNumOrders,omitempty"`
	Maxnumalgoorders int    `json:"maxNumAlgoOrders,omitempty"`
}

type TickerChangeStatics struct {
	Symbol             string `json:"symbol"`
	Pricechange        string `json:"priceChange"`
	Pricechangepercent string `json:"priceChangePercent"`
	Weightedavgprice   string `json:"weightedAvgPrice"`
	Prevcloseprice     string `json:"prevClosePrice"`
	Lastprice          string `json:"lastPrice"`
	Lastqty            string `json:"lastQty"`
	Bidprice           string `json:"bidPrice"`
	Askprice           string `json:"askPrice"`
	Openprice          string `json:"openPrice"`
	Highprice          string `json:"highPrice"`
	Lowprice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	Quotevolume        string `json:"quoteVolume"`
	Opentime           int64  `json:"openTime"`
	Closetime          int64  `json:"closeTime"`
	Firsttradeid       int    `json:"firstId"`
	Lasttradeid        int    `json:"lastId"`
	Tradecount         int    `json:"count"`
}

type OrderBook struct {
	Lastupdateid int        `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}
