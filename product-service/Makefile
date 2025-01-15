## gen swagger
gen-swagger:
	swag init -g ./internal/presentation/http/http.go -o resources/docs --parseDependency true

## go mod tidy
go-mod-tidy:
	export GOPRIVATE=10.96.24.141 
	export GOPROXY=http://10.96.20.152:8083/repository/go_proxy
	export GOSUMDB='sum.golang.org http://10.96.20.152:8083/repository/go_sumdb' 
	go mod tidy

# bind go bin
go-bind:
	export GO_PATH=~/go
	export PATH=$PATH:/$GO_PATH/bin

## mockery
gen-mock:
	mockery --dir=internal/domain --output=./mocks --outpkg=mocks --all
	mockery --dir=internal/usecase/create-transaction --output=./mocks --outpkg=mocks --all
	mockery --dir=internal/usecase/validate-transaction --output=./mocks --outpkg=mocks --all
	mockery --dir=internal/domain --output=./mocks --outpkg=mocks --all
	mockery --dir=internal/usecase/delete-transaction --output=./mocks --outpkg=mocks --all
	mockery --dir=internal/domain --output=./mocks --outpkg=mocks --all
## unit test coverage
gen-converage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=./coverage.out -o coverage.html

go-build:
	GOOS=linux CGO_ENABLED=0 go build -o main ./cmd/main.go

deploy:
	docker container rm product-service
	docker build -t product-service .
	docker run -d -p 7002:7002 --name product-service product-service

