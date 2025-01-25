package collect

import utils "github.com/collect-ui/collect/src/collect/utils"

func GenId(value string) interface{} {
	return utils.GenerateShortUniqueID(value, 32)
}
