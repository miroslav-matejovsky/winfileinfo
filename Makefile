.PHONY: go-deps-ls

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

# publish:
# 	$env:GOPROXY = "proxy.golang.org"
# 	go list -m github.com/miroslav-matejovsky/winfileinfo@v0.1.0
