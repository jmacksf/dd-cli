BINARY_NAME=dd-cli

build:
	GOARCH=amd64 GOOS=darwin  go build -o ${BINARY_NAME}-darwin  ./cmd/${BINARY_NAME}
	GOARCH=amd64 GOOS=linux   go build -o ${BINARY_NAME}-linux   ./cmd/${BINARY_NAME}
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows ./cmd/${BINARY_NAME}

run:
	go run ./cmd/${BINARY_NAME} $(cmd)

build_and_run: build run

clean:
	go clean
	rm -f ${BINARY_NAME}-darwin
	rm -f ${BINARY_NAME}-linux
	rm -f ${BINARY_NAME}-windows

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

dep:
	go mod download

dep_tidy:
	go mod tidy -v

vet:
	go vet ./...%
