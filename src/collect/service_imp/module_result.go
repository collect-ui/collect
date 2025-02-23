package collect

import (
	common "github.com/collect-ui/collect/src/collect/common"
	"github.com/collect-ui/collect/src/collect/config"
	// template "test.mod/src/collect/template"
)

type ModuleResult interface {
	// Result 执行结果
	Result(template *collect.Template, t *TemplateService) *common.Result
	// HandlerData 处理数据
	HandlerData(template *collect.Template, handlerParam *collect.HandlerParam, t *TemplateService) *common.Result
}
