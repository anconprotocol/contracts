package main

import (
	"fmt"
	"os"

	"github.com/anconprotocol/contracts/wasmvm"
	"github.com/anconprotocol/sdk"
	"github.com/anconprotocol/sdk/proofsignature"
	"github.com/second-state/WasmEdge-go/wasmedge"
	dbm "github.com/tendermint/tm-db"
)

func main() {

	dataFolder := ".ancon"
	anconstorage := sdk.NewStorage(dataFolder)
	db := dbm.NewMemDB()

	proofs, _ := proofsignature.NewIavlAPI(anconstorage, nil, db, 2000, 0)

	host := wasmvm.NewHost(anconstorage, proofs)
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

	var type1 = wasmedge.NewFunctionType(
		[]wasmedge.ValType{
			wasmedge.ValType_V128,
		}, []wasmedge.ValType{
			wasmedge.ValType_V128,
		})
	var type2 = wasmedge.NewFunctionType(
		[]wasmedge.ValType{
			wasmedge.ValType_V128,
			wasmedge.ValType_V128,
		}, []wasmedge.ValType{
			wasmedge.ValType_V128,
		})
	fn1 := wasmedge.NewFunction(type2, host.WriteStore, nil, 0)
	wasi.AddFunction("write_store", fn1)

	fn2 := wasmedge.NewFunction(type1, host.ReadStore, nil, 0)
	wasi.AddFunction("read_store", fn2)

	fn3 := wasmedge.NewFunction(type2, host.ReadDagBlock, nil, 0)
	wasi.AddFunction("read_dag_block", fn3)

	fn4 := wasmedge.NewFunction(type1, host.WriteDagBlock, nil, 0)
	wasi.AddFunction("write_dag_block", fn4)
 
	vm.Validate()
	vm.Instantiate()

	f, e := vm.GetFunctionList()
	fmt.Println("%v", f)
	fmt.Println("%v", e)	/// Run bindgen functions
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
	res, err = vm.ExecuteBindgen("execute", wasmedge.Bindgen_return_array, []byte(`query { metadata(cid:"",path:"")  {image}}`))
	if err == nil {
		fmt.Println("Run bindgen -- query:", string(res.([]byte)))
	} else {
		fmt.Println("Run bindgen -- query FAILED")
	}

	vm.Release()

	conf.Release()
}
