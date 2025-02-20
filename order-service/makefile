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
	docker container stop order-service-1 || true
	docker container rm order-service-1 || true
	docker container stop order-service-2 || true
	docker container rm order-service-2 || true
	docker container stop order-service-3 || true
	docker container rm order-service-3 || true
	# Build the Docker image
	docker build -t order-service .
	# Connect the service to the network
	docker network connect thesis-network order-service-1 || true
	docker network connect thesis-network order-service-2 || true
	docker network connect thesis-network order-service-3 || true
	# Run the service on 3 ports with auto-restart enabled
	docker run -d --restart unless-stopped --network=thesis-network -e SERVER_HTTP_PORT=7012 -p 7012:7012 -v ./:/app --name order-service-1 order-service
	docker run -d --restart unless-stopped --network=thesis-network -e SERVER_HTTP_PORT=7013 -p 7013:7013 -v ./:/app --name order-service-2 order-service
	docker run -d --restart unless-stopped --network=thesis-network -e SERVER_HTTP_PORT=7014 -p 7014:7014 -v ./:/app --name order-service-3 order-service


