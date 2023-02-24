package printer

import (
	"net/http"
	"os"
)

func (p *printerHandler) GetChecks(writer http.ResponseWriter, request *http.Request) {
	printerID := request.URL.Query().Get("printer_id")
	if printerID == "" {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("No printer_id query param"))

		return
	}

	checkFileNames, err := p.useCases.GetChecks(request.Context(), printerID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte((err.Error())))

		return
	}
	for _, checkFileName := range checkFileNames {
		bytes, err := os.ReadFile(checkFileName)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte((err.Error())))

			return
		}
		writer.Write(bytes)
	}
	writer.WriteHeader(http.StatusOK)
}
