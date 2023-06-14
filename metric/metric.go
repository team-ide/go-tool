package metric

import (
	"sort"
	"strconv"
	"sync"
	"time"
)

func NewMetric() (res *Metric) {
	res = &Metric{
		locker:       &sync.Mutex{},
		countLocker:  &sync.Mutex{},
		items:        &[]*Item{},
		secondCounts: map[int]*Count{},
		minuteCounts: map[int]*Count{},
	}
	return
}

type Metric struct {
	locker       sync.Locker
	items        *[]*Item
	countLocker  sync.Locker
	count        *Count
	secondCounts map[int]*Count
	minuteCounts map[int]*Count
	countStart   bool
}

func (this_ *Metric) NewItem(workerIndex int, startTime time.Time) (item *Item) {
	item = &Item{
		metric:      this_,
		WorkerIndex: workerIndex,
		StartTime:   startTime.UnixNano(),
	}
	return
}

func (this_ *Metric) GetAndCleanItems() (items *[]*Item) {

	this_.locker.Lock()
	defer this_.locker.Unlock()

	items = this_.items
	this_.items = &[]*Item{}
	return
}

func (this_ *Metric) AddItems(items ...*Item) {

	this_.locker.Lock()
	defer this_.locker.Unlock()

	for _, item := range items {
		*this_.items = append(*this_.items, item)
	}
	return
}

func (this_ *Metric) StopCount() {
	this_.countStart = false

	time.Sleep(time.Second * 1)
	this_.doCount()
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
			// 1 秒进行一次统计
			this_.doCount()
		}
	}()
	return
}

func (this_ *Metric) doCount() {
	items := this_.GetAndCleanItems()

	this_.doCountByDuration(items, int64(time.Second), this_.secondCounts)
	this_.doCountByDuration(items, int64(time.Minute), this_.minuteCounts)

	toCounts := this_.getCountsByMap(this_.minuteCounts)
	if len(toCounts) > 0 {
		this_.count = this_.doCountByCounts(toCounts)
	}
	return
}

func (this_ *Metric) GetCount() (count *Count) {

	count = this_.count
	return
}

func (this_ *Metric) GetSecondCounts() (counts []*Count) {

	counts = this_.getCountsByMap(this_.secondCounts)
	return
}

func (this_ *Metric) GetMinuteCounts() (counts []*Count) {

	counts = this_.getCountsByMap(this_.minuteCounts)
	return
}

func (this_ *Metric) getCountsByMap(countMap map[int]*Count) (counts []*Count) {

	this_.countLocker.Lock()
	defer this_.countLocker.Unlock()

	var ts []int
	for t := range countMap {
		ts = append(ts, t)
	}
	sort.Ints(ts)
	for _, t := range ts {
		//fmt.Println("getCountsByMap t:", t, " count:", countMap[t])
		counts = append(counts, countMap[t])
	}
	return
}

func (this_ *Metric) doCountByDuration(items *[]*Item, duration int64, historyCounts map[int]*Count) {

	this_.countLocker.Lock()
	defer this_.countLocker.Unlock()

	cacheItems := map[int]*[]*Item{}

	for _, item := range *items {
		t := int(item.StartTime / duration)
		if cacheItems[t] == nil {
			cacheItems[t] = &[]*Item{}
		}
		*cacheItems[t] = append(*cacheItems[t], item)
	}
	for t, tItems := range cacheItems {
		count := CountItems(tItems)
		if historyCounts[t] != nil {
			count = this_.doCountByCounts([]*Count{count, historyCounts[t]})
		}
		historyCounts[t] = count
	}

	return
}

type Item struct {
	metric      *Metric
	WorkerIndex int         `json:"workerIndex"` // 工作线程索引  用于并发线程下 每个线程的时间跨度计算
	StartTime   int64       `json:"startTime"`
	EndTime     int64       `json:"endTime"`
	Success     bool        `json:"success"`
	LossTime    int         `json:"lossTime"` // 损耗时长 纳秒 统计时候将去除该部分时长
	Extend      interface{} `json:"extend"`
	UseTime     int         `json:"useTime"`
}

func (this_ *Item) End(endTime time.Time, err error) {
	this_.Success = err == nil
	this_.EndTime = endTime.UnixNano()
	this_.metric.AddItems(this_)
}

func (this_ *Metric) doCountByCounts(toCounts []*Count) (count *Count) {
	count = CountCounts(toCounts)
	return
}

func (this_ *Item) Loss(lossTime int) {
	this_.LossTime = lossTime
}

