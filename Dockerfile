FROM golang:latest

RUN mkdir -p go/src/transportador

WORKDIR /go/src/transportador

COPY . /go/src/transportador

RUN go install transportador

CMD /go/src/transportador

EXPOSE 8080