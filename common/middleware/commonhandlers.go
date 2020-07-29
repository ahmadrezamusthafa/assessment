package commonhandlers

import (
	"context"
	"fmt"
	"github.com/ahmadrezamusthafa/assessment/common/logger"
	"github.com/ahmadrezamusthafa/assessment/common/respwriter"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"net/http/httptest"
	"time"
)

type CommonHandlers struct {
	router *mux.Router
}

func New() *CommonHandlers {
	return &CommonHandlers{
		router: mux.NewRouter(),
	}
}

func (c *CommonHandlers) RecoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Err("Panic: %v", err)
				respWriter := respwriter.New()
				respWriter.ErrorWriter(w, http.StatusInternalServerError, "en", fmt.Errorf("%v", err))
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func (c *CommonHandlers) LoggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.FromString(r.Header.Get("x-request-id"))
		if err != nil {
			id = uuid.NewV4()
		}

		ctx := context.WithValue(r.Context(), "x-request-id", id.String())
		r = r.WithContext(ctx)
		logger.Info("Request #%s - [%s] %q", id.String(), r.Method, r.URL.Path)
		start := time.Now()

		resp := httptest.NewRecorder()
		next.ServeHTTP(resp, r)

		for k, v := range resp.HeaderMap {
			w.Header()[k] = v
		}
		w.WriteHeader(resp.Code)
		s := resp.Body.String()
		code := resp.Code / 1e2
		elapsed := time.Since(start)
		switch code {
		case 2:
			logger.Info("Response #%s - [%s] %q - Result with status - %v - took %s", id.String(), r.Method, r.URL.Path, resp.Code, elapsed)
		case 4:
			logger.Warn("Response #%s - [%s] %q -  %s - Failed with status - %v - took %s", id.String(), r.Method, r.URL.Path, s, resp.Code, elapsed)
		case 5:
			logger.Warn("Response #%s - [%s] %q - %s - Failed with status - %v - took %s", id.String(), r.Method, r.URL.Path, s, resp.Code, elapsed)
		default:
			logger.Info("Response #%s - [%s] %q - %s - with status - %v - took %s", id.String(), r.Method, r.URL.Path, s, resp.Code, elapsed)
		}
		resp.Body.WriteTo(w)
	}
	return http.HandlerFunc(fn)
}

func Chain(handlers ...alice.Constructor) alice.Chain {
	return alice.New(handlers...)
}
