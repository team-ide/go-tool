package datamove

import (
	"fmt"
	"github.com/team-ide/go-dialect/dialect"
	"github.com/team-ide/go-tool/db"
	_ "github.com/team-ide/go-tool/db/db_type_sqlite"
	"github.com/team-ide/go-tool/elasticsearch"
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

func GetDataSourceDb() *DataSourceDb {
	d := &DataSourceDb{
		TableName: "TM_LOG",
	}
	var err error
	d.Service, err = db.New(&db.Config{
		Type:         "sqlite3",
		DatabasePath: testDir + "db",
	})
	if err != nil {
		panic(err)
	}
	//d.ColumnList = []*Column{
	//	{ColumnModel: dialect.ColumnModel{ColumnName: "这是主键"}},
	//	{ColumnModel: dialect.ColumnModel{ColumnName: "这是姓名"}},
	//	{ColumnModel: dialect.ColumnModel{ColumnName: "这是年龄"}},
	//	{ColumnModel: dialect.ColumnModel{ColumnName: "这是密码"}},
	//}
	return d
}

func GetDataSourceDb2() *DataSourceDb {
	d := &DataSourceDb{
		TableName: "TM_LOG",
	}
	var err error
	d.Service, err = db.New(&db.Config{
		Type:         "sqlite3",
		DatabasePath: testDir + "db2",
	})
	if err != nil {
		panic(err)
	}
	_, _ = d.Service.Exec(`DELETE FROM TM_LOG`, nil)
	d.ColumnList = []*Column{
		{ColumnModel: dialect.ColumnModel{ColumnName: "logId"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "loginId"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "userId"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "userName"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "userAccount"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "ip"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "action"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "method"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "param"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "data"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "status"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "error"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "useTime"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "startTime"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "endTime"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "userAgent"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "createTime"}},
	}
	return d
}

func GetDataSourceDbTxt() *DataSourceTxt {
	d := &DataSourceTxt{}
	d.FilePath = testDir + "db.txt"
	d.ColumnList = []*Column{
		{ColumnModel: dialect.ColumnModel{ColumnName: "logId"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "loginId"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "userId"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "userName"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "userAccount"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "ip"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "action"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "method"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "param"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "data"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "status"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "error"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "useTime"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "startTime"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "endTime"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "userAgent"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "createTime"}},
	}
	return d
}

func GetDataSourceDbExcel() *DataSourceExcel {
	d := &DataSourceExcel{}
	d.FilePath = testDir + "db.xlsx"
	d.ColumnList = []*Column{
		{ColumnModel: dialect.ColumnModel{ColumnName: "logId"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "loginId"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "userId"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "userName"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "userAccount"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "ip"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "action"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "method"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "param"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "data"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "status"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "error"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "useTime"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "startTime"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "endTime"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "userAgent"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "createTime"}},
	}
	return d
}

func GetDataSourceEs() *DataSourceEs {
	d := &DataSourceEs{
		IndexName: "index_xxx",
		IdName:    "logId",
	}
	var err error
	d.Service, err = elasticsearch.New(&elasticsearch.Config{
		Url: "http://127.0.0.1:9200/",
	})
	if err != nil {
		panic(err)
	}
	d.ColumnList = []*Column{
		{ColumnModel: dialect.ColumnModel{ColumnName: "logId"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "loginId"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "userId"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "userName"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "userAccount"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "ip"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "action"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "method"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "param"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "data"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "status"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "error"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "useTime"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "startTime"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "endTime"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "userAgent"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "createTime"}},
	}
	return d
}
func GetDataSourceEs2() *DataSourceEs {
	d := &DataSourceEs{
		IndexName: "index_2",
		IdName:    "logId",
	}
	var err error
	d.Service, err = elasticsearch.New(&elasticsearch.Config{
		Url: "http://127.0.0.1:9200/",
	})
	if err != nil {
		panic(err)
	}
	_ = d.Service.DeleteIndex("index_2")
	d.ColumnList = []*Column{
		{ColumnModel: dialect.ColumnModel{ColumnName: "logId"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "loginId"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "userId"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "userName"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "userAccount"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "ip"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "action"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "method"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "param"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "data"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "status"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "error"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "useTime"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "startTime"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "endTime"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "userAgent"}},
		{ColumnModel: dialect.ColumnModel{ColumnName: "createTime"}},
	}
	return d
}

func NewProgress() *Progress {
	p := &Progress{}
	p.BatchNumber = 1000
	return p
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

	p := NewProgress()
	err = DateMove(p, from, to)
	if err != nil {
		panic(err)
	}
	fmt.Println(util.GetStringValue(p))
	fmt.Println(util.GetStringValue(to.Total))
	fmt.Println(util.GetStringValue(to.DataList))

	p = NewProgress()
	err = DateMove(p, from, to)
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

	p := NewProgress()
	err = DateMove(p, from, GetDataSourceTxt())
	if err != nil {
		panic(err)
	}
	fmt.Println(util.GetStringValue(p))

	p = NewProgress()
	err = DateMove(p, from, GetDataSourceTxt())
	if err != nil {
		panic(err)
	}
	fmt.Println(util.GetStringValue(p))
}

func TestTxtToData(t *testing.T) {
	var err error
	from := GetDataSourceTxt()

	to := GetDataSourceData()

	p := NewProgress()
	err = DateMove(p, from, to)
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

	p := NewProgress()
	err = DateMove(p, from, to)
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

	p := NewProgress()
	err = DateMove(p, from, to)
	if err != nil {
		panic(err)
	}
	fmt.Println(util.GetStringValue(p))
}

func TestExcelToExcel2(t *testing.T) {
	var err error
	from := GetDataSourceExcel()

	to := GetDataSourceExcel2()

	p := NewProgress()
	err = DateMove(p, from, to)
	if err != nil {
		panic(err)
	}
	fmt.Println(util.GetStringValue(p))
}

func TestTxt2ToExcel3(t *testing.T) {
	var err error
	from := GetDataSourceTxt2()

	to := GetDataSourceExcel3()

	p := NewProgress()
	err = DateMove(p, from, to)
	if err != nil {
		panic(err)
	}
	fmt.Println(util.GetStringValue(p))
}
func TestExcel2ToTxt3(t *testing.T) {
	var err error
	from := GetDataSourceExcel2()

	to := GetDataSourceTxt3()

	p := NewProgress()
	err = DateMove(p, from, to)
	if err != nil {
		panic(err)
	}
	fmt.Println(util.GetStringValue(p))
}

func TestDbToTxt(t *testing.T) {
	var err error
	from := GetDataSourceDb()

	to := GetDataSourceDbTxt()

	p := NewProgress()
	err = DateMove(p, from, to)
	if err != nil {
		panic(err)
	}
	fmt.Println(util.GetStringValue(p))
}

func TestDbToExcel(t *testing.T) {
	var err error
	from := GetDataSourceDb()

	to := GetDataSourceDbExcel()

	p := NewProgress()
	err = DateMove(p, from, to)
	if err != nil {
		panic(err)
	}
	fmt.Println(util.GetStringValue(p))
}

func TestDbToDb2(t *testing.T) {
	var err error
	from := GetDataSourceDb()

	to := GetDataSourceDb2()

	p := NewProgress()
	err = DateMove(p, from, to)
	if err != nil {
		panic(err)
	}
	fmt.Println(util.GetStringValue(p))
}

func TestExcelToDb2(t *testing.T) {
	var err error
	from := GetDataSourceDbExcel()

	to := GetDataSourceDb2()

	p := NewProgress()
	err = DateMove(p, from, to)
	if err != nil {
		panic(err)
	}
	fmt.Println(util.GetStringValue(p))
}

func TestTxtToDb2(t *testing.T) {
	var err error
	from := GetDataSourceDbExcel()

	to := GetDataSourceDb2()

	p := NewProgress()
	err = DateMove(p, from, to)
	if err != nil {
		panic(err)
	}
	fmt.Println(util.GetStringValue(p))
}

func TestDbToEs(t *testing.T) {
	var err error
	from := GetDataSourceDb()

	to := GetDataSourceEs()

	p := NewProgress()
	err = DateMove(p, from, to)
	if err != nil {
		panic(err)
	}
	fmt.Println(util.GetStringValue(p))
}

func TestEsToEs2(t *testing.T) {
	var err error
	from := GetDataSourceEs()

	to := GetDataSourceEs2()

	p := NewProgress()
	err = DateMove(p, from, to)
	if err != nil {
		panic(err)
	}
	fmt.Println(util.GetStringValue(p))
}
