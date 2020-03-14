package shop_competition

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"reflect"
	"strconv"
)

func ImportProductsCSVHelper(
	ctx context.Context,
	productList [][]string,
	res chan ProductsList,
	done chan struct{},
	errCh chan error) {
	defer func() {
		select {
		case err := <-errCh:
			errCh <- err
		case <-ctx.Done(): // остановка
			done <- struct{}{}
		default:
			done <- struct{}{}
		}
		return
	}()
	tempProductList := NewShopBase()
	for _, rec := range productList {
		name := rec[0] // name
		value, err := strconv.ParseFloat(rec[1], 32)
		if err != nil {
			errCh <- fmt.Errorf("%s -- ошибка парсинга цены %s -- error=%v\n",
				name, rec[1], err)
			return
		}
		price := float32(value) //
		intval, err := strconv.Atoi(rec[2])
		if err != nil {
			errCh <- fmt.Errorf("%s -- ошибка парсинга типа товара %s -- error=%v\n",
				name, rec[2], err)
			return
		}
		productType := ProductType(uint8(intval)) // productType
		err = tempProductList.ProductListWithMutex.AddProduct(name, Product{Price: price, Type: productType})
		if err != nil {
			errCh <- fmt.Errorf("%s -- ошибка добавления товара price=%v type=%v-- error=%v\n",
				name, price, productType, err)
			return
		}
	}
	res <- tempProductList.ProductListWithMutex
	return
}

//ExportProductsCSVHelper
func ExportProductsCSVHelper(ctx context.Context,
	productList []map[string]*Product,
	res chan []byte,
	done chan struct{},
	errCh chan error) {
	defer func() {
		done <- struct{}{}
	}()
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	record := make([]string, 3)
	for _, rec := range productList {
		key0 := (reflect.ValueOf(rec).MapKeys()[0])
		key := fmt.Sprintf("%s", key0)
		record[0] = key
		record[1] = fmt.Sprintf("%.2f", rec[key].Price)
		record[2] = fmt.Sprintf("%v", rec[key].Type)
		err := w.Write(record)
		if err != nil {
			errCh <- err
			return
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		errCh <- err
		return
	}
	res <- buf.Bytes()
	return
}

//ExportAccountsCSVHelper
func ExportAccountsCSVHelper(
	ctx context.Context,
	accountList []map[string]*Account,
	res chan []byte,
	done chan struct{},
	errCh chan error) {

	defer func() {
		done <- struct{}{}
	}()
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	record := make([]string, 3)
	for _, rec := range accountList {
		key0 := (reflect.ValueOf(rec).MapKeys()[0])
		key := fmt.Sprintf("%s", key0)
		record[0] = key
		record[1] = fmt.Sprintf("%.2f", rec[key].Balance)
		record[2] = fmt.Sprintf("%v", rec[key].AccountType)
		err := w.Write(record)
		if err != nil {
			errCh <- err
			return
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		errCh <- err
		return
	}
	res <- buf.Bytes()
	return
}

//ImportAccountsCSVHelper
func ImportAccountsCSVHelper(
	ctx context.Context,
	accountList [][]string,
	res chan AccountsList,
	done chan struct{},
	errCh chan error) {

	defer func() {
		select {
		case err := <-errCh:
			errCh <- err
		case <-ctx.Done(): // остановка
			done <- struct{}{}
		default:
			done <- struct{}{}
		}
		return
	}()
	tempAccountList := NewShopBase()
	for _, rec := range accountList {
		name := rec[0] // name
		intval, err := strconv.Atoi(rec[2])
		if err != nil {
			errCh <- fmt.Errorf("%s -- ошибка парсинга типа аккаунта %s -- error=%v\n",
				name, rec[2], err)
			return
		}
		accountType := AccountType(uint8(intval)) // accountType
		value, err := strconv.ParseFloat(rec[1], 32)
		if err != nil {
			errCh <- fmt.Errorf("%s -- ошибка парсинга баланса %s -- error=%v\n",
				name, rec[1], err)
			return
		}
		balance := float32(value) // Balance
		err = tempAccountList.AccountsListWithMutex.Register(name, accountType)
		if err != nil {
			errCh <- fmt.Errorf("%s -- ошибка регистрации аккаунта %v -- error=%v\n",
				name, accountType, err)
			return
		}
		err = tempAccountList.AccountsListWithMutex.AddBalance(name, balance)
		if err != nil {
			errCh <- fmt.Errorf("%s -- ошибка добавления баланса %v -- error=%v\n",
				name, balance, err)
			return
		}
	}
	res <- tempAccountList.AccountsListWithMutex
	return
}
