//  задание 1
package main

type ItemPrice struct {
	ItemId    int     // УИД товара
	ItemName  string  // название товара
	ItemPrice float32 // цена
}

type Order struct { // УИД заказа
	OrderId  int
	Items    map[int]ItemPrice // элементы [кол-во] товар
	TotalSum float32           // сумма заказа
}

type User struct {
	UserId   int     // УИД пользователя
	UserName string  // Имя пользователя
	Account  float32 // остаток на счету
}

type Bill struct {
	//   UserId OrderId OrderSumm
}

type Orders struct {
}

func main() {

}
