build:
	go build -o bin/raindrop-sync ./cmd/raindrop-sync
	go build -o bin/raindrop-search ./cmd/raindrop-search
	go build -o bin/raindrop-find ./cmd/raindrop-find
	go build -o bin/raindrop-increment ./cmd/raindrop-increment

format:
	gofmt -w cmd/raindrop-find/main.go
	gofmt -w cmd/raindrop-increment/main.go
	gofmt -w cmd/raindrop-search/main.go
	gofmt -w cmd/raindrop-sync/main.go
	gofmt -w internal/raindrop-find/main.go
	gofmt -w internal/raindrop-increment/main.go
	gofmt -w internal/raindrop-search/main.go
	gofmt -w internal/raindrop-sync/main.go
	gofmt -w pkg/file/exists.go
	gofmt -w pkg/raindrop/api.go
	gofmt -w pkg/sqlite3/db.go
