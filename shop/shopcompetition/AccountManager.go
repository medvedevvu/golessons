package shopcompetition

import (
	"fmt"
	ssort "sort"
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
			return fmt.Errorf("пользователь %s уже зарегестрирован",
				username)
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
	return fmt.Errorf("пользователь %s не зарегестрирован ",
		username)
}

//Balance
func (ac AccountList) Balance(username string) (float32, error) {
	for _, item := range ac {
		if item.Name == username {
			return item.Balance, nil
		}
	}
	return 0, fmt.Errorf("пользователь %s не зарегестрирован ",
		username)
}

//GetAccounts
func (ac AccountList) GetAccounts(sort AccountSortType) AccountList {
	sortedAccount := AccountList{}
	switch sort {
	case SortByName, SortByNameReverse:
		vkyes := []string{}
		for _, elem := range ac {
			vkyes = append(vkyes, elem.Name)
		}
		if sort == SortByName {
			ssort.Strings(vkyes)
		} else {
			ssort.Sort(ssort.Reverse(ssort.StringSlice(vkyes)))
		}
		for i := 0; i < len(vkyes); i++ {
			for _, j := range ac {
				if j.Name == vkyes[i] {
					sortedAccount = append(sortedAccount, j)
					break
				}
			}
		}
	case SortByBalance:
		vkyes := []float64{}
		for _, elem := range ac {
			vkyes = append(vkyes, float64(elem.Balance))
		}
		ssort.Sort(ssort.Float64Slice(vkyes))
		for i := 0; i < len(vkyes); i++ {
			for _, j := range ac {
				if float64(j.Balance) == vkyes[i] {
					sortedAccount = append(sortedAccount, j)
					break
				}
			}
		}
	default:
		fmt.Printf("не известный ключ сортировки %d\n", sort)
	}
	return sortedAccount
}
