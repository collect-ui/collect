package collect

import (
	common "github.com/collect-ui/collect/src/collect/common"
	config "github.com/collect-ui/collect/src/collect/config"
)

type ServiceFlowService struct {
	BaseHandler
	BaseFlow
}

// Result 服务流程化
func (s *ServiceFlowService) Result(template *config.Template, ts *TemplateService) *common.Result {
	// 流程返回结果
	flowResult := s.Flow(template, ts, BaseHandlerNode)
	return flowResult
}
