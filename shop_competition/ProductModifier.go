package shop_competition

/*
type ProductModifier interface {
	AddProduct(Product) error
	ModifyProduct(Product) error
	RemoveProduct(name string) error
}
type ProductsList map[string]*Product
*/
// NewProductsList конструктор
func NewProductsList() *ProductsList {
	return &ProductsList{}
}
