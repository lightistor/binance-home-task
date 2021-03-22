package main

import (
	"time"

	"github.com/heptiolabs/healthcheck"
)

func healtcheck() healthcheck.Handler {
	health := healthcheck.NewHandler()

	// App is not ready if can't resolve the upstream dependency in DNS.
	health.AddReadinessCheck(
		"upstream-dep-dns",
		healthcheck.DNSResolveCheck(apiUrl, 50*time.Millisecond))

	// Our app is not happy if we've got more than 100 goroutines running.
	health.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(100))

	return health
}
