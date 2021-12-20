package wasmvm

import (
	"encoding/base64"

	ics23 "github.com/confio/ics23/go"
)

type EncodePackedExistenceProof struct {
	LeafOp        map[int32]int32
	InnerOp       map[int32][]byte
	InnerOpHashOp int32
	Prefix        []byte
	Key           []byte
	Value         []byte
}

func encodePacked(v map[string]interface{}) *EncodePackedExistenceProof {

	var p *ics23.ExistenceProof
	t := v["Proof"].(map[string]interface{})
	t = t["compressed"].(map[string]interface{})
	r = t["entries"].([]interface{})
	innerOp := make(map[int32][]byte, 2)
	leafOp := make(map[int32]int32, 4)

	base64.RawStdEncoding.Decode(innerOp[0], p.Path[0].Prefix)
	base64.RawStdEncoding.Decode(innerOp[1], p.Path[0].Suffix)

	leafOp[0] = ics23.HashOp_value[p.Leaf.Hash.String()]
	leafOp[1] = ics23.HashOp_value[p.Leaf.PrehashKey.String()]
	leafOp[2] = ics23.HashOp_value[p.Leaf.PrehashValue.String()]
	leafOp[3] = ics23.LengthOp_value[p.Leaf.Length.String()]

	var prefix []byte
	var innerOpHash int32
	var key, value []byte

	base64.RawStdEncoding.Decode(key, p.Key)
	base64.RawStdEncoding.Decode(value, p.Value)
	base64.RawStdEncoding.Decode(prefix, p.Leaf.Prefix)
	innerOpHash = ics23.HashOp_value[p.Path[0].Hash.String()]

	return &EncodePackedExistenceProof{
		LeafOp:        leafOp,
		InnerOp:       innerOp,
		InnerOpHashOp: innerOpHash,
		Prefix:        prefix,
		Key:           key,
		Value:         value,
	}

}
