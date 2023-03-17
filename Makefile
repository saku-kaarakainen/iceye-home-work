.PHONY: run
run:
	docker build -t larvis . && docker run -i larvis:latest

local-run:
	go build . && go run . 

.PHONY: test
test:
	go test -v ./...