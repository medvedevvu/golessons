package shop_competition

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

//var (
//	ProductListMain = ProductsList{}
//	globalMutex     sync.Mutex
//)

//CheckAttrsOfProduct проверка атрибутов товара
func (prodList *ProductsList) CheckAttrsOfProduct(productName string,
	product Product, operation OperationType) error {

	if len(strings.Trim(productName, "")) == 0 {
		return fmt.Errorf("у продукта нет названия")
	}
	prodList.Lock() // на момент проверки наличия , заблокируем
	_, ok := (*prodList).Products[productName]
	prodList.Unlock() // разблокируем
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
func (prodList *ProductsList) AddProduct(productName string,
	product Product) error {
	timer := time.NewTimer(time.Second)

	errorChan := make(chan error)
	done := make(chan struct{})
	go func() {
		defer close(done)
		err := prodList.CheckAttrsOfProduct(productName, product, Add)
		if err != nil {
			errorChan <- fmt.Errorf(" Добавление: ошибка проверки аттрибутов  товара %s", err)
			return
		}
		prodList.Lock()
		(*prodList).Products[productName] = &product
		prodList.Unlock()
		errorChan <- nil
		return
	}()
	for {
		select {
		case errorMsg := <-errorChan:
			return errorMsg
		case <-timer.C:
			return errors.New("Превышен интервал ожидания")
		}
	}

}

// ModifyProduct меняем товар в каталоге
func (prodList *ProductsList) ModifyProduct(productName string,
	product Product) error {
	timer := time.NewTimer(time.Second)

	errorChan := make(chan error)
	done := make(chan struct{})
	go func() {
		defer close(done)
		err := prodList.CheckAttrsOfProduct(productName, product, Edit)
		if err != nil {
			errorChan <- fmt.Errorf("Изменение : ошибка проверки аттрибутов  товара %s", err)
			return
		}
		prodList.Lock()
		(*prodList).Products[productName] = &product
		prodList.Unlock()
		errorChan <- nil
		return
	}()

	for {
		select {
		case errorMsg := <-errorChan:
			return errorMsg
		case <-timer.C:
			return errors.New("Превышен интервал ожидания")
		}
	}

}

// RemoveProduct удаляем товар из каталога
func (prodList *ProductsList) RemoveProduct(productName string) error {
	timer := time.NewTimer(time.Second)
	errorChan := make(chan error)
	done := make(chan struct{})
	go func() {
		defer close(done)
		prodList.Lock()
		_, ok := (*prodList).Products[productName]
		prodList.Unlock()
		if !ok {
			errorChan <- fmt.Errorf("Удаление: продукта %s нет в каталоге", productName)
			return
		}
		prodList.Lock()
		delete((*prodList).Products, productName)
		prodList.Unlock()
		errorChan <- nil
		return
	}()
	for {
		select {
		case errorMsg := <-errorChan:
			return errorMsg
		case <-timer.C:
			return errors.New("Превышен интервал ожидания")
		}
	}
}
