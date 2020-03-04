package shop_competition

import (
	"fmt"
	"sync"
	"testing"
)

func InitAccountList() *AccountsList {
	testAccList := NewAccountsList()
	err := testAccList.Register("Kola", AccountNormal)
	if err != nil {
		fmt.Println(err)
	}
	err = testAccList.Register("Vasiy", AccountNormal)
	if err != nil {
		fmt.Println(err)
	}
	err = testAccList.Register("Dram", AccountPremium)
	if err != nil {
		fmt.Println(err)
	}
	err = testAccList.Register("Vortis", AccountPremium)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 1000; i++ {
		s := fmt.Sprintf("User%d", i)
		err = testAccList.Register(s, AccountNormal)
		err = testAccList.AddBalance(s, 99999)
		if err != nil {
			fmt.Println(err)
		}
	}
	return testAccList
}

func InitSmallAccountList() *AccountsList {
	testAccList := NewAccountsList()
	err := testAccList.Register("Kola", AccountNormal)
	if err != nil {
		fmt.Println(err)
	}
	err = testAccList.Register("Vasiy", AccountNormal)
	if err != nil {
		fmt.Println(err)
	}
	err = testAccList.Register("Dram", AccountPremium)
	if err != nil {
		fmt.Println(err)
	}
	err = testAccList.Register("Vortis", AccountPremium)
	if err != nil {
		fmt.Println(err)
	}
	return testAccList
}

func TestInitAccountList(t *testing.T) {
	vtest := InitAccountList()
	if len(*vtest) == 0 {
		t.Fatalf("не выполнена инициализация ")
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		_, ok := (*vtest)["Dram"]
		if !ok {
			t.Fatalf("Init fail with user %s", "Dram")
		}
	}()
	wg.Wait()
}
func Test2WiceRegisterAccountsList(t *testing.T) {
	vtest := InitAccountList()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		err := vtest.Register("Vortis", AccountPremium)
		if err == nil {
			t.Fatalf("Fail with register user %s", "Vortis")
		}
	}()
	go func() {
		defer wg.Done()
		err := vtest.Register("Vortis", AccountPremium)
		if err == nil {
			t.Fatalf("Fail with register twice user %s", "Vortis")
		}
	}()
	wg.Wait()
}

func TestRegisterEmptyNameAccountsList(t *testing.T) {
	vtest := InitAccountList()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := vtest.Register("", AccountPremium)
		if err == nil {
			t.Fatal("Fail with register empty name user ")
		}
	}()
	wg.Wait()
}
func TestAddBalance(t *testing.T) {
	return nil
	vtest := InitAccountList()
	names := map[string]float32{"Kola": 325.12,
		"Vasiy": 900.21, "Dram": 10, "Vortis": 23}

	var wg sync.WaitGroup
	wg.Add(4)
	for key, vals := range names {
		key, vals := key, vals
		go func() {
			defer wg.Done()
			err := vtest.AddBalance(key, vals)
			if err != nil {
				t.Error(err)
			}
		}()
	}

	wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()
		v, err := vtest.Balance("Vasiy")
		if err != nil {
			if fmt.Sprintf("%s", err) != "ok" {
				t.Fatal(err)
			}
		}
		if v != (*vtest)["Vasiy"].Balance {
			t.Fatalf(" %f != %f ", v, (*vtest)["Vasiy"].Balance)
		}
	}()
	wg.Wait()
}

func TestSetBalance(t *testing.T) {
	var wg sync.WaitGroup
	vtest := NewAccountsList()
	vtest.Register("Ada", AccountPremium)
	vtest.Register("Vasiy", AccountPremium)
	vtest.Register("Gladis", AccountNormal)
	vtest.Register("Boris", AccountNormal)
	vtest.Register("Fargus", AccountNormal)
	vtest.Register("Wilams", AccountNormal)

	names := map[string]float32{"Ada": 1111.12,
		"Vasiy": 2222.21, "Boris": 3333, "Gladis": 5555,
		"Fargus": 4444, "Wilams": 555.32}

	wg.Add(6)

	for key, vals := range names {
		go func(key string, vals float32) {
			defer wg.Done()
			err := vtest.AddBalance(key, vals)
			if err != nil {
				t.Fatal(err)
			}
		}(key, vals)
	}
	wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()
		v, err := vtest.Balance("Vasiy")
		if err != nil {
			if fmt.Sprintf("%s", err) != "ok" {
				t.Fatal(err)
			}
		}
		if v != (*vtest)["Vasiy"].Balance {
			t.Fatalf(" %f != %f ", v, (*vtest)["Vasiy"].Balance)
		}
	}()
	wg.Wait()

}

func TestCheckBalance(t *testing.T) {
	//vtest := InitSmallAccountList()
	vtest := InitAccountList()
	var wg sync.WaitGroup

	names := map[string]float32{"Kola": 325.12,
		"Vasiy": 900.21, "Dram": 10, "Vortis": 23}

	for key, vals := range names {
		key, vals := key, vals
		err := vtest.AddBalance(key, vals)
		if err != nil {
			t.Error(err)
		}
	}

	wg.Add(1)

	go func() {
		defer wg.Done()

		vtest.AddBalance("Dram", 100)

		for key, value := range names {

			if key == "Dram" {
				locvalue, err := vtest.Balance(key)
				if err != nil {
					t.Fatalf(" ошибка получения баланса ")
				}
				if locvalue != 110 {
					t.Fatalf(" ошибка добавления баланса ")
				}
			} else {
				locvalue, err := vtest.Balance(key)
				if err != nil {
					t.Fatalf(" ошибка получения баланса ")
				}
				if locvalue != value {
					t.Fatalf(" %s  %v <> %v  ", key, value, locvalue)
				}
			}
		}
	}()
	wg.Wait()
}
