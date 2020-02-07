package shopcompetition

import "testing"


func TestAddProductNormalRight(t *testing.T) {
	prdList := ProductList{}
	err := prdList.AddProduct(Product{Name: "Колбаса", Price: 250.50, Type: ProductNormal})
	t.Log(err.Error())
	if len(prdList) == 0  {
		t.Log("добавление  товара нормального в каталог не работает")
		t.Fail()
	}

}
func TestAddProductPremiumRight(t *testing.T) {
	prdList := ProductList{}
	err := prdList.AddProduct(Product{Name: "Cырок", Price: 1250.50, Type: ProductPremium})
	t.Log(err.Error())
	if len(prdList) == 0  {
		t.Log("добавление товара  премиального в каталог не работает")
		t.Fail()
	}
}
func TestAddProductSampleRight(t *testing.T) {
	prdList := ProductList{}
	err := prdList.AddProduct(Product{Name: "Cалфетка", Price: 1250.50, Type: ProductSample})
	t.Log(err.Error())
	if len(prdList) == 0  {
		t.Log("добавление Sample в каталог не работает")
		t.Fail()
	}
}
func TestAddProductNormalNameEmpty(t *testing.T) {
	prdList := ProductList{}
	err := prdList.AddProduct(Product{Name: " ", Price: 1250.50, Type: ProductNormal})
	t.Log(err.Error())
	if len(prdList) != 0  {
		t.Log("Можно добавить товар нормального типа с пустым именем")
		t.Fail()
	}
}
func TestAddProductPremiumNameEmpty(t *testing.T) {
	prdList := ProductList{}
	err := prdList.AddProduct(Product{Name: " ", Price: 1250.50, Type: ProductPremium})
	t.Log(err.Error())
	if len(prdList) != 0  {
		t.Log("Можно добавить товар премиального  типа с пустым именем")
		t.Fail()
	}
}
func TestAddProductSampleNameEmpty(t *testing.T) {
	prdList := ProductList{}
	err := prdList.AddProduct(Product{Name: " ", Price: 1250.50, Type: ProductSample})
	t.Log(err.Error())
	if len(prdList) != 0  {
		t.Log("Можно добавить товар премиального  типа с пустым именем")
		t.Fail()
	}
}
// Возможно добавить проверки пустоты объекта, slice

func TestAddMinusPriceProductPremium(t *testing.T) {
	prdList := ProductList{}
	err := prdList.AddProduct(Product{Name: "Шампунь", Price: -120.50, Type: ProductPremium})
	t.Log(err.Error())
	if len(prdList) != 0 {
			t.Log(" не работает заглушка по отрицательной цене  у продукта премиум")
			t.Fatal()
		}
	}
func TestAddMinusPriceProductNormal(t *testing.T) {
	prdList := ProductList{}
	err := prdList.AddProduct(Product{Name: "Шампунь", Price: -120.50, Type: ProductNormal})
	t.Log(err.Error())
	if len(prdList) != 0 {
		t.Log(" не работает заглушка по отрицательной цене у нормального продукта ")
		t.Fatal()
	}
}
func TestAddMinusPriceProductSample(t *testing.T) {
	prdList := ProductList{}
	err := prdList.AddProduct(Product{Name: "Шампунь", Price: -120.50, Type: ProductSample})
	t.Log(err.Error())
	if len(prdList) != 0 {
		t.Log(" не работает заглушка по отрицательной цене у нормального продукта ")
		t.Fatal()
	}
}

