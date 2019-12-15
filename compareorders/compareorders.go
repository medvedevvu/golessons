// задание 7 - 8  - 9 - 10
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

/*
сохраним список товаров с ценой во время запроса от пользователя
*/
func seveListwithCost(
	ordersPrice map[float32][]string, // списки товаров с ценами
	itemsPrice map[string]float32, // справочник товаров
	itemsList []string) { // список товаров
	// посмотрим , есть ли такая запись в справочнике список товаров - с ценой
	exists := false
	for _, oList := range ordersPrice {
		if len(itemsList) != len(oList) {
			continue
		}
		for idx := 0; idx < len(itemsList); idx++ {
			if itemsList[idx] == oList[idx] {
				exists = true
			} else {
				exists = false
			}
		}
		if exists {
			break
		}
	}
	if !exists {
		ordersPrice[getOrderCost(itemsPrice, itemsList)] = itemsList
	} else {
		fmt.Printf("Список %s уже есть \n", itemsList)
	}
}

/*
   Регистрация заказа с корректировкой остатка у пользователя
*/
func orderRegister(username string,
	acountList map[string]float32,
	ordersPrice map[float32][]string, // списки товаров с ценами
	itemsPrice map[string]float32,   // справочник товаров
	itemsList []string ,             // список товаров
	billList map[string][]float32    // список счетов 
	) { 
	// проверим пользователя 
	ostatok , ok := acountList[username]
	if !ok {
		fmt.Printf("Пользователь %s не регистрирован !" , username)
	}
	// проверим кредитоспособность
	if ostatok <= 0 {
		fmt.Printf("У пользователя %s остаток на счету %.2f !" , username , ostatok )
	}
	// добавить ветку просмотра 
	totalCost := getOrderCost(itemsPrice, itemsList )
	saldo := ostatok-totalCost
	if saldo >= 0 {
		acountList[username] = saldo
		// сохраним успешный вариант 
		seveListwithCost(ordersPrice, itemsPrice, itemsList ) 
		// сохраним списание 
		billList[username] = append( billList[username] , totalCost )
		fmt.Printf("Списание выполнено , пользователь %s остаток: %.2f списание: %.2f сальдо: %.2f  !" , 
		           username , ostatok , totalCost , saldo )	
	
	} else {
		fmt.Printf("У пользователя %s остаток: %.2f списание: %.2f сальдо: %.2f - не достаточно средств !" , 
		           username , ostatok , totalCost , saldo )	
	}

}

func main() {
	itemsPrice := map[string]float32{"Спички": 1.2,
		"Хлеб":    20.15,
		"Сыр":     200.05,
		"Рыба":    150.45,
		"Сосиски": 300.45}

	ordersPrice := map[float32][]string{} // список товаров с посчитанной общей ценой

	acountList := map[string]float32{ // Список пользователей
		"Вася": 100.0,
		"Коля": 200.0,
		"Дима": 300.0,
		"Петр": 125.0}

	billList := map[string][]float32{ // Список счетов
		"Вася": {0},
		"Коля": {0},
		"Дима": {0},
		"Петр": {0}}

	fmt.Println("----- 5 -----")
	fmt.Println(itemsPrice)
	fmt.Println(addItemsPrice(itemsPrice, "Сосиски", 255.41))
	fmt.Println(addItemsPrice(itemsPrice, "Ветчина", 600.32))
	fmt.Println(itemsPrice)

	fmt.Printf("Цена заказа %.2f\n",
		getOrderCost(itemsPrice, []string{"Хлеб", "Сосиски"}))

	fmt.Println("----- 7 -----")

	seveListwithCost(ordersPrice, itemsPrice, []string{"Хлеб", "Сосиски"})
	seveListwithCost(ordersPrice, itemsPrice, []string{"Хлеб", "Сыр"})
	seveListwithCost(ordersPrice, itemsPrice, []string{"Хлеб", "Рыба"})
	seveListwithCost(ordersPrice, itemsPrice, []string{"Хлеб", "Рыба"})
	seveListwithCost(ordersPrice, itemsPrice, []string{"Хлеб", "Рыба", "Ветчина"})
	fmt.Println(ordersPrice)

	fmt.Println("----- 8 -----")

}
