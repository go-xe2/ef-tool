/*****************************************************************
* Copyright©,2020-2022, email: 279197148@qq.com
* Version: 1.0.0
* @Author: yangtxiang
* @Date: 2020-09-15 10:08
* Description:
*****************************************************************/

package app

// 数据库表字段
type TableColumn struct {
	// 字段名
	Name string `json:"name"`
	// 字段默认值
	Default string `json:"default"`
	// 是否允许为空
	IsNullAble bool `json:"is_nullable"`
	// 字段数据类型
	DataType string `json:"data_type"`
	// 主外键
	ColumnKey string `json:"column_key"`
	// 字段大小
	Size int `json:"size"`
	// decimal长度
	NPrecision int `json:"n_precision"`
	// decimal小数点位数
	NScale int `json:"n_scale"`
	// 字段说明
	Comment       string `json:"comment"`
	AutoIncrement bool   `json:"auto_increment"`
}

type DbTable struct {
	TableName    string `json:"table_name"`
	TableComment string `json:"table_comment"`
	Columns      []*TableColumn
}

func NewTable() *DbTable {
	return &DbTable{
		TableName:    "",
		TableComment: "",
		Columns:      make([]*TableColumn, 0),
	}
}

func (p *DbTable) AddColumn(cl *TableColumn) {
	p.Columns = append(p.Columns, cl)
}
