.PHONY: build get update fmt test ci-test build buf config types ent

APP=litmedia

GO       := GO111MODULE=on CGO_ENABLED=0 GOPRIVATE=github.com/litsoftware/litmedia GOSUMDB=off go
GOTEST   := $(GO) test -gcflags='-l' -p 3 -v -race

FILES    := $(shell find . -name '*.go' -type f)
TESTS    := $(shell find . -name '*.go' -type f)

build:
	$(GO) version
	$(GO) build -o ./bin/litmedia .
	echo "Done!"

get:
	$(GO) version
	$(GO) get ./...
	$(GO) mod verify
	$(GO) mod tidy
	echo "Done!"

update:
	$(GO) list -m -u all
	$(GO) mod verify
	$(GO) mod tidy
	echo "Done!"

buf:
	cd ./cmd/http/proto && buf generate
	echo "Done!"

fmt:
	gofmt -s -l -w $(FILES) $(TESTS)
	echo "Done!"

test:
	$(GOTEST) ./... -args testing
	echo "Done!"

ci-test:
	$(GOTEST) ./... -args testing ci
	echo "Done!"

types:
	# make enum
	cd ./internal/common/types protoc --proto_path=. --go_out=. --go_opt=paths=source_relative ./*.proto
	echo "Done!"

ent:
	$(GO) generate ./internal/ent
	echo "Done!"

config:
	cd ./internal/pkg/config && protoc --proto_path=. --go_out=. --go_opt=paths=source_relative ./*.proto
	echo "Done!"

error:
	cd tools/protoc-gen-go-errors && go build && cp protoc-gen-go-errors /usr/local/bin && cd -
	protoc --proto_path=. --go_out=. --go_opt=paths=source_relative ./internal/errors/*.proto

	# make errors
	protoc --proto_path=. --go_out=. --go_opt=paths=source_relative --go-errors_out=. ./internal/errors/*.proto
	cp github.com/litsoftware/litmedia/internal/errors/*.go ./internal/errors/ && rm -rf github.com

migrate:
	$(GO) run . migrate
