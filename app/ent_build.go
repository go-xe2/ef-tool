/*****************************************************************
* Copyright©,2020-2022, email: 279197148@qq.com
* Version: 1.0.0
* @Author: yangtxiang
* @Date: 2020-09-15 10:54
* Description:
*****************************************************************/

package app

import (
	"fmt"
	"github.com/go-xe2/x/type/t"
	"github.com/go-xe2/x/type/xstring"
	"io"
	"regexp"
	"strings"
	"time"
)

func MakeFileName(prefix string, tbName string) string {
	szFileName := tbName
	if xstring.StartWith(tbName, prefix) {
		szFileName = tbName[len(prefix):]
	}
	szFileName = xstring.Camel2UnderScore(szFileName, "_")
	// 冯
	return szFileName + "_entity.go"
}

func writeString(w io.Writer, str string) {
	if _, err := io.WriteString(w, str); err != nil {
		panic(err)
	}
}

var dbType2EntTypeNames = map[string]string{
	"varchar":   "EFString",
	"char":      "EFString",
	"text":      "EFString",
	"longtext":  "EFString",
	"tinytext":  "EFString",
	"blog":      "EFString",
	"tinyblob":  "EFString",
	"longblob":  "EFString",
	"varbinary": "EFBinary",
	"binary":    "EFBinary",
	"int":       "EFInt",
	"integer":   "EFInt",
	"tinyint":   "EFInt16",
	"smallint":  "EFInt16",
	"bigint":    "EFInt64",
	"bit":       "EFByte",
	"double":    "EFDouble",
	"real":      "EFFloat",
	"numeric":   "EFDouble",
	"decimal":   "EFDouble",
	"date":      "EFDate",
	"datetime":  "EFDate",
	"year":      "EFDate",
	"timestamp": "EFInt64",
}

var dbType2goTypeNames = map[string]string{
	"varchar":   "string",
	"char":      "string",
	"text":      "string",
	"longtext":  "string",
	"tinytext":  "string",
	"blog":      "string",
	"tinyblob":  "string",
	"longblob":  "string",
	"varbinary": "[]byte",
	"binary":    "[]byte",
	"int":       "int",
	"integer":   "int",
	"tinyint":   "int16",
	"smallint":  "int16",
	"bigint":    "int64",
	"bit":       "byte",
	"double":    "float64",
	"real":      "float32",
	"numeric":   "float64",
	"decimal":   "float64",
	"date":      "time.Time",
	"datetime":  "time.Time",
	"year":      "time.Time",
	"timestamp": "int64",
}

var s time.Time

var dbType2EntTypeDefNames = map[string]string{
	"varchar":   "varchar",
	"char":      "char",
	"text":      "text",
	"longtext":  "text",
	"tinytext":  "text",
	"blog":      "blob",
	"tinyblob":  "blob",
	"longblob":  "blob",
	"varbinary": "varbinary",
	"binary":    "varbinary",
	"int":       "int",
	"integer":   "int",
	"tinyint":   "tinyint",
	"smallint":  "tinyint",
	"bigint":    "bigint",
	"bit":       "bit",
	"double":    "double",
	"real":      "double",
	"numeric":   "decimal",
	"decimal":   "decimal",
	"date":      "datetime",
	"datetime":  "datetime",
	"year":      "datetime",
	"timestamp": "timestamp",
}

func dbType2EntDbTypeAttrString(dbType string, size int, nPresision, nScale int, allowNull bool, defValue string, increment bool) string {
	szType, ok := dbType2EntTypeDefNames[dbType]
	if !ok {
		szType = "varchar"
	}
	attr := "@dbType(type=" + szType
	if szType == "varchar" || szType == "char" {
		attr = attr + ", size=" + t.String(size)
	}
	if szType == "decimal" {
		attr = attr + ",size=" + t.String(nPresision)
		attr = attr + ",decimal=" + t.String(nScale)
	}
	if !allowNull {
		attr = attr + ",allowNull=false"
	}
	if defValue != "" {
		attr = attr + ",default=" + defValue
	}
	if increment {
		attr = attr + ",increment=true"
	}
	attr += ")"
	return attr
}

func dbType2EntColType(dbType string) string {
	if s, ok := dbType2EntTypeNames[dbType]; ok {
		return s
	}
	return "EFString"
}

func columnName2StructFieldName(clName string, clPrefix string) string {
	var result = clName
	if clPrefix != "" {
		clNameReg := regexp.MustCompile(`^(` + clPrefix + ")(.+?)$")
		items := clNameReg.FindStringSubmatch(clName)
		if len(items) == 3 {
			result = items[2]
		}
	}
	result = xstring.UnderScore2Camel(result, "_")
	result = xstring.UcFirst(result)
	return result
}

func dbType2goType(dbType string) string {
	if s, ok := dbType2goTypeNames[dbType]; ok {
		return s
	}
	return "string"
}

func column2DefName(cl *TableColumn, clPrefix string) (name string, alias string) {
	name = columnName2StructFieldName(cl.Name, clPrefix)
	alias = xstring.Camel2UnderScore(name, "_")
	if alias == cl.Name {
		alias = ""
	}
	return name, alias
}

