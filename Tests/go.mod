module example.com/Server

go 1.19

replace github.com/berdysh-dev/PHPer4GoLang => /usr/local/GIT/PHPer4GoLang

require github.com/berdysh-dev/PHPer4GoLang v0.0.0-00010101000000-000000000000

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/redis/go-redis/v9 v9.0.4 // indirect
)
