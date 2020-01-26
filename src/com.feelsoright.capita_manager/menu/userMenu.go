package menu

import (
	account "com.feelsoright.capita_manager/capital_model/entity"
	user "com.feelsoright.capita_manager/user_model/entity"
	"fmt"
)

// 当前存储的用户列表数据
var userArray []*user.User

// 当前登录的用户
var loginArray []string

// 首页菜单
func IndexMenu() {
	// 输入的选项
	inputKey := ""
	// 是否退出
	loop := false
	for {
		fmt.Println("-------------------------------个人记账本-----------------------------")
		fmt.Println("                              1: 用户登录")
		fmt.Println("                              2: 创建用户")
		fmt.Println("                              3: 退出记账")
		fmt.Println("请输入《1-3》")
		_, _ = fmt.Scan(&inputKey)
		switch inputKey {
		case "1":
			fmt.Println("您选择用户登录")
			login()
		case "2":
			fmt.Println("您选择创建用户")
			register()
		case "3":
			exit(&loop)
		default:
			fmt.Println("您的输入无效，请重新输入")
		}
		if loop {
			break
		}
	}
}

func login() {
	userName := ""
	password := ""
	fmt.Println("请输入用户名")
	_, _ = fmt.Scan(&userName)
	fmt.Println("请输入密码")
	_, _ = fmt.Scan(&password)
	localUser := findUserByName(userName)
	if localUser == nil || localUser.Password != password {
		fmt.Println("用户名或密码错误")
		return
	}
	userMenu(localUser.Account, localUser)
}

// 根据用户名查询数据
func findUserByName(userName string) *user.User {
	for _, value := range userArray {
		if userName == value.UserName {
			return value
		}
	}
	return nil
}

// 用户登录后的菜单
func userMenu(account *account.Account, user *user.User) {
	// 输入的选项
	inputKey := ""
	// 是否退出
	loop := false
	for {
		fmt.Println("-------------------------------个人记账本-----------------------------")
		fmt.Println("                              1: 收支明细")
		fmt.Println("                              2: 登记收入")
		fmt.Println("                              3: 登记支出")
		fmt.Println("                              4: 退出登录")
		fmt.Println("请输入《1-4》")
		_, _ = fmt.Scan(&inputKey)
		switch inputKey {
		case "1":
			showDetails(account)
		case "2":
			income(account)
		case "3":
			pay(account)
		case "4":
			logout(user, &loop)
		default:
			fmt.Println("您的输入无效，请重新输入")
		}
		if loop {
			break
		}
	}
}

//登出
func logout(user *user.User, flag *bool) {
	exitKey := ""
	fmt.Println("您确定要退出登录吗？y/n")
	for {
		_, _ = fmt.Scan(&exitKey)
		if exitKey == "y" {
			for index, value := range loginArray {
				if value == user.UserName {
					loginArray = append(loginArray[:index], loginArray[index+1:]...)
					break
				}
			}
			fmt.Println("当前用户已登出，请重新登录...")
			*flag = true
			return
		}
		if exitKey == "n" {
			return
		}
		fmt.Println("您的输入有误，请重新输入！")
	}
}

// 注册密码
func register() {
	userName := ""
	for {
		fmt.Println("请输入用户名")
		_, _ = fmt.Scan(&userName)
		localUser := findUserByName(userName)
		if localUser == nil {
			break
		} else {
			fmt.Println("该用户名已存在！")
		}
	}
	password := setPassword()
	userArray = append(userArray, user.NewUser(userName, password, account.NewAccount()))
}

// 用户设定密码
func setPassword() string {
	password := ""
	passwords := ""
	fmt.Println("请输入密码")
	_, _ = fmt.Scan(&password)
	fmt.Println("请输入确认密码")
	_, _ = fmt.Scan(&passwords)
	if password != passwords {
		fmt.Println("当前输入密码不匹配，请重新输入...")
		setPassword()
	}
	return password
}
