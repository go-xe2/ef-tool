/*****************************************************************
* Copyright©,2020-2022, email: 279197148@qq.com
* Version: 1.0.0
* @Author: yangtxiang
* @Date: 2020-09-15 10:17
* Description:
*****************************************************************/

package app

type DbInfo struct {
	DbName string              `json:"db_name"`
	Tables map[string]*DbTable `json:"tables"`
}

func NewDbInfo() *DbInfo {
	return &DbInfo{
		Tables: make(map[string]*DbTable),
	}
}

func (p *DbInfo) AddTable(tb *DbTable) {
	p.Tables[tb.TableName] = tb
}

func (p *DbInfo) AllTables() map[string]*DbTable {
	return p.Tables
}

func (p *DbInfo) HasTable(tbName string) bool {
	if _, ok := p.Tables[tbName]; ok {
		return true
	}
	return false
}
