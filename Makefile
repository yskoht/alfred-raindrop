build:
	GOOS=darwin GOARCH=arm64 go build -o ./bin/raindrop.arm64 ./cmd/raindrop
	GOOS=darwin GOARCH=amd64 go build -o ./bin/raindrop.amd64 ./cmd/raindrop
	lipo -create -output ./bin/raindrop ./bin/raindrop.arm64 ./bin/raindrop.amd64

format:
	gofmt -w ./cmd/raindrop-find/main.go
	gofmt -w ./cmd/raindrop-increment/main.go
	gofmt -w ./cmd/raindrop-reset/main.go
	gofmt -w ./cmd/raindrop-search/main.go
	gofmt -w ./cmd/raindrop-sync/main.go
	gofmt -w ./internal/raindrop-find/main.go
	gofmt -w ./internal/raindrop-increment/main.go
	gofmt -w ./internal/raindrop-reset/main.go
	gofmt -w ./internal/raindrop-search/main.go
	gofmt -w ./internal/raindrop-sync/main.go
	gofmt -w ./pkg/file/exists.go
	gofmt -w ./pkg/raindrop/api.go
	gofmt -w ./pkg/sqlite3/db.go

clean:
	rm -rf ./bin/raindrop-*
