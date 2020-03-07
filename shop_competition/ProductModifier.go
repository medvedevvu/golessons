package shop_competition

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

var (
	productListMain = ProductsList{}
	globalMutex     sync.Mutex
)

//CheckAttrsOfProduct проверка атрибутов товара
func (productsList *ProductsList) CheckAttrsOfProduct(productName string,
	product Product, operation OperationType) error {

	if len(strings.Trim(productName, "")) == 0 {
		return fmt.Errorf("у продукта нет названия")
	}
	globalMutex.Lock() // на момент проверки наличия , заблокируем
	_, ok := (*productsList)[productName]
	globalMutex.Unlock() // разблокируем
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

	lchan := make(chan error)
	done := make(chan struct{})
	go func() {
		defer close(done)
		err := productsList.CheckAttrsOfProduct(productName, product, Add)
		if err != nil {
			lchan <- fmt.Errorf(" Добавление: ошибка проверки аттрибутов  товара %s", err)
			return
		}
		globalMutex.Lock()
		(*productsList)[productName] = &product
		globalMutex.Unlock()
		lchan <- nil
		return
	}()
	for {
		select {
		case localmess := <-lchan:
			return localmess
		case <-timer.C:
			return errors.New("Превышен интервал ожидания")
		}
	}

}

// ModifyProduct меняем товар в каталоге
func (productsList *ProductsList) ModifyProduct(productName string,
	product Product) error {
	timer := time.NewTimer(time.Second)

	lchan := make(chan error)
	done := make(chan struct{})
	go func() {
		defer close(done)
		err := productsList.CheckAttrsOfProduct(productName, product, Edit)
		if err != nil {
			lchan <- fmt.Errorf("Изменение : ошибка проверки аттрибутов  товара %s", err)
			return
		}
		globalMutex.Lock()
		(*productsList)[productName] = &product
		globalMutex.Unlock()
		lchan <- nil
		return
	}()

	for {
		select {
		case localmess := <-lchan:
			return localmess
		case <-timer.C:
			return errors.New("Превышен интервал ожидания")
		}
	}

}

// RemoveProduct удаляем товар из каталога
func (productsList *ProductsList) RemoveProduct(productName string) error {
	timer := time.NewTimer(time.Second)
	lchan := make(chan error)
	done := make(chan struct{})
	go func() {
		defer close(done)
		globalMutex.Lock()
		_, ok := (*productsList)[productName]
		globalMutex.Unlock()
		if !ok {
			lchan <- fmt.Errorf("Удаление: продукта %s нет в каталоге", productName)
			return
		}
		globalMutex.Lock()
		delete(*productsList, productName)
		globalMutex.Unlock()
		lchan <- nil
		return
	}()
	for {
		select {
		case localmess := <-lchan:
			return localmess
		case <-timer.C:
			return errors.New("Превышен интервал ожидания")
		}
	}
}
