package shopcompetition

import (
	"testing"
)

func TestRegister(t *testing.T) {
	accList := AccountList {
		Account{Name: "Колбаса", Balance: 0, AccountType: AccountPremium},
	}
	nameList := []string{"Vasily", "Petr", "Ole"}
	for _, name := range nameList {
		err := accList.Register(name)
		if err != nil {
			t.Errorf(" не смогли зарегестрировать такой пользователь %s уже есть ", "Vasily")
		}
	}
}
func TestRegisterCheckBalanceNewUsers(t *testing.T) {
	accList := new(AccountList)
	name := "Vasily"
	accList.Register(name)
	_, err := accList.Balance(name)
	if err != nil {
		t.Error("Баланс пользователя не нулевой ")
	}
}




func TestWrongRegister(t *testing.T) {
	accList := AccountList {
		Account{Name: "Vasily", Balance: 0, AccountType: AccountPremium},
	}
	name := "Vasily"
	err := accList.Register(name)
	if err != nil {
		t.Logf("не смогли зарегестрировать такой пользователь %s уже есть ", name)
		return
	}
	t.Errorf("не прошла проверка на дублирование имени пользователя %s", name)
}

func TestWrongRegisterEmpty(t *testing.T) {
	accList := new(AccountList)
	name := " "
	err := accList.Register(name)
	if err != nil {
		t.Log("не смогли зарегестрировать пустое имя")
		return
	}
	t.Errorf("не прошла проверка на пустое имя пользователя %s", name)
}


func TestAddBalance(t *testing.T) {

	accList := new(AccountList)
	nameList := []string{"Vasily", "Petr"}
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
		t.Logf(" Пользователю %s пытаемся установить отрицательный  %.2f баланс ",
			name, bl)
		return
	}
	t.Errorf("Не прошла проверка на отрицательный баланс %.2f для пользователя %s",
		bl, name)
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
		Account{Name: "Vasily", AccountType: AccountNormal, Balance: 111.45002},
		Account{Name: "Boris", AccountType: AccountPremium, Balance: 999.99},
		Account{Name: "Alex", AccountType: AccountNormal, Balance: 778.45},
		Account{Name: "Leo", AccountType: AccountNormal, Balance: 111.45},
	})
	vacc := accList.GetAccounts(SortByName)
	if vacc[0].Name != "Alex" {
		t.Errorf("Сортировка SortByName не выполнена")
	} else {
		t.Log(vacc)
	}
	vacc = accList.GetAccounts(SortByNameReverse)
	if vacc[0].Name != "Vasily" {
		t.Errorf("Сортировка SortByNameReverse не выполнена")
	} else {
		t.Log(vacc)
	}
	vacc = accList.GetAccounts(SortByBalance)
	if vacc[0].Name != "Leo" {
		t.Errorf("Сортировка SortByBalance не выполнена")
	} else {
		t.Log(vacc)
	}
	if vacc[0].Name != "Leo" {
		t.Errorf("Сортировка SortByBalance не выполнена")
	} else {
		t.Log(vacc)
	}

}
