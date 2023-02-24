package check

import (
	"net/http"
)

func (c *checkHandler) UpdateChecksStatus(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()

	ids := query["ids"]

	err := c.useCases.SetChecksStatusPrinted(request.Context(), ids)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte((err.Error())))

		return
	}
	writer.WriteHeader(http.StatusOK)
}
