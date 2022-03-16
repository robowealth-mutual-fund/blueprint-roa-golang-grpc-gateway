## build: build golang
.PHONY: build
build:
	go build -o bin/server cmd/server/server.go

## start: run server in development mode
.PHONY: start
start:
	go run cmd/server/server.go

.PHONY: watch
watch:
	CompileDaemon -include=Makefile --build="make build" --command=./bin/server --color=true --log-prefix=false

## pbgen: genrate protobug file
.PHONY: pbgen
pbgen:
	protoc --proto_path=internals/api/v1 --go_out=plugins=grpc:pkg/grpc/health/v1 health.proto
	protoc --proto_path=internals/api/v1 --go_out=plugins=grpc:pkg/api/v1 ping_pong.proto
	protoc --proto_path=internals/api/v1 --proto_path=thirdparty --go_out=plugins=grpc:pkg/api/v1 --grpc-gateway_out=logtostderr=true:pkg/api/v1 --swagger_out=logtostderr=true:swagger product.proto
	protoc-go-inject-tag -input=pkg/api/v1/ping_pong.pb.go
	protoc-go-inject-tag -input=pkg/api/v1/product.pb.go

.PHONY: stringer
stringer:
	stringer -type ErrorCode internals/constants/error_code.go
.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo