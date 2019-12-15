// задание 5
package main

import "fmt"

/* Добавление нового товра в справочник - если есть измениться цена  */
func addItemsPrice(itList map[string]float32,
	itemName string, itemPrice float32) string {
	_, ok := itList[itemName]
	msg := itemName
	if ok { // такой товар есть - обновляем
		msg += "--обновиление--цены--старая:" + fmt.Sprintf("%.2f", itList[itemName]) +
			" новая:" + fmt.Sprintf("%.2f", itemPrice)
	} else {
		msg += "--новый--товар--по цене--:" + fmt.Sprintf("%.2f", itemPrice)
	}
	itList[itemName] = itemPrice
	return msg

}

/* Получить цену заказа по списку товаров - если товара
   нет в справочнике - сообщить об этом пользователю
*/
func getOrderCost(itList map[string]float32, order []string) float32 {
	var ordrCost float32 = 0
	for _, itemName := range order {
		value, ok := itList[itemName]
		if ok {
			ordrCost += value
		} else {
			fmt.Println(" товара >>" + itemName + "<< нет в списке")
		}
	}
	return ordrCost
}

func main() {
	itemsPrice := map[string]float32{"Спички": 1.2,
		"Хлеб":    20.15,
		"Сыр":     200.05,
		"Рыба":    150.45,
		"Сосиски": 300.45}

	fmt.Println(itemsPrice)
	fmt.Println(addItemsPrice(itemsPrice, "Сосиски", 255.41))
	fmt.Println(addItemsPrice(itemsPrice, "Ветчина", 600.32))
	fmt.Println(itemsPrice)

	fmt.Printf("Цена заказа %.2f",
		getOrderCost(itemsPrice, []string{"Хлеб", "Сосиски"}))

}
