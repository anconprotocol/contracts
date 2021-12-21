package wasmvm

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/0xPolygon/polygon-sdk/helper/keccak"
	"github.com/anconprotocol/contracts/adapters/ethereum"
	"github.com/anconprotocol/sdk"
	"github.com/anconprotocol/sdk/proofsignature"
	"github.com/buger/jsonparser"
	_ "github.com/confio/ics23/go"
	"github.com/ipfs/go-graphsync"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/must"
	basicnode "github.com/ipld/go-ipld-prime/node/basic"
	"github.com/ipld/go-ipld-prime/traversal"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/second-state/WasmEdge-go/wasmedge"
)

type Host struct {
	storage *sdk.Storage
	proof   *proofsignature.IavlProofAPI
	gsync   *graphsync.GraphExchange
	adapter *ethereum.OnchainAdapter
}

var (
	WasmFuncType1 *wasmedge.FunctionType = wasmedge.NewFunctionType(
		[]wasmedge.ValType{
			wasmedge.ValType_I32,
			wasmedge.ValType_I32,
			wasmedge.ValType_I32,
		}, []wasmedge.ValType{})

	WasmFuncType2 = wasmedge.NewFunctionType(
		[]wasmedge.ValType{
			wasmedge.ValType_I32,
			wasmedge.ValType_I32,
			wasmedge.ValType_I32,
			wasmedge.ValType_I32,
		}, []wasmedge.ValType{})

	genesis   string = ""
	ROOT_PATH string = "/anconprotocol/onchain"
)

func NewEvmRelayHost(storage sdk.Storage,
	proof *proofsignature.IavlProofAPI,
	evmHostAddr string,
	evmDestAddr string,
	submitPacketWithProofAddr string,
	validatorAddr string) *Host {

	adapter := ethereum.NewOnchainAdapter("0x32A21c1bB6E7C20F547e930b53dAC57f42cd25F6", evmHostAddr, evmDestAddr, submitPacketWithProofAddr, validatorAddr)

	lnk := sdk.CreateCidLink([]byte("merkle tree random root hash"))

	// genesis
	genesis = fmt.Sprintf("%s/%s", ROOT_PATH, lnk.String())
	proof.Service.Set([]byte(genesis), lnk.Bytes())

	return &Host{storage: &storage, proof: proof, adapter: adapter}

}

func (h *Host) GetImports() *wasmedge.ImportObject {

	n := wasmedge.NewImportObject("env")
	fn1 := wasmedge.NewFunction(WasmFuncType2, h.GetProofByCid, nil, 0)
	n.AddFunction("get_proof_by_cid", fn1)

	submit := wasmedge.NewFunction(WasmFuncType2, h.SubmitProof, nil, 0)
	n.AddFunction("submit_proof_onchain", submit)

	verify := wasmedge.NewFunction(WasmFuncType2, h.VerifyProof, nil, 0)
	n.AddFunction("verify_proof_onchain", verify)

	ft := wasmedge.NewFunction(WasmFuncType2, h.FocusedTransformPatch, nil, 0)
	n.AddFunction("focused_transform_patch", ft)

	fn3 := wasmedge.NewFunction(WasmFuncType2, h.ReadDagBlock, nil, 0)
	n.AddFunction("read_dag_block", fn3)

	fn4 := wasmedge.NewFunction(WasmFuncType1, h.WriteDagBlock, nil, 0)
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

	cid := h.storage.Store(ipld.LinkContext{
		LinkPath: ipld.ParsePath(genesis),
	}, n)
	if err != nil {
		return nil, wasmedge.Result_Fail
	}

	var hashed []byte
	r := keccak.Keccak256(hashed, arg1)

	// next
	path := fmt.Sprintf("%s/%s", genesis, cid.String())
	h.proof.Service.Set([]byte(path), r)

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
	var v map[string]interface{}
	json.Unmarshal(arg1, &v)
	nodeType := v["nodeType"].(string)
	link, err := sdk.ParseCidLink(v["cid"].(string))
	path := v["path"].(string)
	previousValue := v["previousValue"].(string)
	nextValue := v["nextValue"].(string)
	if err != nil {
		return nil, wasmedge.Result_Fail
	}

	n, err := h.storage.Load(ipld.LinkContext{}, link)
	if err != nil {
		return nil, wasmedge.Result_Fail
	}

	// patch
	patched, err := traversal.FocusedTransform(
		n,
		datamodel.ParsePath(string(path)),
		func(progress traversal.Progress, prev datamodel.Node) (datamodel.Node, error) {
			if progress.Path.String() == string(path) && must.String(prev) == string(previousValue) {
				nb := prev.Prototype().NewBuilder()
				switch nodeType {
				case "String":
					nb.AssignString((nextValue))
				default:
					nb.AssignBytes([]byte(nextValue))
				}
				return nb.Build(), nil
			}
			return nil, fmt.Errorf("%s not found", path)
		}, false)

	if err != nil {
		return nil, wasmedge.Result_Fail
	}
	cid := h.storage.Store(ipld.LinkContext{}, patched)

	bz := []byte(cid.String())

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
	path := fmt.Sprintf("%s/%s", genesis, cid.String())
	proof, err := h.proof.Service.GetWithProof([]byte(path))
	if err != nil {
		return nil, wasmedge.Result_Fail
	}
	bz := []byte(proof)
	length := uint(len(bz))
	x := i32tob(uint32(len(bz)))
	mem.SetData(bz, uint(params[0].(int32)), length)
	mem.SetData((x), uint(params[3].(int32)), length)
	return nil, wasmedge.Result_Success
}

// #[no_mangle]
// pub fn verify_proof_onchain(key: &str) -> [u8; 1024];

// https://gist.github.com/miguelmota/bc4304bb21a8f4cc0a37a0f9347b8bbb
func EncodeABIPacked(input ...[]byte) []byte {
	return bytes.Join(input, nil)
}

// Host functions
func (h *Host) VerifyProof(data interface{}, mem *wasmedge.Memory, params []interface{}) ([]interface{}, wasmedge.Result) {

	var err error
	arg1, err := mem.GetData(uint(params[1].(int32)), uint(params[2].(int32)))
	if err != nil {
		return nil, wasmedge.Result_Fail
	}
	// gas, err := h.verifier.EstimateGas("verifyProof", abiIcs23Proof)

	abiIcs23Proof := h.adapter.MarshalProof(arg1)
	lnk := strings.Split(string(abiIcs23Proof.Key), "/")
	cid, err := sdk.ParseCidLink(lnk[4])
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

	var hashed []byte
	value := keccak.Keccak256(hashed, []byte(block))
	root, err := h.proof.Service.Hash(&emptypb.Empty{})
	if err != nil {
		return nil, wasmedge.Result_Fail
	}

	currentProofRootHash, err := jsonparser.GetString(root, "hash")

	rootbz, err := base64.StdEncoding.DecodeString(currentProofRootHash)

	ret, err := h.adapter.VerifyProof(abiIcs23Proof, rootbz, value)
	fmt.Println(ret, err)

	bz := []byte("true")
	if err != nil {
		bz = []byte("false")
	}

	length := uint(len(bz))
	x := i32tob(uint32(len(bz)))
	mem.SetData(bz, uint(params[0].(int32)), length)
	mem.SetData((x), uint(params[3].(int32)), length)
	return nil, wasmedge.Result_Success
}
