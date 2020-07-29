package kitarastore

import (
	"github.com/ahmadrezamusthafa/assessment/common/errors"
	"github.com/ahmadrezamusthafa/assessment/common/respwriter"
	"github.com/ahmadrezamusthafa/assessment/config"
	"github.com/ahmadrezamusthafa/assessment/domain/usecase/order"
	"github.com/ahmadrezamusthafa/assessment/domain/usecase/product"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
)

type Handler struct {
	Config         config.Config           `inject:"config"`
	OrderService   *order.OrderService     `inject:"orderService"`
	ProductService *product.ProductService `inject:"productService"`
}

func (h *Handler) AddProduct(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		param      ProductParam
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

	err = h.ProductService.AddProduct(ctx, param.Code, param.Name, param.Qty)
	if err != nil {
		err = errors.AddTrace(err)
		return
	}
	respWriter.SuccessWriter(w, http.StatusOK, param)
}

func (h *Handler) AddProductQuantity(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		param      ProductParam
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

	err = h.ProductService.AddProductQuantity(ctx, param.ID, param.Qty)
	if err != nil {
		err = errors.AddTrace(err)
		return
	}
	respWriter.SuccessWriter(w, http.StatusOK, param)
}

func (h *Handler) DecreaseProductQuantity(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		param      ProductParam
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

	err = h.ProductService.DecreaseProductQuantity(ctx, param.ID, param.Qty)
	if err != nil {
		err = errors.AddTrace(err)
		return
	}
	respWriter.SuccessWriter(w, http.StatusOK, param)
}
