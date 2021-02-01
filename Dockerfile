ARG GO_VERSION=1.15.5
ARG ALPINE_VERSION=latest

FROM golang:${GO_VERSION}-alpine AS builder

ARG SEABOLT_VERSION=v1.7.4

RUN apk add --update --no-cache ca-certificates cmake make g++ openssl-dev openssl-libs-static git curl pkgconfig libcap
RUN git clone -b ${SEABOLT_VERSION} https://github.com/neo4j-drivers/seabolt.git /seabolt
RUN update-ca-certificates 2>/dev/null || true

WORKDIR /seabolt/build

RUN cmake -D CMAKE_BUILD_TYPE=Release -D CMAKE_INSTALL_LIBDIR=lib .. && cmake --build . --target install

RUN mkdir -p /go/src/github.com/qweliant/neo4j 
RUN mkdir /build
ADD . /go/src/github.com/qweliant/neo4j 
WORKDIR /go/src/github.com/qweliant/neo4j 


RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -tags seabolt_static -o /app main.go

# Create alpine runtime image
FROM alpine:${ALPINE_VERSION} as app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app /app

USER 1000

EXPOSE 1738

ENTRYPOINT ["/app"]