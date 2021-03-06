VERSION=1.0.0
SOURCE_FILES=$(shell find . -type d -name vendor -prune -o -type d -path ./cmd -prune -o -type f -name '*.go' -print)
GO_LIST=$(shell go list ./... | grep -v /vendor/)
GOPATH=$(shell echo "$$GOPATH")

all: bin/grpc-gateway bin/grpc-server

dist: build-cross
	cd bin/linux/amd64 && tar zcvf realize_sample-linux-amd64-${VERSION}.tar.gz realize_sample-${VERSION}

bin/grpc-gateway: ${SOURCE_FILES} cmd/grpc-gateway/main.go
	go build -o bin/grpc-gateway cmd/grpc-gateway/main.go

bin/grpc-server: ${SOURCE_FILES} cmd/grpc-server/main.go
	go build -o bin/grpc-server cmd/grpc-server/main.go

proto:
	protoeasy --go --go-import-path github.com/gosagawa/realize_sample/adapter/grpc/proto --out ${GOPATH}/src --grpc --grpc-gateway adapter/grpc/protoDefinition && make import

bundle:
	dep ensure

checkall: 
	fmtcheck importcheck lint vet errcheck misspell

fmt:
	find . -type f -name '*.go' -not -path "./vendor/*" -print0 | xargs -0 gofmt -w

fmtcheck:
	find . -type f -name '*.go' -not -path "./vendor/*" -print0 | xargs -0 gofmt -l | xargs -r false

import:
	find . -type f -name '*.go' -not -path "./vendor/*" -print0 | xargs -0 goimports -w

importcheck:
	find . -type f -name '*.go' -not -path "./vendor/*" -print0 | xargs -0 goimports -l | xargs -r false

lint:
	golint -set_exit_status ${GO_LIST}

vet:
	go vet ${GO_LIST}

errcheck:
	errcheck -blank -ignoretests ${GO_LIST}

misspell:
	find . -type f -name '*.go' -not -path "./vendor/*" -print0 | xargs -0 misspell -error

test:
	go test -v ./...

cover:
	go test -coverprofile=cover.out -v ./...

viewcover:
	go tool cover -html=cover.out

clean:
	rm -rf bin/*
