all: build

build: check fmt
	@go build ./cmd/hl

check:
	@which go >/dev/null 2>&1

clean: fmt
	@rm -f hl

clena: clean

fmt: check
	@go fmt . ./cmd/hl

gen: check
	@go generate
