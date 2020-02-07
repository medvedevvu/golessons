package shopcompetition

import (
	"testing"
)

func TestAccountExport(t *testing.T) {
	accList := new(AccountList)
	nameList := []string{"Vasily", "Petr", "Oleg"}
	for _, name := range nameList {
		accList.Register(name)
	}
	for _, name := range nameList {
		accList.AddBalance(name, 999.23)
	}

	data, err := accList.Export()
	if err != nil {
		t.Fatalf("Сбой маршалинга JSON %s", err)
	}
	t.Logf("%s\n", data)
}

func TestAccountImport(t *testing.T) {
	acc := AccountList{}
	data := []byte(`[{"Name":"Vasily","Balance":999.23,"AccountType":5},{"Name":"Petr","Balance":999.23,"AccountType":5},{"Name":"Oleg","Balance":999.23,"AccountType":5}]`)
	err := acc.Import(data)
	if err != nil {
		t.Fatalf("Сбой демаршалинга JSON %s", err)
	}
	t.Log(acc)
}

func TestProductListExport(t *testing.T) {
	products := []Product{
		Product{Name: "Колбаса", Price: 250.50, Type: ProductNormal},
		Product{Name: "Сыр", Price: 450.23, Type: ProductPremium},
		Product{Name: "Рыба", Price: 300.43, Type: ProductNormal},
		Product{Name: "Зубочистки", Price: 0.0, Type: ProductSample},
	}
	prdList := ProductList{}
	for _, Item := range products {
		prdList.AddProduct(Item)
	}

	data, err := prdList.Export()
	if err != nil {
		t.Fatalf("Сбой маршалинга JSON %s", err)
	}
	t.Logf("%s\n", data)

}

func TestProductListImport(t *testing.T) {
	prdList := ProductList{}
	data := []byte(`[{"Name":"Колбаса","Price":250.5,"Type":0},{"Name":"Сыр","Price":450.23,"Type":1},{"Name":"Рыба","Price":300.43,"Type":0},{"Name":"Зубочистки","Price":0,"Type":2}]`)
	err := prdList.Import(data)
	if err != nil {
		t.Fatalf("Сбой демаршалинга JSON %s", err)
	}
	t.Log(prdList)
}
