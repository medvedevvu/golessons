// задание 5
// тут сделал совершенно не то, что надо !
package main

import (
	"fmt"
)

/*Добавление нового заказа*/
func addNewOrder(shopList map[string]map[string]float32,
	shname string,
	shopListItem map[string]float32) {
	shopList[shname] = shopListItem
}

/*Добавление нового товара заказа - если такой товар есть обновляем цену*/
func addNewOrderItem(shopList map[string]map[string]float32,
	shname string,
	newItem string, newVal float32) string {
	ok := true
	_, ok = shopList[shname]
	if !ok {
		return "такого пользователя нет в базе"
	}
	items := shopList[shname]

	oldVal, ok := items[newItem]
	items[newItem] = newVal
	msg := newItem

	if ok { // такой товар есть - обновляем
		msg += "--обновиление--цены--старая:" + fmt.Sprintf("%f", oldVal) +
			" новая:" + fmt.Sprintf("%f", newVal)
	} else {
		msg += "--новый--товар--по цене--:" + fmt.Sprintf("%f", newVal)
	}
	return msg
}

func getCostItem(shopList map[string]map[string]float32,
	shname string) float64 {
	ok := true
	_, ok = shopList[shname]
	if !ok { // если пользователя нет возвращаем -1
		return -1
	}
	items := shopList[shname]
	total := 0.0
	for _, value := range items {
		total += float64(value)
	}
	return total
}

func main() {
	shopList := make(map[string]map[string]float32)

	addNewOrder(shopList,
		"Покупатель1",
		map[string]float32{"Сыр": 100.3, "Лук": 25.5, "Хлеб": 15.3})
	addNewOrder(shopList,
		"Покупатель2",
		map[string]float32{"Сыр": 122.3, "Котлеты": 345.9, "Чай": 89.3})

	fmt.Println(shopList)
	fmt.Println(getCostItem(shopList, "Покупатель1"))
	fmt.Println(getCostItem(shopList, "Покупатель2"))

	fmt.Println(addNewOrderItem(shopList, "Покупатель2", "Рыба", 11.2))
	fmt.Println(shopList)
	fmt.Println(addNewOrderItem(shopList, "Покупатель2", "Рыба", 33.2))
	fmt.Println(shopList)
}
