build:
	@go build -o bin/goclique

run: build
	@./bin/goclique

clean:
	@-rm bin/*

test:
	@go test -v ./...