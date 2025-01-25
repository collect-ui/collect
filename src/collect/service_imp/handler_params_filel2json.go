package collect

import (
	"encoding/json"
	common "github.com/SelfDown/collect/src/collect/common"
	config "github.com/SelfDown/collect/src/collect/config"
	"io"
)

type File2Json struct {
	BaseHandler
}

func (uf *File2Json) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {

	fileBytes, err := io.ReadAll(ts.File)
	if err != nil {
		return common.NotOk("读取文件失败:" + err.Error())
	}
	var data map[string]interface{}
	json.Unmarshal([]byte(fileBytes), &data)
	if data == nil {
		var dataArr []map[string]interface{}
		json.Unmarshal([]byte(fileBytes), &dataArr)
		return common.Ok(dataArr, "参数处理成功")
	}
	r := common.Ok(data, "处理参数成功")
	return r
}
