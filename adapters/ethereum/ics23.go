package ethereum

import (
	"bytes"
	"context"
	"encoding/hex"
	"math/big"

	"github.com/0xPolygon/polygon-sdk/helper/keccak"
	"github.com/gochain/gochain/v3/accounts/abi/bind"
	"github.com/gochain/web3"
)

type Packet struct {
	ops   []int32
	proof []byte
	root  []byte
	key   []byte
	value []byte
}

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

	packet := &Packet{
		ops: proof.LeafOp,
		proof: encodePacked(
			proof.Prefix,
			proof.InnerOpPrefix,
			proof.InnerOpSuffix,
			i32tob((uint32(proof.InnerOpHashOp))),
		),
		root:  root,
		key:   proof.Key,
		value: value,
	}

	// signedProofData, err := SignedProofAbiMethod().Inputs.Encode(packet)

	// if err != nil {
	// 	return nil, "", fmt.Errorf("packing for signature proof generation failed")
	// }
	client, err := web3.Dial(adapter.HostAddress)

	if err != nil {
		return false, err
	}

	// web3.
	// client
	ethCallSession := EthereumCallerSession{
		Contract: &EthereumCaller{},
		CallOpts: bind.CallOpts{
			Pending:     false,
			From:        [20]byte{},
			BlockNumber: &big.Int{},
			Context:     nil,
		},
	}

	web3.CallConstantFunction(context.Background(), client, AnconVerifier)

	return signedProofData, resultCid, nil
}

func i32tob(val uint32) []byte {
	r := make([]byte, 4)
	for i := uint32(0); i < 4; i++ {
		r[i] = byte((val >> (8 * i)) & 0xff)
	}
	return r
}
