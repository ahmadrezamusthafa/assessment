package server

import (
	"github.com/ahmadrezamusthafa/assessment/common/logger"
	commonhandlers "github.com/ahmadrezamusthafa/assessment/common/middleware"
	"github.com/ahmadrezamusthafa/assessment/config"
	"github.com/ahmadrezamusthafa/assessment/server/http/health"
	"github.com/ahmadrezamusthafa/assessment/server/http/magazinegun"
	"github.com/gorilla/mux"
	"net/http"
)

type HttpServer struct {
	Config      config.Config
	RootHandler *RootHandler
	router      *mux.Router
}

type RootHandler struct {
	Health   *health.Handler      `inject:"healthHandler"`
	Magazine *magazinegun.Handler `inject:"magazineHandler"`
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
	magazinePath := router.PathPrefix("/magazine").Subrouter()
	magazinePath.HandleFunc("/add_magazine", rh.Magazine.AddMagazine).Methods("POST")
	magazinePath.HandleFunc("/add_magazine_bullet", rh.Magazine.AddMagazineBullet).Methods("POST")
	magazinePath.HandleFunc("/attach_magazine", rh.Magazine.AttachMagazine).Methods("GET")
	magazinePath.HandleFunc("/detach_magazine", rh.Magazine.DetachMagazine).Methods("GET")
	magazinePath.HandleFunc("/verify", rh.Magazine.Verify).Methods("GET")
	magazinePath.HandleFunc("/shot", rh.Magazine.ShotBullet).Methods("GET")

	return router
}
