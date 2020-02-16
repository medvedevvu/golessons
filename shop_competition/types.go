package shop_competition

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

//ProductType one of ProductNormal/ProductPremium/ProductSample
type ProductType uint8
type BundleType uint8
type AccountType uint8
type AccountSortType uint8
type OperationType uint8

// ProductsList  все товары map[Name]*Product
type ProductsList map[string]*Product

type Product struct {
	//	Name  string
	Price float32
	Type  ProductType
}

//AccountsOrders Сввяь с Account по имени пользователя
type AccountsOrders map[string][]Order

type Order struct {
	ProductsName    []string //[]Product  список продуктов
	BundlesName     []string // список комплектов
	ProductsPrice   float32  // стоимость товаров    в заказе
	BundlesPrice    float32  // стоимость комплектов в заказе
	TotalOrderPrice float32  // общая стоимость заказа
}

// BundlesList каталог именнованных комплектов
type BundlesList map[string]Bundle

type Bundle struct {
	ProductsName []string //[]Product
	Type         BundleType
	Discount     float32
}

// AccountsList Список всех пользоватиелей map[Name]*Account
type AccountsList map[string]*Account

type Account struct {
	//	Name    string    Вынес в ключевое значение для AccountsList
	Balance float32
	AccountType
}
