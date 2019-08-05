FROM golang:1.13beta1-alpine3.10

RUN apk --update add postgresql-dev gcc libc-dev make libxml2-dev musl-dev libedit-dev

ADD . .

ENV PG_CONFIG=/usr/bin/pg_config
RUN make && make install
RUN go build pg2arrow.go