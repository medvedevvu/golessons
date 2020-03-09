package shop_competition

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func InitSmallProductCatalog() ProductsList {
	ProductListMain = ProductsList{}
	err := ProductListMain.AddProduct("колбаса", Product{Price: 125.23, Type: ProductNormal})
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	err = ProductListMain.AddProduct("водка", Product{Price: 400.23, Type: ProductNormal})
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	err = ProductListMain.AddProduct("сыр", Product{Price: 315.14, Type: ProductPremium})
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	err = ProductListMain.AddProduct("макароны", Product{Price: 47.14, Type: ProductNormal})
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	err = ProductListMain.AddProduct("зубочистка", Product{Price: 0.00, Type: ProductSample})
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	err = ProductListMain.AddProduct("вермишель", Product{Price: 11.20, Type: ProductNormal})
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	err = ProductListMain.AddProduct("хлеб", Product{Price: 32.10, Type: ProductNormal})
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	err = ProductListMain.AddProduct("цветы", Product{Price: 700.10, Type: ProductPremium})
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	err = ProductListMain.AddProduct("шампанское", Product{Price: 150.10, Type: ProductNormal})
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	err = ProductListMain.AddProduct("шоколад", Product{Price: 478.21, Type: ProductPremium})
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	err = ProductListMain.AddProduct("духи", Product{Price: 900.51, Type: ProductPremium})
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	err = ProductListMain.AddProduct("спички", Product{Price: 22.51, Type: ProductNormal})
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	return ProductListMain
}

func InitProductCatalog() ProductsList {
	ProductListMain = ProductsList{}
	err := ProductListMain.AddProduct("колбаса", Product{Price: 125.23, Type: ProductNormal})
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	err = ProductListMain.AddProduct("водка", Product{Price: 400.23, Type: ProductNormal})
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	err = ProductListMain.AddProduct("сыр", Product{Price: 315.14, Type: ProductPremium})
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	err = ProductListMain.AddProduct("макароны", Product{Price: 47.14, Type: ProductNormal})
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	err = ProductListMain.AddProduct("зубочистка", Product{Price: 0.00, Type: ProductSample})
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	err = ProductListMain.AddProduct("вермишель", Product{Price: 11.20, Type: ProductNormal})
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	err = ProductListMain.AddProduct("хлеб", Product{Price: 32.10, Type: ProductNormal})
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	err = ProductListMain.AddProduct("цветы", Product{Price: 700.10, Type: ProductPremium})
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	err = ProductListMain.AddProduct("шампанское", Product{Price: 150.10, Type: ProductNormal})
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	err = ProductListMain.AddProduct("шоколад", Product{Price: 478.21, Type: ProductPremium})
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	err = ProductListMain.AddProduct("духи", Product{Price: 900.51, Type: ProductPremium})
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	err = ProductListMain.AddProduct("спички", Product{Price: 22.51, Type: ProductNormal})
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	for i := 0; i < 1000; i++ {
		err = ProductListMain.AddProduct(fmt.Sprintf("Продукт %d", i), Product{Price: 10.51, Type: ProductNormal})
		if err != nil {
			fmt.Printf("%s\n", err)
		}
	}
	return ProductListMain
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
func TestAddProductsWithWrongValues(t *testing.T) {
	vproductList := InitSmallProductCatalog()
	type ProductDataType struct {
		Name         string
		ProductPrice float32
		ProductType  ProductType
		ErrMessage   string
	}

	type ProductDataArrayType []ProductDataType

	vproductData := ProductDataArrayType{
		{"спички", 10.10, ProductSample, "Добавлен пробник с не нулевой стоимостью"},
		{"зажигалка", -100.14, ProductPremium, "добавлен продукт с отрицательной стоимостью"},
		{"пепельница", 0, ProductPremium, "добавлен продукт с нулевой стоимостью"},
		{"сыр", 315.14, ProductSample, "добавлен одноименный продукт"},
		{"куркума", 0, ProductNormal, "добавлен продукт с нулевой стоимостью"},
	}

	var wg sync.WaitGroup
	wg.Add(len(vproductData))

	for _, product := range vproductData {
		go func(product ProductDataType) {
			defer wg.Done()
			err := vproductList.AddProduct(product.Name,
				Product{Price: product.ProductPrice,
					Type: product.ProductType},
			)
			if err == nil {
				t.Fatalf(product.ErrMessage)
			}
		}(product)
	}
	wg.Wait()
}

func TestOptModifyProducts(t *testing.T) {
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
	// "если добавили все , изменим тип и цену"
	const differentPrice float32 = 154.69
	const differentType ProductType = ProductNormal
	for _, product := range vproductData {
		_, ok := vproductList[product.name]
		if !ok {
			t.Logf(" продукта %s нет в базе \n", product.name)
			t.Fail()
		}
		err := vproductList.ModifyProduct(product.name,
			Product{Price: differentPrice, Type: differentType})
		if err != nil {
			t.Fatalf(" не смогли выполнить обновление %s \n", err)
		}
	}
	// проверим , что сменилось у всех
	for name, value := range vproductList {
		if value.Price != differentPrice &&
			value.Type != differentType {
			t.Fatalf("обновление %s не выполнено", name)
		}
	}
}
func TestModifyProductsWithWrongValues(t *testing.T) {
	vproductList := InitProductCatalog()
	var wg sync.WaitGroup

	const times_ int = 3
	wg.Add(times_)

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

	const times_ int = 2
	wg.Add(times_)

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

func TestCRUDWithProducts(t *testing.T) {
	/*
	 пытаюсь выполнить сразу все CRUD в потоках
	 так как вероятность возникновения ошибки при отсуствии
	 данных по причине их удаления велика убрал Fatal-ы Feil-ы
	*/

	const operationsCount int = 2
	testedAddProduct := Product{Price: 144.25, Type: ProductNormal}
	testedUpdateProduct := Product{Price: 177.33, Type: ProductPremium}
	const testName string = "XXX"
	vproductList := &ProductsList{}

	timer := time.NewTimer(time.Second * 5)
	doneCount := make(chan string, 2) // для операций

	go func() {
		defer func() { doneCount <- "add" }()
		err := vproductList.AddProduct(testName, testedAddProduct)
		if err != nil {
			t.Logf(" не выполнилось добавление %s \n", err)
			return
		}
		t.Log("-- добавление выполнено \n")
	}()

	go func() {
		defer func() { doneCount <- "mdf" }()

	Loop:
		for {
			select {
			case vdone := <-doneCount:
				doneCount <- vdone
				if vdone == "add" {
					break Loop
				}
			default:
			}
		}

		err := vproductList.ModifyProduct(testName, testedUpdateProduct)
		if err != nil {
			t.Logf(" не выполнилось обновление %s \n", err)
			return
		}
		t.Log("-- обновление выполнено \n")
	}()

	go func() {
		defer func() { doneCount <- "dlt" }()
	Loop:
		for {
			select {
			case vdone := <-doneCount:
				doneCount <- vdone
				if vdone == "add" ||
					vdone == "mdf" {
					break Loop
				}
			default:
			}
		}
		err := vproductList.RemoveProduct(testName)
		if err != nil {
			t.Logf(" не выполнилось удаление %s \n", err)
			return
		}
		t.Log("-- удаление выполнено \n")
	}()

	var optinsCount int = 0
Loop:
	for {
		select {
		case <-doneCount:
			optinsCount++
			if optinsCount >= 2 {
				break Loop
			}
		case <-timer.C:
			t.Log(" превышен интервал ожидания ")
			break Loop
		default:
		}
	}

}
