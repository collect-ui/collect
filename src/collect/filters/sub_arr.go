package collect

import utils "collect/src/collect/utils"

func SubArr(arr []map[string]interface{}, index int, field string) []map[string]interface{} {
	item := arr[index]
	subArr, _ := utils.RenderVarToArrMap(field, item)
	return subArr
}
