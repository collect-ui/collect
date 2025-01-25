package collect

import (
	"fmt"
	common "github.com/collect-ui/collect/src/collect/common"
	config "github.com/collect-ui/collect/src/collect/config"
	utils "github.com/collect-ui/collect/src/collect/utils"
	"strings"
)

type ToList struct {
	BaseHandler
}

const Ancestors = "ancestors"

func (uf *ToList) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	params := template.GetParams()
	arr, _ := utils.RenderVarToArrMap(handlerParam.Foreach, params)
	childName := handlerParam.Children
	// 转树形结构
	ancestorsField := handlerParam.Ancestors
	target := treeToList(arr, childName, ancestorsField, []string{})
	for _, child := range target {
		delete(child, childName)
		if !utils.IsValueEmpty(ancestorsField) {
			ancestors := child[Ancestors].([]string)
			if len(ancestors) <= 0 {
				continue
			}
			ancestors = ancestors[:len(ancestors)-1]
			child[Ancestors] = ancestors
			if handlerParam.WithID {
				id := utils.RenderVar(ancestorsField, child).(string)
				parentId := strings.Join(ancestors, "#")

				if !utils.IsValueEmpty(parentId) {
					id = fmt.Sprint(parentId, "#", id)
					parentId = utils.GenerateShortUniqueID(parentId, 32)
				} else {
					parentId = "0"
				}
				child[handlerParam.Id] = utils.GenerateShortUniqueID(id, 32)
				child[handlerParam.Pid] = parentId

			}

		}
	}
	r := common.Ok(target, "处理参数成功")
	return r
}

func treeToList(arr []map[string]interface{}, children string, ancestorsField string, ancestors []string) []map[string]interface{} {
	r := make([]map[string]interface{}, 0)
	for _, aVal := range arr {
		// 如果有要先祖层级
		currentAncestors := make([]string, len(ancestors))
		if !utils.IsValueEmpty(ancestorsField) {
			copy(currentAncestors, ancestors)
			// 将当前节点的 ID 添加到 ancestors
			if name, ok := aVal[ancestorsField].(string); ok {
				currentAncestors = append(currentAncestors, name)
			}
			// 将 ancestors 字段添加到当前节点
			aVal[Ancestors] = currentAncestors
		}
		r = append(r, aVal)
		if _, hasChildren := aVal[children]; hasChildren {
			subTreeList, _ := utils.RenderVarToArrMap(children, aVal)
			subList := treeToList(subTreeList, children, ancestorsField, currentAncestors)
			r = append(r, subList...)
		}

	}
	return r
}
