package logging

import (
	"github.com/logo-user-management/app/render"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Middleware(logger *logrus.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rvr := recover(); rvr != nil {
					logger.Error("Something bad happened\n", rvr)
					render.Respond(w, http.StatusInternalServerError, render.Message("Something Bad Happened"))
				}
			}()

			start := time.Now()
			next.ServeHTTP(w, r)
			logger.WithFields(logrus.Fields{
				"method":   r.Method,
				"path":     r.URL.EscapedPath(),
				"duration": time.Since(start),
			}).Info("Request finished")
		}

		return http.HandlerFunc(fn)
	}
}
