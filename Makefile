PROJECT_DIR = .

.PHONY: proto
proto: clean
	buf generate --exclude-path=third_party/google  --exclude-path=third_party/app-core
	find ./model -type f '(' -name '*.swagger.json' ')' -exec mv -t ./service/swagger/models/ {} +

.PHONY: clean
clean:
	find . -type f '(' -name '*.pb.go' -o -name '*.pb.*.go' -o -name '*.swagger.json' ')' -not -path './vendor/*' -delete

.PHONY: lint-cache-clean
linter-cache-clean:
	golangci-lint cache clean

.PHONY: lint
lint:
	golangci-lint .

.PHONY: test
test:
	go test -v ./...

.PHONY: format
format:
	gofumpt -w .
	golines -w .