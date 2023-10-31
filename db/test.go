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
	"sync"
)

type TestTaskOptions struct {
	*Param
	OwnerName string `json:"ownerName"`
	TableName string `json:"tableName"`

	Username      string                                                `json:"username,omitempty"`
	Password      string                                                `json:"password,omitempty"`
	IsBatch       bool                                                  `json:"isBatch,omitempty"`
	BatchSize     int                                                   `json:"batchSize,omitempty"`
	TestSql       string                                                `json:"testSql,omitempty"`
	GetNextIndex  func() (nextIndex int)                                `json:"-"`
	FormatSqlList func(sqlList *[]string, sqlArgsList *[][]interface{}) `json:"-"`
	OnExec        func(sqlList *[]string, sqlArgsList *[][]interface{}) `json:"-"`
}

func (this_ *Service) NewTestExecutor(options *TestTaskOptions) (testExecutor *TestExecutor, err error) {
	testExecutor = &TestExecutor{
		TestTaskOptions: options,
		workerParam:     make(map[int]*TestWorkerParam),
	}
	var config = *this_.config
	testExecutor.workDb, err = newWorkDb(this_.databaseType, config, options.Username, options.Password, options.OwnerName)
	if err != nil {
		util.Logger.Error("NewTestTask new db pool error", zap.Error(err))
		return
	}
	testExecutor.dia = this_.GetTargetDialect(options.Param)
	return
}

type TestTask struct {
}

type TestExecutor struct {
	task            *TestTask
	workerParam     map[int]*TestWorkerParam
	workerParamLock sync.Mutex
	*TestTaskOptions
	workDb *sql.DB
	dia    dialect.Dialect
}

type TestWorkerParam struct {
	*TestExecutor
	sqlList       []string
	sqlParamsList [][]interface{}
	lock          sync.Mutex

	runtime       *goja.Runtime
	scriptContext map[string]interface{}
}

func (this_ *TestExecutor) Close() {
	if this_.workDb != nil {
		_ = this_.workDb.Close()
	}
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
		err = res.runtime.Set("ownerName", this_.OwnerName)
		if err != nil {
			return
		}
		err = res.runtime.Set("tableName", this_.TableName)
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

func (this_ *TestWorkerParam) GetStringArg(param *task.ExecutorParam, dataIndex int, arg string) (res string, err error) {
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

func (this_ *TestWorkerParam) appendSql(param *task.ExecutorParam, dataIndex int) (err error) {

	this_.lock.Lock()
	defer this_.lock.Unlock()

	var sqlList []string
	var valuesList [][]interface{}
	testSql, err := this_.GetStringArg(param, dataIndex, this_.TestSql)
	if err != nil {
		return
	}
	sqlList = append(sqlList, testSql)
	if this_.FormatSqlList != nil {
		this_.FormatSqlList(&sqlList, &valuesList)
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
	if this_.IsBatch {
		genSize = this_.BatchSize
	}
	if genSize <= 0 {
		return
	}
	err = workerParam.appendSql(param, param.Index)
	if err != nil {
		return
	}
	for i := 1; i < genSize; i++ {
		dataIndex := this_.GetNextIndex()
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

	if this_.OnExec != nil {
		this_.OnExec(&workerParam.sqlList, &workerParam.sqlParamsList)
	}
	var sqlSize = len(workerParam.sqlList)
	var sqlParamsSize = len(workerParam.sqlParamsList)
	if sqlSize > 0 {
		if sqlSize != sqlParamsSize {
			err = errors.New("sql size not equal to sql params size")
			return
		}
		for i := 0; i < sqlSize; i++ {
			_, err = this_.workDb.Exec(workerParam.sqlList[i], workerParam.sqlParamsList[i]...)
		}
	}

	return
}

func (this_ *TestExecutor) After(param *task.ExecutorParam) (err error) {

	return
}
