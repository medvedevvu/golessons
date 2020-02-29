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
		exp = ExportAccountsCSV()
		return
	}()
	wg.Wait()
	fmt.Printf("%v", exp)
	fmt.Println()
}
