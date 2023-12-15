ENV_FILE := .env
ENV = $(shell grep -v "^\#" $(ENV_FILE))
include .env

.PHONY: go-run adr dd commit go-lint locust/up locust/down


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

###### DB

POSTGRES_USER := root
DATABASE      := db
psql:
	$(DOCKER_COMPOSE) exec db psql -U $(POSTGRES_USER) -d $(DATABASE)

GOOSE_DRIVER   := postgres
GOOSE_DBSTRING ?= host=db user=root dbname=db password=p@ssword sslmode=disable
migrate/status:
	$(DOCKER_COMPOSE) run --rm migration status

NAME?=
SERVICE?=
migrate/new:
	echo '-- +goose Up' > app/driver/db/migrations/$(shell ls app/infra/db/migrations | awk -F"_*.sql" 'BEGIN {max=0} {split($$1, a, "_"); if(a[1]>max){max = a[1]}}END{print max+1}')_${NAME}.sql

migrate/up:
	$(DOCKER_COMPOSE) run --rm migration up

##### locust
locust/up:
	cd locust;terraform init;terraform apply --auto-approve

locust/down:
	cd locust;terraform destroy --auto-approve



