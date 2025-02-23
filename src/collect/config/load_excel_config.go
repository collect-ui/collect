package collect

import (
	"encoding/json"
	utils "github.com/collect-ui/collect/src/collect/utils"
)

/*
* 加载data_file 文件内容，处理require 文件的引用
* 将服务的文件路径转换成文件内容
 */
func (t *PluginLoader) LoadExcelConfig(config Plugin, template *Template, routerAll *RouterAll) {
	serviceList := routerAll.GetRegisterServices()
	// 循环服务,将文件路径对应的内容
	for _, service := range serviceList {
		if !utils.IsValueEmpty(service.ExcelConfig) {
			var config ExcelConfig
			json.Unmarshal([]byte(service.ExcelConfigContent), &config)
			for i, sheet := range config.Sheets {
				// 将标题编译成模板
				if !utils.IsValueEmpty(sheet.Title) {
					tpl, err := _load_template(sheet.Title)
					if err != nil {
						template.LogData(err)
						continue
					}
					config.Sheets[i].TitleTpl = tpl
				}
				for j, secondField := range sheet.Fields {
					// 如果没有配置模板，这跳过
					if utils.IsValueEmpty(secondField.Template) {
						continue
					}
					tplName, err := _load_template(secondField.Template)
					if err != nil {
						template.LogData(err)
						continue
					}
					// 将field 编译成模板
					config.Sheets[i].Fields[j].TemplateTpl = tplName

				}
			}
			service.ExcelConfigData = &config

		}
	}

}
