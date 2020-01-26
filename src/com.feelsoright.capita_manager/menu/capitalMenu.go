package menu

import (
	account "com.feelsoright.capita_manager/capital_model/entity"
	"fmt"
)

// 退出
func exit(loop *bool) {
	exitKey := ""
	fmt.Println("您确定要退出吗？y/n")
	for {
		_, _ = fmt.Scan(&exitKey)
		if exitKey == "y" {
			fmt.Println("感谢您的使用，再见！")
			*loop = true
			return
		}
		if exitKey == "n" {
			return
		}
		fmt.Println("您的输入有误，请重新输入！")
	}
}

// 详情
func showDetails(account *account.Account) {
	fmt.Println("-------------------------------当前收支明细-----------------------------")
	if account.Details == "" {
		fmt.Println("当前没有收支明细...来一笔吧！")
	} else {
		fmt.Println("收支\t账户金额\t收支金额\t说  明")
		fmt.Println(account.Details)
	}
}

// 收入
func income(account *account.Account) {
	fmt.Println("您选择登记收入")
	fmt.Println("请输入金额")
	_, _ = fmt.Scan(&account.Money)
	fmt.Println("请输入说明")
	_, _ = fmt.Scan(&account.Note)
	account.Balance += account.Money
	account.Details += fmt.Sprintf("\n收入\t%v\t%v\t\t%v", account.Balance, account.Money, account.Note)
}

// 支出
func pay(account *account.Account) {
	fmt.Println("您选择登记支出")
	fmt.Println("请输入金额")
	_, _ = fmt.Scan(&account.Money)
	if account.Money > account.Balance {
		fmt.Println("账户余额不足")
		return
	}
	fmt.Println("请输入说明")
	_, _ = fmt.Scan(&account.Note)
	account.Balance -= account.Money
	account.Details += fmt.Sprintf("\n支出\t%v\t%v\t\t%v", account.Balance, account.Money, account.Note)
}
