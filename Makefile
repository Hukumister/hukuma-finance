MAIN_PACKAGE_PATH := ./cmd/server
BINARY_NAME := server

build:
	go build -o=/tmp/bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}

clean:
	rm -f /tmp/bin/${BINARY_NAME}

run: build
	/tmp/bin/${BINARY_NAME}