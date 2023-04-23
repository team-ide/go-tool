package util

import (
	"bytes"
	"encoding/json"
)

func JSONDecodeUseNumber(bs []byte, obj interface{}) (err error) {
	d := json.NewDecoder(bytes.NewReader(bs))
	d.UseNumber()
	err = d.Decode(obj)
	return
}
