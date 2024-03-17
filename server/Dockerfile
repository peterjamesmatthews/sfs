FROM golang:1.21 AS base

FROM base AS delve
RUN go install github.com/go-delve/delve/cmd/dlv@latest

FROM base AS dependencies
WORKDIR /sfs
COPY go.mod go.sum ./
RUN go mod download

FROM dependencies AS source
COPY . .

FROM source AS debug
COPY --from=DELVE /go/bin/dlv /usr/local/bin/dlv
EXPOSE 8080
EXPOSE 9090

HEALTHCHECK  CMD curl --fail http://localhost:8080 || exit 1
STOPSIGNAL SIGKILL

CMD ["dlv", "debug", "./cmd/server", "--continue", "--output", "./bin", "--accept-multiclient", "--api-version", "2", "--headless", "--listen=:9090", "--log"]
