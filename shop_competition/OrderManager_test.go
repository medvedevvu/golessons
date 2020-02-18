package shop_competition

import "testing"

func InitEnviroment() (*AccountsList, *ProductsList, *BundlesList, *AccountsOrders) {
	testAccList := NewAccountsList()
	testAccList.Register("Kola", AccountNormal)
	testAccList.Register("Vasiy", AccountNormal)
	testAccList.Register("Dram", AccountPremium)
	testAccList.Register("Vortis", AccountPremium)

	names := map[string]float32{"Kola": 2750.12,
		"Vasiy": 1930.21, "Dram": 5000, "Vortis": 2136.67}

	for key, vals := range names {
		_ = testAccList.AddBalance(key, vals)
	}

	lproductList := NewProductsList()
	_ = lproductList.AddProduct("колбаса", Product{Price: 125.23, Type: ProductNormal})
	_ = lproductList.AddProduct("водка", Product{Price: 400.23, Type: ProductNormal})
	_ = lproductList.AddProduct("сыр", Product{Price: 315.14, Type: ProductPremium})
	_ = lproductList.AddProduct("макароны", Product{Price: 47.14, Type: ProductNormal})
	_ = lproductList.AddProduct("зубочистка", Product{Price: 0.00, Type: ProductSample})
	_ = lproductList.AddProduct("вермишель", Product{Price: 11.20, Type: ProductNormal})
	_ = lproductList.AddProduct("хлеб", Product{Price: 32.10, Type: ProductNormal})
	_ = lproductList.AddProduct("цветы", Product{Price: 30.10, Type: ProductPremium})
	_ = lproductList.AddProduct("шампанское", Product{Price: 150.10, Type: ProductNormal})
	_ = lproductList.AddProduct("шоколад", Product{Price: 478.21, Type: ProductPremium})
	_ = lproductList.AddProduct("духи", Product{Price: 470.51, Type: ProductPremium})
	_ = lproductList.AddProduct("спички", Product{Price: 22.51, Type: ProductNormal})

	vbundleList := NewBundlesList()
	_ = vbundleList.AddBundle("8 марта", "духи", 0.3, "цветы", "шампанское", "шоколад")
	_ = vbundleList.AddBundle("23 февраля", "водка", 0.4, "сыр", "колбаса", "хлеб")
	_ = vbundleList.AddBundle("Новый год", "шампанское", 0.4, "сыр", "колбаса", "шоколад")

	vaccountsOrders := NewAccountsOrders()

	return testAccList, lproductList, vbundleList, vaccountsOrders
}

func TestPlaceOrder(t *testing.T) {
	_, _, _, _ = InitEnviroment()
	vaccountsOrders := GetAccountsOrders()
	vaccountsList := GetAccountsList()

	order := Order{}
	order.ProductsName = []string{"водка", "шампанское", "колбаса"}
	order.BundlesName = []string{"8 марта", "8 марта"}

	money, err := vaccountsList.Balance("Kola")
	if err != nil {
		t.Fatalf("%s\n ошбка получения баланса ", err)
	}
	if money <= 0 {
		t.Fatalf("баланс %f  \n", money)
	}

	err = vaccountsOrders.PlaceOrder("Kola", order)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	money1, err := vaccountsList.Balance("Kola")
	if money <= money1 {
		t.Fatalf(" before %f after %f - не прошло списание ", money, money1)
	}

	t.Logf("---> %v ", vaccountsOrders)
}
