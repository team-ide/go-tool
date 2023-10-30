package db

import (
	"database/sql"
	"errors"
	"github.com/team-ide/go-dialect/dialect"
	"github.com/team-ide/go-tool/javascript"
	"github.com/team-ide/go-tool/task"
	"github.com/team-ide/go-tool/util"
	"github.com/team-ide/goja"
	"go.uber.org/zap"
	"regexp"
	"strings"
	"sync"
)

type TestTaskOptions struct {
	*Param
	OwnerName       string                 `json:"ownerName"`
	TableName       string                 `json:"tableName"`
	ColumnList      []*dialect.ColumnModel `json:"columnList"`
	InsertData      map[string]interface{} `json:"insertData"`
	UpdateData      map[string]interface{} `json:"updateData"`
	UpdateWhereData map[string]interface{} `json:"updateWhereData"`
	DeleteData      map[string]interface{} `json:"deleteData"`

	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	IsBatch   bool   `json:"isBatch,omitempty"`
	BatchSize int    `json:"batchSize,omitempty"`
	TestType  string `json:"testType,omitempty"`
}

func (this_ *Service) NewTestTask(task *task.Task, options *TestTaskOptions) (testTask *TestTask, err error) {
	testTask = &TestTask{
		Task:            task,
		TestTaskOptions: options,
	}
	var config = *this_.config
	testTask.workDb, err = newWorkDb(this_.databaseType, config, options.Username, options.Password, options.OwnerName)
	if err != nil {
		util.Logger.Error("NewTestTask new db pool error", zap.Error(err))
		return
	}
	testTask.dia = this_.GetTargetDialect(options.Param)
	task.OnStop = func() {
		_ = testTask.workDb.Close()
	}
	return
}

type TestTask struct {
	*task.Task
	*TestTaskOptions
	workDb *sql.DB
	dia    dialect.Dialect
}

type TestExecutor struct {
	task            *TestTask
	workerParam     map[int]*TestWorkerParam
	workerParamLock sync.Mutex
}

type TestWorkerParam struct {
	*TestExecutor
	sqlList       []string
	sqlParamsList [][]interface{}
	lock          sync.Mutex

	runtime       *goja.Runtime
	scriptContext map[string]interface{}
}

func (this_ *TestExecutor) getWorkerParam(workerIndex int) (res *TestWorkerParam, err error) {
	this_.workerParamLock.Lock()
	defer this_.workerParamLock.Unlock()

	res = this_.workerParam[workerIndex]
	if res == nil {
		res = &TestWorkerParam{
			TestExecutor: this_,
		}
		res.runtime = goja.New()
		res.scriptContext = javascript.NewContext()
		if len(res.scriptContext) > 0 {
			for key, value := range res.scriptContext {
				err = res.runtime.Set(key, value)
				if err != nil {
					return
				}
			}
		}
		err = res.runtime.Set("workerIndex", workerIndex)
		if err != nil {
			return
		}
		err = res.runtime.Set("ownerName", this_.task.OwnerName)
		if err != nil {
			return
		}
		err = res.runtime.Set("tableName", this_.task.TableName)
		if err != nil {
			return
		}
		this_.workerParam[workerIndex] = res
	}
	return
}

func (this_ *TestWorkerParam) getScriptValue(param *task.ExecutorParam, dataIndex int, script string) (res string, err error) {

	err = this_.runtime.Set("index", dataIndex)
	if err != nil {
		return
	}

	v, err := this_.runtime.RunString(script)
	if err != nil {
		err = errors.New("get script [" + script + "] value error:" + err.Error())
		return
	}
	res = v.String()
	return
}

func (this_ *TestWorkerParam) GetStringArg(param *task.ExecutorParam, dataIndex int, arg string) (res interface{}, err error) {
	if arg == "" {
		res = ""
		return
	}
	text := ""
	var re *regexp.Regexp
	re, _ = regexp.Compile(`[$]+{(.+?)}`)
	indexList := re.FindAllIndex([]byte(arg), -1)
	var lastIndex int = 0
	for _, indexes := range indexList {
		text += arg[lastIndex:indexes[0]]

		lastIndex = indexes[1]

		script := arg[indexes[0]+2 : indexes[1]-1]
		v := ""
		v, err = this_.getScriptValue(param, dataIndex, script)
		if err != nil {
			return
		}
		text += v
	}
	text += arg[lastIndex:]

	res = text
	return
}
func (this_ *TestWorkerParam) getData(param *task.ExecutorParam, dataIndex int, templateData map[string]interface{}) (data map[string]interface{}, err error) {
	if templateData == nil {
		return
	}
	data = map[string]interface{}{}
	for k, v := range templateData {
		putV := v
		if v != nil {
			switch tV := v.(type) {
			case string:
				putV, err = this_.GetStringArg(param, dataIndex, tV)
				if err != nil {
					return
				}
				break
			}
		}
		data[k] = putV
	}

	return
}

