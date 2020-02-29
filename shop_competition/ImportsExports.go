package shop_competition

import (
	"context"
	"fmt"
	"math"
)

/*
ImportProductsCSV([]byte) error,
ExportProdcuctsCSV() []byte,
ImportAccountsCSV([]byte) error,
ExportAccountsCSV() []byte.
*/

//ExportAccountsCSV
func ExportAccountsCSV() []byte {
	accountList := GetAccountsList()

	accountListSlice := []map[string]*Account{}
	for name, elem := range *accountList {
		accountListSlice = append(
			accountListSlice,
			map[string]*Account{name: &Account{Balance: elem.Balance, AccountType: elem.AccountType}})
	}

	var page_size int = 1000
	var pages int = int(len(accountListSlice) / page_size)
	var last_page_add int = int(math.Mod(float64(len(accountListSlice)), float64(page_size)))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	res := make(chan []byte, 1)
	done := make(chan struct{}, pages)
	errCh := make(chan error, 1)

	var start, end int
	for i := 0; i < pages; i++ {
		start = i * page_size
		if i == (pages - 1) {
			end = (i+1)*page_size + last_page_add
		} else {
			end = (i+1)*page_size + last_page_add
		}
		go func(ctx context.Context) {
			select {
			case <-ctx.Done():
				done <- struct{}{}
				return
			default:
			}
			err := ExportAccountsCSVHelper(ctx, accountListSlice[start:end], res, done, errCh)
			if err != nil {
				fmt.Println(err)
			}
		}(ctx)
	}

	return nil
}
