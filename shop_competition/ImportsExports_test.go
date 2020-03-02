package shop_competition

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"sync"
	"testing"
)

func InitAccountListWithBalance() *AccountsList {
	testAccList := NewAccountsList()
	err := testAccList.Register("Kola", AccountNormal)
	err = testAccList.AddBalance("Kola", 102.21)
	if err != nil {
		fmt.Println(err)
	}
	err = testAccList.Register("Vasiy", AccountNormal)
	err = testAccList.AddBalance("Vasiy", 102.21)
	if err != nil {
		fmt.Println(err)
	}
	err = testAccList.Register("Dram", AccountPremium)
	err = testAccList.AddBalance("Dram", 102.21)
	if err != nil {
		fmt.Println(err)
	}
	err = testAccList.Register("Vortis", AccountPremium)
	err = testAccList.AddBalance("Vortis", 102.21)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 1000; i++ {
		s := fmt.Sprintf("User%d", i)
		err = testAccList.Register(s, AccountNormal)
		err = testAccList.AddBalance(s, 100.23)
		if err != nil {
			fmt.Println(err)
		}
	}
	return testAccList
}

func InitAccountListWithWrongBalance() *AccountsList {
	testAccList := NewAccountsList()
	err := testAccList.Register("Kola", AccountNormal)
	err = testAccList.AddBalance("Kola", 102.21)
	if err != nil {
		fmt.Println(err)
	}
	err = testAccList.Register("Vasiy", AccountNormal)
	if err != nil {
		fmt.Println(err)
	}
	err = testAccList.Register("Dram", AccountPremium)
	err = testAccList.AddBalance("Dram", 102.21)
	if err != nil {
		fmt.Println(err)
	}
	err = testAccList.Register("Vortis", AccountPremium)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 100; i++ {
		s := fmt.Sprintf("User%d", i)
		err = testAccList.Register(s, AccountNormal)
		err = testAccList.AddBalance(s, 100.23)
		if err != nil {
			fmt.Println(err)
		}
	}
	return testAccList
}

func TestImportAccountsCSV(t *testing.T) {
	_ = InitAccountList()
	exp := []byte{}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		exp = ExportAccountsCSV()
		return
	}()
	wg.Wait()
	fmt.Println("-------чтение --------")
	r := csv.NewReader(bytes.NewReader(exp))
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	//	for _, record := range records {
	//		fmt.Println(record)
	//	}
	if len(records) <= 0 {
		t.Fatalf("Импорт не исполнен \n")
	}

}

func TestExportAccountsCSV(t *testing.T) {
	_ = InitAccountListWithBalance()
	exp := []byte{}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		exp = ExportAccountsCSV()
		return
	}()
	wg.Wait()

	accountList := GetAccountsList()
	// удалим справочник
	for k := range *accountList {
		delete((*accountList), k)
	}
	// закачаем по новой
	wg.Add(1)
	err := errors.New("")
	go func() {
		defer wg.Done()
		err = ImportAccountsCSV(exp)
		return
	}()
	wg.Wait()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	//	for key, val := range *accountList {
	//		fmt.Printf(" name=%s  AccountType=%v  Balance=%v \n",
	//			key, val.AccountType, val.Balance)
	//	}
	if len(*accountList) <= 0 {
		t.Fatal()
	}
}

func TestWrongBalanceExportAccountsCSV(t *testing.T) {
	_ = InitAccountListWithWrongBalance()
	exp := []byte{}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		exp = ExportAccountsCSV()
		return
	}()
	wg.Wait()

	accountList := GetAccountsList()
	// удалим справочник
	for k := range *accountList {
		delete((*accountList), k)
	}
	// закачаем по новой
	wg.Add(1)
	err := errors.New("")
	go func() {
		defer wg.Done()
		err = ImportAccountsCSV(exp)
		return
	}()
	wg.Wait()
	if err == nil {
		t.Fatalf(" Портировали Vasiy и Vortis с нулевым балансом \n")
	}
}
