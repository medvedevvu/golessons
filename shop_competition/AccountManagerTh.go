package shop_competition

import (
	"errors"
	"fmt"
	sorting "sort"
	"strings"
	"sync"
	"time"
)

// NewAccountsList коструктор
var (
	gaccountsList *AccountsList
	once          sync.Once
	lmu           sync.RWMutex
)

func NewAccountsList() *AccountsList {
	once.Do(func() {
		gaccountsList = &AccountsList{}
	})
	return gaccountsList
}

func GetAccountsList() *AccountsList {
	return gaccountsList
}

// Register - регистрация пользователя

func (accountsList *AccountsList) Register(username string, accounttype AccountType) error {
	const Timeout = time.Second
	timer := time.NewTimer(Timeout)
	res := make(chan string)

	go func() {
		res = (*accountsList).helperRegister(username, accounttype)
	}()

	defer func() {
		close(res)
		timer.Stop()
	}()

	msg := ""
loop:
	for {
		select {
		case resualt := <-res:
			if resualt != "" {
				msg = fmt.Sprintf("%s", resualt)
			}
		case <-timer.C:
			msg = fmt.Sprintf("Exit by timeout")
			break loop
		}
	}
	return errors.New(msg)
}

func (accountsList *AccountsList) helperRegister(username string, accounttype AccountType) chan string {
	cherr := make(chan string)
	if len(strings.Trim(username, "")) == 0 {
		cherr <- fmt.Sprintf("username %s пустое ", username)
		return cherr
	}
	lmu.Lock()
	_, ok := (*accountsList)[username]
	lmu.Unlock()
	if ok {
		cherr <- fmt.Sprintf("такой пользователь %s уже есть ", username)
		return cherr
	}
	lmu.RLock()
	(*accountsList)[username] = &Account{AccountType: accounttype, Balance: 0}
	lmu.RUnlock()
	cherr <- ""
	return cherr
}

// AddBalance - добавим баланс
func (accountsList *AccountsList) AddBalance(username string,
	sum float32) error {
	acc, ok := (*accountsList)[username]
	if !ok {
		return fmt.Errorf("Пользователь %s не найден", username)
	}
	if sum <= 0 {
		return fmt.Errorf("не дoпустимый баланс  %f ", sum)
	}
	acc.Balance += sum
	return nil
}

// Balance - получить баланс
func (accountsList *AccountsList) Balance(username string) (float32, error) {
	acc, ok := (*accountsList)[username]
	if !ok {
		return 0, fmt.Errorf("Пользователь %s не найден", username)
	}
	return acc.Balance, nil
}

// GetAccounts - сортируем аккаунты
func (accountsList AccountsList) GetAccounts(sort AccountSortType) AccountsList {
	outAcc := AccountsList{}
	keys := make([]string, 0, len(accountsList))

	for k := range accountsList {
		keys = append(keys, k)
	}
	keys1 := make([]float64, 0, len(accountsList))
	for _, v := range accountsList {
		keys1 = append(keys1, float64(v.Balance))
	}
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
			outAcc[keys[i]] = accountsList[keys[i]]
			fmt.Printf("%s  %f  %d \n", keys[i], outAcc[keys[i]].Balance, outAcc[keys[i]].AccountType)
		}
	}
	if sort == SortByBalance {
		for i := 0; i < len(keys1); i++ {
			for el, v := range accountsList {
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
