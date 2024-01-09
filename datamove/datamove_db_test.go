package datamove

import (
	"fmt"
	"github.com/team-ide/go-dialect/dialect"
	"github.com/team-ide/go-tool/db"
	_ "github.com/team-ide/go-tool/db/db_type_sqlite"
	"github.com/team-ide/go-tool/util"
	"os"
	"testing"
)

func newTestDataMoveDbToOptions() *Options {
	options := &Options{
		Source: &DataSourceConfig{},
		Target: &DataSourceConfig{},
	}
	options.Key = util.GetUUID()
	options.Source.Type = "database"
	options.Source.DbConfig = &db.Config{
		Type:         "sqlite",
		DatabasePath: "out/db",
	}
	options.AllOwner = true
	options.BatchNumber = 1000

	return options
}

func TestSqlValue(t *testing.T) {
	true_ := new(bool)
	*true_ = true
	sqlInfo := `insert into a xxx\' `
	s, _ := dialect.NewDialect("mysql")
	res := s.SqlValuePack(&dialect.ParamModel{
		AppendSqlValue: true_,
	}, nil, sqlInfo)
	fmt.Println("sqlInfo:", sqlInfo)
	fmt.Println("res:", res)
}

func TestDataMoveDbToTxt(t *testing.T) {
	options := newTestDataMoveDbToOptions()
	options.Target.Type = "txt"
	options.Dir = "out/txt/"
	options.FileNameSplice = "-"
	_ = os.MkdirAll(options.Dir, os.ModePerm)

	task, err := New(options)
	if err != nil {
		panic(err)
	}
	task.Run()
	fmt.Println(util.GetStringValue(task))
}

func TestDataMoveDbToSql(t *testing.T) {
	options := newTestDataMoveDbToOptions()
	options.Target.Type = "sql"
	options.Dir = "out/sql/"
	_ = os.MkdirAll(options.Dir, os.ModePerm)
	options.SqlFileMergeType = "one"
	options.ShouldOwner = true
	options.ShouldTable = true
	task, err := New(options)
	if err != nil {
		panic(err)
	}
	task.Run()
	fmt.Println(util.GetStringValue(task))
}

func TestDataMoveDbToExcel(t *testing.T) {
	options := newTestDataMoveDbToOptions()
	options.Target.Type = "excel"
	options.Dir = "out/excel/"
	options.FileNameSplice = "-"
	_ = os.MkdirAll(options.Dir, os.ModePerm)

	task, err := New(options)
	if err != nil {
		panic(err)
	}
	task.Run()
	fmt.Println(util.GetStringValue(task))
}
