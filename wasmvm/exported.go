package wasmvm

import (
	"fmt"

	"github.com/anconprotocol/sdk"
	"github.com/anconprotocol/sdk/proofsignature"
	"github.com/ipfs/go-graphsync"
	"github.com/ipld/go-ipld-prime"
	"github.com/second-state/WasmEdge-go/wasmedge"
)

type Host struct {
	storage *sdk.Storage
	proof   *proofsignature.IavlProofAPI
	gsync   *graphsync.GraphExchange
}

func NewHost(storage sdk.Storage, proof *proofsignature.IavlProofAPI) *Host {
	return &Host{storage: &storage, proof: proof}
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

	returns := make([]interface{}, 1)
	block, err := sdk.Encode(result)
	if err != nil {
		return nil, wasmedge.Result_Fail
	}

	mem.SetData([]byte(block),  uint(params[0].(int32)), uint(params[1].(int32)))
	returns[0] = block

	return returns, wasmedge.Result_Success
}

// Host functions
func (h *Host) WriteDagBlock(data interface{}, mem *wasmedge.Memory, params []interface{}) ([]interface{}, wasmedge.Result) {
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
