package api_test

import (
	"github.com/JarnoLahti/miniurl-fork/api"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestApi_AddUrl(t *testing.T) {
	const (
		payload        = `{"url": "http://example.com"}`
		expectedBody   = `{"url": "http://example.com", "hash": "test_value"}`
		expectedStatus = http.StatusOK
	)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/url", strings.NewReader(payload))
	rr := httptest.NewRecorder()

	r := httprouter.New()
	h := &strHandler{
		str: "test_value",
	}
	api.Bind(r, h)
	r.ServeHTTP(rr, req)

	assert.Equal(t, expectedStatus, rr.Result().StatusCode)
	body, err := io.ReadAll(rr.Result().Body)
	require.NoError(t, err)
	assert.JSONEq(t, expectedBody, string(body))

}

type strHandler struct {
	str string
}

func (m *strHandler) AddUrl(url string) (hash string, err error) {
	return m.str, nil
}
