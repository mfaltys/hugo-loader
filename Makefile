GOC=go build
GOFLAGS=-a -ldflags '-s'
CGOR=CGO_ENABLED=0
GIT_HASH=$(shell git rev-parse HEAD | head -c 10)

all: hugo-loader

hugo-loader: hugo-loader.go
	$(GOC) hugo-loader.go

dependencies:
	go get gopkg.in/gcfg.v1

run:
	go run hugo-loader

stat:
	$(CGOR) $(GOC) $(GOFLAGS) hugo-loader.go

clean:
	rm -rf hugo-loader
