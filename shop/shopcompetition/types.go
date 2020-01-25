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

	StInProd          ProductCatalogStateErrorType = iota // точное совпадение товара в каталоге
	StNameAndTypeSame                                     // имя и тип "типа" совпали - обновим значения типа и цены
	StTypeNotSame                                         // тип "типа" не совпал
	StDef                                                 // по умолчанию
	StNotFnd                                              // товара нет в каталоге
	StWrongPrice                                          // не верная цена товара
	StNil                                                 // не определное значение
)

//ProductType one of ProductNormal/ProductPremium/ProductSample
type ProductType uint8
type BundleType uint8
type AccountType uint8
type AccountSortType uint8

// тип ошибки при работе с каталогом продуктов
type ProductCatalogStateErrorType uint8

// ошибки при работе с каталогом продуктов
type PCSError struct {
	ErrMsg  string
	ErrCode ProductCatalogStateErrorType
}

func (er PCSError) Error() string {
	return er.ErrMsg
}

func (er PCSError) ErrorCode() ProductCatalogStateErrorType {
	return er.ErrCode
}

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
