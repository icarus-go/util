package model

type ExcelExport interface {
	A1Data() []string          // 表单
	FilePath() string          // 文件保存路径
	DataList() [][]interface{} // 数据
	SheetName() string         // 设置数据保存到哪个Sheet, 如果不指定,就是保存到默认的Sheet1
}

type ExcelImport interface {
	GetCreateData(rows [][]string) []map[string]interface{} // 获取批量插入的数据
	GetTableName() string                                   // 获取数据插入的表名
	SheetName() string                                      // 设置读取哪个Sheet, 如果不指定,则默认读取的是Sheet1
}
