build:
	@go build -o bin/goserver

run: build
	@./bin/goserver

clean:
	@-rm bin/*

test:
	@go test -v ./...