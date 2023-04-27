package util

import (
	"fmt"
	"sort"
	"testing"
)

type testSortString struct {
	Name string `json:"name"`
}

func (this_ *testSortString) GetName() string {
	return this_.Name
}

func TestSortString(t *testing.T) {
	var list []testSortString
	list = append(list, testSortString{Name: "bbb"})
	list = append(list, testSortString{Name: "aaa"})
	list = append(list, testSortString{Name: "ccc"})

	//var sortList StringSortList = list
	fmt.Println("排序前：", list)
	sort.Slice(list, func(i, j int) bool {
		return list[i].Name < list[j].Name //升序  即前面的值比后面的小
	})
	fmt.Println("升序排序：", list)
	sort.Slice(list, func(i, j int) bool {
		return list[i].Name > list[j].Name //降序  即前面的值比后面的大
	})
	fmt.Println("降序排序：", list)
}
