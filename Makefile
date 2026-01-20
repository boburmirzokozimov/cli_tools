.PHONY: run test

ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))

run:
	go run . $(ARGS)

test:
	go test -json ./... | gotestfmt

%:
	@:
