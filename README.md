# binance-home-task

## Tasks

1. Print the top 5 symbols with quote asset `BTC` and the highest volume over the last 24 hours in descending order.
2. Print the top 5 symbols with quote asset `USDT` and the highest number of trades over the last 24 hours in descending order.
3. Using the symbols from Q1, what is the total notional value of the top 200 bids and asks currently on each order book?
4. What is the price spread for each of the symbols from Q2?
5. Every 10 seconds print the result of Q4 and the absolute delta from the previous value for each symbol.
6. Make the output of Q5 accessible by querying http://localhost:8080/metrics using the Prometheus Metrics format.

### Notes

The [exchange information](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md#exchange-information) endpoint returns current exchange trading rules and symbol information. 
Will be used as a starting point to obtain data about available trading symbols.



## Quick Start

```
$ go build -o out/binancehometask && ./out/binancehometask
```

## Principles implemented

### Configuration parameters

The application doesn't store the hardcoded configuration parameters.

The default values are set but can be easily overridden via command line parameters.

In advanced example the config usually would be stored with environment variables.

```
$ ./out/binancehometask -h
Usage of ./out/binancehometask:
  -api-url string
        public Rest API for Binance (default "https://api.binance.com")
  -listen-addres string
        server listen address (default ":8080")
```

### Health check

The application has liveness `/live` and readiness `ready` probe handlers.

A failed liveness check indicates that the app is unhealthy, 
not some upstream dependency, and the app should be destroyed or restarted.

A failed readiness check indicates that the app is currently unable to serve requests
because of an upstream or some transient failure, and the app should no longer receive requests.

See the `health.go` to see the configured checks.

### Logging, Monitoring & Tracing

The Application is configured to log every request.

If request identifier is provided using header `X-Request-Id` it will be
logged as well, otherwise unique string is generated.

Monitoring is exposed in Prometheus format at'/metrics' endpoint.

### No graceful shutdown

The application doesn't use the graceful shutdown for simplicity.

Otherwise `http.Server` type and `sync.WaitGroup` can be utilised.
