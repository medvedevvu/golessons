package shopcompetition

import "testing"

func TestAddProduct(t *testing.T) {
	products := []Product{
		Product{Name: "Колбаса", Price: 250.50, Type: ProductNormal},
		Product{Name: "Сыр", Price: 450.23, Type: ProductPremium},
		Product{Name: "Рыба", Price: 300.43, Type: ProductNormal},
		Product{Name: "Зубочистки", Price: 0.0, Type: ProductSample},
	}
	prdList := ProductList{}
	for _, Item := range products {
		err := prdList.AddProduct(Item)
		if len(prdList) == 0 {
			t.Log("добавление товара в каталог не работает")
			t.Fail()
		}
		t.Log(err.Error())
		t.Log(prdList)
	}
}

func TestAddWrongPriceProduct(t *testing.T) {
	products := []Product{
		Product{Name: "Колбаса", Price: 250.50, Type: ProductNormal},
		Product{Name: "Сыр", Price: -450.23, Type: ProductPremium},
	}
	prdList := ProductList{}
	for _, Item := range products {
		err := prdList.AddProduct(Item)
		if err.ErrorCode() == StWrongPrice {
			t.Logf("%s %d ", err.Error(), err.ErrorCode())
		}
	}
	for _, Item := range prdList {
		if Item.Price < 0 {
			t.Log(" не работает заглушка по отрицательной цене ")
			t.Fatal()
		}
	}
}

func TestModifyProduct(t *testing.T) {
	prdList := ProductList{
		Product{Name: "Колбаса", Price: 250.50, Type: ProductNormal},
	}
	err := prdList.ModifyProduct(Product{Name: "Колбаса",
		Price: 450.50, Type: ProductNormal})
	if err.ErrorCode() == StNil {
		t.Error(" не учтенный вариант обновления ")
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
	if err.ErrorCode() == StWrongPrice {
		t.Logf(" пытаемся обновить на отрицательную цену %.2f ", -450.50)
		return
	} else {
		t.Errorf("%s %d ", err.Error(), err.ErrorCode())
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
