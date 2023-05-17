package context_map

type ModuleInfo struct {
	Name     string       `json:"name"`
	Comment  string       `json:"comment"`
	FuncList []*FuncInfo  `json:"funcList"`
	Service  *ServiceInfo `json:"service"`
}
type ServiceInfo struct {
	Name     string      `json:"name"`
	Comment  string      `json:"comment"`
	Module   string      `json:"module"`
	FuncList []*FuncInfo `json:"funcList"`
}
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
	ModuleList []*ModuleInfo
)

func AddModule(arg *ModuleInfo) {
	ModuleList = append(ModuleList, arg)
}
