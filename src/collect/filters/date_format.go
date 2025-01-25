package collect

import (
	utils "collect/src/collect/utils"
	"time"
)

func DateFormat(timeStr string, from_fmt string) string {
	if utils.IsValueEmpty(from_fmt) {
		// 默认去gitlab 的格式
		//2025-01-10T09:42:05.000+08:00
		from_fmt = time.RFC3339
	}
	t, err := time.Parse(from_fmt, timeStr)
	if err != nil {
		return ""
	}
	// 格式化为目标格式
	return t.Format("2006-01-02 15:04:05")
}
