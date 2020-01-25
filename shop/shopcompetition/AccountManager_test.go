package shopcompetition

import (
	"testing"
)

func TestRegister(t *testing.T) {
	accList := new(AccountList)
	nameList := []string{"Vasily", "Vasily", "Petr", "Oleg"}
	for _, name := range nameList {
		err := accList.Register(name)
		if err != nil {
			t.Errorf(" не смогли зарегестрировать такой пользователь %s уже есть ", "Vasily")
		}
	}
	t.Log(accList)
}

func TestAddBalance(t *testing.T) {

	accList := new(AccountList)
	nameList := []string{"Vasily", "Petr", "Oleg"}
	for _, name := range nameList {
		err := accList.Register(name)
		if err != nil {
			t.Errorf(" не смогли зарегестрировать такой пользователь %s уже есть ", "Vasily")
		}
	}
	for _, name := range nameList {
		err := accList.AddBalance(name, 999.23)
		if err != nil {
			t.Errorf(" не смогли установить баланс пользователю %s", name)
		}
	}
	t.Log(accList)
}

func TestAddWrongBalance(t *testing.T) {
	accList := new(AccountList)
	name := "Vasily"
	err := accList.Register(name)
	var bl float32 = -102.3
	err = accList.AddBalance(name, bl)
	if err != nil {
		t.Errorf(" Пользователю %s пытаемся установить отрицательный  %.2f баланс ",
			name, bl)
	}
	t.Log(accList)
}

func TestBalance(t *testing.T) {
	accList := new(AccountList)
	name := "Vasily"
	accList.Register(name)
	var bl float32 = 102.3
	accList.AddBalance(name, bl)
	bil, err := accList.Balance(name)
	if bil != bl {
		t.Errorf("Ввели %f <> значению %f", bl, bil)
	}
	if err != nil {
		t.Errorf("%v", err)
	}
	t.Logf("%f == %f", bl, bil)
}

func TestGetAccounts(t *testing.T) {
	accList := AccountList([]Account{
		Account{Name: "Vasily", AccountType: AccountNormal, Balance: 251.69},
		Account{Name: "Boris", AccountType: AccountPremium, Balance: 999.99},
		Account{Name: "Alex", AccountType: AccountNormal, Balance: 778.45},
		Account{Name: "Leo", AccountType: AccountNormal, Balance: 111.45},
	})
	vacc := accList.GetAccounts(SortByName)
	t.Log(vacc)
	vacc = accList.GetAccounts(SortByNameReverse)
	t.Log(vacc)
	vacc = accList.GetAccounts(SortByBalance)
	t.Log(vacc)
}
