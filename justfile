default: build

build:
    go build -o sdrangel-mcp ./cmd/sdrangel-mcp

build-all:
    GOOS=linux   GOARCH=amd64 go build -o dist/sdrangel-mcp_linux_amd64   ./cmd/sdrangel-mcp
    GOOS=linux   GOARCH=arm64 go build -o dist/sdrangel-mcp_linux_arm64   ./cmd/sdrangel-mcp
    GOOS=darwin  GOARCH=amd64 go build -o dist/sdrangel-mcp_darwin_amd64  ./cmd/sdrangel-mcp
    GOOS=darwin  GOARCH=arm64 go build -o dist/sdrangel-mcp_darwin_arm64  ./cmd/sdrangel-mcp
    GOOS=windows GOARCH=amd64 go build -o dist/sdrangel-mcp_windows_amd64.exe ./cmd/sdrangel-mcp

install:
    go install ./cmd/sdrangel-mcp

run:
    go run ./cmd/sdrangel-mcp serve

test:
    go test -race -shuffle=on ./...

cover:
    go test -race -shuffle=on -coverprofile=coverage.out ./...
    go tool cover -func=coverage.out

cover-html:
    go test -race -shuffle=on -coverprofile=coverage.out ./...
    go tool cover -html=coverage.out

lint:
    golangci-lint run ./...

fmt:
    gofmt -w .
    goimports -w .

fmt-check:
    @test -z "$(gofmt -l .)" || (echo "gofmt issues found:" && gofmt -l . && exit 1)

vet:
    go vet ./...

sast:
    gosec -fmt sarif -out gosec.sarif ./... || true

vulncheck:
    govulncheck ./...

audit: fmt-check vet lint vulncheck

tidy:
    go mod tidy

update-deps:
    go get -u ./...
    go mod tidy

release-check:
    goreleaser check

check: tidy fmt-check vet build test

clean:
    rm -f sdrangel-mcp coverage.out gosec.sarif
    rm -rf dist/
