package shop_competition

import (
	"errors"
	"fmt"
	"time"
)

//RemoveBundle удалить комплект
func (bndlList *BundlesList) RemoveBundle(name string) error {
	timer := time.NewTimer(time.Second)
	done := make(chan struct{}, 1)
	result := make(chan error, 1)

	go func() {
		defer close(done)
		bndlList.Lock()
		_, ok := bndlList.BundleList[name]
		bndlList.Unlock()
		if !ok {
			result <- fmt.Errorf("Удаление: комплекта %s нет в каталоге", name)
			return
		}
		bndlList.Lock()
		delete(bndlList.BundleList, name)
		bndlList.Unlock()
		result <- nil
		return
	}()

	for {
		select {
		case <-timer.C:
			return errors.New("Превышен интервал ожидания")
		case res := <-result:
			return res
		default:
		}
	}
}

//ChangeDiscount сменить скидку
func (bndlList *BundlesList) ChangeDiscount(name string, discount float32) error {
	timer := time.NewTimer(time.Second)
	done := make(chan struct{})
	result := make(chan error)

	go func() {
		defer close(done)
		bndlList.Lock()
		vtemp, ok := bndlList.BundleList[name]
		bndlList.Unlock()
		if !ok {
			result <- fmt.Errorf("комплект %s не найден в каталоге", name)
			return
		}
		bndlList.Lock()
		vtemp.Discount = discount
		bndlList.Unlock()
		result <- nil
		return
	}()

	for {
		select {
		case <-timer.C:
			return errors.New("Превышен интервал ожидания")
		case res := <-result:
			return res
		default:
		}
	}
}

//AddBundle добавить комплект
func (bndlList *BundlesList) AddBundle(name string,
	main string,
	discount float32,
	shop *ShopBase,
	additional ...string) error {
	timer := time.NewTimer(time.Second)
	done := make(chan struct{})
	result := make(chan error)

	go func() {
		defer close(done)
		defer bndlList.Unlock()

		bndlList.Lock()
		_, ok := bndlList.BundleList[name]
		if ok {
			result <- fmt.Errorf("комплект %s уже есть в каталоге", name)
			return
		}
		// получить каталог товаров
		product, ok := shop.ProductListWithMutex.Products[main]
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
			if shop.ProductListWithMutex.Products[item].Type == ProductSample {
				countSample++ // посчитаем ProductSample
			}
		}
		if countSample > 1 {
			result <- fmt.Errorf("в комплекте может быть только один пробник ")
			return
		}
		if len(additional) == 2 && countSample == 1 {
			bndlList.BundleList[name] = Bundle{ProductsName: additional,
				Type:     BundleSample,
				Discount: 1 - discount,
			}
			result <- nil
			return
		} else {
			bndlList.BundleList[name] = Bundle{ProductsName: additional,
				Type:     BundleNormal,
				Discount: 1 - discount,
			}
		}
		result <- nil
		return
	}()

	for {
		select {
		case <-timer.C:
			return errors.New("Превышен интервал ожидания")
		case res := <-result:
			return res
		default:
		}
	}
}
