package ethereum

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"

	"github.com/0xPolygon/polygon-sdk/helper/keccak"
	"github.com/umbracle/go-web3"
	"github.com/umbracle/go-web3/abi"
)

type Packet struct {
	ops   []int32
	proof []byte
	root  []byte
	key   []byte
	value []byte
}

func SignedProofAbiMethod() *abi.Method {

	// uint256Type, _ := abi.NewType("uint256", "", nil)
	m, err := abi.NewMethod("verifyProof(uint256[] ops,string proof, string root, string key, string value)")

	if err != nil {
		panic(err)
	}

	return m
}

type OnchainAdapter struct {
	From                   string
	HostAddress            string
	DestinationHostAddress string
	VerifierAddress        web3.Address
	SubmitterAddress       web3.Address
	ChainName              string
	ChainID                int
}

func NewOnchainAdapter(from string, evmHostAddr string,
	evmDestAddr string,
	submitPacketWithProofAddr web3.Address,
	validatorAddr web3.Address) *OnchainAdapter {

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
) []byte {

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
	signedProofData, err := SignedProofAbiMethod().Inputs.Encode(packet)

	if err != nil {
		return nil, "", fmt.Errorf("packing for signature proof generation failed")
	}

	return signedProofData, resultCid, nil
}
func i32tob(val uint32) []byte {
	r := make([]byte, 4)
	for i := uint32(0); i < 4; i++ {
		r[i] = byte((val >> (8 * i)) & 0xff)
	}
	return r
}
