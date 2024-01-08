package datamove

import (
	"fmt"
	"github.com/team-ide/go-tool/db"
	_ "github.com/team-ide/go-tool/db/db_type_sqlite"
	"github.com/team-ide/go-tool/util"
	"os"
	"testing"
)

func TestDataMoveDbToTxt(t *testing.T) {
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

	options.Target.Type = "txt"

	options.AllOwner = true
	options.Dir = "out/txt/"
	options.BatchNumber = 1000

	_ = os.MkdirAll(options.Dir, os.ModePerm)

	task, err := New(options)
	if err != nil {
		panic(err)
	}
	task.Run()
	fmt.Println(util.GetStringValue(task))
}

func TestDataMoveDbToSql(t *testing.T) {
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

	options.Target.Type = "sql"

	options.AllOwner = true

	task, err := New(options)
	if err != nil {
		panic(err)
	}
	task.Run()
	fmt.Println(util.GetStringValue(task))
}
