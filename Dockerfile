FROM golang:latest

RUN mkdir -p go/src/transportador

WORKDIR /go/src/transportador

COPY . /go/src/transportador

RUN go get .

CMD go run main.go

EXPOSE 8080