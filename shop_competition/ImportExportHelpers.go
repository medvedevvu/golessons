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
	done chan struct{},
	errCh chan error) {

	go func() {
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
		fmt.Printf("Before \n")
		res <- buf.Bytes()
		fmt.Printf("After \n")

		return
	}()

Loop:
	for {
		select {
		case <-errCh:
			fmt.Println("errCh:")
			break Loop
		case <-res:
			fmt.Println("res")
			break Loop
		case <-ctx.Done():
			fmt.Println("ctx")
			//done <- struct{}{}
			break Loop
		}
	}
	return
}
