package db

import (
	"fmt"
	"github.com/team-ide/go-dialect/dialect"
	"github.com/team-ide/go-tool/util"
	"testing"
)

func TestSqlParamParse(t *testing.T) {
	opts := &TemplateOptions{}
	opts.Dialect, _ = dialect.NewDialect("kingbase")
	opts.StringEmptyUseNull = true
	opts.NumberZeroUseNull = true

	//parser := opts.SqlArgParser("select * from aa where a=${aa} and b=${bb}", &TestUser{})

	parser := opts.SqlParamParser("select * from aa where a=${aa} and b=${cc.Bb}", map[string]any{
		"aa": 1,
		"bb": 2,
		"cc": map[string]any{
			"aa": 11,
			"bb": 22,
		},
	})
	s, args, err := parser.Parse()
	if err != nil {
		panic(err)
	}
	fmt.Println("a sql:", s)
	fmt.Println("a args:", util.GetStringValue(args))
}
