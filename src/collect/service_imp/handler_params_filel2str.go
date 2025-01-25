package collect

import (
	common "collect/src/collect/common"
	config "collect/src/collect/config"
	"io"
)

type File2Str struct {
	BaseHandler
}

func (uf *File2Str) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {

	fileBytes, err := io.ReadAll(ts.File)
	if err != nil {
		return common.NotOk("读取文件失败:" + err.Error())
	}
	r := common.Ok(string(fileBytes), "处理参数成功")
	return r
}
