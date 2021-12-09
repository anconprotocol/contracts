package main

import (
	"fmt"

	"github.com/anconprotocol/contracts/sdk"
)

func main() {

}

//export addMetadata
func AddMetadata(cid, fromOwner, toOwner, metadataCid string) (string, error) {
	jsonmodel, err := sdk.ReadFromStore(cid, "/")
	if err != nil {
		return "", err
	}

	n, err := sdk.FocusedTransform(
		jsonmodel,
		"owner")

	if err != nil {
		return "", fmt.Errorf("")
	}

	// parent update
	n, err = sdk.FocusedTransform(
		n,
		"parent")

	if err != nil {
		return "", fmt.Errorf("focused transform error")
	}

	link, err := sdk.WriteToStore(n)

	if err != nil {
		return "", fmt.Errorf("focused transform error")
	}
	return link, nil
}
