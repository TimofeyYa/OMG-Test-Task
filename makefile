CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
PACKAGE=omg/cmd/app

start:
	go run ${PACKAGE}

build:
	mkdir -p ${BINDIR}
	go build -o ${BINDIR}/app ${PACKAGE}

run-all: build
	sudo docker compose up --force-recreate --build

migrate-up:
	goose -dir ./migrations postgres "postgres://user:password@localhost:5432/omg?sslmode=disable" status
	goose -dir ./migrations postgres "postgres://user:password@localhost:5432/omg?sslmode=disable" up

migrate-down:
	goose -dir ./migrations postgres "postgres://user:password@localhost:5432/omg?sslmode=disable" status
	goose -dir ./migrations postgres "postgres://user:password@localhost:5432/omg?sslmode=disable" down