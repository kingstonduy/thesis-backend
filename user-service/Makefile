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
	docker container stop user-service-1 || true
	docker container rm user-service-1 || true
	docker container stop user-service-2 || true
	docker container rm user-service-2 || true
	docker container stop user-service-3 || true
	docker container rm user-service-3 || true
	# Build the Docker image
	docker build -t user-service .
	# Connect the service to the network
	docker network connect thesis-network user-service-1 || true
	docker network connect thesis-network user-service-2 || true
	docker network connect thesis-network user-service-3 || true
	# Run the service on 3 ports with auto-restart enabled
	docker run -d --restart unless-stopped --network=thesis-network -e SERVER_HTTP_PORT=7022 -p 7022:7022 -v ./:/app --name user-service-1 user-service
	docker run -d --restart unless-stopped --network=thesis-network -e SERVER_HTTP_PORT=7023 -p 7023:7023 -v ./:/app --name user-service-2 user-service
	docker run -d --restart unless-stopped --network=thesis-network -e SERVER_HTTP_PORT=7024 -p 7024:7024 -v ./:/app --name user-service-3 user-service


