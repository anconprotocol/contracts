package sdk

import "encoding/json"

func FocusedTransform(json json.RawMessage, selectPath string) (json.RawMessage, error)
func ReadFromStore(cid, path string) (json.RawMessage, error)
func WriteToStore(node json.RawMessage) (string, error)
