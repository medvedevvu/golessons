package shop_competition

import (
	"fmt"
	"math"
	"sync"
	"testing"
)

func InitAccountList(vlist AccountsList) {
	err := vlist.Register("Kola", AccountNormal)
	if err != nil {
		fmt.Println(err)
	}
	err = vlist.Register("Vasiy", AccountNormal)
	if err != nil {
		fmt.Println(err)
	}
	err = vlist.Register("Dram", AccountPremium)
	if err != nil {
		fmt.Println(err)
	}
	err = vlist.Register("Vortis", AccountPremium)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 10000; i++ {
		s := fmt.Sprintf("User%d", i)
		err = vlist.Register(s, AccountNormal)
		err = vlist.AddBalance(s, 99999)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func InitSmallAccountList(vlist AccountsList) {
	err := vlist.Register("Kola", AccountNormal)
	if err != nil {
		fmt.Println(err)
	}
	err = vlist.Register("Vasiy", AccountNormal)
	if err != nil {
		fmt.Println(err)
	}
	err = vlist.Register("Dram", AccountPremium)
	if err != nil {
		fmt.Println(err)
	}
	err = vlist.Register("Vortis", AccountPremium)
	if err != nil {
		fmt.Println(err)
	}
}

func TestInitAccountList(t *testing.T) {
	accList := NewShopBase().AccountsListWithMutex
	InitAccountList(accList)

	if len(accList.Accounts) == 0 {
		t.Fatalf("не выполнена инициализация ")
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		_, ok := accList.Accounts["Dram"]
		if !ok {
			t.Fatalf("Не найден пользователь %s", "Dram")
		}
	}()
	wg.Wait()
}
func Test2WiceRegisterAccountsList(t *testing.T) {
	accList := NewShopBase().AccountsListWithMutex
	InitAccountList(accList)
	var wg sync.WaitGroup
	var times_ int = 1
	wg.Add(times_)
	go func() {
		defer wg.Done()
		err := accList.Register("Vortis", AccountPremium)
		if err == nil {
			t.Fatalf("Fail with register user %s", "Vortis")
		}
		err = accList.Register("Vortis", AccountPremium)
		if err == nil {
			t.Fatalf("Fail with register twice user %s", "Vortis")
		}
	}()
	wg.Wait()
}

func TestRegisterEmptyNameAccountsList(t *testing.T) {
	accList := NewShopBase().AccountsListWithMutex
	InitAccountList(accList)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := accList.Register("", AccountPremium)
		if err == nil {
			t.Fatal("Fail with register empty name user ")
		}
	}()
	wg.Wait()
}
func TestAddZeroBalance(t *testing.T) {
	vtest := NewShopBase().AccountsListWithMutex
	vtest.Register("Ada", AccountPremium)
	err := vtest.AddBalance("Ada", 0)
	if err != nil {
		t.Logf(" %v ", err)
	}
}

func TestAddMinusBalance(t *testing.T) {
	vtest := NewShopBase().AccountsListWithMutex
	vtest.Register("Ada", AccountPremium)
	err := vtest.AddBalance("Ada", -12)
	if err != nil {
		t.Logf(" %v \n", err)
	}
}

func TestAddBalance(t *testing.T) {
	var wg sync.WaitGroup
	vtest := NewShopBase().AccountsListWithMutex
	vtest.Register("Ada", AccountPremium)
	vtest.Register("Vasiy", AccountPremium)
	vtest.Register("Gladis", AccountNormal)
	vtest.Register("Boris", AccountNormal)
	vtest.Register("Fargus", AccountNormal)
	vtest.Register("Wilams", AccountNormal)
	names := map[string]float32{"Ada": 1111.12,
		"Vasiy": 2222.21, "Boris": 3333, "Gladis": 5555,
		"Fargus": 4444, "Wilams": 555.32}
	var delta float32 = 50
	var Bigdelta float32 = 950
	wg.Add(len(names))
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
	// сделаем копию
	vtestBefore := NewShopBase().AccountsListWithMutex
	for x, value := range vtest.Accounts {
		xvalue := *value
		vtestBefore.Accounts[x] = &xvalue
	}
	// добавим дельту
	wg.Add(2 * len(names))
	for idx := range vtest.Accounts {
		go func(idx string, delta float32) {
			defer wg.Done()
			err := vtest.AddBalance(idx, delta) // маленькую
			if err != nil {
				t.Fatalf("%v\n", err)
			}
		}(idx, delta)

		go func(idx string, delta float32) {
			defer wg.Done()
			err := vtest.AddBalance(idx, Bigdelta) // большую
			if err != nil {
				t.Fatalf("%v\n", err)
			}
		}(idx, delta)
	}
	wg.Wait()
	// сравнили
	for idx := range vtest.Accounts {
		after := vtest.Accounts[idx].Balance
		before := vtestBefore.Accounts[idx].Balance
		a := after * 100
		b := before * 100
		d := (Bigdelta + delta) * 100
		r := math.Round(float64(a) - float64(b))
		if r != float64(d) {
			t.Fatalf("%s %v - %v == %v  \n", idx, after, before, delta)
		}
	}
}

func TestCircleCheckBalance(t *testing.T) {
	vtest := NewShopBase().AccountsListWithMutex
	names := map[string]float32{"Kola": 325.12,
		"Vasiy": 900.21, "Dram": 10, "Vortis": 23}

	var delta float32 = 100

	for key, vals := range names {
		key, vals := key, vals
		err := vtest.Register(key, AccountNormal)
		err = vtest.AddBalance(key, vals)
		if err != nil {
			t.Error(err)
		}
	}

	// сделаем копию
	vtestBefore := NewShopBase().AccountsListWithMutex
	for x, value := range vtest.Accounts {
		xvalue := *value
		vtestBefore.Accounts[x] = &xvalue
	}

	vtest.AddBalance("Dram", delta) // добавим баланс

	for idx, _ := range names {

		if idx == "Dram" {
			after := vtest.Accounts[idx].Balance
			before := vtestBefore.Accounts[idx].Balance
			a := after * 100
			b := before * 100
			d := delta * 100
			r := math.Round(float64(a) - float64(b))
			if r != float64(d) {
				t.Fatalf("ошибка добавления баланса "+
					"%s %v - %v == %v  \n", idx, after, before, delta)
			}
		} else {
			after, err := vtest.Balance(idx)
			before, err := vtestBefore.Balance(idx)
			if err != nil {
				t.Fatalf(" ошибка получения баланса ")
			}
			if after != before {
				t.Fatalf(" %s  %v <> %v  ", idx, after, before)
			}
		}
	}
}
