package main

import (
	_ "github.com/FlorianPerrot/caddy-webp"
	"github.com/caddyserver/caddy/caddy/caddymain"
)

func main() {
	caddymain.EnableTelemetry = false
	caddymain.Run()
}