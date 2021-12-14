package wasmvm

import (
	"fmt"

	"github.com/second-state/WasmEdge-go/wasmedge"
)

// Host functions
func host_add(data interface{}, mem *wasmedge.Memory, params []interface{}) ([]interface{}, wasmedge.Result) {
	/// add: externref, i32, i32 -> i32
	/// call the real add function in externref
	fmt.Println("Go: Entering go host function host_add")

	/// Get the externref
	externref := params[0].(wasmedge.ExternRef)

	/// Get the interface{} from externref
	realref := externref.GetRef()

	/// Cast to the functionp
	realfunc := realref.(func(int32, int32) int32)

	/// Call function
	res := realfunc(params[1].(int32), params[2].(int32))

	/// Set the returns
	returns := make([]interface{}, 1)
	returns[0] = res

	/// Return
	fmt.Println("Go: Leaving go host function host_add")
	return returns, wasmedge.Result_Success
}



