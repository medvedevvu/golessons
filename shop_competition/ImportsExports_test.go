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
	testAccList := &AccountsList{}
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
	testAccList := &AccountsList{}
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

func TestExportAccountsCSV(t *testing.T) {
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
	if len(records) <= 0 {
		t.Fatalf("Импорт не исполнен \n")
	}

}

func TestImportAccountsCSV(t *testing.T) {
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

	accountList := accountsListMain
	// удалим справочник
	for k := range accountList {
		delete((accountList), k)
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
	if len(accountList) <= 0 {
		t.Fatal()
	}
}

func TestWrongBalanceImportAccountsCSV(t *testing.T) {
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

	accountList := accountsListMain
	// удалим справочник
	for k := range accountList {
		delete((accountList), k)
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

func TestExportProdcuctsCSV(_ *testing.T) {
	_ = InitProductCatalog()
	exp := []byte{}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		exp = ExportProdcuctsCSV()
		return
	}()
	wg.Wait()
	fmt.Printf("%v", exp)
}

func TestImportProductsCSV(t *testing.T) {
	_ = InitProductCatalog()
	exp := []byte{}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		exp = ExportProdcuctsCSV()
		return
	}()
	wg.Wait()

	products := productListMain
	// удалим справочник
	for k := range products {
		delete((products), k)
	}
	// закачаем по новой

	wg.Add(1)
	go func() {
		defer wg.Done()
		ImportProductsCSV(exp)
	}()
	wg.Wait()
	globalMutex.Lock()
	records := productListMain
	globalMutex.Unlock()

	len_records := len(records)
	if len_records != 10012 {
		t.Fatalf("Импорт не исполнен = %d \n", len_records)
	}

}

func InitWrongProductCatalog() *ProductsList {
	lproductList := &ProductsList{}
	for i := 0; i < 10000; i++ {
		err := lproductList.AddProduct(fmt.Sprintf("Продукт %d", i), Product{Price: 10.51, Type: ProductNormal})
		if err != nil {
			fmt.Printf("%s\n", err)
		}
	}
	(*lproductList)["XXXXX"] = &Product{Price: -1, Type: ProductNormal}
	return lproductList
}

func TestWrongDataImportProductsCSV(t *testing.T) {
	_ = InitWrongProductCatalog()
	exp := []byte{}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		exp = ExportProdcuctsCSV()
		return
	}()
	wg.Wait()

	products := productListMain
	// удалим справочник
	for k := range products {
		delete((products), k)
	}
	// закачаем по новой

	wg.Add(1)
	err := errors.New("")
	go func() {
		defer wg.Done()
		err = ImportProductsCSV(exp)
	}()
	wg.Wait()

	if err == nil {
		t.Fatalf(" Портировали продукт XXXX с отрицательной ценой \n")
	}

}
