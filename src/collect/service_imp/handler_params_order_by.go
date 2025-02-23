package collect

import (
	common "github.com/collect-ui/collect/src/collect/common"
	config "github.com/collect-ui/collect/src/collect/config"
	utils "github.com/collect-ui/collect/src/collect/utils"
	"github.com/demdxx/gocast"
	"sort"
)

type OrderBy struct {
	BaseHandler
}

func (uf *OrderBy) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	params := template.GetParams()
	arr, _ := utils.RenderVarToArrMap(handlerParam.Foreach, params)
	fields := handlerParam.Fields
	sort.Slice(arr, func(i, j int) bool {
		less := true
		for _, field := range fields {
			fieldName := field.Field
			x := arr[i]
			xValue := utils.RenderVar(fieldName, x)
			y := arr[j]
			yValue := utils.RenderVar(fieldName, y)
			if xValue == yValue {
				continue
			}
			if utils.IsNumber(xValue) && utils.IsNumber(yValue) { // 如果都是数字类型
				less = gocast.ToFloat(xValue) < gocast.ToFloat(yValue)
			} else {
				less = utils.Strval(xValue) < utils.Strval(yValue)
			}

			rule := field.Rule
			if rule == "desc" {
				less = !less
			}
			break

		}
		return less
	})

	r := common.Ok(arr, "处理参数成功")
	return r
}
