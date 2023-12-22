client-dev:
	(cd client && npm run dev)

server-dev:
	go run cmd/server/main.go

preview:
	(cd client && npm run build)
	go run cmd/server/main.go

build:
	(cd client && npm run build)
	go build -o bin/server cmd/server/main.go
