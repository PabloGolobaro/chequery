package check

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func Test_checkHandler_GetCheckPDF(t *testing.T) {
	handler := mockHandler()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(urlGetPDF)
	c.SetParamNames("check_id")
	c.SetParamValues("1")
	// Assertions
	if assert.NoError(t, handler.GetCheckPDF(c)) {
		assert.Equal(t, "pasha\r\n", rec.Body.String())
	}
}

func Test_checkHandler_GetGeneratedChecks(t *testing.T) {
	handler := mockHandler()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, urlGetGenerated, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handler.GetGeneratedChecks(c)) {
		assert.Equal(t, "{\"ids\":[1,3]}\n", rec.Body.String())
	}
}

func Test_checkHandler_UpdateChecksStatus(t *testing.T) {
	handler := mockHandler()

	q := make(url.Values)
	q.Add("id", "1")
	q.Add("id", "3")

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, urlCheck+"?"+q.Encode(), nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handler.UpdateChecksStatus(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
