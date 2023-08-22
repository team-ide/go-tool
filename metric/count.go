package metric

import (
	"sort"
	"strconv"
	"time"
)

var (
	defaultTime         = -99999
	defaultTime64 int64 = -99999
)

func CountItems(itemList *[]*Item, countTop bool) (count *Count) {
	count = &Count{
		MinUseTime: defaultTime,
		MaxUseTime: defaultTime,
		StartTime:  defaultTime64,
		EndTime:    defaultTime64,
	}
	useTimes := &[]int{}

	for _, one := range *itemList {

		if one.UseTime > 0 {
			if count.MinUseTime == defaultTime || one.UseTime < count.MinUseTime {
				count.MinUseTime = one.UseTime
			}
			if count.MaxUseTime == defaultTime || one.UseTime > count.MaxUseTime {
				count.MaxUseTime = one.UseTime
			}
		}

		if one.StartTime > 0 {
			if count.StartTime == defaultTime64 || one.StartTime < count.StartTime {
				count.StartTime = one.StartTime
			}
		}
		if one.EndTime > 0 {
			if count.EndTime == defaultTime64 || one.EndTime > count.EndTime {
				count.EndTime = one.EndTime
			}
		}

		count.UseTime += int64(one.UseTime)
		if countTop {
			*useTimes = append(*useTimes, one.UseTime)
		}
		if one.Success {
			count.SuccessCount++
		} else {
			count.ErrorCount++
		}
	}
	if count.MinUseTime == defaultTime {
		count.MinUseTime = 0
	}
	if count.MaxUseTime == defaultTime {
		count.MaxUseTime = 0
	}
	if count.StartTime == defaultTime64 {
		count.StartTime = time.Now().UnixNano()
	}
	if count.EndTime == defaultTime64 {
		count.EndTime = time.Now().UnixNano()
	}
	count.full(count.UseTime, useTimes)
	return
}

func CountCounts(countList []*Count, countTop bool) (count *Count) {
	count = &Count{
		MinUseTime: defaultTime,
		MaxUseTime: defaultTime,
		StartTime:  defaultTime64,
		EndTime:    defaultTime64,
	}
	useTimes := &[]int{}

	for _, one := range countList {

		if one.MinUseTime > 0 {
			if count.MinUseTime == defaultTime || one.MinUseTime < count.MinUseTime {
				count.MinUseTime = one.MinUseTime
			}
		}

		if one.MaxUseTime > 0 {
			if count.MaxUseTime == defaultTime || one.MaxUseTime > count.MaxUseTime {
				count.MaxUseTime = one.MaxUseTime
			}
		}

		if one.StartTime > 0 {
			if count.StartTime == defaultTime64 || one.StartTime < count.StartTime {
				count.StartTime = one.StartTime
			}
		}
		if one.EndTime > 0 {
			if count.EndTime == defaultTime64 || one.EndTime > count.EndTime {
				count.EndTime = one.EndTime
			}
		}

		count.UseTime += one.UseTime
		if countTop {
			*useTimes = append(*useTimes, *one.useTimes...)
		}

		count.SuccessCount += one.SuccessCount
		count.ErrorCount += one.ErrorCount
	}
	if count.MinUseTime == defaultTime {
		count.MinUseTime = 0
	}
	if count.MaxUseTime == defaultTime {
		count.MaxUseTime = 0
	}
	if count.StartTime == defaultTime64 {
		count.StartTime = time.Now().UnixNano()
	}
	if count.EndTime == defaultTime64 {
		count.EndTime = time.Now().UnixNano()
	}
	count.full(count.UseTime, useTimes)
	return
}

func WorkersCount(countList []*Count, countTop bool) (count *Count) {
	count = &Count{
		MinUseTime: defaultTime,
		MaxUseTime: defaultTime,
		StartTime:  defaultTime64,
		EndTime:    defaultTime64,
	}
	useTimes := &[]int{}

	var workerUseTime = defaultTime64
	for _, one := range countList {

		if one.MinUseTime > 0 {
			if count.MinUseTime == defaultTime || one.MinUseTime < count.MinUseTime {
				count.MinUseTime = one.MinUseTime
			}
		}

		if one.MaxUseTime > 0 {
			if count.MaxUseTime == defaultTime || one.MaxUseTime > count.MaxUseTime {
				count.MaxUseTime = one.MaxUseTime
			}
		}

		if one.StartTime > 0 {
			if count.StartTime == defaultTime64 || one.StartTime < count.StartTime {
				count.StartTime = one.StartTime
			}
		}
		if one.EndTime > 0 {
			if count.EndTime == defaultTime64 || one.EndTime > count.EndTime {
				count.EndTime = one.EndTime
			}
		}

		if one.UseTime > 0 {
			if workerUseTime == defaultTime64 || one.UseTime > workerUseTime {
				workerUseTime = one.UseTime
			}
		}
		count.UseTime += one.UseTime
		if countTop {
			*useTimes = append(*useTimes, *one.useTimes...)
		}

		count.SuccessCount += one.SuccessCount
		count.ErrorCount += one.ErrorCount
	}
	if count.MinUseTime == defaultTime {
		count.MinUseTime = 0
	}
	if count.MaxUseTime == defaultTime {
		count.MaxUseTime = 0
	}
	if workerUseTime == defaultTime64 {
		workerUseTime = 0
	}
	if count.StartTime == defaultTime64 {
		count.StartTime = time.Now().UnixNano()
	}
	if count.EndTime == defaultTime64 {
		count.EndTime = time.Now().UnixNano()
	}
	count.full(workerUseTime, useTimes)
	return
}

