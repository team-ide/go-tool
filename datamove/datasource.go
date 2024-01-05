package datamove

import "github.com/team-ide/go-dialect/dialect"

type DataSource interface {
	Stop()
	ReadStart() (err error)
	Read() (data *Data, err error)
	ReadEnd() (err error)
	WriteStart() (err error)
	Write(data *Data) (err error)
	WriteEnd() (err error)
}

type Data struct {
	DataType   DataType
	SqlList    []string
	DataList   []map[string]interface{}
	ColumnList []*dialect.ColumnModel
	HasNext    bool
}

type DataType int8

const (
	DataTypeEmpty = 0
	DataTypeSql   = 1
	DataTypeMap   = 2
)

type Param struct {
}
