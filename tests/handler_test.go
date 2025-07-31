package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/mourasjames/cep-temp-api/handler"
	"github.com/stretchr/testify/assert"
)

func TestWeatherHandlerInvalidCep(t *testing.T) {
	req := httptest.NewRequest("GET", "/weather?cep=abc123", nil)
	w := httptest.NewRecorder()

	handler.WeatherHandler(w, req)

	assert.Equal(t, 422, w.Code)
	assert.Contains(t, w.Body.String(), "invalid zipcode")
}
