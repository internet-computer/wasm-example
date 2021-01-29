# Go WASM Example

`index.html` is based on `misc/wasm/wasm_exec.html`.

## Build .wasm file.

```shell
GOARCH=wasm GOOS=js go build -o main.wasm
```

## Serve the files.

You can find the result on [localhost:8080](http://localhost:8080/).

```shell
go run server.go
```

## Download Dependencies

```shell
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```

### Links

- [Go + WebAssembly](https://github.com/golang/go/wiki/WebAssembly)
- [syscall/js](https://golang.org/pkg/syscall/js/?GOOS=js&GOARCH=wasm)