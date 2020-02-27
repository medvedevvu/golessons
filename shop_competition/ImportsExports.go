package shop_competition

import (
	"bytes"
	"encoding/csv"
	"fmt"
)

/*
ImportProductsCSV([]byte) error,
ExportProdcuctsCSV() []byte,
ImportAccountsCSV([]byte) error,
ExportAccountsCSV() []byte.
*/

type Mybyte []byte

func (mb *Mybyte) Write(record AccountsList) {
	for name, value := range record {
		str := make([]string, 4)
		str[0] = fmt.Sprintf("%s", name)
		str[1] = fmt.Sprintf("%v", value.AccountType)
		str[2] = fmt.Sprintf("%.2f", value.Balance)
		str[3] = "\n"
		for i := 0; i < len(str); i++ {
			*mb = append(*mb, []byte(str[i])...)
		}
	}
}

//ExportAccountsCSV0
func ExportAccountsCSV0() []byte {
	accountList := GetAccountsList()
	mb := Mybyte{}
	for name, record := range *accountList {
		name, record := name, record
		vb := map[string]*Account{name: record}
		mb.Write(vb)
	}
	return mb
}

//ExportAccountsCSV1
func ExportAccountsCSV1() []byte {
	accountList := GetAccountsList()
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)

	record := make([]string, 3)
	for name, rec := range *accountList {
		name, rec := name, rec
		record[0] = name
		record[1] = fmt.Sprintf("%.2f", rec.Balance)
		record[2] = fmt.Sprintf("%v", rec.AccountType)
		err := w.Write(record)
		if err != nil {
			panic(err)
		}
	}
	w.Flush()

	if err := w.Error(); err != nil {
		panic(err)
	}

	return buf.Bytes()

}

//ExportProdcuctsCSV
func ExportProdcuctsCSV() []byte {
	productList := GetProductList()
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)

	record := make([]string, 3)
	for name, rec := range *productList {
		name, rec := name, rec
		record[0] = name
		record[1] = fmt.Sprintf("%.2f", rec.Price)
		record[2] = fmt.Sprintf("%v", rec.Type)
		err := w.Write(record)
		if err != nil {
			panic(err)
		}
	}
	w.Flush()

	if err := w.Error(); err != nil {
		panic(err)
	}

	return buf.Bytes()
}
