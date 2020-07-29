package magazinegun

import (
	"github.com/ahmadrezamusthafa/assessment/common/errors"
	"github.com/ahmadrezamusthafa/assessment/common/respwriter"
	"github.com/ahmadrezamusthafa/assessment/config"
	"github.com/ahmadrezamusthafa/assessment/domain/usecase/magazinegun"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Handler struct {
	Config          config.Config                `inject:"config"`
	MagazineService *magazinegun.MagazineService `inject:"magazineService"`
}

func (h *Handler) AddMagazine(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		param      Param
		ctx        = r.Context()
		respWriter = respwriter.New()
	)
	defer func() {
		if err != nil {
			respWriter.ErrorWriter(w, errors.GetHttpStatus(err), "en", err)
		}
	}()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = errors.ReadDataError
		return
	}
	err = jsoniter.Unmarshal(body, &param)
	if err != nil {
		err = errors.AddTrace(err)
		return
	}

	err = h.MagazineService.AddMagazine(ctx, param.Name, param.Qty)
	if err != nil {
		err = errors.AddTrace(err)
		return
	}
	respWriter.SuccessWriter(w, http.StatusOK, param)
}

func (h *Handler) AddMagazineBullet(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		param      Param
		ctx        = r.Context()
		respWriter = respwriter.New()
	)
	defer func() {
		if err != nil {
			respWriter.ErrorWriter(w, errors.GetHttpStatus(err), "en", err)
		}
	}()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = errors.ReadDataError
		return
	}
	err = jsoniter.Unmarshal(body, &param)
	if err != nil {
		err = errors.AddTrace(err)
		return
	}

	err = h.MagazineService.AddMagazineBullet(ctx, param.ID, param.Qty)
	if err != nil {
		err = errors.AddTrace(err)
		return
	}
	respWriter.SuccessWriter(w, http.StatusOK, param)
}

func (h *Handler) AttachMagazine(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		id         string
		ctx        = r.Context()
		queries    = r.URL.Query()
		respWriter = respwriter.New()
	)
	defer func() {
		if err != nil {
			respWriter.ErrorWriter(w, errors.GetHttpStatus(err), "en", err)
		}
	}()

	if val, ok := queries["id"]; ok {
		id = val[0]
	}
	err = h.MagazineService.AttachMagazine(ctx, id)
	if err != nil {
		err = errors.AddTrace(err)
		return
	}
	respWriter.SuccessWriter(w, http.StatusOK, id)
}

func (h *Handler) DetachMagazine(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		id         string
		ctx        = r.Context()
		queries    = r.URL.Query()
		respWriter = respwriter.New()
	)
	defer func() {
		if err != nil {
			respWriter.ErrorWriter(w, errors.GetHttpStatus(err), "en", err)
		}
	}()

	if val, ok := queries["id"]; ok {
		id = val[0]
	}
	err = h.MagazineService.DetachMagazine(ctx, id)
	if err != nil {
		err = errors.AddTrace(err)
		return
	}
	respWriter.SuccessWriter(w, http.StatusOK, id)
}

func (h *Handler) Verify(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		ctx        = r.Context()
		respWriter = respwriter.New()
	)
	defer func() {
		if err != nil {
			respWriter.ErrorWriter(w, errors.GetHttpStatus(err), "en", err)
		}
	}()

	magazine, err := h.MagazineService.Verify(ctx)
	if err != nil {
		err = errors.AddTrace(err)
		return
	}
	respWriter.SuccessWriter(w, http.StatusOK, magazine)
}

func (h *Handler) ShotBullet(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		qty        int
		ctx        = r.Context()
		queries    = r.URL.Query()
		respWriter = respwriter.New()
	)
	defer func() {
		if err != nil {
			respWriter.ErrorWriter(w, errors.GetHttpStatus(err), "en", err)
		}
	}()

	if val, ok := queries["qty"]; ok {
		qty, err = strconv.Atoi(val[0])
		if err != nil {
			err = errors.AddTrace(err)
			return
		}
	}

	magazine, err := h.MagazineService.ShotBullet(ctx, qty)
	if err != nil {
		err = errors.AddTrace(err)
		return
	}
	respWriter.SuccessWriter(w, http.StatusOK, magazine)
}
