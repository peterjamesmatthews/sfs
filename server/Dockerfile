FROM golang:1.21-alpine

# install dlv
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# install curl
RUN apk add curl

# install watchexec
RUN apk add watchexec

# install atlas
RUN curl -sSf https://atlasgo.sh | sh

# expose server's port
EXPOSE 8080 

# expose dlv's port
EXPOSE 4000

# stop signal
STOPSIGNAL SIGKILL

# set working directory
WORKDIR /sfs/server

# copy go.mod and go.sum
COPY go.mod go.sum ./

# install dependencies
RUN go mod download

# start server
CMD dlv debug ./cmd/server \
  --headless \
  --listen :4000 \
  --accept-multiclient \
  --continue \
  --api-version 2 \
  --log \
  --output ./__debug_bins/server
