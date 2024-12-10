FROM alpine:latest

RUN apk add curl
RUN curl -fsSL https://raw.githubusercontent.com/pressly/goose/master/install.sh | sh

WORKDIR /src
COPY migrations/ .

ENTRYPOINT ["/usr/local/bin/goose"]
