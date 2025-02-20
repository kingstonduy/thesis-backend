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
	docker container stop cart-service-1 || true
	docker container rm cart-service-1 || true
	docker container stop cart-service-2 || true
	docker container rm cart-service-2 || true
	docker container stop cart-service-3 || true
	docker container rm cart-service-3 || true
	# Build the Docker image
	docker build -t cart-service .
	# Connect the service to the network
	docker network connect thesis-network cart-service-1 || true
	docker network connect thesis-network cart-service-2 || true
	docker network connect thesis-network cart-service-3 || true
	# Run the service on 3 ports with auto-restart enabled
	docker run -d --restart unless-stopped --network=thesis-network -e SERVER_HTTP_PORT=7032 -p 7032:7032 -v ./:/app --name cart-service-1 cart-service
	docker run -d --restart unless-stopped --network=thesis-network -e SERVER_HTTP_PORT=7033 -p 7033:7033 -v ./:/app --name cart-service-2 cart-service
	docker run -d --restart unless-stopped --network=thesis-network -e SERVER_HTTP_PORT=7034 -p 7034:7034 -v ./:/app --name cart-service-3 cart-service


