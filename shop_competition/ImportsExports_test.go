package shop_competition

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"sync"
	"testing"
)

func InitAccountListWithBalance(envList *ShopBase) {
	err := envList.AccountsListWithMutex.Register("Kola", AccountNormal)
	err = envList.AccountsListWithMutex.AddBalance("Kola", 102.21)
	if err != nil {
		fmt.Println(err)
	}
	err = envList.AccountsListWithMutex.Register("Vasiy", AccountNormal)
	err = envList.AccountsListWithMutex.AddBalance("Vasiy", 102.21)
	if err != nil {
		fmt.Println(err)
	}
	err = envList.AccountsListWithMutex.Register("Dram", AccountPremium)
	err = envList.AccountsListWithMutex.AddBalance("Dram", 102.21)
	if err != nil {
		fmt.Println(err)
	}
	err = envList.AccountsListWithMutex.Register("Vortis", AccountPremium)
	err = envList.AccountsListWithMutex.AddBalance("Vortis", 102.21)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 1000; i++ {
		s := fmt.Sprintf("User%d", i)
		err = envList.AccountsListWithMutex.Register(s, AccountNormal)
		err = envList.AccountsListWithMutex.AddBalance(s, 100.23)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func InitAccountListWithWrongBalance(envList *ShopBase) {
	err := envList.AccountsListWithMutex.Register("Kola", AccountNormal)
	err = envList.AccountsListWithMutex.AddBalance("Kola", 102.21)
	if err != nil {
		fmt.Println(err)
	}
	err = envList.AccountsListWithMutex.Register("Vasiy", AccountNormal)
	if err != nil {
		fmt.Println(err)
	}
	err = envList.AccountsListWithMutex.Register("Dram", AccountPremium)
	err = envList.AccountsListWithMutex.AddBalance("Dram", 102.21)
	if err != nil {
		fmt.Println(err)
	}
	err = envList.AccountsListWithMutex.Register("Vortis", AccountPremium)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 100; i++ {
		s := fmt.Sprintf("User%d", i)
		err = envList.AccountsListWithMutex.Register(s, AccountNormal)
		err = envList.AccountsListWithMutex.AddBalance(s, 100.23)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func TestExportAccountsCSV(t *testing.T) {
	shop := NewShopBase()
	accList := shop.AccountsListWithMutex
	InitAccountList(accList)

	exp := []byte{}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		exp = ExportAccountsCSV(shop)
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
	envList := NewShopBase()
	InitAccountListWithBalance(envList)
	exp := []byte{}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		exp = ExportAccountsCSV(envList)
		return
	}()
	wg.Wait()

	accountList := envList.AccountsListWithMutex.Accounts
	// удалим справочник
	for k := range accountList {
		delete((accountList), k)
	}
	// закачаем по новой
	wg.Add(1)
	err := errors.New("")
	go func() {
		defer wg.Done()
		err = ImportAccountsCSV(exp, envList)
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
	envList := NewShopBase()
	InitAccountListWithWrongBalance(envList)
	exp := []byte{}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		exp = ExportAccountsCSV(envList)
		return
	}()
	wg.Wait()

	accountList := envList.AccountsListWithMutex.Accounts
	// удалим справочник
	for k := range accountList {
		delete((accountList), k)
	}
	// закачаем по новой
	wg.Add(1)
	err := errors.New("")
	go func() {
		defer wg.Done()
		err = ImportAccountsCSV(exp, envList)
		return
	}()
	wg.Wait()
	if err == nil {
		t.Fatalf(" %s \n", err)
	}
}

func TestExportProdcuctsCSV(t *testing.T) {
	shop := NewShopBase()
	vals := shop.ProductListWithMutex
	InitProductCatalog(vals)

	exp := []byte{}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		exp = ExportProdcuctsCSV(shop)
		return
	}()
	wg.Wait()
	if len(exp) == 0 {
		t.Fatalf(" ошибка -- кол-во экспорт. записей %d\n", len(exp))
	}
}

func TestImportProductsCSV(t *testing.T) {
	shop := NewShopBase()
	vals := shop.ProductListWithMutex
	InitProductCatalog(vals)

	exp := []byte{}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		exp = ExportProdcuctsCSV(shop)
		return
	}()
	wg.Wait()

	products := shop.ProductListWithMutex.Products
	// удалим справочник
	before_len_records := len(products)
	for k := range products {
		delete((products), k)
	}
	// закачаем по новой
	wg.Add(1)
	go func() {
		defer wg.Done()
		ImportProductsCSV(exp, shop)
	}()
	wg.Wait()
	shop.Lock()
	records := shop.ProductListWithMutex.Products
	len_records := len(records)
	shop.Unlock()

	if len_records != before_len_records {
		t.Fatalf("Импорт не исполнен %d <> %d \n", len_records, before_len_records)
	}

}

func InitWrongProductCatalog(envLst *ShopBase) {
	for i := 0; i < 10000; i++ {
		err := envLst.ProductListWithMutex.AddProduct(fmt.Sprintf("Продукт %d", i), Product{Price: 10.51, Type: ProductNormal})
		if err != nil {
			fmt.Printf("%s\n", err)
		}
	}
	envLst.ProductListWithMutex.Products["XXXXX"] = &Product{Price: -1, Type: ProductNormal}
	envLst.ProductListWithMutex.Products["XXXX1"] = &Product{Price: 0, Type: ProductNormal}
	envLst.ProductListWithMutex.Products["XXXXX2"] = &Product{Price: -10, Type: ProductNormal}
}

func TestWrongDataImportProductsCSV(t *testing.T) {
	envLst := NewShopBase()
	InitWrongProductCatalog(envLst)
	exp := []byte{}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		exp = ExportProdcuctsCSV(envLst)
		return
	}()
	wg.Wait()

	products := envLst.ProductListWithMutex.Products
	// удалим справочник
	for k := range products {
		delete((products), k)
	}
	// закачаем по новой
	wg.Add(1)
	err := errors.New("")
	go func() {
		defer wg.Done()
		err = ImportProductsCSV(exp, envLst)
	}()
	wg.Wait()
	if err == nil {
		t.Fatalf(" Портировали продукты XXXX... с неверными данными \n")
	}
}
