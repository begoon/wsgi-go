ARG GO_VERSION=1.19

FROM golang:${GO_VERSION}-alpine AS build

RUN apk add --no-cache git

WORKDIR /src

COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ ./

RUN CGO_ENABLED=0 go test -v ./pkg/...

RUN \
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -installsuffix 'static' -o /wsgi ./cmd

FROM gcr.io/distroless/static AS final

LABEL maintainer="begoon"
USER nonroot:nonroot

COPY --from=build --chown=nonroot:nonroot /wsgi /wsgi

CMD ["/wsgi"]
