package excel

import (
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"mime/multipart"
	"pmo-test4.yz-intelligence.com/base/utils/excel/model"
	"pmo-test4.yz-intelligence.com/base/utils/excel/rewrite"
	"pmo-test4.yz-intelligence.com/base/utils/optimization"
)

var Excel = new(_excel)

type _excel struct {
	err   error
	file  multipart.File
	excel *excelize.File
	rows  [][]string
}

const (
	DefaultSheet = "Sheet1"
	FirstRow     = "A1"
)

// ExportByInterface 导出xlsx 通过面向接口方式, 数据由后端处理
// Author [SliverHorn](https://github.com/SliverHorn)
func (e *_excel) ExportByInterface(export model.ExcelExport) error {
	a1 := export.A1Data()
	data := export.DataList()
	name := export.SheetName()
	filepath := export.FilePath()
	if filepath == "" || len(a1) == 0 || len(data) == 0 {
		return errors.New("文件名 or A1Data 数据 or DataList数据 不能为空! ")
	}
	excel := excelize.NewFile()
	if name != "" {
		excel.SetSheetName(DefaultSheet, name)
	} else {
		name = DefaultSheet
	}
	if err := excel.SetSheetRow(name, FirstRow, &a1); err != nil {
		return err
	}
	for i, d := range data {
		index := fmt.Sprintf("A%d", i+2)
		if err := excel.SetSheetRow(name, index, &d); err != nil {
			return err
		}
	}
	return excel.SaveAs(filepath)
}

// Parse 通过xlsx导入解析数据
// Author [Kevin-CC](https://github.com/icarus-go)
// formal
//  param header 文件头
//  param _import 接口对象 实现该方法要求的方法 model.ExcelImport
// return
//  param []map[string]interface{} 表数据
//  param error 错误信息
func (e *_excel) Parse(header *multipart.FileHeader, _import model.ExcelImport) ([]map[string]interface{}, error) {
	if e.file, e.err = header.Open(); e.err != nil {
		return nil, e.err
	}
	if e.excel, e.err = excelize.OpenReader(e.file); e.err != nil {
		return nil, e.err
	}
	if sheet := _import.SheetName(); sheet == "" {
		if e.rows, e.err = e.excel.GetRows(DefaultSheet); e.err != nil {
			return nil, e.err
		}
	} else {
		if e.rows, e.err = e.excel.GetRows(sheet); e.err != nil {
			return nil, e.err
		}
	}
	return _import.GetCreateData(e.rows), nil
}

//Rewrite 重写数据到某端，根据业务数据决定, 并且支持分批执行且是否忽略错误
// Author [Kevin-CC](https://github.com/icarus-go)
// formal
//  param entities 数据集
//  param rewrite 重写到哪个端所需执行的方法
//  param pageSize 分批执行多少条的页码
//  param isPage 是否分页
//  param isIgnore 分页后是否忽略每次回写数据时发生的错误
// return
//  param error 重写到某端的错误信息
func (e *_excel) Rewrite(entities []map[string]interface{}, rewrite rewrite.Data, pageSize int, isPage, isIgnore bool) error {
	if !isPage {
		return rewrite(&entities)
	}

	return optimization.Page.Each(len(entities), pageSize, func(start, end int) error {
		temp := entities[start:end]
		if err := rewrite(&temp); err != nil && !isIgnore {
			return err
		}
		return nil
	})
}
