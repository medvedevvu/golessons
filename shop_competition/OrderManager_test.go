package shop_competition

import (
	"errors"
	"sync"
	"testing"
)

func InitEnviroment() (*AccountsList, *ProductsList, *BundlesList, *AccountsOrders) {
	testAccList := NewAccountsList()
	testAccList.Register("Kola", AccountNormal)
	testAccList.Register("Vasiy", AccountNormal)
	testAccList.Register("Dram", AccountPremium)
	testAccList.Register("Vortis", AccountPremium)

	names := map[string]float32{"Kola": 2750.12,
		"Vasiy": 19930.21, "Dram": 5000, "Vortis": 2136.67}

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

func TestAsyncPlaceOrder(t *testing.T) {
	_, _, _, _ = InitEnviroment()
	vaccountsOrders := GetAccountsOrders()
	vaccountsList := GetAccountsList()

	order := Order{}
	order.ProductsName = []string{"водка", "шампанское", "колбаса"}
	order.BundlesName = []string{"8 марта", "8 марта"}

	var wg sync.WaitGroup

	monyKolabefore := make(chan float32, 1)
	monyKolaafter := make(chan float32, 1)
	monyVasiybefore := make(chan float32, 1)
	monyVasiyafter := make(chan float32, 1)

	//err := errors.New("")
	wg.Add(2)
	go func() {
		defer wg.Done()
		val, err := vaccountsList.Balance("Kola")
		monyKolabefore <- val
		if err != nil {
			t.Fatalf("%s\n ошбка получения баланса - user %s -- %f ", err, "Kola", val)
			return
		}
		if val <= 0 {
			t.Fatalf("баланс %f - user %s \n", val, "Kola")
			return
		}
	}()

	go func() {
		defer wg.Done()
		val, err := vaccountsList.Balance("Vasiy")
		monyVasiybefore <- val
		if err != nil {
			t.Fatalf("%s\n ошбка получения баланса - user %s", err, "Vasiy")
			return
		}
		if val <= 0 {
			t.Fatalf("баланс %f - user %s \n", val, "Vasiy")
			return
		}
	}()

	wg.Wait()

	wg.Add(2)
	go func() {
		defer wg.Done()
		err := vaccountsOrders.PlaceOrder("Vasiy", order)
		monyVasiyafter <- 0
		if err != nil {
			t.Fatalf("%s\n", err)
		}
		return
	}()

	go func() {
		defer wg.Done()
		err := vaccountsOrders.PlaceOrder("Kola", order)
		monyKolaafter <- 0
		if err != nil {
			t.Fatalf("%s\n", err)
		}
		return
	}()

	wg.Wait()

	wg.Add(2)
	go func() {
		defer wg.Done()
		val, err := vaccountsList.Balance("Kola")
		if err != nil {
			t.Fail()
		}
		valbefore := <-monyKolabefore
		<-monyKolaafter
		if valbefore <= val {
			t.Fatalf(" before %f after %f - не прошло списание %s \n",
				valbefore, val, "Kola")
		}
		return
	}()

	go func() {
		defer wg.Done()
		val, err := vaccountsList.Balance("Vasiy")
		if err != nil {
			t.Fail()
		}
		valbefore := <-monyVasiybefore
		<-monyVasiyafter
		if valbefore <= val {
			t.Fatalf(" before %f after %f - не прошло списание %s \n",
				valbefore, val, "Vasiy")
		}
		return
	}()
	wg.Wait()
}

func TestPlaceOrder(t *testing.T) {
	_, _, _, _ = InitEnviroment()
	vaccountsOrders := GetAccountsOrders()
	vaccountsList := GetAccountsList()

	order := Order{}
	order.ProductsName = []string{"водка", "шампанское", "колбаса"}
	order.BundlesName = []string{"8 марта", "8 марта"}

	var (
		monyKolabefore  float32
		monyKolaafter   float32
		monyVasiybefore float32
		monyVasiyafter  float32
	)
	err := errors.New("")
	func() {
		monyKolabefore, err = vaccountsList.Balance("Kola")
		if err != nil {
			t.Fatalf("%s\n ошбка получения баланса - user %s -- %f ", err, "Kola", monyKolabefore)
			return
		}
		if monyKolabefore <= 0 {
			t.Fatalf("баланс %f - user %s \n", monyKolabefore, "Kola")
			return
		}
	}()

	func() {
		monyVasiybefore, err = vaccountsList.Balance("Vasiy")
		if err != nil {
			t.Fatalf("%s\n ошбка получения баланса - user %s", err, "Vasiy")
			return
		}
		if monyVasiybefore <= 0 {
			t.Fatalf("баланс %f - user %s \n", monyVasiybefore, "Vasiy")
			return
		}
	}()

	func() {
		err = vaccountsOrders.PlaceOrder("Vasiy", order)
		if err != nil {
			t.Fatalf("%s\n", err)
		}
		return
	}()

	func() {
		err = vaccountsOrders.PlaceOrder("Kola", order)
		if err != nil {
			t.Fatalf("%s\n", err)
		}
		return
	}()

	func() {
		monyKolaafter, err = vaccountsList.Balance("Kola")
		if err != nil {
			t.Fail()
		}
		if monyKolabefore <= monyKolaafter {
			t.Fatalf(" before %f after %f - не прошло списание %s \n",
				monyKolabefore, monyKolaafter, "Kola")
		}
	}()

	func() {
		monyVasiyafter, err = vaccountsList.Balance("Vasiy")
		if err != nil {
			t.Fail()
		}
		if monyVasiybefore <= monyVasiyafter {
			t.Fatalf(" before %f after %f - не прошло списание %s \n",
				monyVasiybefore, monyVasiyafter, "Vasiy")
		}
		return
	}()

}
