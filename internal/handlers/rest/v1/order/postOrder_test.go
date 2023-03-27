package order

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/pablogolobaro/chequery/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_orderHandler_PostOrder(t *testing.T) {
	order := entity.Order{
		PointID: 1,
		Products: []entity.Product{
			{Name: "Meat", Quantity: 3, Price: 145},
			{Name: "vegetables", Quantity: 2, Price: 32},
			{Name: "Juice", Quantity: 1, Price: 48}},
	}

	requestOrder, err := json.Marshal(&OrderCreateRequest{Order: order})
	if err != nil {
		t.Fatal(err)
	}

	handler := mockHandler()

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, urlOrders, bytes.NewReader(requestOrder))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, handler.PostOrder(c)) {
		assert.Equal(t, "{\"ids\":[1,2]}\n", rec.Body.String())
	}
}
