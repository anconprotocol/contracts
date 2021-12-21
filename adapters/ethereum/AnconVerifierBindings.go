// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethereum

import (
	"math/big"
	"strings"

	gochain "github.com/gochain/gochain/v3"
	"github.com/gochain/gochain/v3/accounts/abi"
	"github.com/gochain/gochain/v3/accounts/abi/bind"
	"github.com/gochain/gochain/v3/common"
	"github.com/gochain/gochain/v3/core/types"
	"github.com/gochain/gochain/v3/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = gochain.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ICS23ExistenceProof is an auto generated low-level Go binding around an user-defined struct.
type ICS23ExistenceProof struct {
	Valid bool
	Key   []byte
	Value []byte
	Leaf  ICS23LeafOp
	Path  []ICS23InnerOp
}

// ICS23InnerOp is an auto generated low-level Go binding around an user-defined struct.
type ICS23InnerOp struct {
	Valid  bool
	Hash   uint8
	Prefix []byte
	Suffix []byte
}

// ICS23InnerSpec is an auto generated low-level Go binding around an user-defined struct.
type ICS23InnerSpec struct {
	ChildOrder      []*big.Int
	ChildSize       *big.Int
	MinPrefixLength *big.Int
	MaxPrefixLength *big.Int
	EmptyChild      []byte
	Hash            uint8
}

// ICS23LeafOp is an auto generated low-level Go binding around an user-defined struct.
type ICS23LeafOp struct {
	Valid        bool
	Hash         uint8
	PrehashKey   uint8
	PrehashValue uint8
	Len          uint8
	Prefix       []byte
}

// ICS23ProofSpec is an auto generated low-level Go binding around an user-defined struct.
type ICS23ProofSpec struct {
	LeafSpec  ICS23LeafOp
	InnerSpec ICS23InnerSpec
	MaxDepth  *big.Int
	MinDepth  *big.Int
}

// EthereumABI is the input ABI used to generate the binding from.
const EthereumABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"onlyOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"getIavlSpec\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.LengthOp\",\"name\":\"len\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.LeafOp\",\"name\":\"leafSpec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"childOrder\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"childSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPrefixLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrefixLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"emptyChild\",\"type\":\"bytes\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"}],\"internalType\":\"structICS23.InnerSpec\",\"name\":\"innerSpec\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"maxDepth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDepth\",\"type\":\"uint256\"}],\"internalType\":\"structICS23.ProofSpec\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.LengthOp\",\"name\":\"len\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.LeafOp\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.InnerOp[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structICS23.ExistenceProof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.LengthOp\",\"name\":\"len\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.LeafOp\",\"name\":\"leafSpec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"childOrder\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"childSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPrefixLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrefixLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"emptyChild\",\"type\":\"bytes\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"}],\"internalType\":\"structICS23.InnerSpec\",\"name\":\"innerSpec\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"maxDepth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDepth\",\"type\":\"uint256\"}],\"internalType\":\"structICS23.ProofSpec\",\"name\":\"spec\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"root\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_prefix\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"_leafOpUint\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[][]\",\"name\":\"_innerOp\",\"type\":\"bytes[][]\"},{\"internalType\":\"uint256\",\"name\":\"existenceProofInnerOpHash\",\"type\":\"uint256\"}],\"name\":\"convertProof\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.LengthOp\",\"name\":\"len\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.LeafOp\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.InnerOp[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structICS23.ExistenceProof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"leafOpUint\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes[][]\",\"name\":\"existenceProofInnerOp\",\"type\":\"bytes[][]\"},{\"internalType\":\"uint256\",\"name\":\"existenceProofInnerOpHash\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"existenceProofKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"existenceProofValue\",\"type\":\"bytes\"}],\"name\":\"queryRootCalculation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"leafOpUint\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes[][]\",\"name\":\"existenceProofInnerOp\",\"type\":\"bytes[][]\"},{\"internalType\":\"uint256\",\"name\":\"existenceProofInnerOpHash\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"existenceProofKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"existenceProofValue\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"root\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// Ethereum is an auto generated Go binding around an GoChain contract.
type Ethereum struct {
	EthereumCaller     // Read-only binding to the contract
	EthereumTransactor // Write-only binding to the contract
	EthereumFilterer   // Log filterer for contract events
}

