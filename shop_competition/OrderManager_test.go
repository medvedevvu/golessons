package shop_competition

import (
	"sync"
	"testing"
)

func InitEnviroment(envList *ShopBase) {
	envList.AccountsListWithMutex.Register("Kola", AccountNormal)
	envList.AccountsListWithMutex.Register("Vasiy", AccountNormal)
	envList.AccountsListWithMutex.Register("Dram", AccountPremium)
	envList.AccountsListWithMutex.Register("Vortis", AccountPremium)

	names := map[string]float32{"Kola": 2750.12,
		"Vasiy": 19930.21, "Dram": 5000, "Vortis": 2136.67}

	for key, vals := range names {
		_ = envList.AccountsListWithMutex.AddBalance(key, vals)
	}

	_ = envList.ProductListWithMutex.AddProduct("колбаса", Product{Price: 125.23, Type: ProductNormal})
	_ = envList.ProductListWithMutex.AddProduct("водка", Product{Price: 400.23, Type: ProductNormal})
	_ = envList.ProductListWithMutex.AddProduct("сыр", Product{Price: 315.14, Type: ProductPremium})
	_ = envList.ProductListWithMutex.AddProduct("макароны", Product{Price: 47.14, Type: ProductNormal})
	_ = envList.ProductListWithMutex.AddProduct("зубочистка", Product{Price: 0.00, Type: ProductSample})
	_ = envList.ProductListWithMutex.AddProduct("вермишель", Product{Price: 11.20, Type: ProductNormal})
	_ = envList.ProductListWithMutex.AddProduct("хлеб", Product{Price: 32.10, Type: ProductNormal})
	_ = envList.ProductListWithMutex.AddProduct("цветы", Product{Price: 30.10, Type: ProductPremium})
	_ = envList.ProductListWithMutex.AddProduct("шампанское", Product{Price: 150.10, Type: ProductNormal})
	_ = envList.ProductListWithMutex.AddProduct("шоколад", Product{Price: 478.21, Type: ProductPremium})
	_ = envList.ProductListWithMutex.AddProduct("духи", Product{Price: 470.51, Type: ProductPremium})
	_ = envList.ProductListWithMutex.AddProduct("спички", Product{Price: 22.51, Type: ProductNormal})

	_ = envList.BundlesListWithMutex.AddBundle("8 марта", "духи", 0.3, envList, "цветы", "шампанское", "шоколад")
	_ = envList.BundlesListWithMutex.AddBundle("23 февраля", "водка", 0.4, envList, "сыр", "колбаса", "хлеб")
	_ = envList.BundlesListWithMutex.AddBundle("Новый год", "шампанское", 0.4, envList, "сыр", "колбаса", "шоколад")

}

