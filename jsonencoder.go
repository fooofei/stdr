package stdr

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// jsonEncodeKv encode kv pairs to json string
func jsonEncodeKv(kvList ...interface{}) string {
	vals := make(map[string]interface{}, len(kvList))
	for i := 0; i < len(kvList); i += 2 {
		k, ok := kvList[i].(string)
		if !ok {
			panic(fmt.Sprintf("key is not a string: %s", pretty(kvList[i])))
		}
		var v interface{}
		if i+1 < len(kvList) {
			v = kvList[i+1]
		}
		vals[k] = v
	}
	encoded, err := JSONMarshal(vals)
	if err != nil {
		return err.Error()
	}
	return string(encoded)
}

// JSONMarshal same with json.Marshal but different is
// JSONMarshal will add '\n' at the tail of value
func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	if err != nil {
		return nil, err
	}
	b := buffer.Bytes()
	return bytes.TrimRight(b, "\n"), nil
}
