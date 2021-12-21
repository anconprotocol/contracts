package ethereum

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/0xPolygon/polygon-sdk/helper/keccak"
	"github.com/anconprotocol/contracts/hexutil"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/umbracle/go-web3/abi"
)

// EthereumABI is the input ABI used to generate the binding from.
const EthereumABI = `[{"inputs":[{"internalType":"address","name":"onlyOwner","type":"address"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[],"name":"getIavlSpec","outputs":[{"components":[{"components":[{"internalType":"bool","name":"valid","type":"bool"},{"internalType":"enum ICS23.HashOp","name":"hash","type":"uint8"},{"internalType":"enum ICS23.HashOp","name":"prehash_key","type":"uint8"},{"internalType":"enum ICS23.HashOp","name":"prehash_value","type":"uint8"},{"internalType":"enum ICS23.LengthOp","name":"len","type":"uint8"},{"internalType":"bytes","name":"prefix","type":"bytes"}],"internalType":"struct ICS23.LeafOp","name":"leafSpec","type":"tuple"},{"components":[{"internalType":"uint256[]","name":"childOrder","type":"uint256[]"},{"internalType":"uint256","name":"childSize","type":"uint256"},{"internalType":"uint256","name":"minPrefixLength","type":"uint256"},{"internalType":"uint256","name":"maxPrefixLength","type":"uint256"},{"internalType":"bytes","name":"emptyChild","type":"bytes"},{"internalType":"enum ICS23.HashOp","name":"hash","type":"uint8"}],"internalType":"struct ICS23.InnerSpec","name":"innerSpec","type":"tuple"},{"internalType":"uint256","name":"maxDepth","type":"uint256"},{"internalType":"uint256","name":"minDepth","type":"uint256"}],"internalType":"struct ICS23.ProofSpec","name":"","type":"tuple"}],"stateMutability":"pure","type":"function","constant":true},{"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function","constant":true},{"inputs":[{"components":[{"internalType":"bool","name":"valid","type":"bool"},{"internalType":"bytes","name":"key","type":"bytes"},{"internalType":"bytes","name":"value","type":"bytes"},{"components":[{"internalType":"bool","name":"valid","type":"bool"},{"internalType":"enum ICS23.HashOp","name":"hash","type":"uint8"},{"internalType":"enum ICS23.HashOp","name":"prehash_key","type":"uint8"},{"internalType":"enum ICS23.HashOp","name":"prehash_value","type":"uint8"},{"internalType":"enum ICS23.LengthOp","name":"len","type":"uint8"},{"internalType":"bytes","name":"prefix","type":"bytes"}],"internalType":"struct ICS23.LeafOp","name":"leaf","type":"tuple"},{"components":[{"internalType":"bool","name":"valid","type":"bool"},{"internalType":"enum ICS23.HashOp","name":"hash","type":"uint8"},{"internalType":"bytes","name":"prefix","type":"bytes"},{"internalType":"bytes","name":"suffix","type":"bytes"}],"internalType":"struct ICS23.InnerOp[]","name":"path","type":"tuple[]"}],"internalType":"struct ICS23.ExistenceProof","name":"proof","type":"tuple"},{"components":[{"components":[{"internalType":"bool","name":"valid","type":"bool"},{"internalType":"enum ICS23.HashOp","name":"hash","type":"uint8"},{"internalType":"enum ICS23.HashOp","name":"prehash_key","type":"uint8"},{"internalType":"enum ICS23.HashOp","name":"prehash_value","type":"uint8"},{"internalType":"enum ICS23.LengthOp","name":"len","type":"uint8"},{"internalType":"bytes","name":"prefix","type":"bytes"}],"internalType":"struct ICS23.LeafOp","name":"leafSpec","type":"tuple"},{"components":[{"internalType":"uint256[]","name":"childOrder","type":"uint256[]"},{"internalType":"uint256","name":"childSize","type":"uint256"},{"internalType":"uint256","name":"minPrefixLength","type":"uint256"},{"internalType":"uint256","name":"maxPrefixLength","type":"uint256"},{"internalType":"bytes","name":"emptyChild","type":"bytes"},{"internalType":"enum ICS23.HashOp","name":"hash","type":"uint8"}],"internalType":"struct ICS23.InnerSpec","name":"innerSpec","type":"tuple"},{"internalType":"uint256","name":"maxDepth","type":"uint256"},{"internalType":"uint256","name":"minDepth","type":"uint256"}],"internalType":"struct ICS23.ProofSpec","name":"spec","type":"tuple"},{"internalType":"bytes","name":"root","type":"bytes"},{"internalType":"bytes","name":"key","type":"bytes"},{"internalType":"bytes","name":"value","type":"bytes"}],"name":"verify","outputs":[],"stateMutability":"pure","type":"function","constant":true},{"inputs":[{"internalType":"bytes","name":"key","type":"bytes"},{"internalType":"bytes","name":"value","type":"bytes"},{"internalType":"bytes","name":"_prefix","type":"bytes"},{"internalType":"uint256[]","name":"_leafOpUint","type":"uint256[]"},{"internalType":"bytes","name":"_innerOpPrefix","type":"bytes"},{"internalType":"bytes","name":"_innerOpSuffix","type":"bytes"},{"internalType":"uint256","name":"existenceProofInnerOpHash","type":"uint256"}],"name":"convertProof","outputs":[{"components":[{"internalType":"bool","name":"valid","type":"bool"},{"internalType":"bytes","name":"key","type":"bytes"},{"internalType":"bytes","name":"value","type":"bytes"},{"components":[{"internalType":"bool","name":"valid","type":"bool"},{"internalType":"enum ICS23.HashOp","name":"hash","type":"uint8"},{"internalType":"enum ICS23.HashOp","name":"prehash_key","type":"uint8"},{"internalType":"enum ICS23.HashOp","name":"prehash_value","type":"uint8"},{"internalType":"enum ICS23.LengthOp","name":"len","type":"uint8"},{"internalType":"bytes","name":"prefix","type":"bytes"}],"internalType":"struct ICS23.LeafOp","name":"leaf","type":"tuple"},{"components":[{"internalType":"bool","name":"valid","type":"bool"},{"internalType":"enum ICS23.HashOp","name":"hash","type":"uint8"},{"internalType":"bytes","name":"prefix","type":"bytes"},{"internalType":"bytes","name":"suffix","type":"bytes"}],"internalType":"struct ICS23.InnerOp[]","name":"path","type":"tuple[]"}],"internalType":"struct ICS23.ExistenceProof","name":"","type":"tuple"}],"stateMutability":"pure","type":"function","constant":true},{"inputs":[{"internalType":"uint256[]","name":"leafOpUint","type":"uint256[]"}],"name":"verifyProof","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"pure","type":"function","constant":true}]`

