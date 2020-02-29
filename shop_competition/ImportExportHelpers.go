package shop_competition

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"reflect"
)

//ExportAccountsCSVHelper
func ExportAccountsCSVHelper(
	ctx context.Context,
	accountList []map[string]*Account,
	res chan []byte,
	done chan struct{}, errCh chan error) error {

	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	record := make([]string, 3)

	go func() {
		for _, rec := range accountList {
			key0 := (reflect.ValueOf(rec).MapKeys()[0:1])
			v := reflect.ValueOf(key0)
			key := v.Interface().(string)
			record[0] = key
			record[1] = fmt.Sprintf("%.2f", rec[key].Balance)
			record[2] = fmt.Sprintf("%v", rec[key].AccountType)
			err := w.Write(record)
			if err != nil {
				errCh <- err
			}
		}
		w.Flush()

		if err := w.Error(); err != nil {
			errCh <- err
		}
	}()

	for {
		select {
		case <-ctx.Done():
			done <- struct{}{}
			return nil
		case err := <-errCh:
			return err
		case res <- buf.Bytes():
			return nil
		}
	}

}
