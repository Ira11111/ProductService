include .env
.PHONY: $(MAKECMDGOALS)


COMPILED ?= false

run:
ifeq ($(COMPILED), true)
	ifeq ($(test -f ./bin/sso), 1)
		./bin/sso
	else
		@echo "Binary not found"
	endif
else
	go run ./cmd/products/main.go
endif

build:
	go build -o ./bin/products ./cmd/products/main.go

migrate_up:
	goose -table goose_migrations.goose_db_version up
migrate_down:
	goose -table goose_migrations.goose_db_version down-to 0