build:
	go build -o ./bin/raindrop ./cmd/raindrop

format:
	gofmt -w ./cmd/raindrop/main.go
	gofmt -w ./internal/raindrop-find/main.go
	gofmt -w ./internal/raindrop-increment/main.go
	gofmt -w ./internal/raindrop-reset/main.go
	gofmt -w ./internal/raindrop-search/main.go
	gofmt -w ./internal/raindrop-sync/main.go
	gofmt -w ./pkg/file/exists.go
	gofmt -w ./pkg/raindrop/api.go
	gofmt -w ./pkg/sqlite3/db.go

clean:
	rm -rf ./bin/raindrop\.*

allclean:
	rm -rf ./bin/raindrop*

