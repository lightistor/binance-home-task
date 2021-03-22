package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

type middleware func(http.Handler) http.Handler
type middlewares []middleware

type controller struct {
	logger        *log.Logger
	nextRequestID func() string
}
