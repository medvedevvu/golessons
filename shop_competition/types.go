package shop_competition

import "sync"

const (
	ProductNormal ProductType = iota
	ProductPremium
	ProductSample

	BundleNormal BundleType = iota
	BundleSample

	AccountNormal AccountType = iota
	AccountPremium

	SortByName AccountSortType = iota
	SortByNameReverse
	SortByBalance

	Add OperationType = iota
	Edit
)

//ShopBase
type ShopBase struct {
	ProductListWithMutex    ProductsList
	AccountsListWithMutex   AccountsList
	AccountsOrdersWithMutex AccountsOrders
	BundlesListWithMutex    BundlesList
	sync.RWMutex
}

//ShopBase constructor
func NewShopBase() *ShopBase {
	return &ShopBase{
		AccountsListWithMutex: AccountsList{
			Accounts: make(map[string]*Account),
		},
		ProductListWithMutex: ProductsList{
			Products: make(map[string]*Product),
		},
		AccountsOrdersWithMutex: AccountsOrders{
			AccountOrders: make(map[string][]Order),
		},
		BundlesListWithMutex: BundlesList{
			BundleList: make(map[string]Bundle),
		},
	}
}

//ProductType one of ProductNormal/ProductPremium/ProductSample
type ProductType uint8
type BundleType uint8
type AccountType uint8
type AccountSortType uint8
type OperationType uint8

// ProductsList  все товары map[Name]*Product
type ProductsList struct {
	Products map[string]*Product
	sync.RWMutex
}

type Product struct {
	//	Name  string
	Price float32
	Type  ProductType
}

//AccountsOrders Сввяь с Account по имени пользователя
type AccountsOrders struct {
	AccountOrders map[string][]Order
	sync.RWMutex
}

type Order struct {
	ProductsName    []string //[]Product  список продуктов
	BundlesName     []string // список комплектов
	ProductsPrice   float32  // стоимость товаров    в заказе
	BundlesPrice    float32  // стоимость комплектов в заказе
	TotalOrderPrice float32  // общая стоимость заказа
}

// BundlesList каталог именнованных комплектов
type BundlesList struct {
	BundleList map[string]Bundle
	sync.RWMutex
}

type Bundle struct {
	ProductsName []string //[]Product
	Type         BundleType
	Discount     float32
}

// AccountsList Список всех пользоватиелей map[Name]*Account
type AccountsList struct {
	Accounts map[string]*Account
	sync.RWMutex
}

type Account struct {
	//	Name    string    Вынес в ключевое значение для AccountsList
	Balance float32
	AccountType
}
