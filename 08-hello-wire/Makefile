all: clean install build 
clean:
	rm -f main
	go clean --modcache
	rm -f pkg/wire_gen.go
install:
	go mod tidy
build:
	cd pkg;	wire
	go build pkg/main.go pkg/event.go pkg/message.go pkg/greeter.go pkg/spec.go  pkg/wire_gen.go
debug:
	GIN_MODE=debug ./main
run:
	GIN_MODE=release ./main
