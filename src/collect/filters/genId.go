package collect

import utils "github.com/SelfDown/collect/src/collect/utils"

func GenId(value string) interface{} {
	return utils.GenerateShortUniqueID(value, 32)
}
