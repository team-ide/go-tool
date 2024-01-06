package datamove

import (
	"fmt"
	"github.com/team-ide/go-dialect/dialect"
	"github.com/team-ide/go-tool/util"
	"os"
	"testing"
)

func GetDataSourceData() *DataSourceData {
	d := &DataSourceData{
		Data: &Data{
			DataType: DataTypeData,
		},
	}
	var data []map[string]interface{}
	for i := 0; i < 222222; i++ {
		data = append(data, map[string]interface{}{
			"userId":   i,
			"name":     fmt.Sprintf("名称%d", i),
			"age":      util.RandomInt(1, 115),
			"password": util.RandomString(6, 16),
		})
	}
	d.DataList = data
	d.ColumnList = append(d.ColumnList, &dialect.ColumnModel{
		ColumnName: "userId",
	})
	d.ColumnList = append(d.ColumnList, &dialect.ColumnModel{
		ColumnName: "name",
	})
	d.ColumnList = append(d.ColumnList, &dialect.ColumnModel{
		ColumnName: "age",
	})
	d.ColumnList = append(d.ColumnList, &dialect.ColumnModel{
		ColumnName: "password",
	})
	return d
}

func GetDataSourceTxt() *DataSourceTxt {
	d := &DataSourceTxt{
		DataSourceFile: &DataSourceFile{
			FilePath: testDir + "test.txt",
		},
		ColSeparator: ",",
	}
	return d
}

func GetDataSourceTxt2() *DataSourceTxt {
	d := &DataSourceTxt{
		DataSourceFile: &DataSourceFile{
			FilePath: testDir + "test2.txt",
		},
		ColSeparator: ",",
	}
	return d
}

func GetDataSourceExcel() *DataSourceExcel {
	d := &DataSourceExcel{
		FilePath: testDir + "test2.xlsx",
	}
	return d
}
func TestDataToData(t *testing.T) {
	var err error
	from := GetDataSourceData()
	to := &DataSourceData{}

	param := &Param{
		BatchNumber: 1000,
	}
	param.init()

	err = DateMove(param, from, to, func(progress *DateMoveProgress) {
		//fmt.Println(util.GetStringValue(progress))
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(util.GetStringValue(to))

	err = DateMove(param, from, to, func(progress *DateMoveProgress) {
		//fmt.Println(util.GetStringValue(progress))
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(util.GetStringValue(to))
}

var testDir = "out/"

func TestDataToTxt(t *testing.T) {
	var err error
	from := GetDataSourceData()

	err = os.MkdirAll(testDir, os.ModePerm)
	to := GetDataSourceTxt()

	param := &Param{
		BatchNumber: 1000,
	}
	param.init()

	err = DateMove(param, from, to, func(progress *DateMoveProgress) {
		//fmt.Println(util.GetStringValue(progress))
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(util.GetStringValue(to))
}

func TestTxtToData(t *testing.T) {
	var err error
	from := GetDataSourceTxt()

	to := &DataSourceData{}

	param := &Param{
		BatchNumber: 1000,
	}
	param.init()

	err = DateMove(param, from, to, func(progress *DateMoveProgress) {
		//fmt.Println(util.GetStringValue(progress))
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(util.GetStringValue(to))
}

func TestTxtToTxt2(t *testing.T) {
	var err error
	from := GetDataSourceTxt()

	to := GetDataSourceTxt2()
	to.ColumnNameMapping = map[string]string{
		"userId": "用户主键",
		"name":   "名称",
	}

	param := &Param{
		BatchNumber: 1000,
	}
	param.init()

	err = DateMove(param, from, to, func(progress *DateMoveProgress) {
		//fmt.Println(util.GetStringValue(progress))
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(util.GetStringValue(to))
}

func TestTxt2ToData(t *testing.T) {
	var err error
	from := GetDataSourceTxt2()
	from.ColumnNameMapping = map[string]string{
		"用户主键": "userId",
		"名称":   "name",
	}

	to := &DataSourceData{}

	param := &Param{
		BatchNumber: 1000,
	}
	param.init()

	err = DateMove(param, from, to, func(progress *DateMoveProgress) {
		//fmt.Println(util.GetStringValue(progress))
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(util.GetStringValue(to))
}

func TestTxt2ToExcel(t *testing.T) {
	var err error
	from := GetDataSourceTxt2()
	from.ColumnNameMapping = map[string]string{
		"用户主键": "userId",
		"名称":   "name",
	}

	to := GetDataSourceExcel()

	param := &Param{
		BatchNumber: 1000,
	}
	param.init()

	err = DateMove(param, from, to, func(progress *DateMoveProgress) {
		//fmt.Println(util.GetStringValue(progress))
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(util.GetStringValue(to))
}
