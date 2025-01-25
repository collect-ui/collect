package collect

import (
	common "github.com/SelfDown/collect/src/collect/common"
	config "github.com/SelfDown/collect/src/collect/config"
	utils "github.com/SelfDown/collect/src/collect/utils"
)

/**
* 接收数组
 */
type PropArr struct {
	BaseHandler
}

func (uf *PropArr) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	params := template.GetParams()
	arr, _ := utils.RenderVarToArrMap(handlerParam.Foreach, params)
	li := make([]interface{}, 0)
	for _, item := range arr {

		// 为了处理二级数组
		if handlerParam.AppendItem{
			sub_arr, errMsg := utils.RenderVarToArrMap(handlerParam.Value, item)
			if !utils.IsValueEmpty(errMsg) {
                continue
            }
			for _, v := range sub_arr {
                li = append(li, v)
            }
		}else if !utils.IsValueEmpty(handlerParam.Fields){
			var paramsCopy map[string]interface{}
			if !utils.IsValueEmpty(handlerParam.Item) { // 如果没有配置item 则取本身
				paramsCopy = utils.CopyMap(params)
			}
			for _, field := range handlerParam.Fields {
				for index, dataItem := range arr {
					if !utils.IsValueEmpty(handlerParam.Item) { // 如果配置了item，设置item
						paramsCopy[utils.ItemName] = dataItem
					} else { // 没有配置item取整个item
						paramsCopy = dataItem
					}
					paramsCopy["loop_index"] = index
					//渲染值
					value := utils.RenderTplDataWithType(field.TemplateTpl, paramsCopy, field.Type)
					if utils.IsValueEmpty(value){
						continue
					}
					li = append(li, value)
				}
			}
		} else{
			value := utils.RenderVar(handlerParam.Value, item)
			li = append(li, value)
		}

	}
	return common.Ok(li, "处理成功")
}
