package util

type StringSort interface {
	GetName() string
}

type StringSortList []StringSort

// Len 实现sort.Interface接口的获取元素数量方法
func (m StringSortList) Len() int {
	return len(m)
}

// Less 实现 sort.Interface 接口的比较元素方法
func (m StringSortList) Less(i, j int) bool {
	return m[i].GetName() < m[j].GetName()
}

// Swap 实现sort.Interface接口的交换元素方法
func (m StringSortList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
