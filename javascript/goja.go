package javascript

import "github.com/team-ide/goja"

func Run(str string, context map[string]interface{}) (res interface{}, err error) {
	vm := goja.New()

	if len(context) > 0 {
		for key, value := range context {
			err = vm.Set(key, value)
			if err != nil {
				return
			}
		}
	}
	v, err := vm.RunString(str)
	if err != nil {
		return
	}
	res = v.Export()
	return
}

func RunScript(str string, context map[string]interface{}) (res interface{}, err error) {
	vm := goja.New()

	if len(context) > 0 {
		for key, value := range context {
			err = vm.Set(key, value)
			if err != nil {
				return
			}
		}
	}
	script := `(function (){
` + str + `
})()
`
	v, err := vm.RunScript("", script)
	if err != nil {
		return
	}
	res = v.Export()
	return
}