func TestAsyncPlaceOrder(t *testing.T) {
	envList := NewShopBase()
	InitEnviroment(envList)

	order := Order{}
	order.ProductsName = []string{"водка", "шампанское", "колбаса"}
	order.BundlesName = []string{"8 марта", "8 марта"}

	var wg sync.WaitGroup

	monyKolabefore := make(chan float32, 1)
	monyKolaafter := make(chan float32, 1)
	monyVasiybefore := make(chan float32, 1)
	monyVasiyafter := make(chan float32, 1)

	names := [2]string{"Kola", "Vasiy"}

	wg.Add(len(names))
	go func() {
		defer wg.Done()
		val, err := envList.AccountsListWithMutex.Balance(names[0])
		monyKolabefore <- val
		if err != nil {
			t.Fatalf("%s\n ошбка получения баланса - user %s -- %f ", err, names[0], val)
			return
		}
		if val <= 0 {
			t.Fatalf("баланс %f - user %s \n", val, names[0])
			return
		}
	}()

	go func() {
		defer wg.Done()
		val, err := envList.AccountsListWithMutex.Balance(names[1])
		monyVasiybefore <- val
		if err != nil {
			t.Fatalf("%s\n ошбка получения баланса - user %s", err, names[1])
			return
		}
		if val <= 0 {
			t.Fatalf("баланс %f - user %s \n", val, names[1])
			return
		}
	}()

	wg.Wait()

	wg.Add(len(names))
	go func() {
		defer wg.Done()
		err := envList.AccountsOrdersWithMutex.PlaceOrder(names[1], order, envList)

		monyVasiyafter <- 0
		if err != nil {
			t.Fatalf("%s\n", err)
		}
		return
	}()

	go func() {
		defer wg.Done()
		err := envList.AccountsOrdersWithMutex.PlaceOrder(names[0], order, envList)
		monyKolaafter <- 0
		if err != nil {
			t.Fatalf("%s\n", err)
		}
		return
	}()

	wg.Wait()

	wg.Add(len(names))
	go func() {
		defer wg.Done()
		val, err := envList.AccountsListWithMutex.Balance(names[0])
		if err != nil {
			t.Fail()
		}
		valbefore := <-monyKolabefore
		<-monyKolaafter
		if valbefore <= val {
			t.Fatalf(" before %f after %f - не прошло списание %s \n",
				valbefore, val, names[0])
		}
		return
	}()

	go func() {
		defer wg.Done()
		val, err := envList.AccountsListWithMutex.Balance(names[1])
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

func TestCheckPlaceOrder(t *testing.T) {
	envList := NewShopBase()
	InitEnviroment(envList)

	names := [2]string{"Kola", "Vasiy"}
	order := Order{}
	order.ProductsName = []string{"водка", "шампанское", "колбаса"}
	order.BundlesName = []string{"8 марта", "8 марта"}

	var wg sync.WaitGroup

	var (
		monyKolabefore  float32
		monyVasiybefore float32
	)

	monyKolabefore, err := envList.AccountsListWithMutex.Balance(names[0])
	if err != nil {
		t.Fatalf("%s\n ошбка получения баланса - user %s -- %f ", err, names[0], monyKolabefore)
	}

	monyVasiybefore, err = envList.AccountsListWithMutex.Balance(names[1])
	if err != nil {
		t.Fatalf("%s\n ошбка получения баланса - user %s", err, names[1])
	}

	wg.Add(len(names))
	go func() {
		defer wg.Done()
		err := envList.AccountsOrdersWithMutex.PlaceOrder(names[1], order, envList)
		if err != nil {
			t.Fatalf("%s\n", err)
		}
		return
	}()

	go func() {
		defer wg.Done()
		err := envList.AccountsOrdersWithMutex.PlaceOrder(names[0], order, envList)
		if err != nil {
			t.Fatalf("%s\n", err)
		}
		return
	}()

	wg.Wait()
	func() {
		//	defer wg.Done()
		monyKolaafter, err := envList.AccountsListWithMutex.Balance(names[0])
		if err != nil {
			t.Fail()
		}
		if monyKolabefore <= monyKolaafter {
			t.Fatalf(" before %f after %f - не прошло списание %s \n",
				monyKolabefore, monyKolaafter, names[0])
		}
		return
	}()

	func() {
		//	defer wg.Done()
		monyVasiyafter, err := envList.AccountsListWithMutex.Balance(names[1])
		if err != nil {
			t.Fail()
		}
		if monyVasiybefore <= monyVasiyafter {
			t.Fatalf(" before %f after %f - не прошло списание %s \n",
				monyVasiybefore, monyVasiyafter, names[1])
			return
		}
	}()

}

func TestPlaceOrderAndAddBalanc(t *testing.T) {
	envList := NewShopBase()
	InitEnviroment(envList)

	order := Order{}
	order.ProductsName = []string{"водка", "шампанское", "колбаса"}
	order.BundlesName = []string{"8 марта", "8 марта"}

	var (
		monyKolabefore  float32
		monyKolaafter   float32
		monyVasiybefore float32
		monyVasiyafter  float32
	)
	names := [2]string{"Kola", "Vasiy"}

	var err error
	func() {
		monyKolabefore, err = envList.AccountsListWithMutex.Balance(names[0])
		if err != nil {
			t.Fatalf("%s\n ошбка получения баланса - user %s -- %f ", err, names[0], monyKolabefore)
			return
		}
		if monyKolabefore <= 0 {
			t.Fatalf("баланс %f - user %s \n", monyKolabefore, names[0])
			return
		}
	}()

	func() {
		monyVasiybefore, err = envList.AccountsListWithMutex.Balance(names[1])
		if err != nil {
			t.Fatalf("%s\n ошбка получения баланса - user %s", err, names[1])
			return
		}
		if monyVasiybefore <= 0 {
			t.Fatalf("баланс %f - user %s \n", monyVasiybefore, names[1])
			return
		}
	}()

	func() {
		err = envList.AccountsOrdersWithMutex.PlaceOrder(names[1], order, envList)
		if err != nil {
			t.Fatalf("%s\n", err)
		}
		return
	}()

	func() {
		err = envList.AccountsOrdersWithMutex.PlaceOrder(names[0], order, envList)
		if err != nil {
			t.Fatalf("%s\n", err)
		}
		return
	}()

	func() {
		monyKolaafter, err = envList.AccountsListWithMutex.Balance(names[0])
		if err != nil {
			t.Fail()
		}
		if monyKolabefore <= monyKolaafter {
			t.Fatalf(" before %f after %f - не прошло списание %s \n",
				monyKolabefore, monyKolaafter, names[0])
		}
	}()

	func() {
		monyVasiyafter, err = envList.AccountsListWithMutex.Balance(names[1])
		if err != nil {
			t.Fail()
		}
		if monyVasiybefore <= monyVasiyafter {
			t.Fatalf(" before %f after %f - не прошло списание %s \n",
				monyVasiybefore, monyVasiyafter, names[1])
		}
		return
	}()

}
