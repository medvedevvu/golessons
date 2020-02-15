package shop_competition

import (
	"testing"
)

func InitAccountList() *AccountsList {
	testAccList := NewAccountsList()
	testAccList.Register("Kola")
	testAccList.Register("Vasiy")
	testAccList.Register("Dram")
	testAccList.Register("Vortis")
	return testAccList
}
func TestNewAccountsList(t *testing.T) {
	vtest := InitAccountList()
	_, ok := (*vtest)["Dram"]
	if !ok {
		t.Fatalf("Init fail with user %s", "Dram")
	}
	err := vtest.Register("Boris")
	if err != nil {
		t.Fatalf("Fail with register user %s", "Boris")
	}
	err = vtest.Register("Boris")
	if err == nil {
		t.Fatalf("Fail with register twice user %s", "Boris")
	}
	err = vtest.Register("")
	if err == nil {
		t.Fatalf("Fail with register empty name user %s", "Boris")
	}
	t.Logf("%v", vtest)
}

func TestAddBalance(t *testing.T) {
	vtest := InitAccountList()
	names := map[string]float32{"Kola": 325.12,
		"Vasiy": 900.21, "Dram": 0, "Vortis": -23}

	for key, vals := range names {
		err := vtest.AddBalance(key, vals)
		if err != nil {
			t.Error(err)
		}
	}

	v, err := vtest.Balance("Vasiy")
	if err != nil {
		t.Error(err)
	}
	if v != (*vtest)["Vasiy"].Balance {
		t.Errorf(" %f != %f ", v, (*vtest)["Vasiy"].Balance)
	}
	t.Log("\n FINISH")
}

func TestGetAccounts(t *testing.T) {
	vtest := NewAccountsList()
	vtest.Register("Ada")
	vtest.Register("Vasiy")
	vtest.Register("Gladis")
	vtest.Register("Boris")
	vtest.Register("Fargus")
	vtest.Register("Wilams")

	names := map[string]float32{"Ada": 1111.12,
		"Vasiy": 2222.21, "Boris": 3333, "Gladis": 5555,
		"Fargus": 4444, "Wilams": 555.32}

	for key, vals := range names {
		err := vtest.AddBalance(key, vals)
		if err != nil {
			t.Error(err)
		}
	}

}
