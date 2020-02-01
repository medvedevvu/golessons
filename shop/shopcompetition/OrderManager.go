package shopcompetition

import (
	"fmt"
	"math/big"
)

// GetDSC - определим скидку на продукт
func GetDSC(accType AccountType, prdType ProductType) (disc float32) {
	disc = 1.0
	switch accType {
	case AccountPremium:
		switch prdType {
		case ProductPremium:
			disc = 0.8
		case ProductNormal:
			disc = 1.5
		}
	case AccountNormal:
		switch prdType {
		case ProductPremium:
			disc = 0.95
		case ProductNormal:
			disc = 1
		}
	default:
	}
	return disc
}

// CalculateOrder - Обсчитываем заказ
func (pl ProductList) CalculateOrder(account Account,
	order Order,
	addDiscForBundle float32) (float32, error) {
	// сумма за продукты
	var sumForProducts float32
	if len(order.Products) == 0 && len(order.Bundles) == 0 {
		fmt.Printf("Пустой заказ \n")
		return 0, fmt.Errorf("%d", StEmptyOrder)
	}
	for _, prd := range order.Products {
		fnd := true
		for _, p := range pl {
			if prd.Name == p.Name {
				if prd.Type == ProductSample {
					fmt.Printf("товар %s - пробник можно купить только в наборе  \n", prd.Name)
					fnd = true
					break
				}
				a := big.NewFloat(float64(prd.Price))
				a.Mul(a, big.NewFloat(float64(GetDSC(account.AccountType, prd.Type))))
				s := big.NewFloat(float64(sumForProducts))
				s.Add(s, a)
				//sumForProducts += prd.Price
				sumForProducts, _ = s.Float32()
				fnd = true
				break
			} else {
				fnd = false
			}
		}
		if !fnd {
			fmt.Printf("товар %s не найден -- товары \n", prd.Name)
		}
	}
	var sumForBundles float32
	for _, bndl := range order.Bundles {
		for _, prd := range bndl.Products {
			fnd := true
			for _, p := range pl {
				if prd.Name == p.Name {
					s := big.NewFloat(float64(sumForBundles))
					a := big.NewFloat(float64(prd.Price))
					b := big.NewFloat(1)
					if addDiscForBundle != 0 { // если скидка стоит - это скидка на комплект
						b = big.NewFloat(float64(addDiscForBundle))
					} else {
						b = big.NewFloat(float64(bndl.Discount))
					}
					a.Mul(a, b)
					a.Add(a, s)
					//sumForBundles += (prd.Price * bndl.Discount)
					sumForBundles, _ = a.Float32()
					fnd = true
					break
				} else {
					fnd = false
				}
			}
			if !fnd {
				fmt.Printf("товар %s не найден -- комплекты \n", prd.Name)
			}
		}
	}
	return sumForProducts + sumForBundles, nil
}

// PlaceOrder списание денег за заказ
func (pl ProductList) PlaceOrder(account *Account,
	order Order,
	addDiscForBundle float32) error {
	bill, err := pl.CalculateOrder(*account, order, addDiscForBundle)
	if err != nil {
		return err
	}
	obil := big.NewFloat(float64(bill))
	obal := big.NewFloat(float64(account.Balance))

	switch obal.Cmp(obil) {
	case 0, -1:
		return fmt.Errorf("%d", StNotMany)
	case 1:
		obal.Add(obal, obil.Mul(obil, big.NewFloat(-1.0)))
		account.Balance, _ = obal.Float32()
	}
	return nil
}
