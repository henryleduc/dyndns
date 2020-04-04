FROM golang:1.14-alpine

WORKDIR /go/src/app

ADD . .

RUN go build cmd/main.go

EXPOSE 3000:3000

CMD [ "./main" ]