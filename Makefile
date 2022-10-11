test:
	go test ./...

install-linter:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.50.0

liter: install-linter
	bin/golangci-lint run ./...

run-server:
	go run -race cmd/server/main.go -path=./config.json

run-client:
	go run -race cmd/client/main.go -path=./config.json