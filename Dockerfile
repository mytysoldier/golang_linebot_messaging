FROM golang:alpine

WORKDIR /go/src/golang_linebot_messaging

COPY . .

RUN go build -o app main.go

CMD ["/go/src/golang_linebot_messaging/app"]