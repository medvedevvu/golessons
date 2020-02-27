package shop_competition

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var gbundlesList *BundlesList
var once1 sync.Once

//NewAccountsBundles конструктор
func NewBundlesList() *BundlesList {
	once1.Do(func() {
		gbundlesList = &BundlesList{}
	})
	return gbundlesList
}

//GetBundlesList возвращает каталог товаров
func GetBundlesList() *BundlesList {
	return gbundlesList
}

//RemoveBundle удалить комплект
func (bundlesList *BundlesList) RemoveBundle(name string) error {
	timer := time.NewTimer(time.Second)
	done := make(chan struct{}, 1)
	result := make(chan error, 1)

	go func() {
		defer close(done)
		globalMutex.Lock()
		_, ok := (*bundlesList)[name]
		globalMutex.Unlock()
		if !ok {
			result <- fmt.Errorf("Удаление: комплекта %s нет в каталоге", name)
			return
		}
		globalMutex.Lock()
		delete(*bundlesList, name)
		globalMutex.Unlock()
		result <- nil
		return
	}()

	select {
	case <-timer.C:
		return errors.New("Превышен интервал ожидания")
	case res := <-result:
		return res
	}
}

//ChangeDiscount сменить скидку
func (bundlesList *BundlesList) ChangeDiscount(name string, discount float32) error {
	timer := time.NewTimer(time.Second)
	done := make(chan struct{})
	result := make(chan error)

	go func() {
		defer close(done)
		globalMutex.Lock()
		vtemp, ok := (*bundlesList)[name]
		globalMutex.Unlock()
		if !ok {
			result <- fmt.Errorf("комплект %s не найден в каталоге", name)
			return
		}
		globalMutex.Lock()
		vtemp.Discount = discount
		globalMutex.Unlock()
		result <- nil
		return
	}()

	select {
	case <-timer.C:
		return errors.New("Превышен интервал ожидания")
	case res := <-result:
		return res
	}
}

//AddBundle добавить комплект
func (bundlesList *BundlesList) AddBundle(name string,
	main string,
	discount float32,
	additional ...string) error {
	timer := time.NewTimer(time.Second)
	done := make(chan struct{})
	result := make(chan error)

	go func() {
		defer close(done)
		globalMutex.Lock()
		_, ok := (*bundlesList)[name]
		globalMutex.Unlock()
		if ok {
			result <- fmt.Errorf("комплект %s уже есть в каталоге", name)
			return
		}
		globalMutex.Lock()
		vproductList := GetProductList() // получить каталог товаров
		product, ok := (*vproductList)[main]
		globalMutex.Unlock()
		if !ok {
			result <- fmt.Errorf("товар %s не найден в каталоге товаров", name)
			return
		}
		if product.Type == ProductSample {
			result <- fmt.Errorf("товар %s - пробник не формирует комплект ", name)
			return
		}
		if len(additional) == 0 {
			result <- fmt.Errorf("в комплекте всего один товар ")
			return
		}
		additional = append(additional, main)
		countSample := 0
		for _, item := range additional {
			if (*vproductList)[item].Type == ProductSample {
				countSample++ // посчитаем ProductSample
			}
		}
		if countSample > 1 {
			result <- fmt.Errorf("в комплекте может быть только один пробник ")
			return
		}
		if len(additional) == 2 && countSample == 1 {
			globalMutex.Lock()
			(*bundlesList)[name] = Bundle{ProductsName: additional,
				Type:     BundleSample,
				Discount: 1 - discount,
			}
			globalMutex.Unlock()
			result <- nil
			return
		} else {
			globalMutex.Lock()
			(*bundlesList)[name] = Bundle{ProductsName: additional,
				Type:     BundleNormal,
				Discount: 1 - discount,
			}
			globalMutex.Unlock()
		}
		result <- nil
		return
	}()

	select {
	case <-timer.C:
		return errors.New("Превышен интервал ожидания")
	case res := <-result:
		return res
	}

}
