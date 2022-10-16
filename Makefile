BINARY_NAME=h8-assignment-3
BUILD_DIR=build

.PHONY: build clean

build:
	mkdir -p ${BUILD_DIR}
	GOARCH=amd64 GOOS=linux go build -o ${BUILD_DIR}/${BINARY_NAME}-linux .
	GOARCH=amd64 GOOS=windows go build -o ${BUILD_DIR}/${BINARY_NAME}-windows.exe .

clean:
	go clean
	rm ${BUILD_DIR}/${BINARY_NAME}-linux
	rm ${BUILD_DIR}/${BINARY_NAME}-windows.exe