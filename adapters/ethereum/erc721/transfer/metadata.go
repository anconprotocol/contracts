package transfer

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"

	"github.com/0xPolygon/polygon-sdk/crypto"
	"github.com/0xPolygon/polygon-sdk/helper/keccak"
	"github.com/umbracle/go-web3/abi"
)

type Packet struct {
	metadataCid string
	fromOwner   string
	resultCid   string
	toOwner     string
	toAddress   string
	id          string
	prefix      string
	signature   []byte
}

func SignedProofAbiMethod() *abi.Method {

	// uint256Type, _ := abi.NewType("uint256", "", nil)
	m, err := abi.NewMethod("transferURIWithProof(string metadataCid, string fromOwner, string resultCid, string toOwner, string toAddress, string tokenId, string prefix, bytes signature)")

	if err != nil {
		panic(err)
	}

	return m
}

type OnchainAdapter struct {
	PrivateKey *ecdsa.PrivateKey
	ChainName  string
	ChainID    int
}

func NewOnchainAdapter(pk *ecdsa.PrivateKey) OnchainAdapter {

	return OnchainAdapter{
		PrivateKey: pk,
		ChainName:  "Ethereum",
		ChainID:    5,
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

	signature, err := crypto.Sign(adapter.PrivateKey, hash)
	if err != nil {
		return nil, "", fmt.Errorf("signing failed")
	}

	packet := &Packet{
		metadataCid: metadataCid,
		fromOwner:   fromOwner,
		resultCid:   resultCid,
		toOwner:     toOwner,
		toAddress:   toAddress,
		id:          id,
		prefix:      prefix,
		signature:   signature,
	}
	signedProofData, err := SignedProofAbiMethod().Inputs.Encode(packet)

	if err != nil {
		return nil, "", fmt.Errorf("packing for signature proof generation failed")
	}

	return signedProofData, resultCid, nil
}
