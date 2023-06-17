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
}

func MarkdownTable(counts []*Count, options *Options) (content string) {
	if options == nil {
		options = &Options{}
	}
	content += fmt.Sprintf("|                 任务时间                   | 总/成功/失败    |  执行用时 | 累计用时 |TPS |Avg |Min |Max |T50 |T80 | T90 | T99 |  \n")
	content += fmt.Sprintf("| :------: | :------: | :------: | :------: | :------: |:------: | :------: | :------: | :------: | :------: | :------: | :------: |  \n")

	var s string
	for _, count := range counts {

		content += fmt.Sprintf("|")
		if options.AddHtmlFormat {
			s = " %s <br>-<br> %s <br>-<br> %s |"
		} else {
			s = " %s - %s - %s |"
		}
		content += fmt.Sprintf(s,
			util.TimeFormat(time.UnixMilli(count.StartTime/int64(time.Millisecond)), "2006-01-02 15:04:05.000"),
			util.TimeFormat(time.UnixMilli(count.EndTime/int64(time.Millisecond)), "2006-01-02 15:04:05.000"),
			toTime(count.TotalTime/1000000),
		)
		if options.AddHtmlFormat {
			s = " %d <br> <font color='green'>%d</font> <br> <font color='red'>%d</font> |"
		} else {
			s = " %d / %d / %d |"
		}
		content += fmt.Sprintf(s, count.Count, count.SuccessCount, count.ErrorCount)
		content += fmt.Sprintf(" %s |", toTime(count.ExecuteTime/1000000))
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

func toTime(size int64) (v string) {

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
