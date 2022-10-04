FROM alpine:3.16.2

COPY temporal-server /temporal/temporal-server

WORKDIR /temporal

CMD ["/temporal/temporal-server"]