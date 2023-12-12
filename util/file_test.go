package util

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestLoadDirInfo(t *testing.T) {

	dirInfo, err := LoadDirInfo("/data/linkdood", true)
	if err != nil {
		panic(err)
	}
	bs, _ := json.MarshalIndent(dirInfo, "", "  ")
	fmt.Println(string(bs))
}
