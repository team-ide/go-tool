package metric

import (
	"sort"
	"strconv"
	"sync"
	"time"
)

type Metric struct {
	locker   sync.Mutex
	items    []*Item
	TopCount int `json:"topCount"`
}

func (this_ *Metric) NewItem(workerIndex int, startTime time.Time) (item *Item) {

	item = &Item{
		WorkerIndex: workerIndex,
		StartTime:   startTime.UnixNano(),
	}

	this_.locker.Lock()
	defer this_.locker.Unlock()

	this_.items = append(this_.items, item)
	return
}

func (this_ *Metric) Count() (count *Count) {
	items := this_.items

	count = CountItems(items, this_.TopCount)
	return
}

func (this_ *Metric) CountSecond() (counts []*Count) {
	items := this_.items

	cacheItems := map[int]*[]*Item{}

	var ts []int
	for _, item := range items {
		if item.StartTime == 0 || item.EndTime == 0 {
			continue
		}
		t := int(item.StartTime / int64(time.Second))
		if cacheItems[t] == nil {
			cacheItems[t] = &[]*Item{}
			ts = append(ts, t)
		}
		*cacheItems[t] = append(*cacheItems[t], item)
	}
	sort.Ints(ts)
	for _, t := range ts {
		list := cacheItems[t]
		count := CountItems(*list, this_.TopCount)
		counts = append(counts, count)
	}

	return
}

func (this_ *Metric) CountMinute() (counts []*Count) {
	items := this_.items

	cacheItems := map[int]*[]*Item{}

	var ts []int
	for _, item := range items {
		if item.StartTime == 0 || item.EndTime == 0 {
			continue
		}
		t := int(item.StartTime / int64(time.Minute))
		if cacheItems[t] == nil {
			cacheItems[t] = &[]*Item{}
			ts = append(ts, t)
		}
		*cacheItems[t] = append(*cacheItems[t], item)
	}
	sort.Ints(ts)
	for _, t := range ts {
		list := cacheItems[t]
		count := CountItems(*list, this_.TopCount)
		counts = append(counts, count)
	}

	return
}

type Item struct {
	WorkerIndex int         `json:"workerIndex"` // 工作线程索引  用于并发线程下 每个线程的时间跨度计算
	StartTime   int64       `json:"startTime"`
	EndTime     int64       `json:"endTime"`
	Success     bool        `json:"success"`
	LossTime    int64       `json:"lossTime"` // 损耗时长 纳秒 统计时候将去除该部分时长
	Extend      interface{} `json:"extend"`
	UseTime     int64       `json:"useTime"`
}

func (this_ *Item) End(endTime time.Time, err error) {
	this_.Success = err == nil
	this_.EndTime = endTime.UnixNano()
}

func (this_ *Item) Loss(lossTime int64) {
	this_.LossTime = lossTime
}

func CountItems(itemList ItemList, topCount int) (count *Count) {
	count = &Count{}
	if topCount > 0 {
		count.topCount = topCount
	} else {
		count.topCount = 10
	}
	var MaxUseTime int64 = -1
	var MinUseTime int64 = -1

	count.WorkerLossTime = map[int]int64{}
	count.WorkerStartTime = map[int]int64{}
	count.WorkerEndTime = map[int]int64{}
	for _, item := range itemList {
		if item.StartTime == 0 || item.EndTime == 0 {
			continue
		}
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

		item.UseTime = endNano - startNano

		// 使用时间 需要 减去 损耗时间
		if item.LossTime > 0 {
			count.WorkerLossTime[item.WorkerIndex] += item.LossTime
			item.UseTime = item.UseTime - item.LossTime
		}

		count.UseTime += item.UseTime
		if MaxUseTime < 0 || MaxUseTime < item.UseTime {
			MaxUseTime = item.UseTime
		}
		if MinUseTime < 0 || MinUseTime > item.UseTime {
			MinUseTime = item.UseTime
		}
		if item.Success {
			count.SuccessCount++
		} else {
			count.ErrorCount++
		}
	}
	if MaxUseTime >= 0 {
		count.MaxUseTime = MaxUseTime
	}
	if MinUseTime >= 0 {
		count.MinUseTime = MinUseTime
	}
	if count.StartTime == 0 {
		count.StartTime = time.Now().UnixNano()
	}
	if count.EndTime == 0 {
		count.EndTime = time.Now().UnixNano()
	}
	count.full(itemList)
	return
}

