all: clean install build 
clean:
	rm -f main
	go clean --modcache
install:
	go mod tidy
build:
	go build app/main.go
debug:
	GIN_MODE=debug ./main
run:
	GIN_MODE=release ./main