// EthereumCaller is an auto generated read-only Go binding around an GoChain contract.
type EthereumCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumTransactor is an auto generated write-only Go binding around an GoChain contract.
type EthereumTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumFilterer is an auto generated log filtering Go binding around an GoChain contract events.
type EthereumFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumSession is an auto generated Go binding around an GoChain contract,
// with pre-set call and transact options.
type EthereumSession struct {
	Contract     *Ethereum         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthereumCallerSession is an auto generated read-only Go binding around an GoChain contract,
// with pre-set call options.
type EthereumCallerSession struct {
	Contract *EthereumCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// EthereumTransactorSession is an auto generated write-only Go binding around an GoChain contract,
// with pre-set transact options.
type EthereumTransactorSession struct {
	Contract     *EthereumTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// EthereumRaw is an auto generated low-level Go binding around an GoChain contract.
type EthereumRaw struct {
	Contract *Ethereum // Generic contract binding to access the raw methods on
}

// EthereumCallerRaw is an auto generated low-level read-only Go binding around an GoChain contract.
type EthereumCallerRaw struct {
	Contract *EthereumCaller // Generic read-only contract binding to access the raw methods on
}

// EthereumTransactorRaw is an auto generated low-level write-only Go binding around an GoChain contract.
type EthereumTransactorRaw struct {
	Contract *EthereumTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthereum creates a new instance of Ethereum, bound to a specific deployed contract.
func NewEthereum(address common.Address, backend bind.ContractBackend) (*Ethereum, error) {
	contract, err := bindEthereum(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ethereum{EthereumCaller: EthereumCaller{contract: contract}, EthereumTransactor: EthereumTransactor{contract: contract}, EthereumFilterer: EthereumFilterer{contract: contract}}, nil
}

// NewEthereumCaller creates a new read-only instance of Ethereum, bound to a specific deployed contract.
func NewEthereumCaller(address common.Address, caller bind.ContractCaller) (*EthereumCaller, error) {
	contract, err := bindEthereum(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthereumCaller{contract: contract}, nil
}

// NewEthereumTransactor creates a new write-only instance of Ethereum, bound to a specific deployed contract.
func NewEthereumTransactor(address common.Address, transactor bind.ContractTransactor) (*EthereumTransactor, error) {
	contract, err := bindEthereum(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthereumTransactor{contract: contract}, nil
}

// NewEthereumFilterer creates a new log filterer instance of Ethereum, bound to a specific deployed contract.
func NewEthereumFilterer(address common.Address, filterer bind.ContractFilterer) (*EthereumFilterer, error) {
	contract, err := bindEthereum(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthereumFilterer{contract: contract}, nil
}

// bindEthereum binds a generic wrapper to an already deployed contract.
func bindEthereum(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthereumABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethereum *EthereumRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ethereum.Contract.EthereumCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethereum *EthereumRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethereum.Contract.EthereumTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethereum *EthereumRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethereum.Contract.EthereumTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethereum *EthereumCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ethereum.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethereum *EthereumTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethereum.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethereum *EthereumTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethereum.Contract.contract.Transact(opts, method, params...)
}

// ConvertProof is a free data retrieval call binding the contract method 0xb9406360.
//
// Solidity: function convertProof(bytes key, bytes value, bytes _prefix, uint256[] _leafOpUint, bytes[][] _innerOp, uint256 existenceProofInnerOpHash) pure returns((bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[]))
func (_Ethereum *EthereumCaller) ConvertProof(opts *bind.CallOpts, key []byte, value []byte, _prefix []byte, _leafOpUint []*big.Int, _innerOp [][][]byte, existenceProofInnerOpHash *big.Int) (ICS23ExistenceProof, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "convertProof", key, value, _prefix, _leafOpUint, _innerOp, existenceProofInnerOpHash)

	if err != nil {
		return *new(ICS23ExistenceProof), err
	}

	out0 := *abi.ConvertType(out[0], new(ICS23ExistenceProof)).(*ICS23ExistenceProof)

	return out0, err

}

// ConvertProof is a free data retrieval call binding the contract method 0xb9406360.
//
// Solidity: function convertProof(bytes key, bytes value, bytes _prefix, uint256[] _leafOpUint, bytes[][] _innerOp, uint256 existenceProofInnerOpHash) pure returns((bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[]))
func (_Ethereum *EthereumSession) ConvertProof(key []byte, value []byte, _prefix []byte, _leafOpUint []*big.Int, _innerOp [][][]byte, existenceProofInnerOpHash *big.Int) (ICS23ExistenceProof, error) {
	return _Ethereum.Contract.ConvertProof(&_Ethereum.CallOpts, key, value, _prefix, _leafOpUint, _innerOp, existenceProofInnerOpHash)
}

// ConvertProof is a free data retrieval call binding the contract method 0xb9406360.
//
// Solidity: function convertProof(bytes key, bytes value, bytes _prefix, uint256[] _leafOpUint, bytes[][] _innerOp, uint256 existenceProofInnerOpHash) pure returns((bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[]))
func (_Ethereum *EthereumCallerSession) ConvertProof(key []byte, value []byte, _prefix []byte, _leafOpUint []*big.Int, _innerOp [][][]byte, existenceProofInnerOpHash *big.Int) (ICS23ExistenceProof, error) {
	return _Ethereum.Contract.ConvertProof(&_Ethereum.CallOpts, key, value, _prefix, _leafOpUint, _innerOp, existenceProofInnerOpHash)
}

// GetIavlSpec is a free data retrieval call binding the contract method 0x27dcd78c.
//
// Solidity: function getIavlSpec() pure returns(((bool,uint8,uint8,uint8,uint8,bytes),(uint256[],uint256,uint256,uint256,bytes,uint8),uint256,uint256))
func (_Ethereum *EthereumCaller) GetIavlSpec(opts *bind.CallOpts) (ICS23ProofSpec, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "getIavlSpec")

	if err != nil {
		return *new(ICS23ProofSpec), err
	}

	out0 := *abi.ConvertType(out[0], new(ICS23ProofSpec)).(*ICS23ProofSpec)

	return out0, err

}

// GetIavlSpec is a free data retrieval call binding the contract method 0x27dcd78c.
//
// Solidity: function getIavlSpec() pure returns(((bool,uint8,uint8,uint8,uint8,bytes),(uint256[],uint256,uint256,uint256,bytes,uint8),uint256,uint256))
func (_Ethereum *EthereumSession) GetIavlSpec() (ICS23ProofSpec, error) {
	return _Ethereum.Contract.GetIavlSpec(&_Ethereum.CallOpts)
}

// GetIavlSpec is a free data retrieval call binding the contract method 0x27dcd78c.
//
// Solidity: function getIavlSpec() pure returns(((bool,uint8,uint8,uint8,uint8,bytes),(uint256[],uint256,uint256,uint256,bytes,uint8),uint256,uint256))
func (_Ethereum *EthereumCallerSession) GetIavlSpec() (ICS23ProofSpec, error) {
	return _Ethereum.Contract.GetIavlSpec(&_Ethereum.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ethereum *EthereumCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ethereum *EthereumSession) Owner() (common.Address, error) {
	return _Ethereum.Contract.Owner(&_Ethereum.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ethereum *EthereumCallerSession) Owner() (common.Address, error) {
	return _Ethereum.Contract.Owner(&_Ethereum.CallOpts)
}

// QueryRootCalculation is a free data retrieval call binding the contract method 0x4987dfd3.
//
// Solidity: function queryRootCalculation(uint256[] leafOpUint, bytes prefix, bytes[][] existenceProofInnerOp, uint256 existenceProofInnerOpHash, bytes existenceProofKey, bytes existenceProofValue) view returns(bytes)
func (_Ethereum *EthereumCaller) QueryRootCalculation(opts *bind.CallOpts, leafOpUint []*big.Int, prefix []byte, existenceProofInnerOp [][][]byte, existenceProofInnerOpHash *big.Int, existenceProofKey []byte, existenceProofValue []byte) ([]byte, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "queryRootCalculation", leafOpUint, prefix, existenceProofInnerOp, existenceProofInnerOpHash, existenceProofKey, existenceProofValue)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// QueryRootCalculation is a free data retrieval call binding the contract method 0x4987dfd3.
//
// Solidity: function queryRootCalculation(uint256[] leafOpUint, bytes prefix, bytes[][] existenceProofInnerOp, uint256 existenceProofInnerOpHash, bytes existenceProofKey, bytes existenceProofValue) view returns(bytes)
func (_Ethereum *EthereumSession) QueryRootCalculation(leafOpUint []*big.Int, prefix []byte, existenceProofInnerOp [][][]byte, existenceProofInnerOpHash *big.Int, existenceProofKey []byte, existenceProofValue []byte) ([]byte, error) {
	return _Ethereum.Contract.QueryRootCalculation(&_Ethereum.CallOpts, leafOpUint, prefix, existenceProofInnerOp, existenceProofInnerOpHash, existenceProofKey, existenceProofValue)
}

// QueryRootCalculation is a free data retrieval call binding the contract method 0x4987dfd3.
//
// Solidity: function queryRootCalculation(uint256[] leafOpUint, bytes prefix, bytes[][] existenceProofInnerOp, uint256 existenceProofInnerOpHash, bytes existenceProofKey, bytes existenceProofValue) view returns(bytes)
func (_Ethereum *EthereumCallerSession) QueryRootCalculation(leafOpUint []*big.Int, prefix []byte, existenceProofInnerOp [][][]byte, existenceProofInnerOpHash *big.Int, existenceProofKey []byte, existenceProofValue []byte) ([]byte, error) {
	return _Ethereum.Contract.QueryRootCalculation(&_Ethereum.CallOpts, leafOpUint, prefix, existenceProofInnerOp, existenceProofInnerOpHash, existenceProofKey, existenceProofValue)
}

// Verify is a free data retrieval call binding the contract method 0xb0d264e7.
//
// Solidity: function verify((bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[]) proof, ((bool,uint8,uint8,uint8,uint8,bytes),(uint256[],uint256,uint256,uint256,bytes,uint8),uint256,uint256) spec, bytes root, bytes key, bytes value) pure returns()
func (_Ethereum *EthereumCaller) Verify(opts *bind.CallOpts, proof ICS23ExistenceProof, spec ICS23ProofSpec, root []byte, key []byte, value []byte) error {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "verify", proof, spec, root, key, value)

	if err != nil {
		return err
	}

	return err

}

// Verify is a free data retrieval call binding the contract method 0xb0d264e7.
//
// Solidity: function verify((bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[]) proof, ((bool,uint8,uint8,uint8,uint8,bytes),(uint256[],uint256,uint256,uint256,bytes,uint8),uint256,uint256) spec, bytes root, bytes key, bytes value) pure returns()
func (_Ethereum *EthereumSession) Verify(proof ICS23ExistenceProof, spec ICS23ProofSpec, root []byte, key []byte, value []byte) error {
	return _Ethereum.Contract.Verify(&_Ethereum.CallOpts, proof, spec, root, key, value)
}

// Verify is a free data retrieval call binding the contract method 0xb0d264e7.
//
// Solidity: function verify((bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[]) proof, ((bool,uint8,uint8,uint8,uint8,bytes),(uint256[],uint256,uint256,uint256,bytes,uint8),uint256,uint256) spec, bytes root, bytes key, bytes value) pure returns()
func (_Ethereum *EthereumCallerSession) Verify(proof ICS23ExistenceProof, spec ICS23ProofSpec, root []byte, key []byte, value []byte) error {
	return _Ethereum.Contract.Verify(&_Ethereum.CallOpts, proof, spec, root, key, value)
}

// VerifyProof is a free data retrieval call binding the contract method 0xa91eaa36.
//
// Solidity: function verifyProof(uint256[] leafOpUint, bytes prefix, bytes[][] existenceProofInnerOp, uint256 existenceProofInnerOpHash, bytes existenceProofKey, bytes existenceProofValue, bytes root, bytes key, bytes value) pure returns(bool)
func (_Ethereum *EthereumCaller) VerifyProof(opts *bind.CallOpts, leafOpUint []*big.Int, prefix []byte, existenceProofInnerOp [][][]byte, existenceProofInnerOpHash *big.Int, existenceProofKey []byte, existenceProofValue []byte, root []byte, key []byte, value []byte) (bool, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "verifyProof", leafOpUint, prefix, existenceProofInnerOp, existenceProofInnerOpHash, existenceProofKey, existenceProofValue, root, key, value)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0xa91eaa36.
//
// Solidity: function verifyProof(uint256[] leafOpUint, bytes prefix, bytes[][] existenceProofInnerOp, uint256 existenceProofInnerOpHash, bytes existenceProofKey, bytes existenceProofValue, bytes root, bytes key, bytes value) pure returns(bool)
func (_Ethereum *EthereumSession) VerifyProof(leafOpUint []*big.Int, prefix []byte, existenceProofInnerOp [][][]byte, existenceProofInnerOpHash *big.Int, existenceProofKey []byte, existenceProofValue []byte, root []byte, key []byte, value []byte) (bool, error) {
	return _Ethereum.Contract.VerifyProof(&_Ethereum.CallOpts, leafOpUint, prefix, existenceProofInnerOp, existenceProofInnerOpHash, existenceProofKey, existenceProofValue, root, key, value)
}

// VerifyProof is a free data retrieval call binding the contract method 0xa91eaa36.
//
// Solidity: function verifyProof(uint256[] leafOpUint, bytes prefix, bytes[][] existenceProofInnerOp, uint256 existenceProofInnerOpHash, bytes existenceProofKey, bytes existenceProofValue, bytes root, bytes key, bytes value) pure returns(bool)
func (_Ethereum *EthereumCallerSession) VerifyProof(leafOpUint []*big.Int, prefix []byte, existenceProofInnerOp [][][]byte, existenceProofInnerOpHash *big.Int, existenceProofKey []byte, existenceProofValue []byte, root []byte, key []byte, value []byte) (bool, error) {
	return _Ethereum.Contract.VerifyProof(&_Ethereum.CallOpts, leafOpUint, prefix, existenceProofInnerOp, existenceProofInnerOpHash, existenceProofKey, existenceProofValue, root, key, value)
}
