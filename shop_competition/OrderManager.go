package shop_competition

import (
	"errors"
	"fmt"
	"time"
)

var AccountsOrdersMain = AccountsOrders{}

//PlaceOrder
func (accountsOrders *AccountsOrders) PlaceOrder(username string, order Order) error {
	timer := time.NewTimer(time.Second)

	vaccountsList := AccountsListMain
	vproductsList := ProductListMain
	vboundlesList := BundlesListMain

	errorChan := make(chan error)
	done := make(chan struct{})

	go func() {
		defer close(done)
		globalMutex.Lock()
		vuser, ok := vaccountsList[username]
		if !ok {
			errorChan <- fmt.Errorf(" пользователь %s не регистрирован", username)
			globalMutex.Unlock()
			return
		}
		var productPrice float32
		for _, productName := range order.ProductsName {
			vdiscount := getDiscount(vuser.AccountType, (vproductsList)[productName].Type)
			productPrice += (vproductsList)[productName].Price * vdiscount
		}
		var bundlePrice float32
		for _, bundleName := range order.BundlesName {
			vboundl := (vboundlesList)[bundleName]
			for _, productName := range vboundl.ProductsName {
				bundlePrice += (vproductsList)[productName].Price * vboundl.Discount
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
			globalMutex.Unlock()
			return
		}
		vuser.Balance -= order.TotalOrderPrice
		// запишем в историю списаний
		(*accountsOrders)[username] = append((*accountsOrders)[username], order)
		globalMutex.Unlock()
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
