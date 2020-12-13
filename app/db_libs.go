/*****************************************************************
* Copyright©,2020-2022, email: 279197148@qq.com
* Version: 1.0.0
* @Author: yangtxiang
* @Date: 2020-09-15 10:18
* Description:
*****************************************************************/

package app

import (
	"database/sql"
	"github.com/go-xe2/x/type/t"
)

func LoadDbInfo(conn *sql.DB, dbName string) (*DbInfo, error) {
	result := NewDbInfo()
	rows, err := conn.Query("select table_name, table_comment from information_schema.tables where table_schema=?", dbName)
	if err != nil {
		return nil, err
	}
	var szName string
	var szComment string
	tables := make([]*DbTable, 0)
	for rows.Next() {
		if err := rows.Scan(&szName, &szComment); err != nil {
			return nil, err
		}
		tb := NewTable()
		tb.TableName = szName
		tb.TableComment = szComment
		tables = append(tables, tb)
	}
	// 加载表字段
	//select column_name, column_default, is_nullable,data_type, column_type,column_key, character_maximum_length, numeric_precision, numeric_scale,column_comment from information_schema.COLUMNS where table_name = 'mny_users' and table_schema='mny-users'  order by ordinal_position;
	for _, tb := range tables {
		rows, err := conn.Query("select column_name, column_default, is_nullable,data_type,column_key, character_maximum_length, numeric_precision, numeric_scale,column_comment,extra from information_schema.COLUMNS where table_name = ? and table_schema=?  order by ordinal_position", tb.TableName, dbName)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			cl := &TableColumn{}
			isNullAble := "YES"
			var nprecision interface{}
			var nscale interface{}
			var nSize interface{}
			var def interface{}
			var extra string
			if err := rows.Scan(&cl.Name, &def, &isNullAble, &cl.DataType, &cl.ColumnKey, &nSize, &nprecision, &nscale, &cl.Comment, &extra); err != nil {
				return nil, err
			}

			if def == nil {
				cl.Default = ""
			} else {
				cl.Default = string(def.([]byte))
			}
			if isNullAble == "NO" {
				cl.IsNullAble = false
			} else {
				cl.IsNullAble = true
			}
			if nprecision != nil {
				cl.NPrecision = t.Int(nprecision)
			}
			if nscale != nil {
				cl.NScale = t.Int(nscale)
			}
			if nSize != nil {
				cl.Size = t.Int(nSize)
			}
			cl.AutoIncrement = extra == "auto_increment"
			tb.AddColumn(cl)
		}
		result.AddTable(tb)
	}
	return result, nil
}
