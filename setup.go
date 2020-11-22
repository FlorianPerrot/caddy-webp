package caddywebp

import (
	"github.com/caddyserver/caddy"
	"github.com/caddyserver/caddy/caddyhttp/httpserver"
)

func init() {
	caddy.RegisterPlugin("webp", caddy.Plugin{
		ServerType: "http",
		Action:     Setup,
	})

	httpserver.RegisterDevDirective("webp", "rewrite")
}

func Setup(c *caddy.Controller) error {
	h := handler{}

	httpserver.GetConfig(c).AddMiddleware(func(next httpserver.Handler) httpserver.Handler {
		h.next = next
		return h
	})

	return nil
}
