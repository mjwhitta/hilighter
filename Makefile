GITLAB = hl "gitlab.com/mjwhitta/hilighter"
LOCAL = hl "hilighter"

all: build unlocal

build: local fmt
	@go build ./cmd/hl

clean: unlocal fmt
	@rm -f hl

clena: clean

fmt:
	@go fmt . ./cmd/hl

gen:
	@./scripts/generate_go_funcs

local:
	@find cmd -type f -exec sed -i 's#$(GITLAB)#$(LOCAL)#' {} +
	@rm -f go.mod go.sum

unlocal:
	@find cmd -type f -exec sed -i 's#$(LOCAL)#$(GITLAB)#' {} +
	@git checkout go.mod go.sum
