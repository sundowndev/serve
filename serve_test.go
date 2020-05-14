package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w, err
}

func TestApi(t *testing.T) {
	assert := assert.New(t)
	parseFlags()
	r := newServer()
	RegisterStatic(r, "/", "./fixtures")

	t.Run("Serve", func(t *testing.T) {
		t.Run("Get index.html", func(t *testing.T) {
			res, err := performRequest(r, "GET", "/")

			body, _ := ioutil.ReadAll(res.Body)

			assert.Equal(nil, err, "should be equal")
			assert.Equal(200, res.Result().StatusCode, "should be equal")
			assert.Equal("<html><body><h1>Bonjour!</h1></body></html>", string(body), "should be equal")
		})

		t.Run("404 error - /notfound", func(t *testing.T) {
			res, err := performRequest(r, "GET", "/notfound")

			assert.Equal(nil, err, "should be equal")
			assert.Equal(404, res.Result().StatusCode, "should be equal")
		})
	})
}
