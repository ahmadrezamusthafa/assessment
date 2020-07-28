package magazinegun

import (
	"github.com/ahmadrezamusthafa/assessment/config"
	"github.com/ahmadrezamusthafa/assessment/domain/usecase/magazinegun"
	"net/http"
)

type Handler struct {
	Config          config.Config                `inject:"config"`
	MagazineService *magazinegun.MagazineService `inject:"magazineService"`
}

func (h *Handler) AddMagazine(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) AddMagazineBullet(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) AttachMagazine(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) DetachMagazine(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Verify(w http.ResponseWriter, r *http.Request) {

}
