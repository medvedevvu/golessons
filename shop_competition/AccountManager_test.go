package shop_competition

import (
	"fmt"
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
	return testAccList
}

func TestInitAccountList(t *testing.T) {
	vtest := *InitAccountList()
	if len(vtest) == 0 {
		t.Fatalf("не выполнена инициализация ")
	}
	t.Logf("%v", vtest)

}

func TestNewAccountsList(t *testing.T) {
	vtest := InitAccountList()
	_, ok := (*vtest)["Dram"]
	if !ok {
		t.Fatalf("Init fail with user %s", "Dram")
	}
	err := vtest.Register("Boris", AccountPremium)
	if err != nil {
		t.Fatalf("Fail with register user %s", "Boris")
	}
	err = vtest.Register("Boris", AccountPremium)
	if err == nil {
		t.Fatalf("Fail with register twice user %s", "Boris")
	}
	err = vtest.Register("", AccountPremium)
	if err == nil {
		t.Fatalf("Fail with register empty name user %s", "Boris")
	}
}

func TestAddBalance(t *testing.T) {
	vtest := InitAccountList()
	names := map[string]float32{"Kola": 325.12,
		"Vasiy": 900.21, "Dram": 0, "Vortis": -23}

	for key, vals := range names {
		err := vtest.AddBalance(key, vals)
		if err != nil {
			t.Fatal(err)
		}
	}

	v, err := vtest.Balance("Vasiy")
	if err != nil {
		t.Fatal(err)
	}
	if v != (*vtest)["Vasiy"].Balance {
		t.Fatalf(" %f != %f ", v, (*vtest)["Vasiy"].Balance)
	}
}

func TestSetBalance(t *testing.T) {
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

	for key, vals := range names {
		err := vtest.AddBalance(key, vals)
		if err != nil {
			t.Fatal(err)
		}
	}

}
