package context_map

type ServiceInfo struct {
	Name     string      `json:"name"`
	Comment  string      `json:"comment"`
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
	FuncList    []*FuncInfo
	ServiceList []*ServiceInfo
)

func AddFunc(funcInfo *FuncInfo) {
	FuncList = append(FuncList, funcInfo)
}

func AddService(serviceInfo *ServiceInfo) {
	ServiceList = append(ServiceList, serviceInfo)
}
