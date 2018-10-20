.PHONY: init build clean

DIR := ${CURDIR}

init:
	docker run --rm -v ${DIR}:/go/src/clamor -w /go/src/clamor lushdigital/docker-golang-dep init

build:
	docker run --rm -v ${DIR}:/go/src/clamor -w /go/src/clamor lushdigital/docker-golang-dep ensure
	docker run --rm -v ${DIR}:/go/src/clamor -w /go/src/clamor golang:latest go build -o clamor

clean:
	rm ${DIR}/golang-core
