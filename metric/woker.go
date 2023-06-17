package metric

import (
	"sync"
	"time"
)

// WorkerMetric 单线程 统计
type WorkerMetric struct {
	WorkerIndex       int `json:"workerIndex"` // 工作线程索引  用于并发线程下 每个线程的时间跨度计算
	metric            *Metric
	secondItem        *SecondItem
	secondItems       *[]*SecondItem
	secondItemsLocker sync.Locker
	secondCounts      map[int]*Count
}

var (
	Second = int64(time.Second)
)

func (this_ *WorkerMetric) NewItem(startTime int64) (item *Item) {
	var second = int(startTime / Second)
	if this_.secondItem == nil || second != this_.secondItem.second {
		this_.secondItem = this_.newSecondItem(second)
	}
	item = this_.secondItem.newItem(startTime)

	return
}

func (this_ *WorkerMetric) newSecondItem(second int) (secondItem *SecondItem) {
	secondItem = &SecondItem{
		second:      second,
		itemsLocker: &sync.Mutex{},
		items:       &[]*Item{},
	}
	this_.secondItemsLocker.Lock()
	defer this_.secondItemsLocker.Unlock()

	*this_.secondItems = append(*this_.secondItems, secondItem)
	return
}

func (this_ *WorkerMetric) addSecondItems(secondItems []*SecondItem) {

	this_.secondItemsLocker.Lock()
	defer this_.secondItemsLocker.Unlock()

	*this_.secondItems = append(*this_.secondItems, secondItems...)
	return
}

func (this_ *WorkerMetric) getAndCleanSecondItems() (secondItems *[]*SecondItem) {

	this_.secondItemsLocker.Lock()
	defer this_.secondItemsLocker.Unlock()

	secondItems = this_.secondItems
	this_.secondItems = &[]*SecondItem{}
	return
}

func (this_ *WorkerMetric) countByDuration(nowSecond int, countSecond int) (count *Count, secondCounts map[int]*Count) {

	secondCounts = this_.secondCounts
	secondItems := this_.getAndCleanSecondItems()
	cacheItems := map[int]*[]*Item{}

	// 如果 同一秒 有未结束的 则保留
	var reserves []*SecondItem
	for _, secondItem := range *secondItems {
		if secondItem.second >= nowSecond || secondItem.itemNewSize != secondItem.itemEndSize {
			reserves = append(reserves, secondItem)
		}

		t := secondItem.second / countSecond
		if cacheItems[t] == nil {
			cacheItems[t] = &[]*Item{}
		}
		items := secondItem.getAndCleanItems()
		*cacheItems[t] = append(*cacheItems[t], *items...)
	}
	this_.addSecondItems(reserves)

	for t, tItems := range cacheItems {
		c := CountItems(tItems, this_.metric.countTop)
		if secondCounts[t] != nil {
			c = CountCounts([]*Count{c, secondCounts[t]}, this_.metric.countTop)
		}
		secondCounts[t] = c
	}
	var counts []*Count
	for _, c := range secondCounts {
		counts = append(counts, c)
	}
	count = CountCounts(counts, this_.metric.countTop)
	return
}