type Count struct {
	StartTime       int64          `json:"startTime"` // 纳秒
	EndTime         int64          `json:"endTime"`   // 纳秒
	Count           int            `json:"count"`
	SuccessCount    int            `json:"successCount"`
	ErrorCount      int            `json:"errorCount"`
	TotalTime       int64          `json:"totalTime"` // 执行时长 从 最小 开始时间 到 最大结束时间 的时间差 纳秒
	Total           string         `json:"total"`     // 执行时长 毫秒 保留 2位小数
	UseTime         int64          `json:"useTime"`   // 总调用时长 使用 所有项 的 耗时 相加
	Use             string         `json:"use"`       // 总调用时长 毫秒 保留 2位小数
	Max             string         `json:"max"`       // 最大时间 毫秒 保留 2位小数
	MaxUseTime      int64          `json:"maxTime"`   // 最大时间 纳秒
	Min             string         `json:"min"`       // 最小时间 毫秒 保留 2位小数
	MinUseTime      int64          `json:"minTime"`   // 最小时间 纳秒
	Tps             string         `json:"tps"`       // TPS 总次数 / 执行时长 秒 保留 2位小数
	TpsValue        float64        `json:"tpsValue"`  // TPS 总次数 / 执行时长 秒
	Avg             string         `json:"avg"`       // 平均耗时 总调用时长 / 总次数 毫秒 保留 2位小数
	AvgValue        float64        `json:"avgValue"`  // 平均耗时 总调用时长 / 总次数 毫秒
	T50             string         `json:"t50"`       // TOP 50 表示 百分之 50 的调用超过这个时间 毫秒 保留 2位小数
	T60             string         `json:"t60"`       // TOP 60 表示 百分之 60 的调用超过这个时间 毫秒 保留 2位小数
	T70             string         `json:"t70"`       // TOP 70 表示 百分之 70 的调用超过这个时间 毫秒 保留 2位小数
	T80             string         `json:"t80"`       // TOP 80 表示 百分之 80 的调用超过这个时间 毫秒 保留 2位小数
	T90             string         `json:"t90"`       // TOP 90 表示 百分之 90 的调用超过这个时间 毫秒 保留 2位小数
	T99             string         `json:"t99"`       // TOP 99 表示 百分之 99 的调用超过这个时间 毫秒 保留 2位小数
	TopItems        []*Item        `json:"topItems"`  // 耗时最高的 10 条记录
	WorkerLossTime  map[int]int64  `json:"workerLossTime"`
	WorkerStartTime map[int]int64  `json:"workerStartTime"`
	WorkerEndTime   map[int]int64  `json:"workerEndTime"`
	WorkerTotalTime map[int]int64  `json:"workerTotalTime"`
	WorkerTotal     map[int]string `json:"workerTotal"`
	topCount        int
}

func (this_ *Count) full(itemList ItemList) {

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

	itemSize := itemList.Len()
	if itemSize > 0 {
		sort.Sort(itemList)

		this_.T50 = strconv.FormatFloat(float64(itemList[int(float32(itemSize)*0.50)].UseTime)/millF, 'f', 2, 64)
		this_.T60 = strconv.FormatFloat(float64(itemList[int(float32(itemSize)*0.60)].UseTime)/millF, 'f', 2, 64)
		this_.T70 = strconv.FormatFloat(float64(itemList[int(float32(itemSize)*0.70)].UseTime)/millF, 'f', 2, 64)
		this_.T80 = strconv.FormatFloat(float64(itemList[int(float32(itemSize)*0.80)].UseTime)/millF, 'f', 2, 64)
		this_.T90 = strconv.FormatFloat(float64(itemList[int(float32(itemSize)*0.90)].UseTime)/millF, 'f', 2, 64)
		this_.T99 = strconv.FormatFloat(float64(itemList[int(float32(itemSize)*0.99)].UseTime)/millF, 'f', 2, 64)

		for i := itemSize - 1; i >= 0 && i >= itemSize-this_.topCount; i-- {
			this_.TopItems = append(this_.TopItems, itemList[i])
		}
	}
}

type ItemList []*Item

// Len 实现sort.Interface接口的获取元素数量方法
func (m ItemList) Len() int {
	return len(m)
}

// Less 实现sort.Interface接口的比较元素方法
func (m ItemList) Less(i, j int) bool {
	return m[i].UseTime < m[j].UseTime
}

// Swap 实现sort.Interface接口的交换元素方法
func (m ItemList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
