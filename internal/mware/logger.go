package mware

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

func WithLogging(h http.Handler) http.Handler {
	logFn := func(w http.ResponseWriter, r *http.Request) {

		logger, err := zap.NewDevelopment()

		if err != nil {
			h.ServeHTTP(w, r)
			return
		}
		defer logger.Sync()

		sugar := logger.Sugar()

		start := time.Now()

		uri := r.RequestURI

		method := r.Method

		h.ServeHTTP(w, r)

		duration := time.Since(start)

		sugar.Infoln(
			"uri", uri,
			"method", method,
			"duration", duration,
		)
	}
	return http.HandlerFunc(logFn)
}
