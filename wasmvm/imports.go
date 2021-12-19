package wasmvm

import (
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
	}, []wasmedge.ValType{})

var func3 = wasmedge.NewFunctionType(
	[]wasmedge.ValType{
		wasmedge.ValType_I32,
		wasmedge.ValType_I32,
		wasmedge.ValType_I32,
		wasmedge.ValType_I32,
		wasmedge.ValType_I32,
		wasmedge.ValType_I32,
		wasmedge.ValType_I32,
	}, []wasmedge.ValType{})

var func4 = wasmedge.NewFunctionType(
	[]wasmedge.ValType{
		wasmedge.ValType_I32,
		wasmedge.ValType_I32,
		wasmedge.ValType_I32,
		wasmedge.ValType_I32,
		wasmedge.ValType_I32,
		wasmedge.ValType_I32,
		wasmedge.ValType_I32,
		wasmedge.ValType_I32,
	}, []wasmedge.ValType{})

var func5 = wasmedge.NewFunctionType(
	[]wasmedge.ValType{
		wasmedge.ValType_I32,
		wasmedge.ValType_I32,
		wasmedge.ValType_I32,
		wasmedge.ValType_I32,
		wasmedge.ValType_I32,
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
	fn1 := wasmedge.NewFunction(func2, h.GetProofByCid, nil, 0)
	n.AddFunction("get_proof_by_cid", fn1)

	submit := wasmedge.NewFunction(func3, h.SubmitProof, nil, 0)
	n.AddFunction("submit_proof_onchain", submit)

	verify := wasmedge.NewFunction(func1, h.VerifyProof, nil, 0)
	n.AddFunction("verify_proof_onchain", verify)

	ft := wasmedge.NewFunction(func5, h.FocusedTransformPatch, nil, 0)
	n.AddFunction("focused_transform_patch", ft)

	fn3 := wasmedge.NewFunction(func2, h.ReadDagBlock, nil, 0)
	n.AddFunction("read_dag_block", fn3)

	fn4 := wasmedge.NewFunction(func1, h.WriteDagBlock, nil, 0)
	n.AddFunction("write_dag_block", fn4)

	return n
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
	length := uint(len(bz))
	x := i32tob(uint32(len(bz)))
	mem.SetData(bz, uint(params[0].(int32)), length)
	mem.SetData((x), uint(params[3].(int32)), length)
	return nil, wasmedge.Result_Success
}

func i32tob(val uint32) []byte {
	r := make([]byte, 4)
	for i := uint32(0); i < 4; i++ {
		r[i] = byte((val >> (8 * i)) & 0xff)
	}
	return r
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

// #[no_mangle]
// pub fn submit_proof_onchain(
// 	input: &str,
// 	prev_proof: &str,
// 	cid: &str,
// ) -> [u8; 1024];

// Host functions
func (h *Host) SubmitProof(data interface{}, mem *wasmedge.Memory, params []interface{}) ([]interface{}, wasmedge.Result) {

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

// #[no_mangle]
// pub fn focused_transform_patch(
// 	cid: &str,
// 	path: &str,
// 	prev: &str,
// 	next: &str,
// 	ntype: NodeType,
// ) -> [u8; 1024];
func (h *Host) FocusedTransformPatch(data interface{}, mem *wasmedge.Memory, params []interface{}) ([]interface{}, wasmedge.Result) {

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
	length := uint(len(bz))
	x := i32tob(uint32(len(bz)))
	mem.SetData(bz, uint(params[0].(int32)), length)
	mem.SetData((x), uint(params[3].(int32)), length)
	return nil, wasmedge.Result_Success
}

// #[no_mangle]
// pub fn get_proof_by_cid(key: &str, ret: &i32) -> [u8; 1024];

// Host functions
func (h *Host) GetProofByCid(data interface{}, mem *wasmedge.Memory, params []interface{}) ([]interface{}, wasmedge.Result) {

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
	length := uint(len(bz))
	x := i32tob(uint32(len(bz)))
	mem.SetData(bz, uint(params[0].(int32)), length)
	mem.SetData((x), uint(params[3].(int32)), length)
	return nil, wasmedge.Result_Success
}

// #[no_mangle]
// pub fn verify_proof_onchain(key: &str) -> [u8; 1024];

// Host functions
func (h *Host) VerifyProof(data interface{}, mem *wasmedge.Memory, params []interface{}) ([]interface{}, wasmedge.Result) {

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
