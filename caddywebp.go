package caddywebp

import (
	"bytes"
	"github.com/caddyserver/caddy/caddyhttp/httpserver"
	"github.com/h2non/bimg"
	"net/http"
	"strings"
)

type handler struct {
	next httpserver.Handler
}

func (s handler) ServeHTTP(w http.ResponseWriter, r *http.Request) (int, error) {
	accept := r.Header.Get("Accept")
	if !strings.Contains(accept, "image/webp") {
		return s.next.ServeHTTP(w, r)
	}

	resp := &response{}
	i, err := s.next.ServeHTTP(resp, r)
	if err != nil {
		return i, err
	}

	newImage, err := bimg.NewImage(resp.Body.Bytes()).Convert(bimg.WEBP)
	if err != nil {
		return s.next.ServeHTTP(w, r)
	}

	w.Header().Set("Content-Type", "image/webp")
	_, err = w.Write(newImage)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

type response struct {
	header http.Header
	Body   bytes.Buffer
}

func (s *response) Header() http.Header {
	return http.Header{}
}

func (s *response) Write(data []byte) (int, error) {
	s.Body.Write(data)
	return len(data), nil
}

func (s *response) WriteHeader(i int) {
	return
}
