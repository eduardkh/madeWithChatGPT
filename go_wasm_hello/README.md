# go_wasm_hello

[Awesome - Now is The Best Time to Learn WebAssembly](https://www.youtube.com/watch?v=v8-yeWXCsi4)

> build the binary

```bash
go mod init wasmtest
go run main.go
GOOS=js GOARCH=wasm go build -o main.wasm
```

> serve the files

```bash
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
python3 -m http.server
```
