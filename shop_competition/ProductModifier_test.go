package shop_competition

import (
	"fmt"
	"testing"
)

func InitProductCatalog() *ProductsList {
	lproductList := NewProductsList()
	err := lproductList.AddProduct("колбаса", Product{Price: 125.23, Type: ProductNormal})
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	err = lproductList.AddProduct("водка", Product{Price: 400.23, Type: ProductNormal})
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	err = lproductList.AddProduct("сыр", Product{Price: 315.14, Type: ProductPremium})
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	err = lproductList.AddProduct("макароны", Product{Price: 47.14, Type: ProductNormal})
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	err = lproductList.AddProduct("зубочистка", Product{Price: 0.00, Type: ProductSample})
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	err = lproductList.AddProduct("вермишель", Product{Price: 11.20, Type: ProductNormal})
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	err = lproductList.AddProduct("хлеб", Product{Price: 32.10, Type: ProductNormal})
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	err = lproductList.AddProduct("цветы", Product{Price: 700.10, Type: ProductPremium})
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	err = lproductList.AddProduct("шампанское", Product{Price: 150.10, Type: ProductNormal})
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	err = lproductList.AddProduct("шоколад", Product{Price: 478.21, Type: ProductPremium})
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	err = lproductList.AddProduct("духи", Product{Price: 900.51, Type: ProductPremium})
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	err = lproductList.AddProduct("спички", Product{Price: 22.51, Type: ProductNormal})
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	return lproductList
}

func TestInitProductCatalog(t *testing.T) {
	InitProductCatalog()
	vals := GetProductList()
	if len(*vals) == 0 {
		t.Fatalf("Инициализация не прошла !")
	}
	//t.Log(*vals)
}

func TestAddProduct(t *testing.T) {
	vproductList := InitProductCatalog()
	err := vproductList.AddProduct("спички", Product{Price: 10.10, Type: ProductSample})
	if err == nil {
		t.Fatalf("Добавлен пробник с не нулевой стоимостью ")
	}
	err = vproductList.AddProduct("зажигалка", Product{Price: -100.10, Type: ProductPremium})
	if err == nil {
		t.Fatalf(" добавлен продукт с отрицательной стоимостью ")
	}
	err = vproductList.AddProduct("зажигалка", Product{Price: 0.0, Type: ProductPremium})
	if err == nil {
		t.Fatalf(" добавлен продукт с нулевой стоимостью ")
	}
	err = vproductList.AddProduct("сыр", Product{Price: 315.14, Type: ProductPremium})
	if err == nil {
		t.Fatalf(" добавлен одноименный продукт ")
	}

	err = vproductList.ModifyProduct("сыр", Product{Price: 0, Type: ProductPremium})
	if err == nil {
		t.Fatalf(" проставлена нулевая сумма не у пробника  ")
	}

	err = vproductList.ModifyProduct("сыр", Product{Price: -12.23, Type: ProductPremium})
	if err == nil {
		t.Fatalf(" проставлена отрицательная сумма не у пробника  ")
	}
	err = vproductList.ModifyProduct("зубочистка", Product{Price: 11.10, Type: ProductSample})
	if err == nil {
		t.Fatalf(" проставлена положительная сумма у пробника  ")
	}

}
