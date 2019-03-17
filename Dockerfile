FROM golang:1.12 as builder

WORKDIR /go/src/github.com/zombispormedio/micro

COPY . .

RUN go get -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/micro .

FROM scratch

WORKDIR /root/

COPY --from=builder /go/bin/micro .
