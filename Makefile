BINARY_NAME=user-service

build:
	GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME}-linux main.go
	GOOS=darwin GOARCH=amd64 go build -o ${BINARY_NAME}-darwin main.go
	GOOS=windows GOARCH=amd64 go build -o ${BINARY_NAME}-windows main.go

run: build
	./${BINARY_NAME}-linux

run-mac: build
	./${BINARY_NAME}-darwin

run-win: build
	./${BINARY_NAME}-windows

clean:
	go clean
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-windows