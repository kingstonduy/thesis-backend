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
	-docker container stop product-service || true
	-docker container rm product-service || true
	docker build -t product-service .
	docker run -d -p 7001:7001 --name product-service product-service


