package sql

import (
	"Happy/model/model"
	"Happy/pkg/snowflake"

	"go.uber.org/zap"
)

// 处理用户相关的数据库操作

// IsExist:判断用户是否存在
func IsExist(user string) error {
	sqlString := `SELECT COUNT(*) FROM user WHERE username = ?`
	var count int
	// 查询报错
	if err := SearchRow(dbInstantiate, sqlString, &count, user); err != nil {
		return err
	}
	// 判断是否大于0
	if count > 0 {
		// 证明用户已经存在
		return model.CodeUserExist.Err()
	}
	return nil
}

// IsUserValid:判断用户是否有效
func IsUserValid(user, password string) (*model.User, error) {
	u := new(model.User)
	sqlString := `SELECT user_id,username FROM user Where username = ? and password = ?`
	if err := SearchRow(dbInstantiate, sqlString, u, user, password); err != nil {
		return nil, err
	}
	if u.UID == 0 {
		return nil, model.CodeInvalidPassword.Err()
	}
	return u, nil
}

// InsertUser:插入数据库完成注册
func InsertUser(username, password string) bool {
	// 1.构建sql语句
	sqlString := `INSERT INTO user(user_id,username,password) VALUES(?,?,?)`
	// 2.获取全局id
	userId, err := snowflake.GetID()
	if err != nil {
		zap.L().Error("GetID Error", zap.Error(err))
		return false
	}
	// 获取密文的密码
	pwd := GetEncrypt(password)
	_, _, err = Exec(dbInstantiate, sqlString, userId, username, pwd)
	if err != nil {
		zap.L().Error("Exec Error", zap.Error(err))
		return false
	}
	return true
}

// GetUserName:根据用户ID获取用户name
func GetUserName(uid int64) (*model.User, error) {
	sqlString := `SELECT username FROM user Where user_id = ?`
	u := new(model.User)
	err := SearchRow(dbInstantiate, sqlString, u, uid)
	return u, err
}
