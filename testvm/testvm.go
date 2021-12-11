package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/anconprotocol/contracts/hexutil"
	"github.com/anconprotocol/sdk"
	"github.com/wasmerio/wasmer-go/wasmer"
)

func main() {
	s := sdk.NewStorage(".ancon")
	w := NewVM(s)

	args := make([]interface{}, 3)
	args[0] = "a"
	args[1] = "1"
	args[2] = "2"

	bz, _ := json.Marshal(args)
	h := hexutil.Encode(bz)

	input := ([]byte)(h)
	w.Run(input)
}

// WASM is the ethereum virtual machine
type WASM struct {
	engine *wasmer.Engine
	store  sdk.Storage
}

// NewEVM creates a new WASM
func NewVM(s sdk.Storage) *WASM {
	engine := wasmer.NewEngine()
	return &WASM{store: s, engine: engine}
}

// Name implements the runtime interface
func (e *WASM) Name() string {
	return "wasm"
}

// Run implements the runtime interface
func (e *WASM) Run(v hexutil.Bytes) hexutil.Bytes {

	wasmBytes, _ := ioutil.ReadFile("/home/rogelio/Code/ancon-contracts/functions/testapp/target/wasm32-wasi/debug/testapp.wasm")

	var args []interface{}
	hexbytes, err := hexutil.Decode((string)(v))

	err = json.Unmarshal(hexbytes, &args)
	if err != nil {
		panic(err)
	}

	// targs := make([]interface{}, len(args))
	// for i := 0; i < len(args); i++ {
	// 	targs[i] = cas	t.ToInt32(args[i])
	// }

	store := wasmer.NewStore(e.engine)

	// Compiles the module
	module, err := wasmer.NewModule(store, wasmBytes)

	if err != nil {
		panic(err)
	}
	wasiEnv, _ := wasmer.NewWasiStateBuilder("wasi-program").
		// Choose according to your actual situation
		// Argument("--foo").
		// Environment("ABC", "DEF").
		// MapDirectory("./", ".").
		Finalize()
	importObject, err := wasiEnv.GenerateImportObject(store, module)
	
	if err != nil {
		panic(err)
	}
	hostFunction := wasmer.NewFunction(
		store,
		wasmer.NewFunctionType(wasmer.NewValueTypes(), wasmer.NewValueTypes(wasmer.I32)),
		func(args []wasmer.Value) ([]wasmer.Value, error) {
			return []wasmer.Value{wasmer.NewI32(42)}, nil
		},
	)
	// hostGlobal := wasmer.NewGlobal(
	// 	store,
	// 	wasmer.NewGlobalType(wasmer.NewValueType(wasmer.I32), wasmer.IMMUTABLE),
	// 	wasmer.NewValue(nil, wasmer.AnyRef),
	// )

	importObject.Register(
		"env",
		map[string]wasmer.IntoExtern{
			"focused_transform": hostFunction,
		},
	)
	importObject.Register(
		"go",
		map[string]wasmer.IntoExtern{
			"debug": hostFunction,
		},
		
	)

	instance, err := wasmer.NewInstance(module, importObject)

	if err != nil {
		panic(err)
	}
	start, err := instance.Exports.GetWasiStartFunction()

	if err != nil {
		panic(err)
	}
	start()

	main, err := instance.Exports.GetFunction("addMetadata")

	if err != nil {
		panic(err)
	}
	// Calls that exported function with Go standard values. The WebAssembly
	// types are inferred and values are casted automatically.
	result, err := main((args)...)

	if err != nil {
		panic(err)
	}
	hexvalue, _ := toHex(result)

	// gasCost := vm.GasPolicy.GetCost()

	// // In the case of not enough gas for precompiled execution we return ErrOutOfGas
	// if c.Gas < gasCost {
	// 	return &runtime.ExecutionResult{
	// 		GasLeft: 0,
	// 		Err:     runtime.ErrOutOfGas,
	// 	}
	// }

	// c.Gas = c.Gas - gasCost

	// result := &runtime.ExecutionResult{
	// 	ReturnValue: returnValue,
	// 	GasLeft:     c.Gas,
	// 	Err:         err,
	// }

	return hexvalue
}

func toHex(result interface{}) ([]byte, error) {
	var hexresult hexutil.Bytes

	hexresult, err := json.Marshal(result)

	if err != nil {
		return nil, err
	}

	hexvalue, err := hexresult.MarshalText()

	if err != nil {
		return nil, err
	}
	return hexvalue, nil
}
