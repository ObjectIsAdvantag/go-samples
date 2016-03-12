
default: run

dependencies:
	go get golang.org/x/net/publicsuffix

run:
	go run tropo-session.go

proxy:
	http_proxy=127.0.0.1:60707 go run tropo-session.go

