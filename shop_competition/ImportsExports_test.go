package shop_competition

import (
	"fmt"
	"sync"
	"testing"
)

func TestExportAccountsCSV(_ *testing.T) {
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
		exp = ExportAccountsCSV1()
		return
	}()
	wg.Wait()
	fmt.Printf("%v", exp)
	fmt.Println()
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
