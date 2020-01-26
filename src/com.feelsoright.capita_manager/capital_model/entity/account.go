package entity

type Account struct {
	Id int
	// 记录列表
	Details string
	// 初始余额
	Balance float64
	// 备注信息
	Note string
	// 金额
	Money float64
}

// 工厂模式
func NewAccount() *Account {
	return &Account{
		// 记录列表
		Details: "",
		// 初始余额
		Balance: 1000.0,
		// 备注信息
		Note:  "",
		Money: 0.0,
	}
}
