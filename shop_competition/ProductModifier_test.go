package shop_competition

import (
	"fmt"
	"sync"
	"testing"
)

func InitProductCatalog() ProductsList {
	lproductList := ProductsList{}
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
	for i := 0; i < 10000; i++ {
		err = lproductList.AddProduct(fmt.Sprintf("Продукт %d", i), Product{Price: 10.51, Type: ProductNormal})
		if err != nil {
			fmt.Printf("%s\n", err)
		}
	}
	return lproductList
}

func TestInitProductCatalog(t *testing.T) {
	vals := InitProductCatalog()
	if len(vals) == 0 {
		t.Fatalf("Инициализация не прошла !")
	}
}

func TestAddProduct(t *testing.T) {
	var wg sync.WaitGroup
	vproductList := ProductsList{}
	var vproductData = []struct {
		name         string
		productPrice float32
		productType  ProductType
	}{
		{"колбаса", 125.23, ProductNormal},
		{"сосики", 755.24, ProductNormal},
		{"зефир", 999.24, ProductPremium},
		{"зубочистка", 0, ProductSample},
	}
	wg.Add(len(vproductData))
	errchan := make(chan error, len(vproductData))
	for _, el := range vproductData {
		go func(name string, productPrice float32, productType ProductType) {
			defer wg.Done()
			err := vproductList.AddProduct(name,
				Product{Price: productPrice, Type: productType})
			if err != nil {
				errchan <- fmt.Errorf("продукт %s c ошибкой %s", name, err)
			}
			return
		}(el.name, el.productPrice, el.productType)
	}
	wg.Wait()
	select {
	case erroLog := <-errchan:
		for value := range errchan {
			t.Fatalf("ошибка добавлениея товара %v\n", value)
		}
		t.Fatalf("ошибка добавлениея товара %v\n", erroLog)
	default:
	}
	if len(vproductList) == 0 {
		t.Fatalf("ничего не добавилось \n")
	}
	// "если добавили все , проверим что добавилось"
	for _, product := range vproductData {
		product_local, ok := vproductList[product.name]
		if !ok {
			t.Logf(" продукт %s не добавлен \n", product.name)
			t.Fail()
		}
		if product_local.Price != product.productPrice ||
			product_local.Type != product.productType {
			t.Fatalf("продукт %s  в базе: цена %v  <> исходник: цена %v \n"+
				"  в базе: тип %v  <> исходник: тип %v \n", product.name,
				product_local.Price, product.productPrice,
				product_local.Type, product.productType)
		}
	}
}
func TestRulesAddProduct(t *testing.T) {
	vproductList := InitProductCatalog()

	var wg sync.WaitGroup
	var times_ int = 5
	wg.Add(times_)

	go func() {
		defer wg.Done()

		err := vproductList.AddProduct("спички", Product{Price: 10.10, Type: ProductSample})
		if err == nil {
			t.Fatalf("Добавлен пробник с не нулевой стоимостью ")
		}

	}()
	go func() {
		defer wg.Done()
		err := vproductList.AddProduct("зажигалка", Product{Price: -100.10, Type: ProductPremium})
		if err == nil {
			t.Fatalf(" добавлен продукт с отрицательной стоимостью ")
		}
	}()
	go func() {
		defer wg.Done()
		err := vproductList.AddProduct("зажигалка", Product{Price: 0.0, Type: ProductPremium})
		if err == nil {
			t.Fatalf(" добавлен продукт с нулевой стоимостью ")
		}
	}()
	go func() {
		defer wg.Done()
		err := vproductList.AddProduct("сыр", Product{Price: 315.14, Type: ProductPremium})
		if err == nil {
			t.Fatalf(" добавлен одноименный продукт ")
		}
	}()
	go func() {
		defer wg.Done()
		err := vproductList.AddProduct("куркума", Product{Price: 215.14, Type: ProductPremium})
		if err != nil {
			t.Fatalf(" не добавлен нормальный продукт %s \n", err)
		}
	}()
	wg.Wait()

}

func TestModifyProduct(t *testing.T) {
	vproductList := InitProductCatalog()
	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		err := vproductList.ModifyProduct("сыр", Product{Price: 0, Type: ProductPremium})
		if err == nil {
			t.Fatalf(" проставлена нулевая сумма ")
		}
	}()
	go func() {
		defer wg.Done()
		err := vproductList.ModifyProduct("сыр", Product{Price: -12.23, Type: ProductPremium})
		if err == nil {
			t.Fatalf(" проставлена отрицательная сумма ")
		}
	}()
	go func() {
		defer wg.Done()
		err := vproductList.ModifyProduct("зубочистка", Product{Price: 11.10, Type: ProductSample})
		if err == nil {
			t.Fatalf(" проставлена положительная сумма у пробника  ")
		}
	}()
	wg.Wait()

}

func TestRemoveProduct(t *testing.T) {
	vproductList := InitProductCatalog()
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		err := vproductList.RemoveProduct("шоколад")
		if err != nil {
			t.Fatalf(" ошибка удаления ")
		}
	}()

	go func() {
		defer wg.Done()
		err := vproductList.RemoveProduct("шомпанское")
		if err == nil {
			t.Fatalf(" произошло удаление отсутствующего товар %s ", err)
		}
	}()

	wg.Wait()

}
