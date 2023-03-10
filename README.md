# Experimental code for Chiro API in Golang

The endpoint is built and deployed automatically when its `main` branch is pushed.

<https://chiro-api-go-ixozpwh3jq-nw.a.run.app>

## Goals

To try out the following technologies (the list is being updated frequently):

- [x] MongoDB integration
- [x] Hashing (SHA256, ...)
- [x] GCM/AES encryption
- [x] BASE64
- [x] REST API development in Go
  - [x] unit testing (running test within the image build process under docker)
  - [x] parse the request parameters from the URL and body (JSON)
- [x] Go [Echo](https://echo.labstack.com/) web framework
- [x] Echo middleware:
  - [x] API Key
  - [x] WebSockets
  - [x] unique request id injection
  - [x] logger
    - [x] line-based
    - [x] structured (JSON-based)
  - [x] static files (statuc file or/and website can be built into executable)
  - [x] request/response intercept
  - [x] redirect
  - [x] trailing slash
- [x] Automatic Swagger/OpenAPI documentation with [EchoSwagger](https://pkg.go.dev/github.com/pangpanglabs/echoswagger/v2)
- [x] Dockerize Go applications to tiny images from `scratch` or `distroless` based image. The current final endpoint deployment image size is only 13MB. The image literally contains one single statically linked executable, which embeds everything, including static file or/and a built-in website.
- [x] stress-test the endpoint (up to 5000 concurrent requests) via a concurrent test client `client.go`
