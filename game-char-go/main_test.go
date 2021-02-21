package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/achmadsy/game-char-go/routes"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	main()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/user/1/chars/all", nil)
	routes.Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
