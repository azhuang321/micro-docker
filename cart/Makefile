
GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get github.com/micro/micro/v3/cmd/protoc-gen-micro
.PHONY: proto
proto:
	protoc -I ./proto --go_out=./ --micro_out=./ ./proto/cart/cart.proto
	
.PHONY: build
build:
	go build -o git.imooc.com/cap1573/cart *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t git.imooc.com/cap1573/cart:latest
