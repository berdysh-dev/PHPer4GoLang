GOROOT=/usr/lib/golang
GOPATH=/usr/local/src/Go/GOPATH

SIGN = PHPer4GoLang

all: clean go.mod

go.mod:
	go mod init github.com/berdysh-dev/PHPer4GoLang

clean:
	rm -f go.mod
