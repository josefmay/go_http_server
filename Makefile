build:
	@go build -o bin/go-server

run: build
	@./bin/go-server

clean:
	@-rm bin/*

test:
	@go test -v ./...