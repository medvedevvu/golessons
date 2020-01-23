package shopcompetition

import (
	"errors"
	"fmt"
)

//Info
func Info() {
	fmt.Println("Best shop")
}

//AccountList := []Account{}
type AccountList []Account

//Register
func (ac AccountList) Register(username string) error {
	for _, item := range ac {
		if item.Name == username {
			return errors.New(fmt.Sprintf("пользователь %s",
				username, " уже зарегестрирован !!!"))
		}
	}
	ac = append(ac,
		Account{Name: username, Balance: 0, AccountType: AccountNormal})
	return nil
}

//AddBalance
func (ac AccountList) AddBalance(username string, sum float32) error {
	for _, item := range ac {
		if item.Name == username {
			item.Balance += sum
			return nil
		}
	}
	return errors.New(fmt.Sprintf("пользователь %s",
		username, " не зарегестрирован !!!"))
}

//Balance
func (ac AccountList) Balance(username string) (float32, error) {
	for _, item := range ac {
		if item.Name == username {
			return item.Balance, nil
		}
	}
	return 0, errors.New(fmt.Sprintf("пользователь %s",
		username, " не зарегестрирован !!!"))
}

//GetAccounts
func (ac AccountList) GetAccounts(sort AccountSortType) []Account {
	switch sort {
	case SortByName:
		{

		}
	case SortByNameReverse:
		{

		}
	case SortByBalance:
		{

		}
	}
	return ac
}
