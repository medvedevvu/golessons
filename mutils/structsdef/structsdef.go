package structsdef

// ItemPrice - описание товара
type ItemPrice struct {
	ItemName  string  // название товара
	ItemPrice float32 // цена
}

// Order - описание заказа
type Order struct {
	Items    map[int]ItemPrice // элементы [кол-во] ---> товар
	TotalSum float32           // сумма заказа
}

// User - описание пользователя
type User struct {
	UserName string  // Имя пользователя
	Account  float32 // остаток на счету
}

// Bill - описание счетов пользователя
type Bill struct {
	UserID    float32         // пользователь
	DebitHist map[int]float32 // история списаний id заказа ---> сумма покупки
}
