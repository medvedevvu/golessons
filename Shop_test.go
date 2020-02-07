package shopcompetition

import (
	"testing"
)

func CrtShop(p int) *Shop {

	/*
	  p = 1 данные есть
	  p = 0 данных нет
	*/
	accList := AccountList{}
	prdList := ProductList{}

	if p == 1 {
		nameList := []string{"Vasily", "Petr", "Oleg"}

		for _, name := range nameList {
			accList.Register(name)
		}

		products := []Product{
			Product{Name: "Колбаса", Price: 250.50, Type: ProductNormal},
			Product{Name: "Сыр", Price: 450.23, Type: ProductPremium},
			Product{Name: "Рыба", Price: 300.43, Type: ProductNormal},
			Product{Name: "Зубочистки", Price: 0.0, Type: ProductSample},
		}

		for _, Item := range products {
			prdList.AddProduct(Item)
		}
	}
	return NewShop(accList, prdList)
}

func TestExportShop(t *testing.T) {
	shop := CrtShop(1)

	datAcc, datPrd, errAcc, errPrd := shop.Export()

	if errAcc != nil || errPrd != nil {
		t.Fatalf(" export не работает ")
	}
	t.Logf("%s", datAcc)
	t.Logf("%s", datPrd)
}

func TestImportShop(t *testing.T) {
	shop := CrtShop(0)
	datAcc := []byte(`[{"Name":"Vasily","Balance":0,"AccountType":5},{"Name":"Petr","Balance":0,"AccountType":5},{"Name":"Oleg","Balance":0,"AccountType":5}]`)
	datPrd := []byte(`[{"Name":"Колбаса","Price":250.5,"Type":0},{"Name":"Сыр","Price":450.23,"Type":1},{"Name":"Рыба","Price":300.43,"Type":0},{"Name":"Зубочистки","Price":0,"Type":2}]`)

	errAcc, errPrd := shop.Import(datAcc, datPrd)
	if errAcc != nil || errPrd != nil {
		t.Fatalf(" import не работает ")
	}

	if len(shop.GetAccounts(SortByName)) == 0 || len(shop.ProductList) == 0 {
		t.Fatalf(" не работаеют методы импорта ")
	}

	t.Log(shop.GetAccounts(SortByName))
	t.Log(shop.ProductList)

}

func TestImportShopEmpty(t *testing.T) {
	shop := CrtShop(0)
	datAcc := []byte(` `)
	datPrd := []byte(` `)

	errAcc, errPrd := shop.Import(datAcc, datPrd)

	if errAcc != nil || errPrd != nil {
		t.Fatalf(" import не работает ")
	}

	if len(shop.GetAccounts(SortByName)) == 0 || len(shop.ProductList) == 0 {
		t.Fatalf(" не работаеют методы импорта ")
	}

	t.Log(shop.GetAccounts(SortByName))
	t.Log(shop.ProductList)

}
func TestImportShopEmptySlice(t *testing.T) {
	shop := CrtShop(0)
	datAcc := []byte("[]")
	datPrd := []byte("[]")

	errAcc, errPrd := shop.Import(datAcc, datPrd)

	if errAcc != nil || errPrd != nil {
		t.Fatalf(" import не работает ")
	}

	if len(shop.GetAccounts(SortByName)) == 0 || len(shop.ProductList) == 0 {
		t.Fatalf(" не работаеют методы импорта ")
	}

	t.Log(shop.GetAccounts(SortByName))
	t.Log(shop.ProductList)

}
func TestImportShopEmptyObject(t *testing.T) {
	shop := CrtShop(0)
	datAcc := []byte("{}")
	datPrd := []byte("{}")

	errAcc, errPrd := shop.Import(datAcc, datPrd)

	if errAcc != nil || errPrd != nil {
		t.Fatalf(" import не работает ")
	}

	if len(shop.GetAccounts(SortByName)) == 0 || len(shop.ProductList) == 0 {
		t.Fatalf(" не работаеют методы импорта ")
	}

	t.Log(shop.GetAccounts(SortByName))
	t.Log(shop.ProductList)

}