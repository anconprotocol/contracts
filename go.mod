module github.com/anconprotocol/contracts

go 1.16

replace github.com/libp2p/go-libp2p => github.com/libp2p/go-libp2p v0.14.1

replace github.com/libp2p/go-libp2p-core v0.10.0 => github.com/libp2p/go-libp2p-core v0.9.0

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

require (
	github.com/0xPolygon/polygon-sdk v0.0.0-20211207172349-a9ee5ed12815
	github.com/99designs/gqlgen v0.14.0
	github.com/OneOfOne/xxhash v1.2.8 // indirect
	github.com/Yamashou/gqlgenc v0.0.2
	github.com/anconprotocol/sdk v0.0.0-20211220175629-81de6df5eda5
	github.com/buger/jsonparser v0.0.0-20181115193947-bf1c66bbce23
	github.com/confio/ics23/go v0.6.6
	github.com/deckarep/golang-set v1.7.1 // indirect
	github.com/ethereum/go-ethereum v1.10.13
	github.com/ipfs/go-graphsync v0.9.3
	github.com/ipld/go-ipld-prime v0.14.0
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/second-state/WasmEdge-go v0.9.0
	github.com/tendermint/tm-db v0.6.4
	github.com/umbracle/go-web3 v0.0.0-20211208145232-a62dc1e205cc
	github.com/vektah/gqlparser/v2 v2.2.0
	google.golang.org/protobuf v1.27.1
)
