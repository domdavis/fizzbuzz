docker:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o fizzbuzz .
	docker build -t fizzbuzz .

clean:
	go clean

dependencies:
	go get -u github.com/kardianos/govendor
	$(GOBIN)/govendor sync

build:
	go build

test:
	go test ./...

ci: clean dependencies build test

default: build

.PHONY: test
