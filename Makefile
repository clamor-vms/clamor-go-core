.PHONY: init build clean

DIR := ${CURDIR}

init:
	docker run --rm -v ${DIR}:/go/src/skaioskit -w /go/src/skaioskit lushdigital/docker-golang-dep init

build:
	docker run --rm -v ${DIR}:/go/src/skaioskit -w /go/src/skaioskit lushdigital/docker-golang-dep ensure
	docker run --rm -v ${DIR}:/go/src/skaioskit -w /go/src/skaioskit golang:latest go build -o skaioskit

clean:
	rm ${DIR}/golang-core
