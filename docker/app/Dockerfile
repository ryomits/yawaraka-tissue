FROM golang:1.18.2-alpine3.16

ENV GOOS linux
ENV GOARCH amd64
ENV PATH $PATH:/go/bin/linux_amd64
ENV CGO_ENABLED 0

RUN \
  echo http://dl-cdn.alpinelinux.org/alpine/edge/community/ >> /etc/apk/repositories && \
  apk update --no-cache && \
  apk add --virtual .build \
  gcc \
  make \
  musl-dev \
  libffi-dev \
  ca-certificates \
  tzdata && \
  update-ca-certificates && \
  ln -f -s /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.46.2
RUN go install golang.org/x/tools/cmd/goimports@latest

WORKDIR /app

COPY Makefile .
RUN make setup

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . /app
