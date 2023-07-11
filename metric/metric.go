package metric

import (
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"sort"
	"sync"
	"time"
)

func NewMetric() (res *Metric) {
	res = &Metric{
		locker:              &sync.Mutex{},
		workerMetricsLocker: &sync.Mutex{},
		countLocker:         &sync.Mutex{},
		secondCountCache:    map[int]*Count{},
	}
	return
}

type Metric struct {
	workerMetrics       []*WorkerMetric
	workerMetricsLocker sync.Locker
	locker              sync.Locker
	countLocker         sync.Locker
	count               *Count
	secondCountCache    map[int]*Count
	secondCounts        []*Count
	countStart          bool
	countSecond         int // 统计间隔秒 如 每秒统计 输入 1 默认 10 秒统计
	onCount             func()
	countTop            bool
}

// SetCountTop  是否统计 T99 T90 T80等，高并发下 消耗内存将增加
func (this_ *Metric) SetCountTop(countTop bool) *Metric {
	this_.countTop = countTop

	return this_
}

// SetCountSecond  统计间隔秒 如 每秒统计 输入 1 默认 10 秒统计
func (this_ *Metric) SetCountSecond(countSecond int) *Metric {
	this_.countSecond = countSecond

	return this_
}

// SetOnCount  每次统计完成 调用该方法
func (this_ *Metric) SetOnCount(onCount func()) *Metric {
	this_.onCount = onCount

	return this_
}

func (this_ *Metric) NewWorkerMetric(workerIndex int) (workerMetric *WorkerMetric) {
	workerMetric = &WorkerMetric{
		metric:            this_,
		WorkerIndex:       workerIndex,
		secondItemsLocker: &sync.Mutex{},
		secondCounts:      map[int]*Count{},
		secondItems:       &[]*SecondItem{},
	}
	this_.workerMetricsLocker.Lock()
	defer this_.workerMetricsLocker.Unlock()

	this_.workerMetrics = append(this_.workerMetrics, workerMetric)
	return
}

func (this_ *Metric) GewWorkerMetrics() (workerMetrics []*WorkerMetric) {

	this_.workerMetricsLocker.Lock()
	defer this_.workerMetricsLocker.Unlock()

	workerMetrics = this_.workerMetrics
	return
}

func (this_ *Metric) StopCount() {
	this_.countStart = false

	time.Sleep(time.Millisecond * 100)
	this_.DoCount()
	return
}

func (this_ *Metric) StartCount() {
	if this_.countStart {
		return
	}
	this_.countStart = true

	go func() {
		for this_.countStart {
			time.Sleep(time.Second * 1)
			if !this_.countStart {
				break
			}
			// 1 秒进行一次统计
			this_.DoCount()
		}
	}()
	return
}

func (this_ *Metric) DoCount() {
	this_.countLocker.Lock()
	defer this_.countLocker.Unlock()

	defer func() {
		if e := recover(); e != nil {
			util.Logger.Error("doCount recover error", zap.Any("err", e))
		}
	}()
	workerMetrics := this_.GewWorkerMetrics()
	var nowSecond = int(time.Now().UnixNano() / Second)

	countSecond := this_.countSecond
	if countSecond <= 0 {
		countSecond = 10
	}
	var allSecondCount = map[int][]*Count{}
	var workerCounts []*Count

	for _, workerMetric := range workerMetrics {
		workerCount, secondCount := workerMetric.countByDuration(nowSecond, countSecond)

		for t, c := range secondCount {
			allSecondCount[t] = append(allSecondCount[t], c)
		}
		workerCounts = append(workerCounts, workerCount)
	}

	this_.count = WorkersCount(workerCounts, this_.countTop)

	if this_.count != nil {
		*this_.count.useTimes = []int{}
	}

	var ts []int
	for t := range allSecondCount {
		ts = append(ts, t)
	}

	var secondCounts []*Count
	sort.Ints(ts)
	for _, t := range ts {
		count := WorkersCount(allSecondCount[t], this_.countTop)
		secondCounts = append(secondCounts, count)
	}
	this_.secondCounts = secondCounts

	if this_.onCount != nil {
		this_.onCount()
	}
	return
}

func (this_ *Metric) GetCount() (count *Count) {

	count = this_.count
	if count == nil {
		count = &Count{}
	}
	return
}

func (this_ *Metric) GetSecondCounts() (counts []*Count) {

	counts = this_.secondCounts
	return
}
