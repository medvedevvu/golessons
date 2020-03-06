package shop_competition

import (
	"fmt"
	"sync"
	"testing"
)

func InitBundles() *BundlesList {
	InitProductCatalog()
	vbundleList := &BundlesList{}
	err := vbundleList.AddBundle("8 марта", "духи", 0.3, "цветы", "шампанское", "шоколад")
	if err != nil {
		fmt.Printf("%s", err)
	}
	err = vbundleList.AddBundle("23 февраля", "водка", 0.4, "сыр", "колбаса", "хлеб")
	if err != nil {
		fmt.Printf("%s", err)
	}
	err = vbundleList.AddBundle("Новый год", "шампанское", 0.4, "сыр", "колбаса", "шоколад")
	if err != nil {
		fmt.Printf("%s", err)
	}
	return vbundleList
}

func TestInitBundles(t *testing.T) {
	v := InitBundles()
	if len(*v) == 0 {
		t.Fatal("инициализация не прошла")
	}
}

func TestRemoveBoundle(t *testing.T) {
	lbundleList := InitBundles()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		err := lbundleList.RemoveBundle("XXXX")
		if err == nil {
			t.Fatalf(" %s \n", err)
		}
		return
	}()

	go func() {
		defer wg.Done()
		err := lbundleList.RemoveBundle("Новый год")
		if err != nil {
			t.Fatalf(" %s \n", err)
		}
		return
	}()

	wg.Wait()
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
		err := lbundleList.ChangeDiscount("Новый год", 0.25)
		if err != nil {
			t.Fatalf(" %s \n", err)
		}
		return
	}()
	wg.Wait()
}

func TestAddAndRemoveBoundle(t *testing.T) {
	lbundleList := InitBundles()
	var wg sync.WaitGroup
	wg.Add(3)

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
		err := lbundleList.AddBundle("Новый год", "шампанское", 0.4, "сыр", "колбаса", "шоколад")
		if err != nil {
			t.Fatalf(" %s \n", err)
		}
		return
	}()

	go func() {
		defer wg.Done()
		err := lbundleList.RemoveBundle("23 февраля")
		if err != nil {
			t.Fatalf(" %s \n", err)
		}
		return
	}()
	wg.Wait()
}

func TestAddBoundle(t *testing.T) {
	lbundleList := InitBundles()

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		err := lbundleList.AddBundle("Новый год", "шампанское", 0.4, "сыр", "колбаса", "шоколад")
		if err == nil {
			t.Fatalf("добавили одноименный комплект")
		}
		return
	}()

	go func() {
		defer wg.Done()
		err := lbundleList.AddBundle("Мелочи", "зубочистка", 0.1, "спички", "вермишель")
		if err == nil {
			t.Fatalf("добавили комплект где основа - пробник")
		}
		return
	}()

	go func() {
		defer wg.Done()
		err := lbundleList.AddBundle("Мелочи", "вермишель", 0.1, "зубочистка", "зубочистка")
		if err == nil {
			t.Fatalf("добавили комплект где одни пробники %v", *lbundleList)
		}
		return
	}()

	wg.Wait()
}
