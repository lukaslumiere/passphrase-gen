BINARY=passphrase-gen
LDFLAGS=-ldflags "-s -w"

.PHONY: build install clean

build:
	go build $(LDFLAGS) -o $(BINARY) main.go

install: build
	sudo mv $(BINARY) /usr/local/bin/

clean:
	rm -f $(BINARY)
