package shop_competition

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"math"
	"time"
)

//ImportProductsCSV([]byte)
func ImportProductsCSV(data []byte) error {
	r := csv.NewReader(bytes.NewReader(data))
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	var page_size int = 1000 // 100
	var pages int = int(len(records) / page_size)
	var last_page_add int = int(math.Mod(float64(len(records)), float64(page_size)))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	res := make(chan *ProductsList, pages)
	done := make(chan struct{}, pages)
	errCh := make(chan error, pages)

	var start, end int
	if pages > 0 {
		for i := 0; i < pages; i++ {
			start = i * page_size
			if i == (pages - 1) {
				end = (i+1)*page_size + last_page_add
			} else {
				end = (i + 1) * page_size
			}
			go func(ctx context.Context, start int, end int) {
				ImportProductsCSVHelper(ctx, records[start:end], res, done, errCh)
			}(ctx, start, end)
			time.Sleep(time.Second)
		}
	} else {
		go func(ctx context.Context, last_page_add int) {
			ImportProductsCSVHelper(ctx, records[0:last_page_add], res, done, errCh)
			time.Sleep(time.Second)
		}(ctx, last_page_add)
	}
	result := map[string]*Product{}
Loop:
	for {
		select {
		case res1 := <-res:
			for key, val := range *res1 {
				result[key] = val
			}
		loop:
			for {
				if len(done) == pages {
					break loop
				}
				select {
				case errMsg := <-errCh:
					cancel() // выключаю все горутины
					return errMsg
				default:
				}
			}
			for i := 0; i < len(done)-1; i++ {
				res1 := <-res
				for key, val := range *res1 {
					result[key] = val
				}
			}
			cancel() // выключаю все горутины
			// все прочитались - добовляем в базовую коллекцию
			globalMutex.Lock()
			productsList := GetProductList()
			for key, val := range result {
				(*productsList)[key] = val
			}
			globalMutex.Unlock()
			break Loop
		case errMsg := <-errCh:
			cancel() // выключаю все горутины
			return errMsg
		}
	}
	return nil
}

//ExportProdcuctsCSV
func ExportProdcuctsCSV() []byte {
	globalMutex.Lock()
	productList := GetProductList()
	globalMutex.Unlock()
	productListSlice := []map[string]*Product{}
	for name, elem := range *productList {
		productListSlice = append(
			productListSlice,
			map[string]*Product{name: &Product{Price: elem.Price, Type: elem.Type}})
	}

	var page_size int = 100 // 100
	var pages int = int(len(productListSlice) / page_size)
	var last_page_add int = int(math.Mod(float64(len(productListSlice)),
		float64(page_size)))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	res := make(chan []byte, pages)
	done := make(chan struct{}, pages)
	errCh := make(chan error, pages)
	var start, end int
	if pages > 0 {
		for i := 0; i < pages; i++ {
			start = i * page_size
			if i == (pages - 1) {
				end = (i+1)*page_size + last_page_add
			} else {
				end = (i + 1) * page_size
			}
			go func(ctx context.Context, start int, end int) {
				ExportProductsCSVHelper(ctx, productListSlice[start:end], res, done, errCh)
			}(ctx, start, end)
			time.Sleep(time.Second)
		}
	} else {
		go func(ctx context.Context, last_page_add int) {
			ExportProductsCSVHelper(ctx, productListSlice[0:last_page_add], res, done, errCh)
			time.Sleep(time.Second)
		}(ctx, last_page_add)
	}
	result := []byte{}
Loop:
	for {
		select {
		case res1 := <-res:
			for _, temp := range res1 {
				result = append(result, temp)
			}

		loop:
			for {
				if len(done) == pages {
					break loop
				}
			}
			for i := 0; i < len(done)-1; i++ {
				res1 := <-res
				result = append(result, res1...)
			}
			break Loop
		case errMsg := <-errCh:
			fmt.Println(errMsg)
			break Loop
		}
	}
	return result
}

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

	var page_size int = 100 // 100
	var pages int = int(len(accountListSlice) / page_size)
	var last_page_add int = int(math.Mod(float64(len(accountListSlice)), float64(page_size)))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	res := make(chan []byte, pages)
	done := make(chan struct{}, pages)
	errCh := make(chan error, pages)
	var start, end int
	if pages > 0 {
		for i := 0; i < pages; i++ {
			start = i * page_size
			if i == (pages - 1) {
				end = (i+1)*page_size + last_page_add
			} else {
				end = (i + 1) * page_size
			}
			go func(ctx context.Context, start int, end int) {
				ExportAccountsCSVHelper(ctx, accountListSlice[start:end], res, done, errCh)
			}(ctx, start, end)
			time.Sleep(time.Second)
		}
	} else {
		go func(ctx context.Context, last_page_add int) {
			ExportAccountsCSVHelper(ctx, accountListSlice[0:last_page_add], res, done, errCh)
			time.Sleep(time.Second)
		}(ctx, last_page_add)
	}
	result := []byte{}
Loop:
	for {
		select {
		case res1 := <-res:
			for _, temp := range res1 {
				result = append(result, temp)
			}

		loop:
			for {
				if len(done) == pages {
					break loop
				}
			}
			for i := 0; i < len(done)-1; i++ {
				res1 := <-res
				result = append(result, res1...)
			}
			break Loop
		case errMsg := <-errCh:
			fmt.Println(errMsg)
			break Loop
		}
	}

	return result
}

//ImportAccountsCSV
func ImportAccountsCSV(data []byte) error {
	r := csv.NewReader(bytes.NewReader(data))
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	var page_size int = 100 // 100
	var pages int = int(len(records) / page_size)
	var last_page_add int = int(math.Mod(float64(len(records)), float64(page_size)))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	res := make(chan *AccountsList, pages)
	done := make(chan struct{}, pages)
	errCh := make(chan error, pages)

	var start, end int
	if pages > 0 {
		for i := 0; i < pages; i++ {
			start = i * page_size
			if i == (pages - 1) {
				end = (i+1)*page_size + last_page_add
			} else {
				end = (i + 1) * page_size
			}
			go func(ctx context.Context, start int, end int) {
				ImportAccountsCSVHelper(ctx, records[start:end], res, done, errCh)
			}(ctx, start, end)
			time.Sleep(time.Second)
		}
	} else {
		go func(ctx context.Context, last_page_add int) {
			ImportAccountsCSVHelper(ctx, records[0:last_page_add], res, done, errCh)
			time.Sleep(time.Second)
		}(ctx, last_page_add)
	}
	result := map[string]*Account{}
Loop:
	for {
		select {
		case res1 := <-res:
			for key, val := range *res1 {
				result[key] = val
			}
		loop:
			for {
				if len(done) == pages {
					break loop
				}
				select {
				case errMsg := <-errCh:
					cancel() // выключаю все горутины
					return errMsg
				default:
				}
			}
			for i := 0; i < len(done)-1; i++ {
				res1 := <-res
				for key, val := range *res1 {
					result[key] = val
				}
			}
			cancel() // выключаю все горутины
			// все прочитались - добовляем в базовую коллекцию
			globalMutex.Lock()
			accountList := GetAccountsList()
			for key, val := range result {
				(*accountList)[key] = val
			}
			globalMutex.Unlock()
			break Loop
		case errMsg := <-errCh:
			cancel() // выключаю все горутины
			return errMsg
		}
	}
	return nil
}
