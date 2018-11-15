package controllers_test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"tax-calculator/core"
	"testing"
)

func TestCommonController(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()
	core.Globals.Router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Equal(t, "Welcome to Tax Calculator API", resp.Body.String())
}
