package wasmvm

import (
	"encoding/base64"

	"github.com/buger/jsonparser"

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

var template = `{
	"proof": {
		"proofs": [
			{
				"Proof": {
					"exist": {
						"key": "YmFndXFlZXJhcDVwZHlmend6ZDV4NnR2cGJibjRsZno0bWg2cnd2dGhvbmJ3ZzRrbnBkcGt5dDNucmhrcQ==",
						"leaf": {
							"hash": 1,
							"prehash_value": 1,
							"length": 1,
							"prefix": "AAIC"
						},
						"path": [
							{
								"hash": 1,
								"prefix": "AgQCIGhqEPIiQrR2tMcmliOUwD/Yq+51sHW7EIDc5BAgCtIpIA=="
							}
						]
					}
				}
			}
		]
	},
	"value": null
}`

func encodePacked(v []byte) *EncodePackedExistenceProof {
	// "key": "YmFndXFlZXJhcDVwZHlmend6ZDV4NnR2cGJibjRsZno0bWg2cnd2dGhvbmJ3ZzRrbnBkcGt5dDNucmhrcQ==",
	// "leaf": {
	// 	"hash": 1,
	// 	"prehash_value": 1,
	// 	"length": 1,
	// 	"prefix": "AAIC"
	// },
	// "path": [
	// 	{
	// 		"hash": 1,
	// 		"prefix": "AgQCIGhqEPIiQrR2tMcmliOUwD/Yq+51sHW7EIDc5BAgCtIpIA=="
	// 	}
	// ]

	existPayload, _, _, err := jsonparser.Get(v,"proof","proofs", "[0]", "Proof", "[0]", "exist")
	if err != nil {
		return nil
	}
	_key, _, _, _ := jsonparser.Get(existPayload, "key")
	k, _ := base64.RawStdEncoding.DecodeString(string(_key))
	_value, _, _, _ := jsonparser.Get(existPayload, "value")
	val, _ := base64.RawStdEncoding.DecodeString(string(_value))
	_leaf, _, _, _ := jsonparser.Get(existPayload, "leaf")
	hash, _ := jsonparser.GetInt(_leaf, "hash")
	prehashValue, _ := jsonparser.GetInt(_leaf, "prehash_value")
	length, _ := jsonparser.GetInt(_leaf, "length")
	pre, _ := jsonparser.GetString(_leaf, "prefix")
	p := &ics23.ExistenceProof{
		Key:   k,
		Value: val,
		Leaf:  &ics23.LeafOp{
			Hash:         ics23.HashOp(hash),
			PrehashKey:   ics23.HashOp(hash),
			PrehashValue: ics23.HashOp(prehashValue),
			Length:       ics23.LengthOp(length),
			Prefix:       []byte(pre),
		},
		Path:  []*ics23.InnerOp{},
	}
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
