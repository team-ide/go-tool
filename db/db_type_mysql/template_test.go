package db_type_mysql

import (
	"fmt"
	"github.com/team-ide/go-dialect/worker"
	"github.com/team-ide/go-tool/db"
	"github.com/team-ide/go-tool/util"
	"reflect"
	"testing"
)

type TestUser struct {
	UserId int64  `json:"userId" column:"user_id"`
	AA     string `json:"a" column:"Name"`
}

func TestSelect(t *testing.T) {
	ser, err := db.New(&db.Config{
		Type:     "mysql",
		Host:     "127.0.0.1",
		Port:     3306,
		Username: "root",
		Password: "123456",
		Database: "test",
	})
	if err != nil {
		panic(err)
	}
	opts := &db.TemplateOptions{
		Service: ser,
	}

	var b TestUser

	query(opts, b)
	query(opts, &b)
}
func query[T any](opts *db.TemplateOptions, obj T) {
	fmt.Println("query by:", reflect.TypeOf(obj))
	template := db.WarpTemplate(obj, opts)
	findOne, err := template.Query("select * from tb_user", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("findOne:", reflect.TypeOf(findOne))
	fmt.Println("findOne:", util.GetStringValue(findOne))
	findLise, err := template.Query("select * from tb_user", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("findLise:", reflect.TypeOf(findLise))
	fmt.Println("findLise:", util.GetStringValue(findLise))
	findPage, err := template.QueryPage("select * from tb_user ", nil, &worker.Page{})
	if err != nil {
		panic(err)
	}
	fmt.Println("findPage:", reflect.TypeOf(findPage))
	fmt.Println("findPage:", util.GetStringValue(findPage))
}
