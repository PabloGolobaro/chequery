package order

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func Test_orderHandler_PostOrder(t *testing.T) {
	type product struct {
		Label    string  `json:"label,omitempty"`
		Quantity int     `json:"quantity,omitempty"`
		Price    float32 `json:"price,omitempty"`
	}

	order := struct {
		Products []product `json:"products,omitempty"`
	}{
		Products: []product{
			{Label: "Meat", Quantity: 3, Price: 145.4},
			{Label: "vegetables", Quantity: 2, Price: 32.5},
			{Label: "Juice", Quantity: 1, Price: 48},
		},
	}

	marshal, err := json.Marshal(&order)
	if err != nil {
		t.Fatal(err)
	}

	requestOrder, err := json.Marshal(&OrderCreateRequest{Order: string(marshal)})
	if err != nil {
		t.Fatal(err)
	}

	handler := mockHandler()

	v := make(url.Values)

	v.Set("point_id", "18462")

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, urlOrders+"/?"+v.Encode(), bytes.NewReader(requestOrder))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, handler.PostOrder(c)) {
		assert.Equal(t, "{\"ids\":[1,2]}\n", rec.Body.String())
	}
}
