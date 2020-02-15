package shop_competition

import (
	"fmt"
	"reflect"
	"strings"
)

/*
type ProductModifier interface {
	AddProduct(productName string, product Product) error
	ModifyProduct(productName string, product Product) error
	RemoveProduct(productName string) error
}
type ProductsList map[string]*Product
*/
// NewProductsList конструктор
func NewProductsList() *ProductsList {
	return &ProductsList{}
}

//CheckAttrsOfProduct проверка атрибутов товара
func (productsList *ProductsList) CheckAttrsOfProduct(productName string,
	product Product, operation OperationType) error {

	if len(strings.Trim(productName, "")) == 0 {
		return fmt.Errorf("у продукта нет названия")
	}
	_, ok := (*productsList)[productName]

	if operation == Add {
		if ok {
			return fmt.Errorf("продукт %s уже есть", productName)
		}
	} else { // Edit
		if !ok {
			return fmt.Errorf("продукта %s нет в каталоге", productName)
		}
	}

	if product.Price < 0 {
		return fmt.Errorf("у продукта %s не допуситимая цена %f",
			productName, product.Price)
	}

	if product.Price == 0 && product.Type != ProductSample {
		return fmt.Errorf("0 цена только у пробников !")
	}

	if product.Type <= 0 {
		return fmt.Errorf("у продукта %s не верное значение типа %d",
			productName, product.Type)
	}
	if reflect.TypeOf(product.Type) != reflect.TypeOf(ProductNormal) {
		return fmt.Errorf("у продукта %s не допустимый тип %v",
			productName, product.Type)
	}

	return nil
}

// AddProduct добавляем товар в каталог
func (productsList *ProductsList) AddProduct(productName string,
	product Product) error {
	err := productsList.CheckAttrsOfProduct(productName, product, Add)
	if err != nil {
		return fmt.Errorf(" Добавление: ошибка проверки аттрибутов  товара %s", err)
	}
	(*productsList)[productName] = &product
	return nil
}

// ModifyProduct меняем товар в каталоге
func (productsList *ProductsList) ModifyProduct(productName string,
	product Product) error {
	err := productsList.CheckAttrsOfProduct(productName, product, Edit)
	if err != nil {
		return fmt.Errorf("Изменение : ошибка проверки аттрибутов  товара %s", err)
	}
	(*productsList)[productName] = &product
	return nil
}

// RemoveProduct удаляем товар из каталога
func (productsList *ProductsList) RemoveProduct(productName string) error {
	_, ok := (*productsList)[productName]
	if !ok {
		return fmt.Errorf("Удаление: продукта %s нет в каталоге", productName)
	}
	delete(*productsList, productName)
	return nil
}
