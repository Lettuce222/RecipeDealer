FROM golang:1.16.3-buster
WORKDIR /go/src/app
ADD . ./
ENV GO111MODULE=on\
    GOPROXY=https://proxy.golang.org,direct
RUN go get gorm.io/driver/postgres\
           gorm.io/gorm
CMD go mod tidy && go run main.go