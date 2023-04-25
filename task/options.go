package task

type Options struct {
	Key string `json:"key"` // Key 任务Key，同时只能存在一个

	Worker        int   `json:"worker"`   // 任务 并发数
	Duration      int   `json:"duration"` // 任务 执行持续时长  分钟
	durationMilli int64 // 时长 毫秒数
	Frequency     int   `json:"frequency"` // 任务 执行次数，和 时长 互斥，只能一个生效，优先级高于 时长

	Executor Executor `json:"-"`
}
