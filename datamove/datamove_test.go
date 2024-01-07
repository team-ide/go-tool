package datamove

import (
	"fmt"
	"github.com/team-ide/go-dialect/dialect"
	"github.com/team-ide/go-tool/util"
	"os"
	"testing"
)

func GetDataSourceData() *DataSourceData {
	d := &DataSourceData{}
	d.ColumnList = []*Column{
		{ColumnModel: dialect.ColumnModel{ColumnName: "userId"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "name"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "age"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "password"}},
	}
	return d
}

func GetDataSourceTxt() *DataSourceTxt {
	d := &DataSourceTxt{}
	d.FilePath = testDir + "txt.txt"
	d.ColSeparator = ","
	d.ColumnList = []*Column{
		{ColumnModel: dialect.ColumnModel{ColumnName: "userId"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "name"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "age"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "password"}},
	}

	return d
}

func GetDataSourceTxt2() *DataSourceTxt {
	d := &DataSourceTxt{}
	d.FilePath = testDir + "txt2.txt"
	d.ColSeparator = ","
	d.ColumnList = []*Column{
		{ColumnModel: dialect.ColumnModel{ColumnName: "这是主键"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "这是姓名"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "这是年龄"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "这是密码"}},
	}
	return d
}

func GetDataSourceTxt3() *DataSourceTxt {
	d := &DataSourceTxt{}
	d.FilePath = testDir + "txt3.txt"
	d.ColSeparator = ","
	d.ColumnList = []*Column{
		{ColumnModel: dialect.ColumnModel{ColumnName: "这是主键"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "这是姓名"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "这是年龄"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "这是密码"}},
	}
	return d
}

func GetDataSourceExcel() *DataSourceExcel {
	d := &DataSourceExcel{
		FilePath: testDir + "excel.xlsx",
	}
	d.ColumnList = []*Column{
		{ColumnModel: dialect.ColumnModel{ColumnName: "userId"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "name"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "age"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "password"}},
	}
	return d
}

func GetDataSourceExcel2() *DataSourceExcel {
	d := &DataSourceExcel{
		FilePath: testDir + "excel2.xlsx",
	}
	d.ColumnList = []*Column{
		{ColumnModel: dialect.ColumnModel{ColumnName: "这是主键"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "这是姓名"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "这是年龄"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "这是密码"}},
	}
	return d
}

func GetDataSourceExcel3() *DataSourceExcel {
	d := &DataSourceExcel{
		FilePath: testDir + "excel3.xlsx",
	}
	d.ColumnList = []*Column{
		{ColumnModel: dialect.ColumnModel{ColumnName: "这是主键"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "这是姓名"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "这是年龄"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "这是密码"}},
	}
	return d
}
func TestDataToData(t *testing.T) {
	var err error
	from := GetDataSourceData()

	for i := 0; i < 22; i++ {
		from.DataList = append(from.DataList, map[string]interface{}{
			"userId":   util.NextId(),
			"name":     fmt.Sprintf("名称%d", i),
			"age":      util.RandomInt(1, 115),
			"password": util.RandomString(6, 16),
		})
	}
	to := GetDataSourceData()

	param := &Param{
		BatchNumber: 1000,
	}
	param.init()

	var p *DateMoveProgress
	err = DateMove(param, from, to, func(progress *DateMoveProgress) {
		p = progress
		//fmt.Println(util.GetStringValue(progress))
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(util.GetStringValue(p))
	fmt.Println(util.GetStringValue(to.Total))
	fmt.Println(util.GetStringValue(to.DataList))

	err = DateMove(param, from, to, func(progress *DateMoveProgress) {
		p = progress
		//fmt.Println(util.GetStringValue(progress))
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(util.GetStringValue(p))
	fmt.Println(util.GetStringValue(to.Total))
	fmt.Println(util.GetStringValue(to.DataList))
}

var testDir = "out/"

func TestDataToTxt(t *testing.T) {
	var err error
	from := GetDataSourceData()

	for i := 0; i < 22; i++ {
		from.DataList = append(from.DataList, map[string]interface{}{
			"userId":   util.NextId(),
			"name":     fmt.Sprintf("名称%d", i),
			"age":      util.RandomInt(1, 115),
			"password": util.RandomString(6, 16),
		})
	}

	err = os.MkdirAll(testDir, os.ModePerm)

	param := &Param{
		BatchNumber: 1000,
	}
	param.init()

	var p *DateMoveProgress
	err = DateMove(param, from, GetDataSourceTxt(), func(progress *DateMoveProgress) {
		p = progress
		//fmt.Println(util.GetStringValue(progress))
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(util.GetStringValue(p))

	err = DateMove(param, from, GetDataSourceExcel(), func(progress *DateMoveProgress) {
		p = progress
		//fmt.Println(util.GetStringValue(progress))
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(util.GetStringValue(p))
}

func TestTxtToData(t *testing.T) {
	var err error
	from := GetDataSourceTxt()

	to := GetDataSourceData()

	param := &Param{
		BatchNumber: 1000,
	}
	param.init()

	var p *DateMoveProgress
	err = DateMove(param, from, to, func(progress *DateMoveProgress) {
		p = progress
		//fmt.Println(util.GetStringValue(progress))
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(util.GetStringValue(p))
	fmt.Println(util.GetStringValue(to.Total))
	fmt.Println(util.GetStringValue(to.DataList))
}

func TestExcelToData(t *testing.T) {
	var err error
	from := GetDataSourceExcel()

	to := GetDataSourceData()

	param := &Param{
		BatchNumber: 1000,
	}
	param.init()

	var p *DateMoveProgress
	err = DateMove(param, from, to, func(progress *DateMoveProgress) {
		p = progress
		//fmt.Println(util.GetStringValue(progress))
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(util.GetStringValue(p))
	fmt.Println(util.GetStringValue(to.Total))
	fmt.Println(util.GetStringValue(to.DataList))
}

func TestTxtToTxt2(t *testing.T) {
	var err error
	from := GetDataSourceTxt()

	to := GetDataSourceTxt2()

	param := &Param{
		BatchNumber: 1000,
	}
	param.init()

	var p *DateMoveProgress
	err = DateMove(param, from, to, func(progress *DateMoveProgress) {
		p = progress
		//fmt.Println(util.GetStringValue(progress))
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(util.GetStringValue(p))
}

func TestExcelToExcel2(t *testing.T) {
	var err error
	from := GetDataSourceExcel()

	to := GetDataSourceExcel2()

	param := &Param{
		BatchNumber: 1000,
	}
	param.init()

	var p *DateMoveProgress
	err = DateMove(param, from, to, func(progress *DateMoveProgress) {
		p = progress
		//fmt.Println(util.GetStringValue(progress))
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(util.GetStringValue(p))
}

func TestTxt2ToExcel3(t *testing.T) {
	var err error
	from := GetDataSourceTxt2()

	to := GetDataSourceExcel3()

	param := &Param{
		BatchNumber: 1000,
	}
	param.init()

	var p *DateMoveProgress
	err = DateMove(param, from, to, func(progress *DateMoveProgress) {
		p = progress
		//fmt.Println(util.GetStringValue(progress))
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(util.GetStringValue(p))
}
func TestExcel2ToTxt3(t *testing.T) {
	var err error
	from := GetDataSourceExcel2()

	to := GetDataSourceTxt3()

	param := &Param{
		BatchNumber: 1000,
	}
	param.init()

	var p *DateMoveProgress
	err = DateMove(param, from, to, func(progress *DateMoveProgress) {
		p = progress
		//fmt.Println(util.GetStringValue(progress))
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(util.GetStringValue(p))
}