func TestAddNullPriceProductNormal(t *testing.T) {
	prdList := ProductList{}
	err := prdList.AddProduct(Product{Name: "Шампунь", Price: 0, Type: ProductNormal})
	t.Log(err.Error())
	if len(prdList) != 0 {
		t.Log(" Возможно нулевая цена продукта Normal")
		t.Fatal()
	}
}
func TestAddNullPriceProductPremium(t *testing.T) {
	prdList := ProductList{}
	err := prdList.AddProduct(Product{Name: "Шампунь", Price: 0, Type: ProductPremium})
	t.Log(err.Error())
	if len(prdList) != 0 {
		t.Log(" Возможно нулевая цена продукта Premium")
		t.Fatal()
	}
}
func TestAddNullPriceProductSample(t *testing.T) {
	prdList := ProductList{}
	err := prdList.AddProduct(Product{Name: "Шампунь", Price: 0, Type: ProductSample})
	t.Log(err.Error())
	if len(prdList) == 0 {
		t.Log(" Невозможно установить нулевую цена продукта Premium")
		t.Fatal()
	}
}
func TestAddNotNullPriceProductSample(t *testing.T) {
	prdList := ProductList{}
	err := prdList.AddProduct(Product{Name: "Шампунь", Price: 10, Type: ProductSample})
	t.Log(err.Error())
	if len(prdList) != 0 {
		t.Log("Возможна не нулевая цена продукта Sample")
		t.Fatal()
	}
}
func TestAddProductNormalReiteration(t *testing.T) {
	prdList := ProductList{
		Product{Name: "Колбаса", Price: 250.50, Type: ProductNormal},
	}
	err := prdList.AddProduct(Product{Name: "Колбаса", Price: 250.50, Type: ProductNormal})
	t.Log(err.Error())
	if len(prdList) == 2  {
		t.Log("добавление  товара с одинаковым именем")
		t.Fail()
	}

}
func TestAddProductPremiumReiteration(t *testing.T) {
	prdList := ProductList{
		Product{Name: "Колбаса", Price: 250.50, Type: ProductPremium},
	}
	err := prdList.AddProduct(Product{Name: "Колбаса", Price: 250.50, Type: ProductPremium})
	t.Log(err.Error())
	if len(prdList) == 2  {
		t.Log("добавление  товара с одинаковым именем")
		t.Fail()
	}

}
func TestAddProductSampleReiteration(t *testing.T) {
	prdList := ProductList{
		Product{Name: "Колбаса", Price: 0, Type: ProductSample},
	}
	err := prdList.AddProduct(Product{Name: "Колбаса", Price: 0, Type: ProductSample})
	t.Log(err.Error())
	if len(prdList) == 2  {
		t.Log("добавление  товара с одинаковым именем")
		t.Fail()
	}

}
func TestModifyProductPrice(t *testing.T) {
	prdList := ProductList{
		Product{Name: "Колбаса", Price: 250.50, Type: ProductNormal},
	}
	err := prdList.ModifyProduct(Product{Name: "Колбаса",
		Price: 450.50, Type: ProductNormal})
	t.Log(prdList)
	if err.ErrorCode() == StNil {
		t.Error(" не учтенный вариант обновления ")
		t.Fatal()
		return
	}
	t.Logf("%s %d ", err.Error(), err.ErrorCode())
}
func TestModifyProductChangeTypePremium(t *testing.T) {
	prdList := ProductList{
		Product{Name: "Колбаса", Price: 250.50, Type: ProductNormal},
	}
	err := prdList.ModifyProduct(Product{Name: "Колбаса",
		Price: 450.50, Type: ProductPremium})
	t.Log(prdList)
	if err.ErrorCode() == StNil {
		t.Error(" Не возможно сменить тип продукта")
		t.Fatal()
		return
	}
	t.Logf("%s %d ", err.Error(), err.ErrorCode())
}
func TestModifyProductChangeTypeSample(t *testing.T) {
	prdList := ProductList{
		Product{Name: "Колбаса", Price: 250.50, Type: ProductNormal},
	}
	err := prdList.ModifyProduct(Product{Name: "Колбаса",
		Price: 0, Type: ProductSample})
	t.Log(prdList)
	if err.ErrorCode() == StNil {
		t.Error(" Не возможно сменить тип продукта")
		t.Fatal()
		return
	}
	t.Logf("%s %d ", err.Error(), err.ErrorCode())
}
func TestModifyNullPriceProductNormal(t *testing.T) {
	prdList := ProductList{
		Product{Name: "Колбаса", Price: 250.50, Type: ProductNormal},
	}
	err := prdList.ModifyProduct(Product{Name: "Колбаса",
		Price: 0, Type: ProductNormal})
	t.Log(prdList)
	if err.ErrorCode() != StWrongPrice {
		t.Error(" ProductNormal можно изменить цену на нулевую")
		t.Fatal()
		return
	}
	t.Logf("%s %d ", err.Error(), err.ErrorCode())
}
func TestModifyNullPriceProductPremium(t *testing.T) {
	prdList := ProductList{
		Product{Name: "Колбаса", Price: 250.50, Type: ProductPremium},
	}
	err := prdList.ModifyProduct(Product{Name: "Колбаса",
		Price: 0, Type: ProductPremium})
	t.Log(prdList)
	if err.ErrorCode() != StWrongPrice {
		t.Error(" ProductPremium можно изменить цену на нулевую")
		t.Fatal()
		return
	}
	t.Logf("%s %d ", err.Error(), err.ErrorCode())
}
func TestModifyProductWrongPrice(t *testing.T) {
	prdList := ProductList{
		Product{Name: "Колбаса", Price: 250.50, Type: ProductNormal},
	}
	err := prdList.ModifyProduct(Product{Name: "Колбаса",
		Price: -450.50, Type: ProductNormal})
	t.Log(prdList)
	if err.ErrorCode() == StWrongPrice {
		t.Logf(" пытаемся обновить на отрицательную цену %.2f ", -450.50)
		return
	} else {
		t.Errorf("%s %d ", err.Error(), err.ErrorCode())
	}
}
func TestModifyProductWrongPriceSample(t *testing.T) {
	prdList := ProductList{
		Product{Name: "Колбаса", Price: 0, Type: ProductSample},
	}
	err := prdList.ModifyProduct(Product{Name: "Колбаса",
		Price: 2, Type: ProductNormal})
	t.Log(prdList)
	if err.ErrorCode() == StWrongPrice {
		t.Log(" Нельзя изменить цену наненулевую")
		return
	} else {
		t.Log("Можно изменить цену Sample на не нулевую")
	}
}
func TestRemoveProduct(t *testing.T) {
	prdList := ProductList{
		Product{Name: "Колбаса", Price: 250.50, Type: ProductNormal},
	}
	err := prdList.RemoveProduct("Колбаса")
	if err.ErrorCode() != StDef {
		t.Errorf("%s %d ", err.Error(), err.ErrorCode())
		return
	}
	t.Logf("%s %d ", err.Error(), err.ErrorCode())
}
