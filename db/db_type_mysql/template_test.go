package db_type_mysql

import (
	"context"
	"fmt"
	"github.com/team-ide/go-tool/db"
	"github.com/team-ide/go-tool/util"
	"reflect"
	"testing"
)

type TestUser struct {
	UserId   int64  `json:"userId" column:"user_id"`
	Name     string `json:"name" column:"name"`
	Account  string `json:"account" column:"account"`
	Password string `json:"password" column:"password"`
	CreateAt int64  `json:"createAt" column:"create_at"`
	DeleteAt int64  `json:"deleteAt" column:"delete_at"`
	Status   int    `json:"status" column:"status"`
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

	ser2, err := db.New(&db.Config{
		Type:     "mysql",
		Host:     "127.0.0.1",
		Port:     3306,
		Username: "root",
		Password: "123456",
		Database: "test_2",
	})
	if err != nil {
		panic(err)
	}
	opts := &db.TemplateOptions{
		Service: ser,
	}
	opts2 := &db.TemplateOptions{
		Service: ser2,
	}

	var b TestUser
	template := db.WarpTemplate[*TestUser](opts)
	template2 := db.WarpTemplate[*TestUser](opts2)
	ctx, err := db.OpenTxContext(context.Background())
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = db.EndTxContext(ctx, err != nil)
	}()
	_, err = template.Insert(ctx, "tb_user", &TestUser{
		UserId: util.NextId(), Name: "x", Account: "xx", Password: "xxx", CreateAt: util.GetNowMilli(),
	})
	if err != nil {
		panic(err)
	}
	_, err = template.Insert(ctx, "tb_user", &TestUser{
		UserId: util.NextId(), Name: "x", Account: "xx", Password: "xxx", CreateAt: util.GetNowMilli(),
	})
	if err != nil {
		panic(err)
	}
	_, err = template2.Insert(ctx, "tb_user", &TestUser{
		UserId: util.NextId(), Name: "x", Account: "xx", Password: "xxx", CreateAt: util.GetNowMilli(),
	})
	if err != nil {
		panic(err)
	}
	_, err = template2.Insert(ctx, "tb_user", &TestUser{
		UserId: util.NextId(), Name: "x", Account: "xx", Password: "xxx", CreateAt: util.GetNowMilli(),
	})
	if err != nil {
		panic(err)
	}

	query(opts, b)
	query(opts, &b)

	var m map[string]interface{}

	query(opts, m)
	query(opts, &m)
}
func query[T any](opts *db.TemplateOptions, obj T) {
	fmt.Println("query by:", reflect.TypeOf(obj))
	template := db.WarpTemplate(obj, opts)
	findOne, err := template.SelectOne(context.Background(), "select * from tb_user where user_id=1", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("findOne:", reflect.TypeOf(findOne))
	fmt.Println("findOne:", util.GetStringValue(findOne))
	findLise, err := template.SelectList(context.Background(), "select * from tb_user", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("findLise:", reflect.TypeOf(findLise))
	fmt.Println("findLise:", util.GetStringValue(findLise))
	findPage, err := template.SelectPageBean(context.Background(), "select * from tb_user", nil, 10, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println("findPage:", reflect.TypeOf(findPage))
	fmt.Println("findPage:", util.GetStringValue(findPage))
}
