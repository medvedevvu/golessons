package shop_competition

import "fmt"

var gbundlesList *BundlesList

//NewAccountsBundles конструктор
func NewBundlesList() *BundlesList {
	gbundlesList = &BundlesList{}
	return gbundlesList
}

//GetBundlesList возвращает каталог товаров
func GetBundlesList() *BundlesList {
	return gbundlesList
}

//RemoveBundle удалить комплект
func (bundlesList *BundlesList) RemoveBundle(name string) error {
	_, ok := (*bundlesList)[name]
	if !ok {
		return fmt.Errorf("Удаление: комплекта %s нет в каталоге", name)
	}
	delete(*bundlesList, name)
	return nil
}

//ChangeDiscount сменить скидку
func (bundlesList *BundlesList) ChangeDiscount(name string, discount float32) error {
	vtemp, ok := (*bundlesList)[name]
	if !ok {
		return fmt.Errorf("комплект %s не найден в каталоге", name)
	}
	vtemp.Discount = discount
	return nil
}

//AddBundle добавить комплект
func (bundlesList *BundlesList) AddBundle(name string,
	main string,
	discount float32,
	additional ...string) error {
	_, ok := (*bundlesList)[name]
	if ok {
		return fmt.Errorf("комплект %s уже есть в каталоге", name)
	}
	vproductList := GetProductList() // получить каталог товаров
	product, ok := (*vproductList)[main]
	if !ok {
		return fmt.Errorf("товар %s не найден в каталоге товаров", name)
	}
	if product.Type == ProductSample {
		return fmt.Errorf("товар %s - пробник не формирует комплект ", name)
	}
	if len(additional) == 0 {
		return fmt.Errorf("в комплекте всего один товар ")
	}
	additional = append(additional, main)
	countSample := 0
	for _, item := range additional {
		if (*vproductList)[item].Type == ProductSample {
			countSample++ // посчитаем ProductSample
		}
	}
	if countSample > 1 {
		return fmt.Errorf("в комплекте может быть только один пробник ")
	}
	if len(additional) == 2 && countSample == 1 {
		(*bundlesList)[name] = Bundle{ProductsName: additional,
			Type:     BundleSample,
			Discount: 1 - discount,
		}
	} else {
		(*bundlesList)[name] = Bundle{ProductsName: additional,
			Type:     BundleNormal,
			Discount: 1 - discount,
		}
	}
	return nil
}
