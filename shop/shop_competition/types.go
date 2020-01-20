package shopcompetition

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
	Name  string
	Price float32
	Type  ProductType
}

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
	Name    string
	Balance float32
	AccountType
}
