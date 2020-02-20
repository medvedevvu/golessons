package shop_competition

//Shop - сборный интерфейс магазина. Объект реализующий этот интерфейс будет тестироваться.
//Если реализованы не все методы, оставить заглушки.
type Shop interface {
	ProductModifier
	AccountManager
	OrderManager
	Importer
	Exporter
}

//ProductModifier - интерфейс для работы со списком продуктов магазина
type ProductModifier interface {
	AddProduct(productName string, product Product) error
	ModifyProduct(productName string, product Product) error
	RemoveProduct(productName string) error
	// проверка атрибутов товара
	CheckAttrsOfProduct(productName string, product Product, operation OperationType) error
	GetProductList() *ProductsList
}

//AccountManager - интерфейс для работы с пользователями.
type AccountManager interface {
	Register(username string, accounttype AccountType) error
	AddBalance(username string, sum float32) error
	Balance(username string) (float32, error)
	GetAccounts(sort AccountSortType) AccountsList //[]Account
}

//OrderManager - интерфейс для работы заказами. Рассчитать заказ и купить.
type OrderManager interface {
	PlaceOrder(username string, order Order) error
	//CalculateOrder(order Order) (float32, error)
}

//BundleManager - интерфейс для работы с наборами.
type BundleManager interface {
	AddBundle(name string,
		main string, // название основного продукта
		discount float32, additional ...string, // все остальные товары комлекта
	) error
	ChangeDiscount(name string, discount float32) error
	RemoveBundle(name string) error
	GetBundlesList() *BundlesList
}

//Exporter - интерфейс для получения полного состояния магазина.
type Exporter interface {
	Export() ([]byte, error)
}

//Importer - интерфейс для загрузки состояния магазина. Принимает формат который возвращает Exporter.
type Importer interface {
	Import(data []byte) error
}

// ShopInfo - краткая информация о магазине
func ShopInfo() string {
	return "Shop-commercial"
}