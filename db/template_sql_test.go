package db

import (
	"fmt"
	"github.com/team-ide/go-dialect/dialect"
	"github.com/team-ide/go-tool/util"
	"testing"
)

type TestUser struct {
	UserId int64  `json:"userId" column:"user_id"`
	AA     string `json:"a" column:"Name"`
}

func TestSql(t *testing.T) {
	opts := &TemplateOptions{}
	opts.Dialect, _ = dialect.NewDialect("kingbase")
	opts.StringEmptyUseNull = true
	opts.NumberZeroUseNull = true
	var u *TestUser
	temp := WarpTemplate(u, opts)

	insertSqlA, insertSqlAArgs := temp.GetInsertSql("tb_user", &TestUser{
		UserId: 11, AA: "xx",
	})
	fmt.Println("a sql:", insertSqlA)
	fmt.Println("a args:", util.GetStringValue(insertSqlAArgs))

	insertSqlA, insertSqlAArgs = temp.GetInsertSql("tb_user", map[string]any{
		"user_id": 1,
		"name":    "张三",
	})
	fmt.Println("a sql:", insertSqlA)
	fmt.Println("a args:", util.GetStringValue(insertSqlAArgs))
	insertSqlB, insertSqlBArgs := temp.GetBatchInsertSql("tb_user", []*TestUser{
		{UserId: 11, AA: "xx"},
		{UserId: 22, AA: ""},
		{UserId: 0, AA: "xx"},
	})
	fmt.Println("a sql:", insertSqlB)
	fmt.Println("a args:", util.GetStringValue(insertSqlBArgs))

	var up = &TestUser{
		UserId: 11, AA: "xx",
	}
	updateSqlA, updateSqlAArgs, err := temp.GetUpdateSql("tb_user", up, temp.SqlParamParser("user_id=${userId}", up), "user_id")
	if err != nil {
		panic(err)
	}
	fmt.Println("a sql:", updateSqlA)
	fmt.Println("a args:", util.GetStringValue(updateSqlAArgs))
}
