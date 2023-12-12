.PHONY: go-run adr dd commit go-lint


setup-tools:
	@if [ -z `command -v air` ]; then go install github.com/cosmtrek/air@latest ; fi
	@if ! [ -x /usr/local/bin/golangci-lint ]; then brew install golangci/tap/golangci-lint ; fi
	@if ! [ -x ${GOPATH}/bin/wire ]; then go install github.com/google/wire/cmd/wire@latest ;fi

##### exec
run:
	go run ./cmd/main.go

##### DI
wire:
	wire gen

##### scaffing

ADR_COUNT:=$(shell find docs/ADR -type f | wc -l | tr -d ' ') 
DD_COUNT:=$(shell find docs/DesignDog -type f | wc -l | tr -d ' ') 
adr:
	npx scaffdog generate ADR --output 'docs/ADR' --answer 'number:${ADR_COUNT}'

dd:
	npx scaffdog generate DD --output 'docs/DesignDog' --answer 'number:${DD_COUNT}'

##### git

commit:
	npx git-cz

##### tools

go-lint:
	cd go;golangci-lint run --concurrency 2  


