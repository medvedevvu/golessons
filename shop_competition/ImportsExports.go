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
	globalMutex.Lock()
	accountList := GetAccountsList()
	globalMutex.Unlock()
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
	//done := make(chan struct{}, pages)
	done := make(chan struct{}, 1)
	errCh := make(chan error, 1)

	var start, end int
	if pages > 0 {
		for i := 0; i < pages; i++ {
			start = i * page_size
			if i == (pages - 1) {
				end = (i+1)*page_size + last_page_add
			} else {
				end = (i+1)*page_size + last_page_add
			}
			go func(ctx context.Context, start int, end int) {
				select {
				case <-ctx.Done():
					done <- struct{}{}
					return
				default:
				}
				ExportAccountsCSVHelper(ctx, accountListSlice[start:end], res, done, errCh)
			}(ctx, start, end)
		}
	} else {
		go func(ctx context.Context, last_page_add int) {
			select {
			case <-ctx.Done():
				done <- struct{}{}
				return
			default:
			}
			ExportAccountsCSVHelper(ctx, accountListSlice[0:last_page_add], res, done, errCh)
		}(ctx, last_page_add)
	}

	result := []byte{}
Loop:
	for {
		select {
		/*		case <-done:
				fmt.Println("DONE")
				break Loop*/
		case res := <-res:
			for _, temp := range res {
				result = append(result, temp)
			}
			result = append(result, res...)
			break Loop
		case errMsg := <-errCh:
			fmt.Println(errMsg)
			break Loop
		}
	}

	return result
}
