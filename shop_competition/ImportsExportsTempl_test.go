package shop_competition

import (
	"fmt"
	"sync"
	"testing"
)

func XXXTestExportAccountsCSVt(_ *testing.T) {
	_ = InitAccountList()
	exp := []byte{}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		exp = ExportAccountsCSV0()
		return
	}()
	wg.Wait()
	fmt.Printf("%v", exp)
	fmt.Println()
	wg.Add(1)
	go func() {
		defer wg.Done()
		exp = ExportAccountsCSVt()
		return
	}()
	wg.Wait()
	fmt.Printf("%v", exp)
	fmt.Println()
}

func TestXXXExportProdcuctsCSVt(_ *testing.T) {
	_ = InitProductCatalog()
	exp := []byte{}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		exp = ExportProdcuctsCSVt()
		return
	}()
	wg.Wait()
	fmt.Printf("%v", exp)
}

func XXXTestImportProductsCSVt(_ *testing.T) {
	_ = InitProductCatalog()
	exp := []byte{}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		exp = ExportProdcuctsCSVt()
		return
	}()
	wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()
		ImportProductsCSVt(exp)
	}()
	wg.Wait()

}
