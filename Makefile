default: build

test:
	$(shell which go) test ./...

build:
	$(shell which go) build

install: build
	mkdir -p ~/.tflint.d/plugins
	mv ./tflint-ruleset-exec ~/.tflint.d/plugins
