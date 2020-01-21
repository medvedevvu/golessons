package structsdef2

import "fmt"

// ItemPrice - описание товара в дальнейшем map[ItemName] *ItemPrice
type ItemPrice struct {
	//ItemName  string  // название товара
	ItemPrice float32 // цена
	ItemType  int     //  0 - обычный товар 1 - премиум товар 2 - пробник с нулевой ценой
}

// Order - описание заказа
type Order struct {
	Items     []string // список товаров
	TotalSum  float32  // сумма заказа
	OrderType int
	//  0 - товар 1 - набор  2 - (набор + пробник или товар + пробник )
}

// User - описание пользователя - в дальнейшем отображение map[userName]*User
type User struct {
	Email    string  // электронная почта пользователя - полагаю что уникальная
	Account  float32 // остаток на счету
	UserType int     // 0 - обычный пользователь 1 - премиум ползователь
}

// SetDiscount - устанавливаем скидку на заказ в зависимости от пользователя
// типов товаров и заказов
func (order *Order) SetDiscount(userName string,
	acountList map[string]*User,
	itemsPrice map[string]*ItemPrice,
	addDisc float32, // скидывать предполагаю как множитель - если не задан или 0 - то 1
) (float32, float32, float32) {
	// если пользователь обычный
	// если пользователь премиум
	if addDisc <= 0 { // заглушка на значение по умолчанию.
		addDisc = 1
	}

	vUser, ok := acountList[userName]

	if !ok {
		return 0, 0, 0
	}

	var grandTotal float32
	switch vUser.UserType {
	case 0:
		{
			for _, itemName := range order.Items {
				item, ok := itemsPrice[itemName]
				if ok {
					switch item.ItemType {
					case 2:
						{
							// пробники пропускаем
							continue
						}
					case 1:
						{
							grandTotal += item.ItemPrice * 1.5
						}
					default:
						{
							grandTotal += item.ItemPrice
						}
					}
					//TotalSumNoDiscount += item.ItemPrice
				} else {
					fmt.Printf("Товара %s нет в каталоге \n", itemName)
				}
			}

		}
	case 1:
		{

			for _, itemName := range order.Items {
				item, ok := itemsPrice[itemName]
				if ok {
					switch item.ItemType {
					case 2:
						{
							// пробники пропускаем
							continue
						}
					case 1:
						{
							grandTotal += item.ItemPrice * 0.8
						}
					default:
						{
							grandTotal += item.ItemPrice * 0.95
						}
					}
					//TotalSumNoDiscount += item.ItemPrice
				} else {
					fmt.Printf("Товара %s нет в каталоге \n", itemName)
				}
			}
		}
	}
	// посмотрим - набор это или нет - если набор пытаемся сделать дополнительную скидку
	TotalSumNoDiscount := order.TotalSum       // без скидки
	TotalSumNoDiscountByComplect := grandTotal // со скидкой - комплект не учтен
	order.TotalSum = grandTotal
	if order.OrderType == 1 {
		order.TotalSum = order.TotalSum * addDisc // со скидкой + комплект
	}
	// будем возвращать сумму без скидки , без скидки по комлекту , выставленную к оплате
	return TotalSumNoDiscount, TotalSumNoDiscountByComplect, order.TotalSum
}
