package shop_competition

import (
	"fmt"
	sorting "sort"
	"strings"
	"time"
)

// Register - регистрация пользователя
func (accList *AccountsList) Register(username string, accounttype AccountType) error {
	done := make(chan struct{})
	errmsg := make(chan error, 1)
	timer := time.NewTimer(time.Second)
	go func() {
		defer close(done)
		defer accList.Unlock()

		accList.Lock()
		if len(strings.Trim(username, "")) == 0 {
			errmsg <- fmt.Errorf("username %s пустое ", username)
			return
		}
		_, ok := (*accList).Accounts[username]
		if ok {
			errmsg <- fmt.Errorf("такой пользователь %s уже есть ", username)
			return
		}
		(*accList).Accounts[username] = &Account{AccountType: accounttype, Balance: 0}
		errmsg <- nil
		return
	}()

	for {
		select {
		case <-done:
			return nil
		case <-timer.C:
			errmsg <- fmt.Errorf("Превышен интервал ожидания")
		case err := <-errmsg:
			return err
		default:
		}
	}
}

// AddBalance - добавим баланс
func (accList *AccountsList) AddBalance(username string,
	sum float32) error {
	done := make(chan struct{})
	errmsg := make(chan error, 1)
	timer := time.NewTimer(time.Second)

	go func() {
		defer close(done)
		accList.Lock()
		acc, ok := (*accList).Accounts[username]
		if !ok {
			errmsg <- fmt.Errorf("Пользователь %s не найден", username)
			return
		}
		if sum <= 0 {
			errmsg <- fmt.Errorf("не дoпустимый баланс  %f ", sum)
			return
		}
		acc.Balance += sum
		accList.Unlock()
		errmsg <- nil
		return
	}()

	for {
		select {
		case <-done:
			return nil
		case <-timer.C:
			return fmt.Errorf("Превышен интервал ожидания")
		case err := <-errmsg:
			return err
		default:
		}
	}
}

// Balance - получить баланс
func (accList *AccountsList) Balance(username string) (float32, error) {
	type vmsgType struct {
		balance float32
		errmsg  error
	}

	done := make(chan struct{})
	vmsg := make(chan vmsgType, 1)
	timer := time.NewTimer(time.Second)

	go func() {
		defer close(done)

		accList.Lock()
		acc, ok := (*accList).Accounts[username]
		accBalance := acc.Balance
		accList.Unlock()
		if !ok {
			vmsg <- vmsgType{balance: 0,
				errmsg: fmt.Errorf("Пользователь %s не найден", username)}
			return
		}
		vmsg <- vmsgType{balance: accBalance, errmsg: nil}
		return
	}()

	for {
		select {
		case <-done:
			bal := <-vmsg
			return bal.balance, bal.errmsg
		case <-timer.C:
			bal := vmsgType{balance: 0,
				errmsg: fmt.Errorf("Превышен интервал ожидания")}
			return bal.balance, bal.errmsg
		default:
		}
	}
}

// GetAccounts - сортируем аккаунты
func (accList AccountsList) GetAccounts(sort AccountSortType) map[string]*Account {

	outAcc := map[string]*Account{}

	accList.Lock()
	keys := make([]string, 0, len(accList.Accounts))

	for k := range accList.Accounts {
		keys = append(keys, k)
	}
	keys1 := make([]float64, 0, len(accList.Accounts))
	for _, v := range accList.Accounts {
		keys1 = append(keys1, float64(v.Balance))
	}
	accList.Unlock()
	switch sort {
	case SortByName:
		sorting.Strings(keys)
	case SortByNameReverse:
		sorting.Sort(sorting.Reverse(sorting.StringSlice(keys)))
	case SortByBalance:
		sorting.Float64s(keys1)
	}

	if sort == SortByName || sort == SortByNameReverse {
		for i := 0; i < len(keys); i++ {
			outAcc[keys[i]] = accList.Accounts[keys[i]]
			fmt.Printf("%s  %f  %d \n", keys[i],
				outAcc[keys[i]].Balance,
				outAcc[keys[i]].AccountType)
		}
	}
	if sort == SortByBalance {
		for i := 0; i < len(keys1); i++ {
			for el, v := range accList.Accounts {
				if v.Balance == float32(keys1[i]) {
					outAcc[el] = v
					fmt.Printf("%s  %f  %d \n", el,
						v.Balance,
						v.AccountType)
					break
				}
			}
		}
	}
	return outAcc
}
