package shop_competition

import (
	"bytes"
	"encoding/csv"
	"fmt"
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
	exp = ExportAccountsCSV(shop)
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
	exp := ExportAccountsCSV(envList)
	accountList := envList.AccountsListWithMutex.Accounts
	// удалим справочник
	for k := range accountList {
		delete((accountList), k)
	}
	// закачаем по новой
	err := ImportAccountsCSV(exp, envList)
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
	exp := ExportAccountsCSV(envList)
	accountList := envList.AccountsListWithMutex.Accounts
	// удалим справочник
	for k := range accountList {
		delete((accountList), k)
	}
	// закачаем по новой
	err := ImportAccountsCSV(exp, envList)
	if err == nil {
		t.Fatalf(" %s \n", err)
	}
}

func TestExportProdcuctsCSV(t *testing.T) {
	shop := NewShopBase()
	vals := shop.ProductListWithMutex
	InitProductCatalog(vals)
	exp := ExportProdcuctsCSV(shop)
	if len(exp) == 0 {
		t.Fatalf(" ошибка -- кол-во экспорт. записей %d\n", len(exp))
	}
}

func TestImportProductsCSV(t *testing.T) {
	shop := NewShopBase()
	vals := shop.ProductListWithMutex
	InitProductCatalog(vals)
	exp := ExportProdcuctsCSV(shop)
	//fmt.Printf("%v ", exp)
	products := shop.ProductListWithMutex.Products
	// удалим справочник
	before_len_records := len(products)
	for k := range products {
		delete((products), k)
	}
	err := ImportProductsCSV(exp, shop)
	if err != nil {
		t.Fatalf(" ошибка импорта %v\n", err)
	}
	shop.Lock()
	records := shop.ProductListWithMutex.Products
	len_records := len(records)
	shop.Unlock()
	if len_records != before_len_records {
		t.Fatalf("Импорт не исполнен %d <> %d \n", len_records, before_len_records)
	}
}

func InitWrongProductCatalog(envLst *ShopBase) {
	for i := 0; i < 1000; i++ {
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
	exp := ExportProdcuctsCSV(envLst)
	products := envLst.ProductListWithMutex.Products
	// удалим справочник
	for k := range products {
		delete((products), k)
	}
	// закачаем по новой
	err := ImportProductsCSV(exp, envLst)
	if err == nil {
		t.Fatalf(" Портировали продукты XXXX... с неверными данными \n")
	}
}
