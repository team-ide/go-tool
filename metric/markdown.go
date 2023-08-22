package metric

import (
	"fmt"
	"github.com/team-ide/go-tool/util"
	"time"
)

func MarkdownTableByCounts(counts []*Count) (content string) {
	return MarkdownTable(counts, &Options{
		AddHtmlFormat: true,
	})
}

type Options struct {
	AddHtmlFormat bool
	WarnUseTime   int64 // 超过该时间 出现警告  为毫秒数
	Columns       []*TableColumn
}

type TableColumn struct {
	Label  string `json:"label"`
	GetCol func(count *Count, options *Options) (res string)
}

var (
	TableColumnName = &TableColumn{
		Label: "名称",
		GetCol: func(count *Count, options *Options) (res string) {
			res = count.Name
			return
		},
	}
	TableColumnTime = &TableColumn{
		Label: "任务时间",
		GetCol: func(count *Count, options *Options) (res string) {
			if options.AddHtmlFormat {
				res = " %s <br>-<br> %s"
			} else {
				res = " %s - %s"
			}
			res = fmt.Sprintf(res,
				util.TimeFormat(time.UnixMilli(count.StartTime/int64(time.Millisecond)), "2006-01-02 15:04:05.000"),
				util.TimeFormat(time.UnixMilli(count.EndTime/int64(time.Millisecond)), "2006-01-02 15:04:05.000"),
			)
			return
		},
	}
	TableColumnCount = &TableColumn{
		Label: "总/成功/失败",
		GetCol: func(count *Count, options *Options) (res string) {
			if options.AddHtmlFormat {
				res = "%d <br> <font color='green'>%d</font> <br> <font color='red'>%d</font>"
			} else {
				res = "%d / %d / %d"
			}
			res = fmt.Sprintf(res, count.Count, count.SuccessCount, count.ErrorCount)
			return
		},
	}
	TableColumnTotalTime = &TableColumn{
		Label: "任务用时",
		GetCol: func(count *Count, options *Options) (res string) {
			res = ToTimeStr(count.TotalTime / 1000000)
			return
		},
	}
	TableColumnExecuteTime = &TableColumn{
		Label: "执行用时",
		GetCol: func(count *Count, options *Options) (res string) {
			res = ToTimeStr(count.ExecuteTime / 1000000)
			return
		},
	}
	TableColumnUseTime = &TableColumn{
		Label: "累计用时",
		GetCol: func(count *Count, options *Options) (res string) {
			res = ToTimeStr(count.UseTime / 1000000)
			return
		},
	}
	TableColumnTps = &TableColumn{
		Label: "TPS",
		GetCol: func(count *Count, options *Options) (res string) {
			res = count.Tps
			return
		},
	}
	TableColumnAvg = &TableColumn{
		Label: "Avg",
		GetCol: func(count *Count, options *Options) (res string) {
			res = GetTableTimeOut(count.AvgValue, count.Avg, options)
			return
		},
	}
	TableColumnMin = &TableColumn{
		Label: "Min",
		GetCol: func(count *Count, options *Options) (res string) {
			res = GetTableTimeOut(float64(count.MinUseTime)/float64(time.Millisecond), count.Min, options)
			return
		},
	}
	TableColumnMax = &TableColumn{
		Label: "Max",
		GetCol: func(count *Count, options *Options) (res string) {
			res = GetTableTimeOut(float64(count.MaxUseTime)/float64(time.Millisecond), count.Max, options)
			return
		},
	}
	TableColumnT50 = &TableColumn{
		Label: "T50",
		GetCol: func(count *Count, options *Options) (res string) {
			res = GetTableTimeOut(util.StringToFloat64(count.T50), count.T50, options)
			return
		},
	}
	TableColumnT80 = &TableColumn{
		Label: "T80",
		GetCol: func(count *Count, options *Options) (res string) {
			res = GetTableTimeOut(util.StringToFloat64(count.T80), count.T80, options)
			return
		},
	}
	TableColumnT90 = &TableColumn{
		Label: "T90",
		GetCol: func(count *Count, options *Options) (res string) {
			res = GetTableTimeOut(util.StringToFloat64(count.T90), count.T90, options)
			return
		},
	}
	TableColumnT99 = &TableColumn{
		Label: "T99",
		GetCol: func(count *Count, options *Options) (res string) {
			res = GetTableTimeOut(util.StringToFloat64(count.T99), count.T99, options)
			return
		},
	}
	defaultColumns = []*TableColumn{
		TableColumnTime,
		TableColumnCount,
		TableColumnTotalTime,
		TableColumnTps,
		TableColumnAvg,
		TableColumnMin,
		TableColumnMax,
		TableColumnT50,
		TableColumnT80,
		TableColumnT90,
		TableColumnT99,
	}
)

func GetTableTimeOut(v float64, s string, options *Options) string {
	if options.AddHtmlFormat && options.WarnUseTime > 0 && int64(v) >= options.WarnUseTime {
		return fmt.Sprintf("<font color='red'>%s</font>", s)
	}
	return s
}

func MarkdownTable(counts []*Count, options *Options) (content string) {
	if options == nil {
		options = &Options{}
	}
	columns := options.Columns
	if len(columns) == 0 {
		columns = defaultColumns
	}
	content += "|"
	for _, column := range columns {
		content += " " + column.Label + " |"
	}
	content += "\n"
	content += "|"
	for _, _ = range columns {
		content += " :------: |"
	}
	content += "\n"

	for _, count := range counts {
		content += "|"
		for _, column := range columns {
			content += " " + column.GetCol(count, options) + " |"
		}
		content += "\n"
	}

	return
}

type tS struct {
	Size int64
	Unit string
}

var (
	tList = []*tS{
		{Size: 1000 * 60 * 60 * 24, Unit: "天"},
		{Size: 1000 * 60 * 60, Unit: "时"},
		{Size: 1000 * 60, Unit: "分"},
		{Size: 1000, Unit: "秒"},
	}
)

func ToTimeStr(size int64) (v string) {

	var timeV = size

	for _, s := range tList {
		if timeV >= s.Size {
			tV := timeV / s.Size
			timeV -= tV * s.Size
			v += fmt.Sprintf("%d%s", tV, s.Unit)
		}
	}
	if timeV > 0 {
		v += fmt.Sprintf("%d%s", timeV, "毫秒")
	}
	return
}
