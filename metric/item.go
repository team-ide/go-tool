package metric

import "sync"

type SecondItem struct {
	second      int
	items       *[]*Item
	itemsLocker sync.Locker
	itemNewSize int
	itemEndSize int
}

func (this_ *SecondItem) getAndCleanItems() (items *[]*Item) {
	this_.itemsLocker.Lock()
	defer this_.itemsLocker.Unlock()

	items = this_.items
	size := len(*items)
	this_.itemNewSize -= size
	this_.itemEndSize -= size
	this_.items = &[]*Item{}
	return
}

func (this_ *SecondItem) newItem(startTime int64) (item *Item) {
	item = &Item{
		secondItem: this_,
		StartTime:  startTime,
	}

	this_.itemsLocker.Lock()
	defer this_.itemsLocker.Unlock()

	this_.itemNewSize++
	return
}
func (this_ *SecondItem) addItem(item *Item) {
	this_.itemsLocker.Lock()
	defer this_.itemsLocker.Unlock()

	this_.itemNewSize++
	this_.itemEndSize++
	*this_.items = append(*this_.items, item)
	return
}
func (this_ *SecondItem) endItem(item *Item) {
	this_.itemsLocker.Lock()
	defer this_.itemsLocker.Unlock()

	this_.itemEndSize++
	*this_.items = append(*this_.items, item)
	return
}

type Item struct {
	secondItem *SecondItem
	StartTime  int64       `json:"startTime"`
	EndTime    int64       `json:"endTime"`
	Success    bool        `json:"success"`
	Extend     interface{} `json:"extend"`
	UseTime    int         `json:"useTime"`
}

func (this_ *Item) End(useTime int, endTime int64, err error) {
	this_.Success = err == nil
	this_.UseTime = useTime
	this_.EndTime = endTime

	this_.secondItem.endItem(this_)
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