func buildColDefineString(cl *TableColumn, clPrefix string) string {
	// 处理字段名
	clName := columnName2StructFieldName(cl.Name, clPrefix)
	// 数据类型转实体字段数据类型
	clType := dbType2EntColType(cl.DataType)
	result := fmt.Sprintf("%s %s", clName, clType)
	fdAlias := xstring.Camel2UnderScore(clName, "_")
	if fdAlias == cl.Name {
		fdAlias = ""
	}
	if cl.ColumnKey == "PRI" {
		result += "`ef:\"@field(name='" + cl.Name + "',primary=true"
	} else {
		result += "`ef:\"@field(name='" + cl.Name + "'"
	}
	if fdAlias != "" {
		result += ",alias='" + fdAlias + "'"
	}
	result += ")"
	dbTypeStr := dbType2EntDbTypeAttrString(cl.DataType, cl.Size, cl.NPrecision, cl.NScale, cl.IsNullAble, cl.Default, cl.AutoIncrement)
	if dbTypeStr != "" {
		result += ";" + dbTypeStr
	}
	result += "\"`"
	return result
}

func buildRecordStructFieldStr(cl *TableColumn, clPrefix string) string {
	clName := columnName2StructFieldName(cl.Name, clPrefix)
	clType := dbType2goType(cl.DataType)
	clAlias := xstring.Camel2UnderScore(clName, "_")
	szResult := fmt.Sprintf("%s %s", clName, clType)
	if clAlias != "" {
		szResult += "\t`json:\"" + clAlias + ",omitempty\"`"
	}
	return szResult
}

var entFieldTryGetValueFuncNames = map[string]string{
	"EFString": "TryStr()",
	"EFBool":   "TryBool()",
	"EFByte":   "TryByte()",
	"EFInt":    "TryInt()",
	"EFInt8":   "TryInt8()",
	"EFInt16":  "TryInt16()",
	"EFInt32":  "TryInt32()",
	"EFInt64":  "TryInt64()",
	"EFUint":   "TryUint()",
	"EFUint8":  "TryUint8()",
	"EFUint16": "TryUint16()",
	"EFUint32": "TryUint32()",
	"EFUint64": "TryUint64()",
	"EFFloat":  "TryFloat()",
	"EFDouble": "TryDouble()",
	"EFBinary": "TryBytes()",
	"EFDate":   "TryDate()",
}

func buildRecordFieldAssignValue(cl *TableColumn, owner string, rec string, clPrefix string) string {
	clType := dbType2EntColType(cl.DataType)
	fdName := columnName2StructFieldName(cl.Name, clPrefix)
	valFun, ok := entFieldTryGetValueFuncNames[clType]
	if !ok {
		valFun = "TryGetVal()"
	}
	result := fmt.Sprintf("\tif v, ok := %s.%s.%s; ok {\n", owner, fdName, valFun)
	result += fmt.Sprintf("\t\t%s.%s = v\n", rec, fdName)
	result += "\t}\n"
	return result
}

