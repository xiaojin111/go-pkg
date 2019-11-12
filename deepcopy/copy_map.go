package deepcopy

import (
	"bytes"
	"encoding/gob"
)

func init() {
	gob.Register(map[string]interface{}{})
}

// CopyMap 对 map[string]interface{} 类型深拷贝
func CopyMap(m map[string]interface{}) (map[string]interface{}, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)
	err := enc.Encode(m)
	if err != nil {
		return nil, err
	}
	var copy map[string]interface{}
	err = dec.Decode(&copy)
	if err != nil {
		return nil, err
	}
	return copy, nil
}
