package collect

import utils "collect/src/collect/utils"

func GenId(value string) interface{} {
	return utils.GenerateShortUniqueID(value, 32)
}
