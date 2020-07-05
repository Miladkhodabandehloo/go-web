FROM golang:latest

WORKDIR /go/src/go-web

COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download

COPY ./ ./

EXPOSE 8000

RUN go build -o /go/bin/go-web

CMD ["/go/bin/go-web"]