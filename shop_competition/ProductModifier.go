package shop_competition

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

var gproductList *ProductsList
var once3 sync.Once

// NewProductsList конструктор
func NewProductsList() *ProductsList {
	once3.Do(func() {
		gproductList = &ProductsList{}
	})
	return gproductList
}

// GetProductList возвращает каталог товаров
func GetProductList() *ProductsList {
	return gproductList
}

//CheckAttrsOfProduct проверка атрибутов товара
func (productsList *ProductsList) CheckAttrsOfProduct(productName string,
	product Product, operation OperationType) error {
	var localmutex sync.Mutex
	if len(strings.Trim(productName, "")) == 0 {
		return fmt.Errorf("у продукта нет названия")
	}
	localmutex.Lock() // на момент проверки наличия , заблокируем
	_, ok := (*productsList)[productName]
	localmutex.Unlock() // разблокируем
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
	if product.Price != 0 && product.Type == ProductSample {
		return fmt.Errorf("цена у пробников только 0 !")
	}
	return nil
}

// AddProduct добавляем товар в каталог
func (productsList *ProductsList) AddProduct(productName string,
	product Product) error {
	timer := time.NewTimer(time.Second)
	var localmutex sync.Mutex
	mthread := func() chan string {
		lchan := make(chan string)
		done := make(chan struct{})
		go func() {
			defer close(done)
			err := productsList.CheckAttrsOfProduct(productName, product, Add)
			if err != nil {
				lchan <- fmt.Sprintf(" Добавление: ошибка проверки аттрибутов  товара %s", err)
				return
			}
			localmutex.Lock()
			(*productsList)[productName] = &product
			localmutex.Unlock()
			lchan <- ""
			return
		}()
		return lchan
	}
	res := mthread()

	select {
	case localmess := <-res:
		if localmess == "" {
			return nil
		} else {
			return errors.New(localmess)
		}
	case <-timer.C:
		return errors.New("Превышен интервал ожидания")
	}

}

// ModifyProduct меняем товар в каталоге
func (productsList *ProductsList) ModifyProduct(productName string,
	product Product) error {
	timer := time.NewTimer(time.Second)
	var localmutex sync.Mutex
	mthread := func() chan string {
		lchan := make(chan string)
		done := make(chan struct{})
		go func() {
			defer close(done)
			err := productsList.CheckAttrsOfProduct(productName, product, Edit)
			if err != nil {
				lchan <- fmt.Sprintf("Изменение : ошибка проверки аттрибутов  товара %s", err)
				return
			}
			localmutex.Lock()
			(*productsList)[productName] = &product
			localmutex.Unlock()
			lchan <- ""
			return
		}()
		return lchan
	}
	res := mthread()

	select {
	case localmess := <-res:
		if localmess == "" {
			return nil
		} else {
			return errors.New(localmess)
		}
	case <-timer.C:
		return errors.New("Превышен интервал ожидания")
	}
}

// RemoveProduct удаляем товар из каталога
func (productsList *ProductsList) RemoveProduct(productName string) error {
	timer := time.NewTimer(time.Second)
	var localmutex sync.Mutex
	mthread := func() chan string {
		lchan := make(chan string)
		done := make(chan struct{})
		go func() {
			defer close(done)
			_, ok := (*productsList)[productName]
			if !ok {
				lchan <- fmt.Sprintf("Удаление: продукта %s нет в каталоге", productName)
				return
			}
			localmutex.Lock()
			delete(*productsList, productName)
			localmutex.Unlock()
			lchan <- ""
			return
		}()
		return lchan
	}
	res := mthread()
	select {
	case localmess := <-res:
		if localmess == "" {
			return nil
		} else {
			return errors.New(localmess)
		}
	case <-timer.C:
		return errors.New("Превышен интервал ожидания")
	}
}
