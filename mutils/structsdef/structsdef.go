package structsdef

// ItemPrice - описание товара
type ItemPrice struct {
	ItemName  string  // название товара
	ItemPrice float32 // цена
}

// Order - описание заказа
type Order struct {
	Items    []string // список товаров
	TotalSum float32  // сумма заказа
}

// User - описание пользователя
type User struct {
	UserName string  // Имя пользователя
	Account  float32 // остаток на счету
}
