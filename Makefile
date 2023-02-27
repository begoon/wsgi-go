all: test docker-build docker-build docker-run

local: test serve

test:
	go test ./pkg/...

serve:
	go run ./cmd

build:
	go build -o wsgi ./cmd

docker-build:
	docker build -t wsgi-go --platform linux/amd64 .
	
docker-run:
	docker run -it --rm -p 8000:80 -e PORT=80 wsgi-go

build-linux:
	GOOS=darwin GOARCH=amd64 go build -o wsgi ./cmd/

client:
	go run ./cmd 100 http://127.0.0.1:8000/api/process

docker-tag:
	docker tag wsgi-go europe-docker.pkg.dev/iproov-chiro/chiro/chiro-api-go:latest

docker-push:
	docker push europe-docker.pkg.dev/iproov-chiro/chiro/chiro-api-go:latest

push:
	git push origin main
	git push iproov main
