package github

import (
	"bytes"
	"github.com/caddyserver/caddy"
	"github.com/caddyserver/caddy/caddyhttp/httpserver"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"github.com/h2non/bimg"
)

func init() {
	log.Println("RegisterPlugin")
	caddy.RegisterPlugin("webp", caddy.Plugin{
		ServerType: "http",
		Action:     setup,
	})
}

func setup(c *caddy.Controller) error {
	log.Println("setupFunc")
	h := handler{}
	//for c.Next() {
	//
	//}
	httpserver.GetConfig(c).AddMiddleware(func(next httpserver.Handler) httpserver.Handler {
		h.next = next
		return h
	})
	return nil
}

type handler struct {
	next httpserver.Handler
}

func (s handler) ServeHTTP(w http.ResponseWriter, r *http.Request) (int, error) {
	accept := r.Header.Get("Accept")
	if !(strings.Contains(accept, "image/webp") || strings.Contains(accept, "*/*")) {
		return s.next.ServeHTTP(w, r)
	}
	resp := &response{}
	i, err := s.next.ServeHTTP(resp, r)
	if err != nil {
		return i, err
	}

	newImage, err := bimg.NewImage(resp.Body.Bytes()).Convert(bimg.WEBP)
	if err != nil {
		log.Error(err)
		return s.next.ServeHTTP(w, r)
	}

	w.Header().Set("Content-Type", "image/webp")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(newImage)
	if err != nil {
		log.Error(err)
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
