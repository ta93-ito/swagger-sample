FROM golang:1.16.5 AS dev

WORKDIR /go/src
ADD . /go/src

RUN go mod download

RUN go build -o /bin/main ./api/main.go

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b /bin

CMD ["./bin/air"]
