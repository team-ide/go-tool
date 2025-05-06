package util

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestLoadDirInfo(t *testing.T) {

	dirInfo, err := LoadDirInfo("/data/linkdood", true)
	if err != nil {
		panic(err)
	}
	bs, _ := ObjToMarshalIndent(dirInfo, "", "  ")
	fmt.Println(string(bs))
}

func TestDirD(t *testing.T) {
	dir1 := `D:\工作\全君项目\升级前jar解压\97vserver-base-1.0-SNAPSHOT.jar.src`
	dir2 := `D:\工作\全君项目\升级前jar解压\vserver-base-1.0-SNAPSHOT.jar.src`

	dir1 = FormatPath(dir1) + "/"
	dir2 = FormatPath(dir2) + "/"
	eqDir(dir1, dir2)
}

func eqDir(dir1, dir2 string) {

	fs1, err := os.ReadDir(dir1)
	if err != nil {
		panic(err)
	}
	fs1Map := map[string]os.DirEntry{}
	for _, f := range fs1 {
		fs1Map[f.Name()] = f
	}

	fs2, err := os.ReadDir(dir2)
	if err != nil {
		panic(err)
	}
	fs2Map := map[string]os.DirEntry{}
	for _, f := range fs2 {
		fs2Map[f.Name()] = f
	}

	for k1 := range fs1Map {
		f1 := fs1Map[k1]
		_, find := fs2Map[k1]
		if !find {
			fmt.Println("dir2 not found " + dir2 + k1)
			continue
		}
		if f1.IsDir() {
			eqDir(dir1+k1+"/", dir2+k1+"/")
			continue
		}
		bs1, _ := os.ReadFile(dir1 + k1)
		bs2, _ := os.ReadFile(dir2 + k1)

		str1 := string(bs1)
		str2 := string(bs2)

		if strings.Contains(str1, "/* Location:") {
			str1 = strings.Split(str1, "/* Location:")[0]
		}
		if strings.Contains(str2, "/* Location:") {
			str2 = strings.Split(str2, "/* Location:")[0]
		}
		if str1 != str2 {
			fmt.Println("dir1 not eq dir2 " + dir1 + k1)
			continue
		}
	}

}
