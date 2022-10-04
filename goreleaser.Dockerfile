FROM temporalio/base-server:1.10.0

WORKDIR /home/temporal

ENV TEMPORAL_HOME /home/temporal

RUN addgroup -g 1000 temporal
RUN adduser -u 1000 -G temporal -D temporal
USER temporal

COPY temporal-server /home/temporal/temporal-server

CMD ["/home/temporal/temporal-server"]