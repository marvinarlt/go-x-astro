dev:
	npm run dev
	go run cmd/server/main.go

preview:
	npm run build
	go run cmd/server/main.go

build:
	npm run build
	go build -o bin/server cmd/server/main.go
