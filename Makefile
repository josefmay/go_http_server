build:
	@go build -o bin/goserver cmd/api/main.go

run: build
	@./bin/goserver

clean:
	@-rm bin/*

test:
	@go test -v ./...