func (this_ *TestWorkerParam) appendSql(param *task.ExecutorParam, dataIndex int) (err error) {

	this_.lock.Lock()
	defer this_.lock.Unlock()

	var sqlList []string
	var valuesList [][]interface{}
	switch strings.ToLower(this_.task.TestType) {
	case "insert":
		var insertData map[string]interface{}
		insertData, err = this_.getData(param, dataIndex, this_.task.InsertData)
		if err != nil || insertData == nil {
			return
		}
		sqlList, valuesList, _, _, err = this_.task.dia.DataListInsertSql(
			this_.task.Param.ParamModel, this_.task.OwnerName, this_.task.TableName, this_.task.ColumnList,
			[]map[string]interface{}{insertData},
		)
		if err != nil {
			return
		}
		break
	case "update":
		var updateData map[string]interface{}
		updateData, err = this_.getData(param, dataIndex, this_.task.UpdateData)
		if err != nil || updateData == nil {
			return
		}
		var updateWhereData map[string]interface{}
		updateWhereData, err = this_.getData(param, dataIndex, this_.task.UpdateWhereData)
		if err != nil || updateWhereData == nil {
			return
		}
		sqlList, valuesList, err = this_.task.dia.DataListUpdateSql(
			this_.task.Param.ParamModel, this_.task.OwnerName, this_.task.TableName, this_.task.ColumnList,
			[]map[string]interface{}{updateData}, []map[string]interface{}{updateWhereData},
		)
		if err != nil {
			return
		}
		break
	case "delete":
		var deleteData map[string]interface{}
		deleteData, err = this_.getData(param, dataIndex, this_.task.DeleteData)
		if err != nil || deleteData == nil {
			return
		}
		sqlList, valuesList, err = this_.task.dia.DataListDeleteSql(
			this_.task.Param.ParamModel, this_.task.OwnerName, this_.task.TableName, this_.task.ColumnList,
			[]map[string]interface{}{deleteData},
		)
		if err != nil {
			return
		}
		break
	}

	this_.sqlList = append(this_.sqlList, sqlList...)
	this_.sqlParamsList = append(this_.sqlParamsList, valuesList...)

	return
}

func (this_ *TestExecutor) initParam(param *task.ExecutorParam) (err error) {
	workerParam, err := this_.getWorkerParam(param.WorkerIndex)
	if err != nil {
		return
	}
	workerParam.sqlParamsList = [][]interface{}{}
	workerParam.sqlList = []string{}

	param.Extend = workerParam

	var genSize = 1
	if this_.task.IsBatch {
		genSize = this_.task.BatchSize
	}
	if genSize <= 0 {
		return
	}
	for i := 0; i < genSize; i++ {
		dataIndex := this_.task.GetNextIndex()
		if dataIndex < 0 {
			break
		}
		err = workerParam.appendSql(param, dataIndex)
		if err != nil {
			return
		}
	}
	return
}

func (this_ *TestExecutor) Before(param *task.ExecutorParam) (err error) {

	err = this_.initParam(param)

	return
}

func (this_ *TestExecutor) Execute(param *task.ExecutorParam) (err error) {

	workerParam := param.Extend.(*TestWorkerParam)

	var sqlSize = len(workerParam.sqlList)
	var sqlParamsSize = len(workerParam.sqlParamsList)
	if sqlSize > 0 {
		if sqlSize != sqlParamsSize {
			err = errors.New("sql size not equal to sql params size")
			return
		}
		for i := 0; i < sqlSize; i++ {
			_, err = this_.task.workDb.Exec(workerParam.sqlList[i], workerParam.sqlParamsList[i]...)
		}
	}

	return
}

func (this_ *TestExecutor) After(param *task.ExecutorParam) (err error) {

	return
}
