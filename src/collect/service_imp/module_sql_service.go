package collect

import (
	common "github.com/collect-ui/collect/src/collect/common"

	"database/sql"
	"fmt"
	handler_template "github.com/collect-ui/collect/src/collect/service_imp/handler_template"
	utils "github.com/collect-ui/collect/src/collect/utils"
	"github.com/demdxx/gocast"
	"log"
	text_template "text/template"

	config "github.com/collect-ui/collect/src/collect/config"
)

type SqlService struct {
	BaseHandler
}
type SqlResult struct {
	Data  []map[string]interface{}
	Error error
}

func runSQL(db *sql.DB, sql string, realValues []interface{}, result chan<- *SqlResult) {
	data, err := sqlToData(db, sql, realValues...)
	r := SqlResult{Data: data, Error: err}
	result <- &r
}
func (s *SqlService) Result(template *config.Template, ts *TemplateService) *common.Result {
	r := common.Result{}
	var err error
	var db *sql.DB
	if utils.IsValueEmpty(template.DataSource) {
		db, err = s.GetDatasource()
		if err != nil {
			msg := err.Error()
			return r.NotOk(msg)
		}
	} else {
		dataSource := utils.RenderVarOrValue(template.DataSource, template.GetParams())
		if utils.IsValueEmpty(dataSource) {
			db, err = s.GetDatasource()
		} else {
			db, err = s.GetOtherDatasource(dataSource.(string))
		}

		if err != nil {
			msg := err.Error()
			return r.NotOk(msg)
		}
	}

	//获取文件数据
	// fileData := template.GetFileData()
	params := template.GetParams()
	if template.Log {
		template.LogData("服务请求参数:")
		template.LogData(params)
	}
	if template.FileDataTpl == nil {
		template.LogData("data_file 不存在，请检查sql 文件路径")
	}

	// 生成执行SQL和参数
	sql, realValues := getSQLByTpl(template.FileDataTpl, params)
	// 执行SQL
	if template.Log {
		template.LogData("执行数据SQL:")
		template.LogData(sql)
		template.LogData("数据SQL参数:")
		template.LogData(realValues)
	}
	var maps []map[string]interface{}

	var count int64
	// 判断是否运行总数sql
	runCountStr := template.Count
	runCount := true
	if !utils.IsValueEmpty(runCountStr) {
		runCount = gocast.ToBool(utils.RenderVar(runCountStr, params))
	}
	if template.CountFileDataTpl != nil && runCount {
		// count 设置不分页
		params[template.Pagination] = false
		// 生成执行SQL和参数
		countSql, countRealValues := getSQLByTpl(template.CountFileDataTpl, params)
		// 执行SQL
		if template.Log {
			template.LogData("执行count SQL:")
			template.LogData(countSql)
			template.LogData("count SQL参数:")
			template.LogData(countRealValues)
		}

		//多线程执行开始 start
		resultChanList := make([]chan *SqlResult, 2)
		ch1 := make(chan *SqlResult)
		ch2 := make(chan *SqlResult)
		resultChanList[0] = ch1
		resultChanList[1] = ch2
		resultList := make([]*SqlResult, 2)
		// 多线程执行
		go runSQL(db, sql, realValues, resultChanList[0])
		go runSQL(db, countSql, countRealValues, resultChanList[1])
		for index, result := range resultChanList {
			resultData := <-result
			resultList[index] = resultData
		}
		// 执行结果
		maps = resultList[0].Data
		error := resultList[0].Error
		countMaps := resultList[1].Data
		countError := resultList[1].Error
		// 多线程执行结束 end
		// 如果多线程执行有问题，可以把上面代码屏蔽，执行下面3行
		//mapsSimple, error := sqlToData(db, sql, realValues...)
		//maps = mapsSimple
		//countMaps, countError := sqlToData(db, countSql, countRealValues...)
		if error != nil {
			return common.NotOk("执行sql报错:\n" + sql + "\n详情:" + error.Error())
		}
		// 执行结果
		if countError != nil {
			return common.NotOk("执行sql报错:\n" + countSql + "\n详情:" + countError.Error())
		}
		if len(countMaps) != 0 {
			countData := countMaps[0]
			var countValue interface{}

			if !utils.IsEmpty("count", countData) { // 获取小写的count
				countValue = utils.GetSafeData("count", countData)
			} else if !utils.IsEmpty("COUNT", countData) { //获取大写的count
				countValue = utils.GetSafeData("COUNT", countData)
			} else { //获取第一个key 的值
				countValue = utils.GetMapValues(countData)[0]

			}
			count = gocast.ToInt64(countValue)

		}

	} else { //只有sql直接运行
		mapsSimple, error := sqlToData(db, sql, realValues...)
		maps = mapsSimple
		if error != nil {
			return common.NotOk("执行sql报错:\n" + sql + "\n详情:" + error.Error())
		}
	}
	t := r.OkWithCount(maps, "执行成功", count)
	return t
}

func getSQLByTpl(tpl *text_template.Template, params map[string]interface{}) (string, []interface{}) {
	// 渲染第一次，将二级变量处理成一级变量。第一遍，根据模块转换
	t := handler_template.NewSqlTemplateByTpl(tpl)
	sql, sqlParams, _ := t.Content2Sql(params, true)
	// 渲染第二次,获取实际值，第二步根据模板转换的结果，重新渲染
	t = handler_template.NewSqlTemplate(sql)
	sql, _, realValues := t.Content2Sql(sqlParams, false)
	return sql, realValues
}

func getSQL(sqlData string, params map[string]interface{}) (string, []interface{}) {

	// 渲染第一次，将二级变量处理成一级变量
	t := handler_template.NewSqlTemplate(sqlData)
	sql, sqlParams, _ := t.Content2Sql(params, true)
	// 渲染第二次,获取实际值
	t = handler_template.NewSqlTemplate(sql)
	sql, _, realValues := t.Content2Sql(sqlParams, false)
	return sql, realValues
}

func sqlToData(db *sql.DB, sqlTemplate string, params ...any) ([]map[string]interface{}, error) {
	rows, err := db.Query(sqlTemplate, params...)
	if err != nil {
		fmt.Println("出错了", err)
		return nil, err
	}
	//转换成map
	maps := convertMaps(rows)
	return maps, nil
}

func convertMaps(rows *sql.Rows) []map[string]interface{} {

	colNames, _ := rows.Columns()
	colTypes, _ := rows.ColumnTypes()
	var cols = make([]interface{}, len(colNames))
	for i := 0; i < len(colNames); i++ {
		cols[i] = new(interface{})
	}
	var maps = make([]map[string]interface{}, 0)
	for rows.Next() {
		err := rows.Scan(cols...)
		if err != nil {
			log.Fatal(err.Error())
		}
		var rowMap = make(map[string]interface{})
		for i := 0; i < len(colNames); i++ {
			rowMap[colNames[i]] = convertRowByCol(colTypes[i].DatabaseTypeName(), *(cols[i].(*interface{})))
		}
		maps = append(maps, rowMap)
	}
	return maps

}
func convertRowByCol(colType string, value any) any {
	return utils.CastValue(value, colType)

}
