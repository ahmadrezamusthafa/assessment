package respwriter

import (
	"github.com/ahmadrezamusthafa/assessment/common/errors"
	"github.com/json-iterator/go"
	"net/http"
	"time"
)

func New() *Module {
	return &Module{
		start: time.Now(),
	}
}

func (m *Module) SuccessWriter(writer http.ResponseWriter, status int, data interface{}) {
	response := Response{
		ProcessTime: time.Since(m.start).Seconds(),
		Data:        data,
		IsSuccess:   true,
	}
	jsonByte, err := jsoniter.Marshal(response)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	writer.Write(jsonByte)
}

func (m *Module) ErrorWriter(writer http.ResponseWriter, status int, lang string, err error) {
	assessmentError := errors.ParseError(lang, err)
	response := Response{
		ProcessTime: time.Since(m.start).Seconds(),
		IsSuccess:   false,
		Error: ErrorResponse{
			Code:    assessmentError.Code,
			Message: assessmentError.Detail,
			Traces:  assessmentError.Traces,
		},
	}
	jsonByte, err := jsoniter.Marshal(response)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	writer.Write(jsonByte)
}