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
)

//ProductType one of ProductNormal/ProductPremium/ProductSample
type ProductType uint8
type BundleType uint8
type AccountType uint8
type AccountSortType uint8

type Product struct {
	//	Name  string
	Price float32
	Type  ProductType
}

// ProductsList  все товары map[Name]*Product
type ProductsList map[string]*Product

type Order struct {
	Products []Product
	Bundles  []Bundle
}

type Bundle struct {
	Products []Product
	Type     BundleType
	Discount float32
}

type Account struct {
	//	Name    string    Вынес в ключевое значение для AccountsList
	Balance float32
	AccountType
}

// AccountsList Список всех пользоватиелей map[Name]*Account
type AccountsList map[string]*Account
