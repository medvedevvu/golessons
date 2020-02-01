package shopcompetition

import "fmt"

//ProductList каталог товаров
type ProductList []Product

// RemoveProduct - удаляем продукт
func (pl ProductList) RemoveProduct(name string) PCSError {
	state := StNotFnd
	for i, ipl := range pl {
		if ipl.Name == name {
			copy(pl[i:], pl[i+1:])    // Shift a[i+1:] left one index.
			pl[len(pl)-1] = Product{} // Erase last element (write zero value).
			pl = pl[:len(pl)-1]       // Truncate slice.
			state = StDef
		}
	}
	switch state {
	case StNotFnd:
		return PCSError{ErrCode: StNotFnd,
			ErrMsg: fmt.Sprintf("товар %s не найден в каталоге ", name)}
	case StDef:
		return PCSError{ErrCode: StDef,
			ErrMsg: fmt.Sprintf("товар %s удален из каталога ", name)}
	default:
		return PCSError{ErrCode: StNil, ErrMsg: ""}
	}
}

// ModifyProduct обновляем товар
func (pl *ProductList) ModifyProduct(pr Product) PCSError {
	state := StNotFnd
	for idx, item := range *pl {
		if pr.Price < 0 {
			state = StWrongPrice
			break
		}
		if item.Name == pr.Name {
			(*pl)[idx].Price = pr.Price
			(*pl)[idx].Type = pr.Type
			state = StDef
			break
		}
	}
	switch state {
	case StDef:
		return PCSError{ErrCode: StDef,
			ErrMsg: fmt.Sprintf("товар %v успешно обновлен ", pr)}
	case StNotFnd:
		return PCSError{ErrCode: StNotFnd,
			ErrMsg: fmt.Sprintf("товар %v не найден в каталоге - операция пропущена", pr)}
	case StWrongPrice:
		return PCSError{ErrCode: StWrongPrice,
			ErrMsg: fmt.Sprintf("товар %v имеет не верную цену", pr)}
	default:
		return PCSError{ErrCode: StNil, ErrMsg: ""}
	}

}

// AddProduct добавляем товар
func (pl *ProductList) AddProduct(pr Product) PCSError {
	state := StDef
	for idx, ipl := range *pl {
		if pr.Price < 0 {
			state = StWrongPrice
			break
		}
		if ipl.Name == pr.Name &&
			ipl.Price == pr.Price &&
			ipl.Type == pr.Type {
			state = StInProd
			break
		}
		if ipl.Name == pr.Name {
			switch interface{}(pr.Type).(type) {
			case ProductType:
				state = StNameAndTypeSame
				(*pl)[idx].Price = pr.Price
				(*pl)[idx].Type = pr.Type
			default:
				state = StTypeNotSame
			}
		}
		break
	}
	switch state {
	case StInProd:
		//return fmt.Errorf("товар %v уже есть каталоге- ничего не делаем", pr)
		return PCSError{ErrCode: StInProd,
			ErrMsg: fmt.Sprintf("товар %v уже есть каталоге- ничего не делаем", pr)}
	case StNameAndTypeSame:
		//		return fmt.Errorf("одноименный товар %v обновлен по цене и типу", pr)
		return PCSError{ErrCode: StNameAndTypeSame,
			ErrMsg: fmt.Sprintf("одноименный товар %v обновлен по цене и типу", pr)}
	case StTypeNotSame:
		//		return fmt.Errorf("товар %v имеет не верный тип - операция пропущена", pr)
		return PCSError{ErrCode: StTypeNotSame,
			ErrMsg: fmt.Sprintf("товар %v имеет не верный тип - операция пропущена", pr)}
	case StWrongPrice:
		//	return fmt.Errorf("товар %v имеет не верную цену", pr)
		return PCSError{ErrCode: StWrongPrice,
			ErrMsg: fmt.Sprintf("товар %v имеет не верную цену", pr)}
	case StDef:
		*pl = append(*pl, pr)
		//		return fmt.Errorf("товар %v добавлен в каталог", pr)
		return PCSError{ErrCode: StDef,
			ErrMsg: fmt.Sprintf("товар %v добавлен в каталог", pr)}
	default:
		return PCSError{ErrCode: StNil, ErrMsg: ""}
	}
}
