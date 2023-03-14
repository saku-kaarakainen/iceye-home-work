.PHONY: run
run:
	docker build -t larvis . && docker run larvis:latest

.PHONY: test
test:
	go test -v ./...