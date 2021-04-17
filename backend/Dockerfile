FROM golang:1.16.3-buster
WORKDIR /go/src/app
ADD . ./
ENV GO111MODULE=on\
    GOPROXY=https://proxy.golang.org,direct
RUN go get -u -t\
         # DB
            gorm.io/driver/postgres\
            gorm.io/gorm\
         # HTTP Server  
            github.com/gin-gonic/gin\
         #  Go for VSCode  
            golang.org/x/tools/gopls\
         # ライブローディング
            github.com/cosmtrek/air