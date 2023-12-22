# Proof of Concept: Go X Astro

This concept combines GO and Astro by embedding the static files, that are generated by Astro, into the binary that GO creates. The main idea is to build the frontend and backend code into a **single binary**, while maintaining the possibility of dynamic server side rendered html.

## Developing

Run the following command to start separated servers for frontend and backend:

```sh
make dev
```

Or without Makefile:

```sh
npm run dev && go run cmd/server/main.go
```

## Preview

To build a preview version that already combines the frontend and backend, run:

```sh
make preview
```

Or without Makefile:

```sh
npm run build && go run cmd/server/main.go
```

## Building

Finally to build everything into a single binary, run:

```sh
make build
```

Or without Makefile:

```sh
npm run build && go build -o bin/server cmd/server/main.go
```
