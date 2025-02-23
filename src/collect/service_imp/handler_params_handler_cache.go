package collect

import (
	common "github.com/collect-ui/collect/src/collect/common"
	config "github.com/collect-ui/collect/src/collect/config"
	cacheHandler "github.com/collect-ui/collect/src/collect/service_imp/cache_handler"
	utils "github.com/collect-ui/collect/src/collect/utils"
)

// 处理缓存
type HandlerCache struct {
	BaseHandler
}

func (hc *HandlerCache) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {

	params := template.GetParams()
	method := handlerParam.Method
	handler := cacheHandler.CacheHandler{}
	fieldList := hc.GetFieldNames(handlerParam, params)

	if method == cacheHandler.CacheGetName { // 获取缓存，并且结束
		dataKey := handler.GetCacheKey(handlerParam.Room, fieldList, params)
		data, ok := handler.Get(dataKey)
		if ok {
			result := data.(common.Result)
			//result.SetFinish(true)
			return &result
		}
		//else
		//{
		//	return common.NotOk("未获取到缓存变量")
		//}

	} else if method == cacheHandler.CacheSetName { // 单个设置缓存
		dataKey := handler.GetCacheKey(handlerParam.Room, fieldList, params)
		result := template.GetResult()
		var ok = false
		field := handlerParam.Field

		if utils.IsValueEmpty(field) { //如果有field 证明是在handler_params 中，没有field是在cache中
			ok = handler.Set(dataKey, *result, handlerParam.Second)
		} else {
			data := utils.RenderVar(field, params)
			result = common.Ok(data, "缓存数据")
			ok = handler.Set(dataKey, *result, handlerParam.Second)
		}
		if !ok {
			template.LogData("缓存设置失败" + dataKey)
		} else {
			handler.Wait()
		}
	} else if method == cacheHandler.BulkSetCache {
		foreach, _ := utils.RenderVarToArrMap(handlerParam.Foreach, params)
		for _, item := range foreach {
			itemMap := make(map[string]interface{})
			itemMap[handlerParam.Item] = item
			dataKey := handler.GetCacheKey(handlerParam.Room, fieldList, itemMap)
			data := utils.RenderVar(handlerParam.Field, itemMap)
			result := common.Ok(data, "缓存中批量获取数据")
			ok := handler.Set(dataKey, *result, handlerParam.Second)
			if !ok {
				template.LogData("缓存设置失败" + dataKey)
			}
		}
		handler.Wait()

	} else {
		return common.NotOk("缓存设置：" + method + " 方法不存在")
	}

	r := common.Ok(nil, "处理参数成功")
	return r
}
