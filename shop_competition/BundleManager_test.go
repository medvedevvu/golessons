package shop_competition

import (
	"fmt"
	"testing"
)

func InitBundles() *BundlesList {
	InitProductCatalog()
	vbundleList := NewBundlesList()
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

func TestAddBoundle(t *testing.T) {
	lbundleList := InitBundles()

	err := lbundleList.AddBundle("Новый год", "шампанское", 0.4, "сыр", "колбаса", "шоколад")
	if err == nil {
		t.Fatalf("добавили одноименный комплект")
	}
	err = lbundleList.AddBundle("Мелочи", "зубочистка", 0.1, "спички", "вермишель")
	if err == nil {
		t.Fatalf("добавили комплект где основа - пробник")
	}

	err = lbundleList.AddBundle("Мелочи", "вермишель", 0.1, "зубочистка", "зубочистка")
	if err == nil {
		t.Fatalf("добавили комплект где одни пробники %v", *lbundleList)
	}

}
