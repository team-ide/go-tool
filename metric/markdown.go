package metric

import (
	"fmt"
	"github.com/team-ide/go-tool/util"
	"time"
)

func MarkdownTableByCounts(counts []*Count) (content string) {
	content += fmt.Sprintf("| 任务时间 | 总/成功/失败 |执行用时|累计用时 |TPS |Avg |Min |Max |T50 |T80 | T90 | T99 |  \n")
	content += fmt.Sprintf("| :------: | :------: |:------: |:------:|:------: |:------: |:------: |:------: |:------: |:------: | :------: | :------: |  \n")

	for _, count := range counts {

		content += fmt.Sprintf("|")
		content += fmt.Sprintf(" %s <br>-<br> %s |",
			util.TimeFormat(time.UnixMilli(count.StartTime/int64(time.Millisecond)), "2006-01-02 15:04:05.000"),
			util.TimeFormat(time.UnixMilli(count.EndTime/int64(time.Millisecond)), "2006-01-02 15:04:05.000"),
		)
		content += fmt.Sprintf(" %d <br> <font color='green'>%d</font> <br> <font color='red'>%d</font> |", count.Count, count.SuccessCount, count.ErrorCount)
		content += fmt.Sprintf(" %s |", toTime(count.TotalTime/1000000))
		content += fmt.Sprintf(" %s |", toTime(count.UseTime/1000000))
		content += fmt.Sprintf(" %s |", count.Tps)
		content += fmt.Sprintf(" %s |", count.Avg)
		content += fmt.Sprintf(" %s |", count.Min)
		content += fmt.Sprintf(" %s |", count.Max)
		content += fmt.Sprintf(" %s |", count.T50)
		content += fmt.Sprintf(" %s |", count.T80)
		content += fmt.Sprintf(" %s |", count.T90)
		content += fmt.Sprintf(" %s |", count.T99)
		content += fmt.Sprintf("\n")
	}
	return
}

type tS struct {
	Size float64
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

func toTime(size int64) (v string) {

	var timeUnit string
	var timeV = float64(size)

	for _, s := range tList {
		if timeUnit == "" && timeV >= s.Size {
			timeV = timeV / s.Size
			timeUnit = s.Unit
		}
	}
	if timeUnit == "" {
		timeUnit = "毫秒"
	}
	v = fmt.Sprintf("%.2f%s", timeV, timeUnit)
	return
}
