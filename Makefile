build:
	go build -o ./bin/sprat

run: build
	./bin/sprat

test:
	go test -v ./...