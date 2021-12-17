package wasmvm

import (
	"fmt"

	"github.com/anconprotocol/sdk"
	"github.com/anconprotocol/sdk/proofsignature"
	"github.com/ipfs/go-graphsync"
	"github.com/ipld/go-ipld-prime"
	basicnode "github.com/ipld/go-ipld-prime/node/basic"
	"github.com/second-state/WasmEdge-go/wasmedge"
)

type Host struct {
	storage *sdk.Storage
	proof   *proofsignature.IavlProofAPI
	gsync   *graphsync.GraphExchange
}

var func1 = wasmedge.NewFunctionType(
	[]wasmedge.ValType{
		wasmedge.ValType_I32,
		wasmedge.ValType_I32,
		wasmedge.ValType_I32,
	}, []wasmedge.ValType{})
var func2 = wasmedge.NewFunctionType(
	[]wasmedge.ValType{
		wasmedge.ValType_I32,
		wasmedge.ValType_I32,
		wasmedge.ValType_I32,
		wasmedge.ValType_I32,
		wasmedge.ValType_I32,
	}, []wasmedge.ValType{})

func NewHost(storage sdk.Storage, proof *proofsignature.IavlProofAPI) *Host {
	return &Host{storage: &storage, proof: proof}

}

func (h *Host) GetImports() *wasmedge.ImportObject {

	n := wasmedge.NewImportObject("env")
	fn1 := wasmedge.NewFunction(func2, h.WriteStore, nil, 0)
	n.AddFunction("write_store", fn1)

	fn2 := wasmedge.NewFunction(func1, h.ReadStore, nil, 0)
	n.AddFunction("read_store", fn2)

	fn3 := wasmedge.NewFunction(func2, h.ReadDagBlock, nil, 0)
	n.AddFunction("read_dag_block", fn3)

	fn4 := wasmedge.NewFunction(func1, h.WriteDagBlock, nil, 0)
	n.AddFunction("write_dag_block", fn4)

	return n
}

// Host functions
func (h *Host) WriteStore(data interface{}, mem *wasmedge.Memory, params []interface{}) ([]interface{}, wasmedge.Result) {
	/// add: externref, i32, i32 -> i32
	/// call the real add function in externref
	fmt.Println("Go: Entering go host function write_store")

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

// Host functions
func (h *Host) ReadStore(data interface{}, mem *wasmedge.Memory, params []interface{}) ([]interface{}, wasmedge.Result) {
	/// add: externref, i32, i32 -> i32
	/// call the real add function in externref
	fmt.Println("Go: Entering go host function read_store")

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

// Host functions
func (h *Host) ReadDagBlock(data interface{}, mem *wasmedge.Memory, params []interface{}) ([]interface{}, wasmedge.Result) {

	arg1, err := mem.GetData(uint(params[1].(int32)), uint(params[2].(int32)))
	if err != nil {
		return nil, wasmedge.Result_Fail
	}
	/// Call function
	// arg2, err := mem.GetData(uint(params[3].(int32)), uint(params[4].(int32)))
	// if err != nil {
	// 	return nil, wasmedge.Result_Fail
	// }

	cid, err := sdk.ParseCidLink(string(arg1))
	if err != nil {
		return nil, wasmedge.Result_Fail
	}

	// path := string(arg2)

	result, err := h.storage.Load(ipld.LinkContext{}, cid)
	if err != nil {
		return nil, wasmedge.Result_Fail
	}

	block, err := sdk.Encode(result)
	if err != nil {
		return nil, wasmedge.Result_Fail
	}

	bz := []byte(block)
	mem.SetData(bz, uint(params[0].(int32)), uint(len(bz)))

	return nil, wasmedge.Result_Success
}

// Host functions
func (h *Host) WriteDagBlock(data interface{}, mem *wasmedge.Memory, params []interface{}) ([]interface{}, wasmedge.Result) {

	arg1, err := mem.GetData(uint(params[1].(int32)), uint(params[2].(int32)))
	if err != nil {
		return nil, wasmedge.Result_Fail
	}
	/// Call function
	// arg2, err := mem.GetData(uint(params[3].(int32)), uint(params[4].(int32)))
	// if err != nil {
	// 	return nil, wasmedge.Result_Fail
	// }

	n, err := sdk.Decode(basicnode.Prototype.Any, (string(arg1)))
	cid := h.storage.Store(ipld.LinkContext{}, n)
	if err != nil {
		return nil, wasmedge.Result_Fail
	}

	bz := []byte(cid.String())

	err = mem.SetData(bz, uint(params[0].(int32)), uint(len(bz)))
	if err != nil {
		return nil, wasmedge.Result_Fail
	}

	return nil, wasmedge.Result_Success
}
