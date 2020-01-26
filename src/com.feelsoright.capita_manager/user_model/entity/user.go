package entity

import (
	account "com.feelsoright.capita_manager/capital_model/entity"
)

type User struct {
	Id       int
	Name     string
	UserName string
	Password string
	// 当前用户账户数据
	Account *account.Account
}

// 工厂模式
func NewUser(userName string, password string, account *account.Account) *User {
	return &User{
		UserName: userName,
		Password: password,
		Account:  account,
	}
}
