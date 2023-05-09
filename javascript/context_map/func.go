package context_map

type FuncInfo struct {
	Name     string         `json:"name"`
	Comment  string         `json:"comment"`
	Params   []*FuncVarInfo `json:"params"`
	Return   *FuncVarInfo   `json:"return"`
	HasError bool           `json:"hasError"`
	Func     interface{}    `json:"-"`
}

type FuncVarInfo struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
	Type    string `json:"type"`
}

var (
	FuncList []*FuncInfo
)

func AddFunc(funcInfo *FuncInfo) {
	FuncList = append(FuncList, funcInfo)
}
