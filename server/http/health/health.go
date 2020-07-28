package health

import (
	"github.com/ahmadrezamusthafa/assessment/config"
	"net/http"
)

type Handler struct {
	Config config.Config `inject:"config"`
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
