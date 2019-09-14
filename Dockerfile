FROM golang:latest

RUN apt-get update && apt-get install -y \
    dnsutils && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /go/src/app
COPY main.go .

RUN go get -d -v ./... && \
    go install -v ./...

CMD app
