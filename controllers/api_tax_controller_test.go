package controllers_test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"tax-calculator/core"
	"testing"
)

func TestTaxAPIFindAllController(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/tax", nil)
	resp := httptest.NewRecorder()
	core.Globals.Router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.NotEqual(t, "null", resp.Body.String())
}

func TestTaxAPILoadController(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/tax/1", nil)
	resp := httptest.NewRecorder()
	core.Globals.Router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.NotEqual(t, "null", resp.Body.String())
}