// 输出表实体定义
func BuildEntity(out io.Writer, pkName string, tbPrefix, clPrefix string, tb *DbTable) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	// 处理实体名称字符串
	tbName := tb.TableName

	if xstring.StartWith(tbName, tbPrefix) {
		tbName = tbName[len(tbPrefix):]
	}
	var dbTableName = tbName
	tbName = xstring.UnderScore2Camel(tbName, "_")
	tbName = xstring.UcFirst(tbName)

	entName := tbName + "Entity"
	writeString(out, fmt.Sprintf("package %s\n", pkName))
	writeString(out, "import (\n")
	writeString(out, "\t. \"github.com/go-xe2/x/xf/ef/xq/xentity\"\n")
	writeString(out, "\t. \"github.com/go-xe2/x/xf/ef/xqi\"\n")
	writeString(out, ")\n")
	writeString(out, "\n")
	columns := tb.Columns

	// 生成数据结构定义
	recName := tbName + "Rec"
	if tb.TableComment != "" {
		writeString(out, "// "+tb.TableComment+"\n")
	}
	writeString(out, fmt.Sprintf("type %s struct {\n", recName))
	// 生成字段定义
	for _, cl := range columns {
		s := buildRecordStructFieldStr(cl, clPrefix)
		if cl.Comment != "" {
			writeString(out, "\t"+s+"\t// "+cl.Comment+"\n")
		} else {
			writeString(out, "\t"+s+"\n")
		}
	}
	writeString(out, "}\n")
	writeString(out, "\n")

	writeString(out, "// table:"+tb.TableName+"\n")
	if tb.TableComment != "" {
		writeString(out, "// "+tb.TableComment+"\n")
	}
	// 开始写实体定义
	writeString(out, fmt.Sprintf("type %s struct {\n", entName))
	// 写实体字段列表
	writeString(out, "\t*TEntity\n")
	writeString(out, "\tnameFields map[string]EntField\n")

	for _, cl := range columns {
		str := buildColDefineString(cl, clPrefix)
		if cl.Comment != "" {
			writeString(out, "\t"+str+"\t// "+cl.Comment+"\n")
		} else {
			writeString(out, "\t"+str+"\n")
		}
	}
	// 实体定义结束
	writeString(out, "}\n")
	writeString(out, "\n")

	// 生成实体工厂类型
	var entClsName = entName + "Class"
	var dbTableAlias = ""
	dbTableAlias = xstring.Camel2UnderScore(tbName, "_")
	items := strings.Split(dbTableAlias, "_")
	if len(items) < 2 {
		dbTableAlias = tbName[:2]
		if len(tbName) > 2 {
			dbTableAlias = dbTableAlias + string(tbName[len(tbName)-1])
		}
		dbTableAlias = xstring.LowerCase(dbTableAlias)
	} else {
		tmp := ""
		for _, s := range items {
			tmp += string(s[0])
		}
		dbTableAlias = xstring.LowerCase(tmp)
	}
	writeString(out, fmt.Sprintf("var %s = ClassOfEntity(func() Entity {\n", entClsName))
	writeString(out, fmt.Sprintf("\tvar inst Entity = new(%s)\n", entName))
	writeString(out, "\treturn inst\n")
	writeString(out, fmt.Sprintf("}, []XqAttribute{MakeEntityAttr(\"%s\", \"%s\")})\n", dbTableName, dbTableAlias))
	writeString(out, "\n")

	// 写实体方法Implement
	writeString(out, "// 设置实体继承类实例, 供构架内部调用\n")
	writeString(out, fmt.Sprintf("func (ent *%s) Implement(supper interface{}) {\n", entName))
	writeString(out, fmt.Sprintf("\tif v, ok := supper.(*TEntity); ok {\n"))
	writeString(out, fmt.Sprintf("\t\tent.TEntity = v\n"))
	writeString(out, "\t}\n")
	writeString(out, "}\n")
	writeString(out, "\n")

	// 写实体方法Supper
	writeString(out, "// 继承的父类\n")
	writeString(out, fmt.Sprintf("func (ent *%s) Supper() Entity {\n", entName))
	writeString(out, "\treturn ent.TEntity\n")
	writeString(out, "}\n")
	writeString(out, "\n")

	// 写实体方法Constructor
	writeString(out, "// 实体构造方法\n")
	writeString(out, fmt.Sprintf("func (ent *%s) Constructor(attrs []XqAttribute, inherited ...interface{}) interface{} {\n", entName))
	writeString(out, "\tent.Supper().Constructor(attrs, inherited...)\n")
	writeString(out, "\treturn ent\n")
	writeString(out, "}\n")
	writeString(out, "\n")

	// 写实体方法String
	writeString(out, fmt.Sprintf("func (ent *%s) String() string {\n", entName))
	writeString(out, fmt.Sprintf("\treturn \"%s\"\n", tbName))
	writeString(out, "}\n")
	writeString(out, "\n")

	// 写实体Record方法
	writeString(out, fmt.Sprintf("func (ent *%s) Record() *%s {\n", entName, recName))
	writeString(out, "\tif !ent.IsOpen() {\n")
	writeString(out, fmt.Sprintf("\t\treturn &%s{}\n", recName))
	writeString(out, "\t}\n")

	writeString(out, fmt.Sprintf("\tresult := &%s{}\n", recName))

	// 写字段赋值
	for _, cl := range columns {
		str := buildRecordFieldAssignValue(cl, "ent", "result", clPrefix)
		writeString(out, str)
	}

	writeString(out, "\treturn result\n")

	writeString(out, "}\n")

	// 写实体NameFields方法
	writeString(out, fmt.Sprintf("func (ent *%s) NameFields() map[string]EntField {\n", entName))
	writeString(out, fmt.Sprintf("\tif ent.nameFields == nil {\n"))
	writeString(out, "\t\tent.nameFields = map[string]EntField{\n")
	// 写字段赋值
	for _, cl := range columns {
		fieldDef, fdAlias := column2DefName(cl, clPrefix)
		fieldName := fdAlias
		if fieldName == "" {
			fieldName = fieldDef
		}
		writeString(out, fmt.Sprintf("\t\t\t\"%s\": ent.%s,\n", fieldName, fieldDef))
	}
	writeString(out, "\t\t}\n")
	writeString(out, "\t}\n")
	writeString(out, "\treturn ent.nameFields\n")
	writeString(out, "}\n")
	return nil
}

// 生成所有实体
func BuildEntities(getOut func(tbName string) (io.WriteCloser, error), pkName string, tbPrefix string, clPrefix string, tables *DbInfo) error {
	items := tables.AllTables()
	for k, tb := range items {
		io, err := getOut(k)
		if err != nil {
			return err
		}
		if err := BuildEntity(io, pkName, tbPrefix, clPrefix, tb); err != nil {
			_ = io.Close()
			return err
		}
		_ = io.Close()
	}
	return nil
}
