PWD := $(shell pwd)

## gen swagger
gen-swagger:
	swag init -g ./internal/presentation/http/http.go -o resources/docs --parseDependency true
	
# bind go bin
go-bind:
	export GO_PATH=~/go
	export PATH=$PATH:/$GO_PATH/bin

go-build:
	GOOS=linux CGO_ENABLED=0 go build -o main ./cmd/main.go

deploy:
	GOOS=linux CGO_ENABLED=0 go build -o main ./cmd/main.go
	
	# Stop and remove existing containers
	docker container stop product-service-1 || true
	docker container rm product-service-1 || true
	docker container stop product-service-2 || true
	docker container rm product-service-2 || true
	docker container stop product-service-3 || true
	docker container rm product-service-3 || true

	# Build the Docker image
	docker build -t product-service .

	# Connect the service to the network
	docker network connect kong-net product-service-1 || true
	docker network connect kong-net product-service-2 || true
	docker network connect kong-net product-service-3 || true

	# Run the service on 3 ports with auto-restart enabled
	docker run -d --restart unless-stopped --network=kong-net -e SERVER_HTTP_PORT=7002 -p 7002:7002 -v ./:/app --name product-service-1 product-service
	docker run -d --restart unless-stopped --network=kong-net -e SERVER_HTTP_PORT=7003 -p 7003:7003 -v ./:/app --name product-service-2 product-service
	docker run -d --restart unless-stopped --network=kong-net -e SERVER_HTTP_PORT=7004 -p 7004:7004 -v ./:/app --name product-service-3 product-service
