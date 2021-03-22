package main

import (
	"flag"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

type key int

const (
	requestIDKey key = 0
)

var (
	apiUrl        string
	listenAddress string
)

func main() {

	flag.StringVar(&apiUrl, "api-url", "https://api.binance.com", "public Rest API for Binance")
	flag.StringVar(&listenAddress, "listen-addres", ":8080", "server listen address")
	flag.Parse()

	c := &controller{logger: log.New(), nextRequestID: func() string { return strconv.FormatInt(time.Now().UnixNano(), 36) }}

	router := http.NewServeMux()

	router.Handle("/metrics", promhttp.Handler())

	health := healtcheck()
	router.HandleFunc("/live", health.LiveEndpoint)
	router.HandleFunc("/ready", health.ReadyEndpoint)

	log.WithField("listen-addres", listenAddress).Info("Starting HTTP server")
	log.Fatal(http.ListenAndServe(listenAddress, (middlewares{c.tracing, c.logging}).apply(router)))
}

func (mws middlewares) apply(hdlr http.Handler) http.Handler {
	if len(mws) == 0 {
		return hdlr
	}
	return mws[1:].apply(mws[0](hdlr))
}
