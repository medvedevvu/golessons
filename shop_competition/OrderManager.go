package shop_competition

import (
	"errors"
	"fmt"
	"time"
)

//var AccountsOrdersMain = AccountsOrders{}

//PlaceOrder
func (accOrders *AccountsOrders) PlaceOrder(username string, order Order, shop *ShopBase) error {
	timer := time.NewTimer(time.Second)

	errorChan := make(chan error)
	done := make(chan struct{})

	go func() {
		defer close(done)

		accOrders.Lock()
		vuser, ok := shop.AccountsListWithMutex.Accounts[username]
		if !ok {
			errorChan <- fmt.Errorf(" пользователь %s не регистрирован", username)
			accOrders.Unlock()
			return
		}
		var productPrice float32
		for _, productName := range order.ProductsName {
			vdiscount := getDiscount(vuser.AccountType,
				shop.ProductListWithMutex.Products[productName].Type,
			)
			productPrice += shop.ProductListWithMutex.Products[productName].Price * vdiscount
		}
		var bundlePrice float32
		for _, bundleName := range order.BundlesName {
			vboundl := shop.BundlesListWithMutex.BundleList[bundleName]
			for _, productName := range vboundl.ProductsName {
				bundlePrice += shop.ProductListWithMutex.Products[productName].Price * vboundl.Discount
			}
		}

		order.BundlesPrice = bundlePrice
		order.ProductsPrice = productPrice
		order.TotalOrderPrice = bundlePrice + productPrice

		if (vuser.Balance - order.TotalOrderPrice) <= 0 {
			errorChan <- fmt.Errorf(" %s : остаток %f - списание %f = %f - мало на счету",
				username,
				vuser.Balance,
				order.TotalOrderPrice,
				vuser.Balance-order.TotalOrderPrice)
			accOrders.Unlock()
			return
		}
		vuser.Balance -= order.TotalOrderPrice
		// запишем в историю списаний
		accOrders.AccountOrders[username] = append(accOrders.AccountOrders[username], order)
		accOrders.Unlock()
		errorChan <- nil
		return
	}()

	for {
		select {
		case errorMsg := <-errorChan:
			return errorMsg
		case <-timer.C:
			return errors.New("Превышен интервал ожидания")
		default:
		}
	}
}

func getDiscount(accountType AccountType, productType ProductType) (discount float32) {
	discount = 1.0
	switch accountType {
	case AccountPremium:
		switch productType {
		case ProductPremium:
			discount = 0.8
		case ProductNormal:
			discount = 1.5
		}
	case AccountNormal:
		switch productType {
		case ProductPremium:
			discount = 0.95
		case ProductNormal:
			discount = 1
		}
		//default:
	}
	return discount
}