func CountItems(itemList *[]*Item) (count *Count) {
	count = &Count{}
	useTimes := &[]int{}

	count.WorkerLossTime = map[int]int64{}
	count.WorkerStartTime = map[int]int64{}
	count.WorkerEndTime = map[int]int64{}
	for _, item := range *itemList {
		startNano := item.StartTime
		endNano := item.EndTime

		if count.StartTime == 0 || startNano < count.StartTime {
			count.StartTime = item.StartTime
		}
		if count.EndTime == 0 || endNano > count.EndTime {
			count.EndTime = item.EndTime
		}

		if count.WorkerStartTime[item.WorkerIndex] == 0 || startNano < count.WorkerStartTime[item.WorkerIndex] {
			count.WorkerStartTime[item.WorkerIndex] = item.StartTime
		}
		if count.WorkerEndTime[item.WorkerIndex] == 0 || endNano > count.WorkerEndTime[item.WorkerIndex] {
			count.WorkerEndTime[item.WorkerIndex] = item.EndTime
		}

		item.UseTime = int(endNano - startNano)

		// 使用时间 需要 减去 损耗时间
		if item.LossTime > 0 {
			count.WorkerLossTime[item.WorkerIndex] += int64(item.LossTime)
			item.UseTime = item.UseTime - item.LossTime
		}

		count.UseTime += int64(item.UseTime)
		*useTimes = append(*useTimes, item.UseTime)
		if item.Success {
			count.SuccessCount++
		} else {
			count.ErrorCount++
		}
	}

	if count.StartTime == 0 {
		count.StartTime = time.Now().UnixNano()
	}
	if count.EndTime == 0 {
		count.EndTime = time.Now().UnixNano()
	}
	count.full(useTimes)
	return
}

func CountCounts(countList []*Count) (count *Count) {
	count = &Count{}
	useTimes := &[]int{}

	count.WorkerLossTime = map[int]int64{}
	count.WorkerStartTime = map[int]int64{}
	count.WorkerEndTime = map[int]int64{}
	for _, item := range countList {
		startNano := item.StartTime
		endNano := item.EndTime

		if count.StartTime == 0 || startNano < count.StartTime {
			count.StartTime = item.StartTime
		}
		if count.EndTime == 0 || endNano > count.EndTime {
			count.EndTime = item.EndTime
		}

		for k, v := range item.WorkerStartTime {
			if count.WorkerStartTime[k] == 0 || v < count.WorkerStartTime[k] {
				count.WorkerStartTime[k] = v
			}
		}
		for k, v := range item.WorkerEndTime {
			if count.WorkerEndTime[k] == 0 || v > count.WorkerEndTime[k] {
				count.WorkerEndTime[k] = v
			}
		}

		count.UseTime += item.UseTime
		*useTimes = append(*useTimes, *item.useTimes...)

		count.SuccessCount += item.SuccessCount
		count.ErrorCount += item.ErrorCount
	}
	if count.StartTime == 0 {
		count.StartTime = time.Now().UnixNano()
	}
	if count.EndTime == 0 {
		count.EndTime = time.Now().UnixNano()
	}
	count.full(useTimes)
	return
}

type Count struct {
	StartTime    int64   `json:"startTime"` // 纳秒
	EndTime      int64   `json:"endTime"`   // 纳秒
	Count        int     `json:"count"`
	SuccessCount int     `json:"successCount"`
	ErrorCount   int     `json:"errorCount"`
	TotalTime    int64   `json:"totalTime"` // 执行时长 从 最小 开始时间 到 最大结束时间 的时间差 纳秒
	Total        string  `json:"total"`     // 执行时长 毫秒 保留 2位小数
	UseTime      int64   `json:"useTime"`   // 总调用时长 使用 所有项 的 耗时 相加
	Use          string  `json:"use"`       // 总调用时长 毫秒 保留 2位小数
	Max          string  `json:"max"`       // 最大时间 毫秒 保留 2位小数
	MaxUseTime   int     `json:"maxTime"`   // 最大时间 纳秒
	Min          string  `json:"min"`       // 最小时间 毫秒 保留 2位小数
	MinUseTime   int     `json:"minTime"`   // 最小时间 纳秒
	Tps          string  `json:"tps"`       // TPS 总次数 / 执行时长 秒 保留 2位小数
	TpsValue     float64 `json:"tpsValue"`  // TPS 总次数 / 执行时长 秒
	Avg          string  `json:"avg"`       // 平均耗时 总调用时长 / 总次数 毫秒 保留 2位小数
	AvgValue     float64 `json:"avgValue"`  // 平均耗时 总调用时长 / 总次数 毫秒
	T50          string  `json:"t50"`       // TOP 50 表示 百分之 50 的调用超过这个时间 毫秒 保留 2位小数
	T60          string  `json:"t60"`       // TOP 60 表示 百分之 60 的调用超过这个时间 毫秒 保留 2位小数
	T70          string  `json:"t70"`       // TOP 70 表示 百分之 70 的调用超过这个时间 毫秒 保留 2位小数
	T80          string  `json:"t80"`       // TOP 80 表示 百分之 80 的调用超过这个时间 毫秒 保留 2位小数
	T90          string  `json:"t90"`       // TOP 90 表示 百分之 90 的调用超过这个时间 毫秒 保留 2位小数
	T99          string  `json:"t99"`       // TOP 99 表示 百分之 99 的调用超过这个时间 毫秒 保留 2位小数
	//TopItems        []*Item        `json:"topItems"`  // 耗时最高的 10 条记录
	WorkerLossTime  map[int]int64  `json:"workerLossTime"`
	WorkerStartTime map[int]int64  `json:"workerStartTime"`
	WorkerEndTime   map[int]int64  `json:"workerEndTime"`
	WorkerTotalTime map[int]int64  `json:"workerTotalTime"`
	WorkerTotal     map[int]string `json:"workerTotal"`
	//topCount        int
	useTimes *[]int
}