type Count struct {
	Name         string  `json:"name"`
	StartTime    int64   `json:"startTime"` // 纳秒
	EndTime      int64   `json:"endTime"`   // 纳秒
	Count        int     `json:"count"`
	SuccessCount int     `json:"successCount"`
	ErrorCount   int     `json:"errorCount"`
	TotalTime    int64   `json:"totalTime"`   // 执行时长包括额外开销 从 最小 开始时间 到 最大结束时间 的时间差 纳秒
	Total        string  `json:"total"`       // 执行时长包括额外开销 毫秒 保留 2位小数
	ExecuteTime  int64   `json:"executeTime"` // 执行时长包括额外开销 从 最小 开始时间 到 最大结束时间 的时间差 纳秒
	Execute      string  `json:"execute"`     // 执行时长包括额外开销 毫秒 保留 2位小数
	UseTime      int64   `json:"useTime"`     // 总调用时长 使用 所有项 的 耗时 相加
	Use          string  `json:"use"`         // 总调用时长 毫秒 保留 2位小数
	Max          string  `json:"max"`         // 最大时间 毫秒 保留 2位小数
	MaxUseTime   int     `json:"maxTime"`     // 最大时间 纳秒
	Min          string  `json:"min"`         // 最小时间 毫秒 保留 2位小数
	MinUseTime   int     `json:"minTime"`     // 最小时间 纳秒
	Tps          string  `json:"tps"`         // TPS 总次数 / 执行时长 秒 保留 2位小数
	TpsValue     float64 `json:"tpsValue"`    // TPS 总次数 / 执行时长 秒
	Avg          string  `json:"avg"`         // 平均耗时 总调用时长 / 总次数 毫秒 保留 2位小数
	AvgValue     float64 `json:"avgValue"`    // 平均耗时 总调用时长 / 总次数 毫秒
	T50          string  `json:"t50"`         // TOP 50 表示 百分之 50 的调用超过这个时间 毫秒 保留 2位小数
	T60          string  `json:"t60"`         // TOP 60 表示 百分之 60 的调用超过这个时间 毫秒 保留 2位小数
	T70          string  `json:"t70"`         // TOP 70 表示 百分之 70 的调用超过这个时间 毫秒 保留 2位小数
	T80          string  `json:"t80"`         // TOP 80 表示 百分之 80 的调用超过这个时间 毫秒 保留 2位小数
	T90          string  `json:"t90"`         // TOP 90 表示 百分之 90 的调用超过这个时间 毫秒 保留 2位小数
	T99          string  `json:"t99"`         // TOP 99 表示 百分之 99 的调用超过这个时间 毫秒 保留 2位小数
	useTimes     *[]int
}

func (this_ *Count) full(executeTime int64, useTimes *[]int) {
	this_.useTimes = useTimes
	itemSize := len(*useTimes)
	if itemSize > 0 {
		sort.Sort(sort.IntSlice(*useTimes))
	}
	millF := float64(1000000)
	secF := float64(1000000000)

	// 耗时 纳秒
	this_.TotalTime = this_.EndTime - this_.StartTime
	this_.Total = strconv.FormatFloat(float64(this_.TotalTime)/millF, 'f', 2, 64)
	this_.Use = strconv.FormatFloat(float64(this_.UseTime)/millF, 'f', 2, 64)
	this_.ExecuteTime = executeTime
	this_.Execute = strconv.FormatFloat(float64(this_.ExecuteTime)/millF, 'f', 2, 64)
	// 总次数
	this_.Count = this_.SuccessCount + this_.ErrorCount
	// 计算 TPS
	if executeTime > 0 {
		this_.TpsValue = float64(this_.Count) / (float64(this_.TotalTime) / secF)
	}
	this_.Tps = strconv.FormatFloat(this_.TpsValue, 'f', 2, 64)

	// 计算 平均
	if this_.Count > 0 {
		this_.AvgValue = float64(this_.UseTime) / float64(this_.Count) / millF
	}
	this_.Avg = strconv.FormatFloat(this_.AvgValue, 'f', 2, 64)

	this_.Min = strconv.FormatFloat(float64(this_.MinUseTime)/millF, 'f', 2, 64)
	this_.Max = strconv.FormatFloat(float64(this_.MaxUseTime)/millF, 'f', 2, 64)

	this_.T50 = strconv.FormatFloat(0, 'f', 2, 64)
	this_.T60 = strconv.FormatFloat(0, 'f', 2, 64)
	this_.T70 = strconv.FormatFloat(0, 'f', 2, 64)
	this_.T80 = strconv.FormatFloat(0, 'f', 2, 64)
	this_.T80 = strconv.FormatFloat(0, 'f', 2, 64)
	this_.T90 = strconv.FormatFloat(0, 'f', 2, 64)
	this_.T99 = strconv.FormatFloat(0, 'f', 2, 64)
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
