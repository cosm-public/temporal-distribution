FROM golang:1.19 as builder

WORKDIR /temporal

COPY go.mod go.sum main.go Makefile ./

RUN make temporal-server

FROM alpine:3.16.2

COPY --from=builder /temporal/bin/temporal-server ./

CMD ["./temporal-server"]