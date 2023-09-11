APP_NAME = currency-convertor

.PHONY: all build clean

all: build

build:
	@mkdir -p ./build/bin
	go build -o ./build/bin/$(APP_NAME) ./cmd/app

clean:
	@rm -rf ./build

## lint:                              run golangci-lint with .golangci.yml config file
lint:
	@./build/bin/golangci-lint run --config ./.golangci.yml

## lintci-deps:                       (re)installs golangci-lint to build/bin/golangci-lint
lintci-deps:
	rm -f ./build/bin/golangci-lint
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./build/bin v1.54.0