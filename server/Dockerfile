FROM golang:1.21-bullseye

# set working directory
WORKDIR /sfs/server

# expose port
EXPOSE 8080

# stop signal
STOPSIGNAL SIGKILL

# copy go.mod and go.sum
COPY go.mod go.sum ./

# install dependencies
RUN go mod download

# start server
CMD go run ./cmd/server
