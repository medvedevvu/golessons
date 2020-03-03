package shop_competition

import (
	"bytes"
	"encoding/csv"
	"fmt"
)

/*
ImportProductsCSVt([]byte) error,
ExportProdcuctsCSVt() []byte,
ImportAccountsCSVt([]byte) error,
ExportAccountsCSVt() []byte.
*/

type Mybyte []byte

func (mb *Mybyte) Write(record AccountsList) {
	for name, value := range record {
		str := make([]string, 4)
		str[0] = fmt.Sprintf("%s", name)
		str[1] = fmt.Sprintf("%v", value.AccountType)
		str[2] = fmt.Sprintf("%.2f", value.Balance)
		str[3] = " "
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

//ExportAccountsCSVt
func ExportAccountsCSVt() []byte {
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

//ExportProdcuctsCSVt
func ExportProdcuctsCSVt() []byte {
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

//ImportProductsCSVt
func ImportProductsCSVt(data []byte) error {
	r := csv.NewReader(bytes.NewReader(data))
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	for _, record := range records {
		fmt.Println(record)
	}
	return nil
}
