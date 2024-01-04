PROGRAM := demo

.DEFAULT_GOAL := $(PROGRAM)
GOPATH := $(shell go env GOPATH)

GOIMPORTS := $(GOPATH)/bin/goimports
DLV := $(GOPATH)/bin/dlv
STATICCHECK := $(GOPATH)/bin/staticcheck

$(GOIMPORTS):
	go install golang.org/x/tools/cmd/goimports@latest

$(DLV):
	go install github.com/go-delve/delve/cmd/dlv@latest

$(STATICCHECK):
	go install honnef.co/go/tools/cmd/staticcheck@latest

fmt: $(GOIMPORTS)
	go fmt
	$(GOIMPORTS) -l -w .

lint: fmt $(STATICCHECK)
	staticcheck
.PHONY: lint

vet: fmt
	go vet
.PHONY: vet

$(PROGRAM): vet
	go mod tidy
	go build -gcflags="all=-N -l" -o $(PROGRAM)
.PHONY: $(PROGRAM)

run: $(PROGRAM)
	./$(PROGRAM)
.PHONY: run

debug: $(PROGRAM) $(DLV)
	dlv --listen=:2345 --headless=true --api-version=2 exec $(PROGRAM)
.PHONY: debug

test: vet
	go test -v ./...