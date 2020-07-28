package server

import (
	"github.com/ahmadrezamusthafa/assessment/common/logger"
	commonhandlers "github.com/ahmadrezamusthafa/assessment/common/middleware"
	"github.com/ahmadrezamusthafa/assessment/config"
	"github.com/ahmadrezamusthafa/assessment/server/http/health"
	"github.com/gorilla/mux"
	"net/http"
)

type HttpServer struct {
	Config      config.Config
	RootHandler *RootHandler
	router      *mux.Router
}

type RootHandler struct {
	Health *health.Handler `inject:"healthHandler"`
}

func (svr *HttpServer) Serve() {
	route := createRouter(svr.RootHandler)
	commonHandlers := commonhandlers.New()
	middleware := commonhandlers.Chain(
		commonHandlers.RecoverHandler,
		commonHandlers.LoggingHandler,
	)

	logger.Info("Http server is serving on :" + svr.Config.HttpPort)
	err := http.ListenAndServe(":"+svr.Config.HttpPort, middleware.Then(route))
	if err != nil {
		logger.Err(err.Error())
	}
}

func createRouter(rh *RootHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health", rh.Health.HealthCheck).Methods("GET")
	return router
}
