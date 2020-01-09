package structsdef1

// ItemPrice - описание товара
type ItemPrice struct {
	ItemName  string  // название товара
	ItemPrice float32 // цена
	ItemType  int     //  0 - обычный товар 1 - премиум товар 2 - пробник с нулевой ценой
}

// Order - описание заказа
type Order struct {
	Items     []string // список товаров
	TotalSum  float32  // сумма заказа
	OrderType int      //  0 - обычный заказ 1 - комплект  2 - комплект + пробник
}

// SetDiscount - устанавливаем скидку на заказ в зависимотси от пользователя
// типов товаров и заказов
func (order *Order) SetDiscount(user User, itemsPrice map[int]*ItemPrice,
	addDisc float32) (float32, float32, float32) {
	var TotalSumNoDiscount float32
	// если пользователь обычный
	// если пользователь премиум
	switch user.UserType {
	case 0:
		{
			for i := 0; i < len(order.Items); i++ {
				for x := 0; x < len(itemsPrice); x++ {
					if order.Items[i] == itemsPrice[x].ItemName {
						if itemsPrice[x].ItemType == 2 {
							// пробники пропускаем
							continue
						}
						if itemsPrice[x].ItemType == 1 {
							order.TotalSum += itemsPrice[x].ItemPrice * 1.5
						} else {
							order.TotalSum += itemsPrice[x].ItemPrice
						}
						TotalSumNoDiscount += itemsPrice[x].ItemPrice
						break // все нашли переходим к другому товару из списка
					}
				}
			}
		}
	case 1:
		{
			for i := 0; i < len(order.Items); i++ {
				for x := 0; x < len(itemsPrice); x++ {
					if order.Items[i] == itemsPrice[x].ItemName {
						if itemsPrice[x].ItemType == 2 {
							// пробники пропускаем
							continue
						}
						if itemsPrice[x].ItemType == 1 {
							order.TotalSum += itemsPrice[x].ItemPrice * 0.8
						} else {
							order.TotalSum += itemsPrice[x].ItemPrice * 0.95
						}
						TotalSumNoDiscount += itemsPrice[x].ItemPrice
						break // все нашли переходим к другому товару из списка
					}
				}
			}
		}
	}
	// посмотрим - набор это или нет - если набор пытаемся сделать дополнительную скидку
	TotalSumNoDiscountByCoplect := order.TotalSum
	if order.OrderType == 1 {
		order.TotalSum = order.TotalSum * addDisc
	}
	// будем возвращать сумму без скидки , без скидки по комлекту , выставленную к оплате
	return TotalSumNoDiscount, TotalSumNoDiscountByCoplect, order.TotalSum
}

// User - описание пользователя
type User struct {
	UserName string  // Имя пользователя
	Account  float32 // остаток на счету
	UserType int     // 0 - обычный пользователь 1 - премиум ползователь
}
