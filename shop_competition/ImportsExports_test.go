package shop_competition

import (
	"bytes"
	"encoding/csv"
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

func TestImportAccountsCSV(_ *testing.T) {
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
	for _, record := range records {
		fmt.Println(record)
	}

}
