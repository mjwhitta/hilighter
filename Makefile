all: build

build: fmt
	@go build ./cmd/hl

clean: fmt
	@rm -f hl

clena: clean

fmt:
	@go fmt . ./cmd/hl

gen:
	@./scripts/generate_go_funcs
