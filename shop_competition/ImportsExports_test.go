package shop_competition

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
)

func InitAccountListWithBalance() AccountsList {
	AccountsListMain = AccountsList{}
	err := AccountsListMain.Register("Kola", AccountNormal)
	err = AccountsListMain.AddBalance("Kola", 102.21)
	if err != nil {
		fmt.Println(err)
	}
	err = AccountsListMain.Register("Vasiy", AccountNormal)
	err = AccountsListMain.AddBalance("Vasiy", 102.21)
	if err != nil {
		fmt.Println(err)
	}
	err = AccountsListMain.Register("Dram", AccountPremium)
	err = AccountsListMain.AddBalance("Dram", 102.21)
	if err != nil {
		fmt.Println(err)
	}
	err = AccountsListMain.Register("Vortis", AccountPremium)
	err = AccountsListMain.AddBalance("Vortis", 102.21)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 1000; i++ {
		s := fmt.Sprintf("User%d", i)
		err = AccountsListMain.Register(s, AccountNormal)
		err = AccountsListMain.AddBalance(s, 100.23)
		if err != nil {
			fmt.Println(err)
		}
	}
	return AccountsListMain
}

func InitAccountListWithWrongBalance() AccountsList {
	AccountsListMain = AccountsList{}
	err := AccountsListMain.Register("Kola", AccountNormal)
	err = AccountsListMain.AddBalance("Kola", 102.21)
	if err != nil {
		fmt.Println(err)
	}
	err = AccountsListMain.Register("Vasiy", AccountNormal)
	if err != nil {
		fmt.Println(err)
	}
	err = AccountsListMain.Register("Dram", AccountPremium)
	err = AccountsListMain.AddBalance("Dram", 102.21)
	if err != nil {
		fmt.Println(err)
	}
	err = AccountsListMain.Register("Vortis", AccountPremium)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 100; i++ {
		s := fmt.Sprintf("User%d", i)
		err = AccountsListMain.Register(s, AccountNormal)
		err = AccountsListMain.AddBalance(s, 100.23)
		if err != nil {
			fmt.Println(err)
		}
	}
	return AccountsListMain
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
	r := csv.NewReader(bytes.NewReader(exp))
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	if len(records) <= 0 {
		t.Fatalf("Экспорт не исполнен \n")
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

	accountList := AccountsListMain
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

	accountList := AccountsListMain
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
		t.Fatalf(" %s \n", err)
	}
}

func TestExportProdcuctsCSV(t *testing.T) {
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
	if len(exp) == 0 {
		t.Fatalf(" ошибка -- кол-во экспорт. записей %d\n", len(exp))
	}
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

	products := ProductListMain
	// удалим справочник
	before_len_records := len(products)
	for k := range products {
		delete((products), k)
	}
	// закачаем по новой
	stopCh := make(chan struct{}, 1)
	wg.Add(2)

	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 100)
		stopCh <- struct{}{}
	}()

	go func() {
		defer wg.Done()
		err := ImportProductsCSV(exp, stopCh)
		if err != nil {
			t.Logf("%s\n", err)
		}
	}()
	wg.Wait()
	globalMutex.Lock()
	records := ProductListMain
	globalMutex.Unlock()

	len_records := len(records)
	if len_records != before_len_records {
		t.Logf("Импорт не исполнен %d <> %d \n", len_records, before_len_records)
	} else {
		t.Fatalf("не выполнена отмена импорта ")
	}

}

func InitWrongProductCatalog() ProductsList {
	ProductListMain = ProductsList{}
	for i := 0; i < 10000; i++ {
		err := ProductListMain.AddProduct(fmt.Sprintf("Продукт %d", i), Product{Price: 10.51, Type: ProductNormal})
		if err != nil {
			fmt.Printf("%s\n", err)
		}
	}
	ProductListMain["XXXXX"] = &Product{Price: -1, Type: ProductNormal}
	ProductListMain["XXXX1"] = &Product{Price: 0, Type: ProductNormal}
	ProductListMain["XXXXX2"] = &Product{Price: -10, Type: ProductNormal}
	return ProductListMain
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

	products := ProductListMain
	// удалим справочник
	for k := range products {
		delete((products), k)
	}
	// закачаем по новой
	stopCh := make(chan struct{})
	wg.Add(1)
	err := errors.New("")
	go func() {
		defer wg.Done()
		err = ImportProductsCSV(exp, stopCh)
	}()
	wg.Wait()
	if err == nil {
		t.Fatalf(" Портировали продукты XXXX... с неверными данными \n")
	}
}
