FROM golang

WORKDIR $GOPATH/app

ADD . .

RUN go get && go build -o /app
