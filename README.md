# Caddy Webp
Automatically convert many images format to webp format.  
This caddy plugin use `github.com/h2non/bimg` for fast high-level image processing using [libvips](https://github.com/jcupitt/libvips) via C bindings.

## Prerequisites
- [libvips](https://github.com/libvips/libvips) 8.3+ (8.8+ recommended)

## Example
On Caddyfile
```Caddyfile
localhost {
    webp # Enable plugin

    root * /var/www/
    encode gzip
    file_server
}
```

With [Dockerfile](example-caddy-build/Dockerfile)
