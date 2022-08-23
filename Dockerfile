FROM golang:1.16-alpine as download-dependencies

RUN apk update && \
    apk add git && \
    apk add build-base && \
    apk add bash

WORKDIR /sbdb-cad-api-test

RUN mkdir libs/
RUN mkdir tests/
RUN mkdir types/
RUN mkdir utils/

COPY libs/ ./libs
COPY types/ ./types
COPY utils/ ./utils
COPY tests/ ./tests
COPY go.mod ./
COPY go.sum ./
COPY test.sh ./
RUN go mod download

RUN pwd
RUN ls

ENTRYPOINT [ "/bin/bash", "test.sh", "runDocker" ]


