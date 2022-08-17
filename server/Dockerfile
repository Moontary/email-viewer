#
# Build the Go server
#
# build stage
ARG ALPINE_VERSION=3.16

FROM golang:1.19-alpine${ALPINE_VERSION} AS builder

ARG SRC=/build/

COPY . ${SRC}

WORKDIR ${SRC}

RUN go mod download && go mod verify

RUN go build -o /bin/email cmd/main.go

# final stage

FROM alpine:${ALPINE_VERSION}

COPY --from=builder /bin /bin

EXPOSE 9090/tcp

CMD [ "email" ]