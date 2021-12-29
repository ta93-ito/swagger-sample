FROM golang:1.16.5 as builder

WORKDIR /go/src

COPY ${PWD} /go/src

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64

RUN go build \
    -o /go/bin/main \
    api/main.go

FROM busybox

COPY --from=builder /go/bin/main .

CMD ["./main"]
