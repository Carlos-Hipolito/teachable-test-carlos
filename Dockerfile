FROM golang:1.20.1-alpine
ADD . /go/src/app
WORKDIR /go/src/app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go get
RUN go build -o /teachable-test

EXPOSE 8080

CMD [ "/teachable-test"]