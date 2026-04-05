.PHONY: build fmt clean

build:
	go build -o gopermutor main.go

fmt:
	gofmt -s -w main.go

clean:
	go clean -cache
	rm gopermutor
