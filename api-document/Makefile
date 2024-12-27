## gen swagger
gen-swagger:
	swag init -g ./handlers/cart.go -o resources/docs --parseDependency true
	swag init -g ./handlers/product.go -o resources/docs --parseDependency true
	swag init -g ./handlers/user.go -o resources/docs --parseDependency true
	swag init -g ./handlers/order.go -o resources/docs --parseDependency true
	swag init -g ./handlers/comment.go -o resources/docs --parseDependency true

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