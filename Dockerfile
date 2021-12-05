### GOLANG ###
ARG GOLANG_VERSION
ARG ALPINE_VERSION

FROM golang:${GOLANG_VERSION}-alpine${ALPINE_VERSION} AS builder

RUN set -eux \
  && apk update && apk --no-cache add \
    ca-certificates \
    postgresql-dev \
    libzip-dev \
    git \
	icu-libs \
	bash \
	autoconf \
	g++ \
	make \
	icu-dev \
	gcc \
	libtool \
	rabbitmq-c \
	rabbitmq-c-dev

# Install beego & bee
RUN go get -u github.com/gin-gonic/gin

RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.42.1

RUN golangci-lint --version

# Godep for vendoring
RUN go get github.com/tools/godep

# Recompile the standard library without CGO
RUN CGO_ENABLED=0 go install -a std

RUN go mod tidy

WORKDIR /app

EXPOSE 8085
