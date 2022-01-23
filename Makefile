LDFLAGS := '-s -w'
GOFLAGS := CGO_ENABLED=0
GO := $(GOFLAGS) go
build:
	$(GO) build -ldflags $(LDFLAGS)
install: build
	$(GO) install .

buildarm64:
	GOARCH=arm64 $(GO) build

buildamd64:
	GOARCH=amd64 $(GO) build
