FROM golang:1.16 AS builder
ENV GOPROXY https://goproxy.cn,direct
ARG VERSION=0.0.10
WORKDIR $GOPATH/src/app
COPY . $GOPATH/src/app
# RUN go build -o main main.go
RUN go build -o main -ldflags="-X 'main.version=${VERSION}'" main.go

FROM debian:stable-slim
COPY main $GOPATH/src/app/main 
ENV PATH="/go/bin:${PATH}"
CMD ["main"]