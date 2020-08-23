/******
** @创建时间 : 2020/8/11 19:48
** @作者 : SongZhiBin
******/
package sql

import (
	"Happy/settings"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"strings"
)

// 这里使用sqlx进行链接
// 目前支持mysql和sqlserver
// 根据配置文件的driver配置

var dbInstantiate *sqlx.DB

// Init:初始化
func Init() error {
	var dns string
	var err error
	// 根据不同的driver设置不同的dns
	switch strings.ToUpper(settings.GetString("DB.Driver")) {
	case "MYSQL":
		dns = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
			settings.GetString("DB.User"),
			settings.GetString("DB.Password"),
			settings.GetString("DB.Host"),
			settings.GetInt("DB.Port"),
			settings.GetString("DB.DBName"))
		// 连接数据库
		dbInstantiate, err = sqlx.Connect("mysql", dns)
	case "MSSQL", "SQLSERVER":
		dns = fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;port=%d;encrypt=disable",
			settings.GetString("DB.Host"),
			settings.GetString("DB.DBName"),
			settings.GetString("DB.User"),
			settings.GetString("DB.Password"),
			settings.GetInt("DB.Port"))
		// 连接数据库
		dbInstantiate, err = sqlx.Connect("mssql", dns)
	}
	if err != nil {
		// 失败后记录日志 并返回错误
		zap.L().Error(fmt.Sprintf("No connect DB:%s", err))
		return err
	}
	// 设置对打连接数
	dbInstantiate.SetMaxOpenConns(settings.GetInt("DB.MaxOpenCons"))
	dbInstantiate.SetMaxIdleConns(settings.GetInt("DB.MaxIdleCons"))
	return nil
}

// Close:关闭时回收资源
func Close() {
	_ = dbInstantiate.Close()
}

// 封装的一些方法

// 增删改查

// SearchRow:查询单行 传入对应的结构体实例解析
// dbInstantiate:数据库对象
// sql: 查询的sql语句
// object:对应映射的结构体
// parameter:参数
func SearchRow(dbInstantiate *sqlx.DB, sqlString string, object interface{}, parameter ...interface{}) error {
	err := dbInstantiate.Get(object, sqlString, parameter...)
	if err != nil {
		zap.L().Error(fmt.Sprintf("SearchRow Error:%s", err))
		return err
	}
	return nil
}

// SearchAll:查询多行 传入对应的结构体切片实例解析
// dbInstantiate:数据库对象
// sql: 查询的sql语句
// objectSlice:对应映射的结构体切片
// parameter:参数
func SearchAll(dbInstantiate *sqlx.DB, sql string, objectSlice interface{}, parameter ...interface{}) error {
	err := dbInstantiate.Select(objectSlice, sql, parameter...)
	if err != nil {
		zap.L().Error("SearchAll Error", zap.Error(err))
		return err
	}
	return nil
}

// Exec:插入更新和删除
// dbInstantiate:数据库对象
// sql: 查询的sql语句
// parameter:参数
// 返回分别为
// eType:返回类型 0 为异常 1位插入 2 为增删
// res:根据eType的不同分别是插入的id值 或 影响行数
// error:错误信息
func Exec(dbInstantiate *sqlx.DB, sql string, parameter ...interface{}) (int, int, error) {
	res, err := dbInstantiate.Exec(sql, parameter...)
	if err != nil {
		zap.L().Error(fmt.Sprintf("Exec Error SQL%s", sql), zap.Error(err))
		return 0, 0, err
	}
	theID, _ := res.LastInsertId()
	if theID != 0 {
		// 证明是插入语句
		return 1, int(theID), nil
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		zap.L().Info(fmt.Sprintf("SQL Exec RowAffected is 0 SQL:%s,Param:%s", sql, parameter))
	}
	return 2, int(n), nil
}

// ExecAffair:插入更新和删除增加事务
// dbInstantiate:数据库对象
// sql: 查询的sql语句
// parameter:参数
// 返回分别为
// eType:返回类型 0 为异常 1位插入 2 为增删
// res:根据eType的不同分别是插入的id值 或 影响行数
// error:错误信息
func ExecAffair(dbInstantiate *sqlx.DB, sql string, parameter ...interface{}) (int, int, error) {
	tx, err := dbInstantiate.Beginx()
	if err != nil {
		zap.L().Error(fmt.Sprintf("Affair Begin Error SQL:%s Error:%s", sql, err))
		return 0, 0, err
	}
	// 注册异常处理
	defer func() {
		if p := recover(); p != nil {
			// 如果捕获到panic 则回滚
			_ = tx.Rollback()
			zap.L().Error(fmt.Sprintf("Panic Recover SQL:%s Painc%s", sql, p))
			return
		} else if err != nil {
			// 如果error不为空 则回滚
			_ = tx.Rollback()
			zap.L().Error(fmt.Sprintf("Exec Error SQL%s Error:%s", sql, err))
			return
		} else {
			// 最终 提交
			_ = tx.Commit()
			return
		}
	}()
	res, err := dbInstantiate.Exec(sql, parameter...)
	if err != nil {
		return 0, 0, err
	}
	theID, _ := res.LastInsertId()
	if theID != 0 {
		// 证明是插入语句
		return 1, int(theID), nil
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		// 表示没有影响任何数据 此次无效将其回滚
		zap.L().Info(fmt.Sprintf("SQL Exec RowAffected is 0 SQL:%s,Param:%s", sql, parameter))
		err = fmt.Errorf("SQL Exec RowAffected is 0 SQL:%s,Param:%s", sql, parameter)
		return 0, 0, err
	}
	return 2, int(n), nil
}

// NewPrepare:创建一个预处理对象
func NewPrepare(sql string) (*sqlx.Stmt, error) {
	return dbInstantiate.Preparex(sql)
}

// SSearchRow
// SdbInstantiate:预处理对象
// object:对应映射的结构体
// parameter:参数
func SSearchRow(dbInstantiate *sqlx.Stmt, object interface{}, parameter ...interface{}) error {
	err := dbInstantiate.Get(object, parameter...)
	if err != nil {
		zap.L().Error("SearchRow Error", zap.Error(err))
		return err
	}
	return nil
}

// SSearchAll
// dbInstantiate:预处理对象
// objectSlice:对应映射的结构体切片
// parameter:参数
func SSearchAll(dbInstantiate *sqlx.Stmt, objectSlice interface{}, parameter ...interface{}) error {
	err := dbInstantiate.Select(objectSlice, parameter...)
	if err != nil {
		zap.L().Error("SearchAll Error", zap.Error(err))
		return err
	}
	return nil
}

// SExec:插入更新和删除
// dbInstantiate:预处理对象
// parameter:参数
// 返回分别为
// eType:返回类型 0 为异常 1位插入 2 为增删
// res:根据eType的不同分别是插入的id值 或 影响行数
// error:错误信息
func SExec(dbInstantiate *sqlx.Stmt, parameter ...interface{}) (int, int, error) {
	res, err := dbInstantiate.Exec(parameter...)
	if err != nil {
		zap.L().Error("Exec Error Error", zap.Error(err))
		return 0, 0, err
	}
	theID, _ := res.LastInsertId()
	if theID != 0 {
		// 证明是插入语句
		return 1, int(theID), nil
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		zap.L().Info(fmt.Sprintf("SQL Exec RowAffected is 0 Param:%s", parameter))
	}
	return 2, int(n), nil
}

// SExecAffair:插入更新和删除增加事务
// 预处理不支持事务...
