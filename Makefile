BINDIR  := $(CURDIR)/bin
BINNAME ?= gopermutor

LDFLAGS     ?= -w -s
GOFLAGS     ?=
CGO_ENABLED ?= 0

SRC := main.go

.PHONY: build fmt clean

build:
	CGO_ENABLED=$(CGO_ENABLED) go build -trimpath -ldflags '$(LDFLAGS)' -o '$(BINDIR)/$(BINNAME)' $(SRC)

fmt:
	gofmt -s -w $(SRC)

clean:
	rm -rf $(BINDIR)
