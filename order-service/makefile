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
	-docker container stop order-service || true
	-docker container rm order-service || true
	docker build -t order-service .
	docker network connect kong-net product-service || true
	docker run -d --network=kong-net -p 7005:7005 -v ./:/app --name order-service order-service


