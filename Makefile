all: build docker-build docker-run

serve:
	go run cmd/main.go cmd/client.go serve

docker-build:
	docker build -t wsgi-go .
	
docker-run:
	docker run -it --rm -p 8000:80 -e PORT=80 wsgi-go

build:
	GOOS=linux GOARCH=amd64 go build -o exe cmd/main.go cmd/client.go

client:
	go run cmd/main.go cmd/client.go 100 http://127.0.0.1:8000/api/process
