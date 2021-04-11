FROM golang:1.16.3-buster
WORKDIR /go/src/app
ENV GO111MODULE=on\
    GOPROXY=https://proxy.golang.org,direct
RUN go get github.com/go-redis/redis/v8@latest &&\
    go install github.com/pilu/fresh@latest
CMD ["go", "run", "main.go"]