GOROOT=/usr/lib/golang
GOPATH=/usr/local/src/Go/GOPATH

PHP=/bin/php

SIGN = Server

SRCS = ${SIGN}.go

all: clean tidy build start

go.mod:
	@go mod init example.com/${SIGN}
	@echo "" >> go.mod
	@echo "replace github.com/berdysh-dev/PHPer4GoLang => /usr/local/GIT/PHPer4GoLang" >> go.mod

tidy: go.mod
	go mod tidy

build: ${SIGN}

${SIGN}: ${SIGN}.go
	@go build ${SRCS}

start: ${SIGN}
	@./${SIGN}

clean:
	@rm -rf pkg go.mod go.sum ${SIGN}

run:
	@go run ${SRCS}

php:    test.php
	${PHP} test.php

