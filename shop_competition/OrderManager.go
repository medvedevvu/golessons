package shop_competition

import "fmt"

var gaccountsOrders *AccountsOrders

//NewAccountsOrders конструктор
func NewAccountsOrders() *AccountsOrders {
	gaccountsOrders = &AccountsOrders{}
	return gaccountsOrders
}

// GetAccountsOrders получить список список Account -> Orders
func GetAccountsOrders() *AccountsOrders {
	return gaccountsOrders
}

// PlaceOrder
func (accountsOrders *AccountsOrders) PlaceOrder(username string, order Order) error {
	vaccountsList := GetAccountsList()
	vproductsList := GetProductList()
	vboundlesList := GetBundlesList()
	vuser, ok := (*vaccountsList)[username]
	if !ok {
		return fmt.Errorf(" пользователь %s не регистрирован", username)
	}
	var productPrice float32
	for _, productName := range order.ProductsName {
		vdiscount := getDiscount(vuser.AccountType, (*vproductsList)[productName].Type)
		productPrice += (*vproductsList)[productName].Price * vdiscount
	}
	var bundlePrice float32
	for _, bundleName := range order.BundlesName {
		vboundl := (*vboundlesList)[bundleName]
		for _, productName := range vboundl.ProductsName {
			bundlePrice += (*vproductsList)[productName].Price * vboundl.Discount
		}
	}

	order.BundlesPrice = bundlePrice
	order.ProductsPrice = productPrice
	order.TotalOrderPrice = bundlePrice + productPrice

	if (vuser.Balance - order.TotalOrderPrice) <= 0 {
		return fmt.Errorf(" %s : остаток %f - списание %f = %f - мало на счету",
			username,
			vuser.Balance,
			order.TotalOrderPrice,
			vuser.Balance-order.TotalOrderPrice)
	}
	vuser.Balance = vuser.Balance - order.TotalOrderPrice
	// запишем в историю списаний
	(*accountsOrders)[username] = append((*accountsOrders)[username], order)
	return nil

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
	default:
	}
	return discount
}
