FROM golang

WORKDIR $GOPATH/src/app

ADD . .

RUN go get && go build -o /app
