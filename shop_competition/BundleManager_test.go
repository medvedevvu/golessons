package shop_competition

import (
	"fmt"
	"sync"
	"testing"
)

func InitBundles() BundlesList {
	InitProductCatalog()
	BundlesListMain := BundlesList{}
	err := BundlesListMain.AddBundle("8 марта", "духи", 0.3, "цветы", "шампанское", "шоколад")
	if err != nil {
		fmt.Printf("%s", err)
	}
	err = BundlesListMain.AddBundle("23 февраля", "водка", 0.4, "сыр", "колбаса", "хлеб")
	if err != nil {
		fmt.Printf("%s", err)
	}
	err = BundlesListMain.AddBundle("Новый год", "шампанское", 0.4, "сыр", "колбаса", "шоколад")
	if err != nil {
		fmt.Printf("%s", err)
	}
	return BundlesListMain
}

func TestInitBundles(t *testing.T) {
	v := InitBundles()
	if len(v) == 0 {
		t.Fatal("инициализация не прошла")
	}
}

func TestSimpleRemoveBoundle(t *testing.T) {
	/*
	 пришлось добавить Simple - а то стартовал вместе м нижним тестом
	*/
	lbundleList := InitBundles()
	var wg sync.WaitGroup
	bundles := [2]string{"XXXX", "Новый год"}
	wg.Add(len(bundles))

	_, ok := lbundleList[bundles[0]]
	if ok {
		t.Fatalf(" инит. среды не верный - не должно быть комплекта  %s \n", bundles[0])
	}
	_, ok = lbundleList[bundles[1]]
	if !ok {
		t.Fatalf(" инит. среды не верный - должtн быть комплект %s \n", bundles[1])
	}

	go func() { // удаляем не существующий комплект
		defer wg.Done()
		err := lbundleList.RemoveBundle(bundles[0])
		if err == nil {
			t.Fatalf(" %s \n", err)
		}
		return
	}()

	go func() { // удаляем существующий комплект
		defer wg.Done()
		err := lbundleList.RemoveBundle(bundles[1])
		if err != nil {
			t.Fatalf(" %s \n", err)
		}
		return
	}()

	wg.Wait()

	_, ok = lbundleList[bundles[0]]
	if ok {
		t.Fatalf(" не должно быть комплекта  %s \n", bundles[0])
	}
	_, ok = lbundleList[bundles[1]]
	if ok {
		t.Fatalf(" не должно быть - комплект %s удален \n", bundles[1])
	}

}

func TestRemoveBoundleAndChangeDiscount(t *testing.T) {
	lbundleList := InitBundles()
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		err := lbundleList.RemoveBundle("Новый год")
		if err != nil {
			t.Fatalf(" %s \n", err)
		}
		return
	}()

	go func() {
		defer wg.Done()
		var err error
		err = lbundleList.ChangeDiscount("Новый год", 0.25)
		if err != nil {
			t.Logf(" %s \n", err)
		}
		return
	}()
	wg.Wait()
}

func TestAddAndRemoveBoundle(t *testing.T) {
	lbundleList := InitBundles()
	var wg sync.WaitGroup

	bundles := [3]string{"Новый год", "Новый год", "23 февраля"}

	for idx := 0; idx < len(bundles); idx++ {
		_, ok := lbundleList[bundles[idx]]
		if !ok {
			t.Fatalf(" до теста: комплект %s должен быть в базе \n", bundles[idx])
		}
	}

	wg.Add(len(bundles))

	go func() {
		defer wg.Done()
		err := lbundleList.RemoveBundle(bundles[0])
		if err != nil {
			t.Fatalf(" %s \n", err)
		}
		return
	}()

	go func() {
		defer wg.Done()
		err := lbundleList.AddBundle(bundles[0], "шампанское", 0.4, "сыр", "колбаса", "шоколад")
		if err != nil {
			t.Fatalf(" %s \n", err)
		}
		return
	}()

	go func() {
		defer wg.Done()
		err := lbundleList.RemoveBundle(bundles[2])
		if err != nil {
			t.Fatalf(" %s \n", err)
		}
		return
	}()
	wg.Wait()
	for idx := 0; idx < len(bundles); idx++ {
		_, ok := lbundleList[bundles[idx]]
		if idx == 0 && !ok {
			t.Fatalf("после теста: комплект %s должен быть в базе \n", bundles[idx])
		}
		if idx == 1 && !ok {
			t.Fatalf("после теста: комплект %s должен быть в базе \n", bundles[idx])
		}
		if idx > 1 && ok {
			t.Fatalf("после теста: комплект %s не должен быть в базе \n", bundles[idx])
		}
	}

}

func TestAddWrongBoundle(t *testing.T) {
	lbundleList := InitBundles()

	var wg sync.WaitGroup

	bundles := [3]string{"Новый год", "Мелочи", "Мелочи1"}

	for idx := 0; idx < len(bundles); idx++ {
		_, ok := lbundleList[bundles[idx]]
		if idx == 0 && !ok {
			t.Fatalf(" комплект %s должен быть в базе \n", bundles[idx])
		}
		if idx != 0 && ok {
			t.Fatalf(" комплект %s не должен быть в базе \n", bundles[idx])
		}
	}

	wg.Add(len(bundles))

	go func() {
		defer wg.Done()
		err := lbundleList.AddBundle(bundles[0], "шампанское", 0.4, "сыр", "колбаса", "шоколад")
		if err == nil {
			t.Fatalf("добавили одноименный комплект %s \n", bundles[0])
		}
		return
	}()

	go func() {
		defer wg.Done()
		err := lbundleList.AddBundle(bundles[1], "зубочистка", 0.1, "спички", "вермишель")
		if err == nil {
			t.Fatalf("добавили комплект %s где основа - пробник\n", bundles[1])
		}
		return
	}()

	go func() {
		defer wg.Done()
		err := lbundleList.AddBundle(bundles[2],
			"вермишель", 0.1, "зубочистка", "зубочистка")
		if err == nil {
			t.Fatalf("добавили комплект %s где одни пробники \n", bundles[2])
		}
		return
	}()
	wg.Wait()

	for idx := 0; idx < len(bundles); idx++ {
		_, ok := lbundleList[bundles[idx]]
		if idx == 0 && !ok {
			t.Fatalf(" после теста: комплект %s должен быть в базе \n", bundles[idx])
		}
		if idx != 0 && ok {
			t.Fatalf(" после теста: комплект %s не должен быть в базе \n", bundles[idx])
		}
	}

}
