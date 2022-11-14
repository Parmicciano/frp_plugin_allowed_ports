export PATH := $(GOPATH)/bin:$(PATH)
export GO111MODULE=on
LDFLAGS := -s -w

all: fmt build

build: frps_allowed_ports

fmt:
	go fmt ./...

frps_allowed_ports:
	env CGO_ENABLED=0 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/frps_allowed_ports ./cmd/frps_allowed_ports

clean:
	rm -f ./bin/frps_allowed_ports
