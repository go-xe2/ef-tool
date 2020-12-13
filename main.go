/*****************************************************************
* Copyright©,2020-2022, email: 279197148@qq.com
* Version: 1.0.0
* @Author: yangtxiang
* @Date: 2020-09-15 09:55
* Description:
*****************************************************************/

package main

import (
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xe2/x/os/xfile"
	"github.com/go-xe2/x/os/xlog"
	"github.com/go-xe2/x/xf/ef/xq/xdatabase"
	"io"
	"mnyun.net/ef-tool/app"
	"os"
)

var (
	argDbHost       = flag.String("h", "127.0.0.1", "mysql数据库地址,默认127.0.0.1")
	argDbPort       = flag.Int("hp", 3306, "mysql数据端口, 默认3306")
	argDbUser       = flag.String("u", "root", "mysql账号名, 默认root")
	argDbPwd        = flag.String("p", "", "mysql数据库密码, 默认为空")
	argDbName       = flag.String("db", "", "数据库名")
	argTable        = flag.String("t", "", "要生成实体的表名")
	argTablePrefix  = flag.String("tp", "", "数据库表名前缀")
	argColumnPrefix = flag.String("cp", ".?_", "数据库字段名前缀正则表达式,默认.?_")
	argOutDir       = flag.String("o", "", "实体模型文件输出路径, 默认在当前目录下的entities目录")
	argPkName       = flag.String("pk", "entities", "实体包名, 默认entities")
)

func main() {
	flag.Parse()
	// 初始化数据库
	if *argDbName == "" {
		xlog.Info("数据库名不能为空")
		return
	}
	cfg := xdatabase.NewMysqlConfig6(*argDbHost, *argDbPort, *argDbUser, *argDbPwd, *argDbName, "tcp")
	db := xdatabase.DB()
	if err := db.Connection(cfg); err != nil {
		xlog.Info(err)
		return
	}
	dbTables, err := app.LoadDbInfo(db.GetConn().GetQueryDB(), *argDbName)
	if err != nil {
		xlog.Info(err)
		return
	}
	dir := "./entities"
	if *argOutDir != "" {
		dir = *argOutDir
	}
	if !xfile.Exists(dir) {
		if err := xfile.Mkdir(dir); err != nil {
			xlog.Info(err)
			return
		}
	}
	if *argTable != "" {
		// 生成表个表实体
		tb, ok := dbTables.AllTables()[*argTable]
		if !ok {
			xlog.Infof("库中不存在表%s", *argTable)
			return
		}
		szFileName := xfile.Join(dir, app.MakeFileName(*argTablePrefix, tb.TableName))
		f, err := xfile.OpenWithFlag(szFileName, os.O_CREATE|os.O_TRUNC|os.O_RDWR)
		if err != nil {
			xlog.Info("创建文件出错:", err)
			return
		}
		defer f.Close()
		if err := app.BuildEntity(f, *argPkName, *argTablePrefix, *argColumnPrefix, tb); err != nil {
			xlog.Info("生成实体出错:", err)
			return
		}
		xlog.Info("生成实体成功，输出文件:", szFileName)
		return
	}
	// 生成整个库的实体
	if err := app.BuildEntities(func(tbName string) (io.WriteCloser, error) {
		szFileName := app.MakeFileName(*argTablePrefix, tbName)
		f, err := xfile.OpenWithFlag(xfile.Join(dir, szFileName), os.O_CREATE|os.O_TRUNC|os.O_RDWR)
		return f, err
	}, *argPkName, *argTablePrefix, *argColumnPrefix, dbTables); err != nil {
		xlog.Info("生成实体出错:", err)
	} else {
		xlog.Info("生成所有实体成功， 输出路径:", dir)
	}
}
