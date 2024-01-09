package datamove

import (
	"errors"
	"fmt"
	"github.com/team-ide/go-tool/task"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"strings"
)

func New(options *Options) (t *task.Task, err error) {
	progress := NewProgress(options)

	if options.Dir == "" {
		err = errors.New(fmt.Sprintf("存储目录为空"))
		return
	}
	exists, err := util.PathExists(options.Dir)
	if err != nil {
		util.Logger.Error("检测存储目录异常", zap.Any("dir", options.Dir), zap.Error(err))
		return
	}
	if !exists {
		err = errors.New(fmt.Sprintf("目录[%s]未创建", options.Dir))
		util.Logger.Error(fmt.Sprintf("目录[%s]未创建", options.Dir))
		return
	}
	options.Dir = util.FormatPath(options.Dir)
	if !strings.HasSuffix(options.Dir, "/") {
		options.Dir += "/"
	}

	if progress.BatchNumber <= 0 {
		progress.BatchNumber = 1
	}

	executor := &Executor{
		Progress: progress,
	}
	t, err = task.New(&task.Options{
		Key:       progress.Key,
		Frequency: 1,
		Worker:    1,
		Executor:  executor,
	})
	if err != nil {
		return
	}
	return

}
func NewProgress(options *Options) *Progress {
	res := &Progress{
		Options: options,
	}
	res.OwnerCount = &ProgressCount{}
	res.TableCount = &ProgressCount{}
	res.ReadCount = &ProgressCount{}
	res.WriteCount = &ProgressCount{}
	res.IndexCount = &ProgressCount{}
	res.TopicCount = &ProgressCount{}
	return res
}

type Progress struct {
	*Options

	DataTotal int64 `json:"dataTotal"`

	OwnerTotal int64          `json:"ownerTotal"`
	OwnerCount *ProgressCount `json:"ownerCount"`

	TableTotal int64          `json:"tableTotal"`
	TableCount *ProgressCount `json:"tableCount"`

	ReadCount  *ProgressCount `json:"readCount"`
	WriteCount *ProgressCount `json:"writeCount"`

	IndexTotal int64          `json:"indexTotal"`
	IndexCount *ProgressCount `json:"indexCount"`

	TopicTotal int64          `json:"topicTotal"`
	TopicCount *ProgressCount `json:"topicCount"`

	isEnd        bool
	isStopped    bool
	dataMoveStop *bool
}

func (this_ *Progress) ShouldStop() bool {
	if this_.dataMoveStop != nil {
		if *this_.dataMoveStop {
			return true
		}
	}
	return this_.isEnd || this_.isStopped
}

type ProgressCount struct {
	Total   int64    `json:"total"`
	Error   int64    `json:"error"`
	Success int64    `json:"success"`
	Errors  []string `json:"errors"`
}

func (this_ *ProgressCount) AddSuccess(size int64) {
	this_.Total += size
	this_.Success += size
}

func (this_ *ProgressCount) AddError(size int64, err error) {
	this_.Total += size
	this_.Error += size
	if err != nil {
		util.Logger.Error("progress error", zap.Error(err))
		this_.Errors = append(this_.Errors, err.Error())
	}
}

type Executor struct {
	*Progress
}

func (this_ *Executor) Before(param *task.ExecutorParam) (err error) {
	return
}

func (this_ *Executor) Execute(param *task.ExecutorParam) (err error) {
	err = this_.execute()
	return
}

func (this_ *Executor) After(param *task.ExecutorParam) (err error) {
	return
}
