.PHONY: go-deps-ls deps fmt lint test publish

go-deps-ls:
	go list -u -m -f '{{if not .Indirect}}{{.}}{{end}}' all

deps:
	go mod download
	go mod tidy

fmt:
	go fmt ./...

lint: fmt deps
	golangci-lint run ./... --timeout 5m

test: lint
	gotestsum --format testdox -- -v ./...

# Usage: make publish VERSION=v0.1.0
publish:
	git tag $(VERSION)
	git push origin $(VERSION)
	go list -m github.com/miroslav-matejovsky/winfileinfo@$(VERSION)
