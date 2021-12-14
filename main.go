package main

import (
	"fmt"
	"os"

	"github.com/second-state/WasmEdge-go/wasmedge"
)

func main() {
	/// Expected Args[0]: program name (./bindgen_funcs)
	/// Expected Args[1]: wasm or wasm-so file (rust_bindgen_funcs_lib_bg.wasm))

	/// Set not to print debug info
	wasmedge.SetLogErrorLevel()

	/// Create configure
	var conf = wasmedge.NewConfigure(wasmedge.WASI)

	/// Create VM with configure
	var vm = wasmedge.NewVMWithConfig(conf)


	/// Init WASI
	var wasi = vm.GetImportObject(wasmedge.WASI)
	wasi.InitWasi(
		os.Args[1:],     /// The args
		os.Environ(),    /// The envs
		[]string{".:."}, /// The mapping preopens
	)

	/// Instantiate wasm
	file := "/home/rogelio/Code/ancon-contracts/contracts/metadata/pkg/metadata_lib_bg.wasm"
	vm.LoadWasmFile(file)
	vm.Validate()
	vm.Instantiate()
	// f , e:= vm.GetFunctionList()
	// fmt.Println("%v", f)
	// fmt.Println("%v", e)
	// fn1 := wasmedge.NewFunction(vm.GetFunctionType("write_store"),host_add,nil,444444)
	// wasi.AddFunction("write_store",fn1)

	/// Run bindgen functions
	var res interface{}
	var err error
	// /// create_line: array, array, array -> array (inputs are JSON stringified)
	// res, err = vm.ExecuteBindgen("create_line", wasmedge.Bindgen_return_array, []byte("{\"x\":1.5,\"y\":3.8}"), []byte("{\"x\":2.5,\"y\":5.8}"), []byte("A thin red line"))
	// if err == nil {
	// 	fmt.Println("Run bindgen -- create_line:", string(res.([]byte)))
	// } else {
	// 	fmt.Println("Run bindgen -- create_line FAILED")
	// }
	// /// say: array -> array
	// res, err = vm.ExecuteBindgen("say", wasmedge.Bindgen_return_array, []byte("bindgen funcs test"))
	// if err == nil {
	// 	fmt.Println("Run bindgen -- say:", string(res.([]byte)))
	// } else {
	// 	fmt.Println("Run bindgen -- say FAILED")
	// }
	// /// obfusticate: array -> array
	// res, err = vm.ExecuteBindgen("obfusticate", wasmedge.Bindgen_return_array, []byte("A quick brown fox jumps over the lazy dog"))
	// if err == nil {
	// 	fmt.Println("Run bindgen -- obfusticate:", string(res.([]byte)))
	// } else {
	// 	fmt.Println("Run bindgen -- obfusticate FAILED")
	// }
	// /// lowest_common_multiple: i32, i32 -> i32
	// res, err = vm.ExecuteBindgen("lowest_common_multiple", wasmedge.Bindgen_return_i32, int32(123), int32(2))
	// if err == nil {
	// 	fmt.Println("Run bindgen -- lowest_common_multiple:", res.(int32))
	// } else {
	// 	fmt.Println("Run bindgen -- lowest_common_multiple FAILED")
	// }
	// /// sha3_digest: array -> array
	// res, err = vm.ExecuteBindgen("sha3_digest", wasmedge.Bindgen_return_array, []byte("This is an important message"))
	// if err == nil {
	// 	fmt.Println("Run bindgen -- sha3_digest:", res.([]byte))
	// } else {
	// 	fmt.Println("Run bindgen -- sha3_digest FAILED")
	// }
	/// keccak_digest: array -> array
	res, err = vm.ExecuteBindgen("query", wasmedge.Bindgen_return_array, []byte("query { metadata()}"))
	if err == nil {
		fmt.Println("Run bindgen -- query:", string(res.([]byte)))
	} else {
		fmt.Println("Run bindgen -- query FAILED")
	}

	vm.Release()


	conf.Release()
}

// Host functions
func host_add(data interface{}, mem *wasmedge.Memory, params []interface{}) ([]interface{}, wasmedge.Result) {
	/// add: externref, i32, i32 -> i32
	/// call the real add function in externref
	fmt.Println("Go: Entering go host function host_add")

	/// Get the externref
	externref := params[0].(wasmedge.ExternRef)

	/// Get the interface{} from externref
	realref := externref.GetRef()

	/// Cast to the functionp
	realfunc := realref.(func(int32, int32) int32)

	/// Call function
	res := realfunc(params[1].(int32), params[2].(int32))

	/// Set the returns
	returns := make([]interface{}, 1)
	returns[0] = res

	/// Return
	fmt.Println("Go: Leaving go host function host_add")
	return returns, wasmedge.Result_Success
}



