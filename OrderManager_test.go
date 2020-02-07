package shopcompetition

import (
	"fmt"
	"testing"
)

var acc Account = Account{Name: "Viktor",
	Balance:     9000.32,
	AccountType: AccountPremium}

var prl []Product = []Product{
	Product{Name: "Фуагра", Price: 333.14, Type: ProductPremium},
	Product{Name: "Шомпанское", Price: 150.23, Type: ProductNormal},
	Product{Name: "Зубочистки", Price: 0.0, Type: ProductSample},
}

var order Order = Order{
	[]Product{Product{Name: "Колбаса", Price: 451.23, Type: ProductNormal},
		Product{Name: "Сыр", Price: 445.27, Type: ProductPremium},
		Product{Name: "Зубочистки", Price: 0.0, Type: ProductSample},
	},
	[]Bundle{
		{Products: prl,
			Type:     BundleNormal,
			Discount: 0.75,
		},
	},
}

var prdl ProductList = []Product{
	Product{Name: "Фуагра", Price: 333.14, Type: ProductPremium},
	Product{Name: "Шомпанское", Price: 150.23, Type: ProductNormal},
	Product{Name: "Колбаса", Price: 451.23, Type: ProductNormal},
	Product{Name: "Сыр", Price: 445.27, Type: ProductPremium},
	Product{Name: "Зубочистки", Price: 0.0, Type: ProductSample},
}

func TestCalculateOrder(t *testing.T) {
	ordSumm, err := prdl.CalculateOrder(acc, order, 0)
	if err != nil {
		t.Error(err)
	}
	if ordSumm != 1395.5885 {
		t.Error("Не верный счет заказа")
	}
	t.Log(ordSumm)
}

func TestPlaceOrder(t *testing.T) {
	t.Log(acc)
	err := prdl.PlaceOrder(&acc, order, 0)
	if err != nil {
		if err.Error() == fmt.Sprintf("%d", StNotMany) {
			t.Error("Не хватает денег на счете")
			return
		} else {
			t.Log(err.Error())
			t.Fail()
			return
		}
	}
	t.Log(acc)
}

func TestPlaceOrderSmallMoney(t *testing.T) {
	t.Log(acc)
	acc.Balance = 0
	err := prdl.PlaceOrder(&acc, order, 0)
	if err != nil {
		if err.Error() == fmt.Sprintf("%d", StNotMany) {
			t.Log("Не хватает денег на счете ")
			return
		} else {
			t.Log(err.Error())
			t.Fail()
			return
		}
	}
	t.Log(acc)
}

var AccountCheckSumm Account = Account{Name: "Viktor",
	Balance:     9000.32,
	AccountType: AccountPremium}


var OrderCheckSumm Order = Order{
	[]Product{Product{Name: "Колбаса", Price: 0, Type: ProductNormal},
		Product{Name: "Сыр", Price: 0, Type: ProductPremium},
		Product{Name: "Зубочистки", Price: 0.0, Type: ProductSample},
	},
	[]Bundle{
		{Products: prl,
			Type:     BundleNormal,
			Discount: 0.75,
		},
	},
}

var ProductListCheckSumm ProductList = []Product{
	Product{Name: "Фуагра", Price: 0, Type: ProductPremium},
	Product{Name: "Шомпанское", Price: 0, Type: ProductNormal},
	Product{Name: "Колбаса", Price: 0, Type: ProductNormal},
	Product{Name: "Сыр", Price: 0, Type: ProductPremium},
	Product{Name: "Зубочистки", Price: 0, Type: ProductSample},
	Product{Name: "Соль", Price: 0, Type: ProductSample},
}

func TestCalculateOrderCheckSumNull(t *testing.T) {
	ordSum, err := prdl.CalculateOrder(AccountCheckSumm, OrderCheckSumm, 0)
	if err != nil {
		t.Error(err)
	}
	if ordSum != 0 {
		t.Error("Не верный счет заказа")
	}
	t.Log(ordSum)
}
//Проверка как работает с неправильными ценниками
func TestPlaceOrderCheckSum(t *testing.T) {
	err := prdl.PlaceOrder(&AccountCheckSumm, OrderCheckSumm, 0)
	if err != nil {
		if err.Error() == fmt.Sprintf("%d", StNotMany) {
			t.Error("Не хватает денег на счете")
			return
		} else {
			t.Log(err.Error())
			t.Fail()
			return
		}
	}
}