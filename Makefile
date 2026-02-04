build:
	go build -o bin/caching-proxy ./cmd/caching-proxy

run:
	go run ./cmd/caching-proxy -port=8080 -origin=http://dummyjson.com/products

start:
	./bin/caching-proxy --port 8080 --origin http://dummyjson.com 