package shop_competition

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"reflect"
	"strconv"
)

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
	res chan *AccountsList,
	done chan struct{},
	errCh chan error) {

	defer func() {
		select {
		case err := <-errCh:
			errCh <- err
		default:
			done <- struct{}{}
		}
		return
	}()

	tempAccountList := &AccountsList{}
	for _, rec := range accountList {
		name := rec[0] // name
		intval, err := strconv.Atoi(rec[1])
		if err != nil {
			errCh <- fmt.Errorf("%s -- ошибка парсинга типа аккаунта %s",
				name, rec[1])
			return
		}
		accountType := AccountType(uint8(intval)) // accountType
		value, err := strconv.ParseFloat(rec[2], 32)
		if err != nil {
			errCh <- fmt.Errorf("%s -- ошибка парсинга баланса %s",
				name, rec[2])
			return
		}
		balance := float32(value) // Balance
		err = tempAccountList.Register(name, accountType)
		if err != nil {
			errCh <- fmt.Errorf("%s -- ошибка регистрации аккаунта %v",
				name, accountType)
			return
		}
		err = tempAccountList.AddBalance(name, balance)
		if err != nil {
			errCh <- fmt.Errorf("%s -- ошибка добавления баланса %v",
				name, balance)
			return
		}
	}
	res <- tempAccountList
	return
}
