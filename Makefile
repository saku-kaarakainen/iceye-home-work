.PHONY: run
run:
	docker build -t larvis . && docker run -i larvis:latest

.PHONY: test
test:
	go test -v ./...