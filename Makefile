# .PHONY: test
# test:
# 	go test -v ./... ./internal/...

.PHONY: run
run:
	docker build -t larvis . && docker run larvis:latest