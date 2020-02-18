FROM golang:latest

RUN mkdir -p go/src/transportador

WORKDIR /go/src/transportador

COPY . /go/src/transportador

RUN go get github.com/go-kit/kit/endpoint

RUN go get github.com/go-kit/kit/log

RUN go get github.com/go-kit/kit/transport/http

RUN go get github.com/gorilla/mux

RUN go get github.com/lib/pq

CMD go run main.go

EXPOSE 8080