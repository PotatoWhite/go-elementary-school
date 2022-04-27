all: clean install build 
clean:
	rm -f main
	go clean --modcache
	rm -f app/wire_gen.go
install:
	go mod tidy
build:
	cd app;	wire
	go build app/main.go app/wire_gen.go
debug:
	GIN_MODE=debug ./main
run:
	GIN_MODE=release ./main
