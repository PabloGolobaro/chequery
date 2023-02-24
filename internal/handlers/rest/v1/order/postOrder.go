package order

import (
	"encoding/json"
	"io"
	"net/http"
)

func (o *orderHandler) PostOrder(writer http.ResponseWriter, request *http.Request) {
	body := request.Body
	defer body.Close()

	orderBytes, err := io.ReadAll(body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte((err.Error())))

		return
	}

	if !json.Valid(orderBytes) {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte((err.Error())))

		return
	}

	err = o.checkUseCases.CreateChecks(request.Context(), string(orderBytes))
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte((err.Error())))

		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte("Order registered successfully"))
}
