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