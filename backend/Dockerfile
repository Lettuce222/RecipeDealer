FROM golang:1.16.3-buster
WORKDIR /go/src/app
ADD . ./
ENV GO111MODULE=on\
    GOPROXY=https://proxy.golang.org,direct
RUN go get -u -t -v\
        # DB
           gorm.io/driver/postgres\
           gorm.io/gorm\
        # HTTP Server  
           github.com/gin-gonic/gin\

        #  Go for VSCode  
           github.com/nsf/gocode\
           github.com/rogpeppe/godef\
           github.com/zmb3/gogetdoc\
           golang.org/x/lint/golint\
           github.com/ramya-rao-a/go-outline\
           sourcegraph.com/sqs/goreturns\
           golang.org/x/tools/cmd/gorename\
           github.com/tpng/gopkgs\
           github.com/acroca/go-symbols\
           golang.org/x/tools/cmd/guru\
           github.com/cweill/gotests/...\
           golang.org/x/tools/cmd/godoc\
           github.com/fatih/gomodifytags\
           github.com/josharian/impl\

         # ライブローディング
            github.com/cosmtrek/air
CMD go mod tidy && air -d