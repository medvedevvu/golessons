package shopcompetition

//Shop - сборный интерфейс магазина. Объект реализующий этот интерфейс будет тестироваться.
//Если реализованы не все методы, оставить заглушки.

/*type Shop interface {
	ProductModifier
	AccountManager
	OrderManager
	Importer
	Exporter
}*/

type Shop struct {
	AccountList
	ProductList
}

//ProductModifier - интерфейс для работы со списком продуктов магазина
type ProductModifier interface {
	AddProduct(Product) PCSError        // error
	ModifyProduct(Product) PCSError     // error
	RemoveProduct(name string) PCSError // error
}

//AccountManager - интерфейс для работы с пользователями.
type AccountManager interface {
	Register(username string) error
	AddBalance(username string, sum float32) error
	Balance(username string) (float32, error)
	GetAccounts(sort AccountSortType) []Account
}

//OrderManager - интерфейс для работы заказами. Рассчитать заказ и купить.
type OrderManager interface {
	CalculateOrder(account Account, order Order, addDisc float32) (float32, error) // addDisc - дополнительная скидка
	//PlaceOrder(username string, order Order) error //
	PlaceOrder(account Account, order Order) error //
}

//BundleManager - интерфейс для работы с наборами.
type BundleManager interface {
	AddBundle(name string, main Product, discount float32, additional ...Product) error
	ChangeDiscount(name string, discount float32) error
	RemoveBundle(name string) error
}

//Exporter - интерфейс для получения полного состояния магазина.
type Exporter interface {
	Export() ([]byte, error)
}

//Importer - интерфейс для загрузки состояния магазина. Принимает формат который возвращает Exporter.
type Importer interface {
	Import(data []byte) error
}
