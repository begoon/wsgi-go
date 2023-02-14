all: build docker-build docker-run

serve:
	go run cmd/main.go cmd/client.go serve

docker-build:
	docker build -t wsgi-go --platform linux/amd64 .
	
docker-run:
	docker run -it --rm -p 8000:80 -e PORT=80 wsgi-go

build:
	GOOS=linux GOARCH=amd64 go build -o exe cmd/main.go cmd/client.go

client:
	go run cmd/main.go cmd/client.go 100 http://127.0.0.1:8000/api/process

test:
	go test ./wsgi

docker-tag:
	docker tag wsgi-go europe-docker.pkg.dev/iproov-chiro/chiro/chiro-api-go:latest

docker-push:
	docker push europe-docker.pkg.dev/iproov-chiro/chiro/chiro-api-go:latest
