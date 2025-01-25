package collect

import (
	utils "github.com/collect-ui/collect/src/collect/utils"
	"time"
)

func UnixTime2Datetime(unit int64) string {
	t := time.Unix(unit, 0)
	return utils.DateFormatDefault(t)
}
