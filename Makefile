build:
	go build -o bin/caching-proxy ./cmd/caching-proxy

run:
	go run ./cmd/caching-proxy -port=8080 -origin=https://dummyjson.com

start:
	./bin/caching-proxy --port 8080 --origin https://dummyjson.com 