package shopcompetition

import "fmt"

/*
AddProduct(Product) error
ModifyProduct(Product) error
RemoveProduct(name string) error
*/
//ProductList каталог товаров
type ProductList []Product

// RemoveProduct - удаляем продукт
func (pl ProductList) RemoveProduct(name string) error {
	isexists := exists4
	for i, ipl := range pl {
		if ipl.Name == name {
			copy(pl[i:], pl[i+1:])    // Shift a[i+1:] left one index.
			pl[len(pl)-1] = Product{} // Erase last element (write zero value).
			pl = pl[:len(pl)-1]       // Truncate slice.
			isexists = exists3
		}
	}
	switch isexists {
	case exists4:
		return fmt.Errorf(" товар %s не найден в каталоге ", name)
	default:
		return nil
	}

}

// ModifyProduct обновляем товар
func (pl ProductList) ModifyProduct(pr Product) error {
	return pl.AddProduct(pr)
}

// AddProduct добавляем товар
func (pl ProductList) AddProduct(pr Product) error {
	isexists := exists3
	for _, ipl := range pl {
		if ipl.Name == pr.Name && ipl.Price == pr.Price && ipl.Type == pr.Type {
			isexists = exists0
			break
		}
		if ipl.Name == pr.Name {
			switch interface{}(pr.Type).(type) {
			case ProductType:
				isexists = exists1
				pr.Price = ipl.Price
				pr.Type = ipl.Type
			default:
				isexists = exists2
			}
		}
		break
	}
	switch isexists {
	case exists0:
		return fmt.Errorf("товар %v уже есть каталоге- ничего не делаем", pr)
	case exists1:
		return fmt.Errorf("одноименный товар %v обновлен по цене и типу", pr)
	case exists2:
		return fmt.Errorf("товар %v имеет не верный тип - операция пропущена", pr)
	}
	return nil
}
