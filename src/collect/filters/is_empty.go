package collect

import utils "github.com/collect-ui/collect/src/collect/utils"

func IsEmpty(value interface{}) bool {
	return utils.IsValueEmpty(value)
}
