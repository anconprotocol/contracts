all: wasm-example

wasm-example:   contract.go
	GOOS=js GOARCH=wasm go build -o wasm-example   contract.go