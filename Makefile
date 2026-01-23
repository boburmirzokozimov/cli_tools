.PHONY: run test

ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))

run:
	go run ./cmd $(ARGS)

test:
	go test -json ./... | gotestfmt

%:
	@:
