package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/anconprotocol/contracts/wasmvm"
	"github.com/anconprotocol/sdk"
	"github.com/anconprotocol/sdk/proofsignature"
	"github.com/second-state/WasmEdge-go/wasmedge"
	dbm "github.com/tendermint/tm-db"
)

func main() {
// to_address, cid, from(which eventually will be the to_addres), to_data(json[query,variables,operation_name]), resolve has to be json blob
	dataFolder := ".ancon"
	anconstorage := sdk.NewStorage(dataFolder)
	db := dbm.NewMemDB()

	proofs, _ := proofsignature.NewIavlAPI(anconstorage, nil, db, 2000, 0)
	homeChain := "http://localhost:8545"
	destinationChain := "http://localhost:8546"

	verifier := "0x71E56696Eb1A1d0b0e96A01A03DA7481e0008F3F"
	submitter := "0x29F4BA75B8BD3CF70a853271E0351e9dA4112AC3"
	// todo: implement root updater and chain interface
	host := wasmvm.NewEvmRelayHost(anconstorage, proofs, homeChain, destinationChain, submitter, verifier)
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

	// a := wasmedge.NewImportObject("ancon")
	/// Instantiate wasm
	file := "/home/rogelio/Code/ancon-contracts/contracts/metadata/pkg/metadata_lib_bg.wasm"
	vm.LoadWasmFile(file)
	vm.RegisterImport(host.GetImports())
	// wasi.InitWasi(

	vm.Validate()
	vm.Instantiate()
	var res interface{}
	var err error

	payload := []byte(`{ "name":"" , "image":"", "description":""}`)
	res, err = vm.ExecuteBindgen("store", wasmedge.Bindgen_return_array, payload)
	if err == nil {
		fmt.Println("Run bindgen -- store:", string(res.([]byte)))
	} else {
		fmt.Println("Run bindgen -- store FAILED")
	}

	cid := strings.Trim(string(res.([]byte)), "\x00")
	sprintRes := fmt.Sprintf(`query { metadata(cid:"%s", path:"/") { image } }`, cid)
	fmt.Println("%s", sprintRes)

	q := []byte(sprintRes)

	res, err = vm.ExecuteBindgen("execute", wasmedge.Bindgen_return_array, q)

	fmt.Println(string(res.([]byte)))

	args := fmt.Sprintf(`mutation {
		transfer(input:{path: "%s", cid: "%s", owner:"%s", newOwner:"%s"}){
		  cid
		}
	  }
	  `, "/", cid, "alice", "bob")
	res, err = vm.ExecuteBindgen("execute", wasmedge.Bindgen_return_array, []byte(args))

	fmt.Println(string(res.([]byte)))

	vm.Release()

	conf.Release()
}
