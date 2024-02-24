BINARY_NAME := server

build:
	go build -o=/tmp/bin/${BINARY_NAME} .

clean:
	rm -f /tmp/bin/${BINARY_NAME}

run: build
	/tmp/bin/${BINARY_NAME}