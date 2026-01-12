APP=pipewatch

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./...

race:
	go test -race ./...

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

junit:
	go test -v 2>&1 ./... | go-junit-report -set-exit-code > junit.xml

run:
	PIPEWATCH_ADDR=:8080 go run ./cmd/pipewatch