type OnchainAdapter struct {
	From                   string
	HostAddress            string
	DestinationHostAddress string
	VerifierAddress        string
	SubmitterAddress       string
	ChainName              string
	ChainID                int
}

func NewOnchainAdapter(from string, evmHostAddr string,
	evmDestAddr string,
	submitPacketWithProofAddr string,
	validatorAddr string) *OnchainAdapter {

	return &OnchainAdapter{
		From:                   from,
		HostAddress:            evmHostAddr,
		DestinationHostAddress: evmDestAddr,
		VerifierAddress:        validatorAddr,
		SubmitterAddress:       submitPacketWithProofAddr,
		ChainName:              "Ethereum",
		ChainID:                5,
	}
}

// https://gist.github.com/miguelmota/bc4304bb21a8f4cc0a37a0f9347b8bbb
func encodePacked(input ...[]byte) []byte {
	return bytes.Join(input, nil)
}

func encodeBytesString(v string) []byte {
	decoded, err := hex.DecodeString(v)
	if err != nil {
		panic(err)
	}
	return decoded
}

func (adapter *OnchainAdapter) ApplyRequestWithProof(
	ctx context.Context,
	metadataCid string,
	resultCid string,
	fromOwner string,
	toOwner string,
	toAddress string,
	tokenId string,
	prefix string,
) ([]byte, string, error) {

	id := (tokenId)
	var proof []byte
	keccak.Keccak256(proof, encodePacked(
		// Current metadata cid
		[]byte(metadataCid),
		// Current owner (opaque)
		[]byte(fromOwner),
		// Updated metadata cid
		[]byte(resultCid),
		// New owner address
		[]byte(toOwner),
		// Token Address
		[]byte(toAddress),
		// Token Id
		[]byte(id),
		// Contract Prefix
		[]byte(prefix)))

	unsignedProofData := encodePacked(
		[]byte("\x19Ethereum Signed Message:\n32"),
		// Proof
		proof)

	var hash []byte
	keccak.Keccak256(hash, unsignedProofData)

	return nil, resultCid, nil
}

func (adapter *OnchainAdapter) VerifyProof(
	proof *EncodePackedExistenceProof,
	root []byte,
	value []byte,
) (bool, error) {

	client, err := ethclient.Dial(adapter.HostAddress)
	//client.SetChainID(adapter.ChainID)
	if err != nil {
		return false, err
	}

	methods, err := abi.NewABI((EthereumABI))
	if err != nil {
		return false, err
	}

	res, err := adapter.CallConstantFunction(context.Background(), client, methods, adapter.VerifierAddress, "verifyProof",
		proof.LeafOp,
		// encodePacked(
		// 	proof.Prefix,
		// 	proof.InnerOpPrefix,
		// 	proof.InnerOpSuffix,
		// 	i32tob((uint32(proof.InnerOpHashOp))),
		// ),
		// root,
		// proof.Key,
		// value,
	)
	if err != nil {
		return false, err
	}

	fmt.Println(res)
	return true, nil
}

func i32tob(val uint32) []byte {
	r := make([]byte, 4)
	for i := uint32(0); i < 4; i++ {
		r[i] = byte((val >> (8 * i)) & 0xff)
	}
	return r
}

// CallConstantFunction executes a contract function call without submitting a transaction.
func (adapter *OnchainAdapter) CallConstantFunction(ctx context.Context, client *ethclient.Client, myabi *abi.ABI, address string, functionName string, params ...interface{}) ([]interface{}, error) {
	if address == "" {
		return nil, errors.New("no contract address specified")
	}
	fn := myabi.Methods[functionName]
	input, err := fn.Inputs.Encode(params)
	if err != nil {
		return nil, err
	}
	// input, err := myabi.Pack(functionName, goParams...)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to pack values: %v", err)
	// }
	toAddress := common.HexToAddress(address)

	callMsg := ethereum.CallMsg{
		To:   &toAddress,
		Data: input,
	}
	//es, err := client.EstimateGas(context.Background(), callMsg)
	//callMsg.GasPrice = big.NewInt(int64(es))
	res, err := client.CallContract(ctx, callMsg, nil)
	if err != nil {
		return nil, err
	}
	// TODO: calling a function on a contract errors on unpacking, it should probably know it's not a contract before hand if it can
	// fmt.Printf("RESPONSE: %v\n", string(res))
	vals, err := fn.Outputs.Decode(res)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack values from %s: %v", hexutil.Encode(res), err)
	}
	return vals.([]interface{}), nil
}
