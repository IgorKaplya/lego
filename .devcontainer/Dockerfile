FROM alpine:latest

ENV PATH=$PATH:/usr/local/go/bin

RUN apk add --no-cache git

RUN set -eux \
    && wget -O /tmp/go.tar.gz "https://go.dev/dl/go1.25.5.linux-amd64.tar.gz" \
    && tar -C /usr/local -xzf /tmp/go.tar.gz \
    && rm /tmp/go.tar.gz \
    && go version \
    && go install github.com/go-delve/delve/cmd/dlv@v1.26.0