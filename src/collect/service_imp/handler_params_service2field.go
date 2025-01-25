package collect

import (
	common "collect/src/collect/common"
	config "collect/src/collect/config"
	utils "collect/src/collect/utils"
	"time"
)

type Service2Field struct {
	BaseHandler
}

/**
* 只做了拼接参数，未做渲染
 */
func (uf *Service2Field) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	params := template.GetParams()
	serviceParam := utils.GetServiceParam(handlerParam.Service, params, handlerParam.AppendParam)
	if handlerParam.AppendItemParam {
		itemMap, msg := utils.RenderVarToMap(handlerParam.Item, params)
		if utils.IsValueEmpty(msg) {
			for key, value := range itemMap {
				serviceParam[key] = value
			}
		} else {
			return common.NotOk(msg)
		}

	}
	if handlerParam.LoopMax > 0 {
		var i int64 // 定义 i 为 int64 类型
		var r2 *common.Result
		for i = 0; i < handlerParam.LoopMax; i++ {
			r2 = ts.ResultInner(serviceParam)
			if !r2.Success {
				return r2
			}
			saveField := handlerParam.SaveField //处理存储
			if !utils.IsValueEmpty(saveField) { //如果有保存字段，则处理
				template.AddParam(saveField, r2.GetData())
			}
			// 处理判断服务是否运行正确
			tpl := handlerParam.TemplateTpl
			if tpl != nil { // 如果template 模板不为空，则渲染值
				success := utils.RenderTplBool(tpl, params)
				if success { // 成功
					break
				}
			}
			// 请求完成等待
			if handlerParam.Second > 0 {
				tmp := time.Second * time.Duration(handlerParam.Second)
				time.Sleep(tmp)
			} else {
				time.Sleep(time.Second * 1) //
			}
		}
		return r2
	} else {
		r2 := ts.ResultInner(serviceParam)
		return r2
	}

}