func (this_ *Count) full(useTimes *[]int) {
	this_.useTimes = useTimes
	itemSize := len(*useTimes)
	if itemSize > 0 {
		sort.Sort(sort.IntSlice(*useTimes))
		this_.MinUseTime = (*useTimes)[0]
		this_.MaxUseTime = (*useTimes)[itemSize-1]
	}
	millF := float64(1000000)
	secF := float64(1000000000)

	// 耗时 纳秒
	this_.TotalTime = this_.EndTime - this_.StartTime
	// 真正消耗时间  需要 减去 损耗时间  这里减去某个工作线程的 损耗时间
	for _, one := range this_.WorkerLossTime {
		if one > 0 {
			this_.TotalTime -= one
			break
		}
	}
	this_.WorkerTotalTime = make(map[int]int64)
	this_.WorkerTotal = make(map[int]string)
	for workerIndex, startTime := range this_.WorkerStartTime {
		if this_.WorkerEndTime[workerIndex] == 0 {
			continue
		}
		this_.WorkerTotalTime[workerIndex] = this_.WorkerEndTime[workerIndex] - startTime
		this_.WorkerTotalTime[workerIndex] -= this_.WorkerLossTime[workerIndex]
		this_.WorkerTotal[workerIndex] = strconv.FormatFloat(float64(this_.WorkerTotalTime[workerIndex])/millF, 'f', 2, 64)
	}

	if this_.TotalTime > 0 {
		this_.Total = strconv.FormatFloat(float64(this_.TotalTime)/millF, 'f', 2, 64)
	}
	if this_.UseTime > 0 {
		this_.Use = strconv.FormatFloat(float64(this_.UseTime)/millF, 'f', 2, 64)
	}
	// 总次数
	this_.Count = this_.SuccessCount + this_.ErrorCount
	// 计算 TPS
	if this_.TotalTime > 0 && this_.Count > 0 {
		this_.TpsValue = float64(this_.Count) / (float64(this_.TotalTime) / secF)
		this_.Tps = strconv.FormatFloat(this_.TpsValue, 'f', 2, 64)
	}

	// 计算 平均
	if this_.UseTime > 0 && this_.Count > 0 {
		this_.AvgValue = float64(this_.UseTime) / float64(this_.Count) / millF
		this_.Avg = strconv.FormatFloat(this_.AvgValue, 'f', 2, 64)
	}

	if this_.MinUseTime > 0 {
		this_.Min = strconv.FormatFloat(float64(this_.MinUseTime)/millF, 'f', 2, 64)
	}
	if this_.MaxUseTime > 0 {
		this_.Max = strconv.FormatFloat(float64(this_.MaxUseTime)/millF, 'f', 2, 64)
	}

	if itemSize > 0 {

		this_.T50 = strconv.FormatFloat(float64((*useTimes)[int(float32(itemSize)*0.50)])/millF, 'f', 2, 64)
		this_.T60 = strconv.FormatFloat(float64((*useTimes)[int(float32(itemSize)*0.60)])/millF, 'f', 2, 64)
		this_.T70 = strconv.FormatFloat(float64((*useTimes)[int(float32(itemSize)*0.70)])/millF, 'f', 2, 64)
		this_.T80 = strconv.FormatFloat(float64((*useTimes)[int(float32(itemSize)*0.80)])/millF, 'f', 2, 64)
		this_.T90 = strconv.FormatFloat(float64((*useTimes)[int(float32(itemSize)*0.90)])/millF, 'f', 2, 64)
		this_.T99 = strconv.FormatFloat(float64((*useTimes)[int(float32(itemSize)*0.99)])/millF, 'f', 2, 64)

		//for i := itemSize - 1; i >= 0 && i >= itemSize-this_.topCount; i-- {
		//	this_.TopItems = append(this_.TopItems, (*itemList)[i])
		//}
	}
}

type ItemList []*Item

// Len 实现sort.Interface接口的获取元素数量方法
func (m *ItemList) Len() int {
	return len(*m)
}

// Less 实现sort.Interface接口的比较元素方法
func (m *ItemList) Less(i, j int) bool {
	return (*m)[i].UseTime < (*m)[j].UseTime
}

// Swap 实现sort.Interface接口的交换元素方法
func (m *ItemList) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}
