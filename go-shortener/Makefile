PROGRAM=shortener
GOFLAGS= -tags netgo

default: run

clean:
	rm -f $(PROGRAM) $(PROGRAM).exe $(PROGRAM).zip $(PROGRAM).debug

build: clean goget
	go build $(GOFLAGS) *.go

dev:
	go run $(GOFLAGS) *.go

run:
	go run $(GOFLAGS) *.go &
	lt -p 8081 -s shortener
