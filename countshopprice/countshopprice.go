//  задание 1
package main

import (
	"fmt"
	mu "mutils/structsdef"
)

/* Добавление нового товра в справочник - если есть измениться цена  */
func addItemsPrice(itemList map[int]mu.ItemPrice, item mu.ItemPrice) string {
	/* проверим наличие в каталоге */
	fnd := false
	vIdx := 0
	msg := item.ItemName
	var oldPrice float32 = 0
	for cidx, citem := range itemList {
		vIdx = cidx
		if citem.ItemName == item.ItemName {
			oldPrice = citem.ItemPrice
			fnd = true
			break
		}
	}
	if fnd {
		msg += "--обновиление--цены--старая:" + fmt.Sprintf("%.2f", oldPrice) +
			" новая:" + fmt.Sprintf("%.2f", item.ItemPrice)
		itemList[vIdx] = item
	} else {
		itemList[vIdx+1] = item
		msg += "--новый--товар--по цене--:" + fmt.Sprintf("%.2f", item.ItemPrice)
	}
	return msg
}

/* Получить цену заказа по списку товаров - если товара
   нет в справочнике - сообщить об этом пользователю
   вернуть заказ с посчитанной ценой
*/
func getOrderCost(itemList map[int]mu.ItemPrice, shopList mu.Order) float32 {
	var ordrCost float32 = 0
	for _, shopName := range shopList.Items { // бегу по списку товаров в заказе
		fond := false
		for _, item := range itemList { // бегу по списку товаров в каталоге
			if shopName == item.ItemName {
				ordrCost += item.ItemPrice
				fond = true
				break
			}
		}
		if !fond {
			fmt.Println(" товара >>" + shopName + "<< нет в каталоге")
		}
	}
	return ordrCost
}

/*
сохраним список товаров с ценой во время запроса от пользователя
*/

func compStrArr(in1, in2 []string) bool {
	if len(in1) != len(in2) {
		return false
	}
	rez := false
	for i := 0; i < len(in1); i++ {
		rez = (in1[i] == in2[i])
	}
	return rez
}

func seveListwithCost(
	ordersPrice []mu.Order, // списки товаров с ценами
	itemsPrice map[int]mu.ItemPrice, // справочник товаров
	itemsList mu.Order) mu.Order { // список товаров заказа
	// посмотрим , есть ли такая запись в справочнике список товаров - с ценой
	if ordersPrice == nil {
		itemsList.TotalSum = getOrderCost(itemsPrice, itemsList)
		ordersPrice = append(ordersPrice, itemsList)
		return itemsList
	}

	exists := false
	tmpordersPrice := mu.Order{}
	for oidx, oItem := range ordersPrice {
		tmpordersPrice = ordersPrice[oidx]
		if exists = compStrArr(oItem.Items, itemsList.Items); exists { // такой список товаров уже есть
			break
		}
	}
	if !exists {
		itemsList.TotalSum = getOrderCost(itemsPrice, itemsList)
		ordersPrice = append(ordersPrice, itemsList)
		return itemsList
	} else {
		fmt.Printf("Список %s уже есть \n", tmpordersPrice.Items)
		return tmpordersPrice
	}
}

func main() {
	itemsPrice := map[int]mu.ItemPrice{} // каталог товаров
	// --- положим немного данных в каталог
	itemsPrice[0] = mu.ItemPrice{ItemName: "Спички", ItemPrice: 1.2}
	itemsPrice[1] = mu.ItemPrice{ItemName: "Хлеб", ItemPrice: 20.15}
	itemsPrice[2] = mu.ItemPrice{ItemName: "Сыр", ItemPrice: 200.05}
	itemsPrice[3] = mu.ItemPrice{ItemName: "Рыба", ItemPrice: 150.45}
	itemsPrice[4] = mu.ItemPrice{ItemName: "Сосиски", ItemPrice: 300.45}

	fmt.Println("----- добавление товара в каталог -----")
	fmt.Println(itemsPrice)

	fmt.Println(addItemsPrice(itemsPrice, mu.ItemPrice{ItemName: "Сосиски", ItemPrice: 255.41}))
	fmt.Println(addItemsPrice(itemsPrice, mu.ItemPrice{ItemName: "Ветчина", ItemPrice: 600.32}))
	fmt.Println(itemsPrice)

	fmt.Println("----- получить цену заказа -----")
	vTempOrder := mu.Order{[]string{"Хлеб", "Сосиски", "Салями"}, 0}
	fmt.Printf("Цена заказа %.2f\n", getOrderCost(itemsPrice, vTempOrder))

	acountList := map[int]mu.User{} // каталог пользователей
	// --- положим немного данных о пользователях
	acountList[0] = mu.User{UserName: "Вася", Account: 800.0}
	acountList[1] = mu.User{UserName: "Коля", Account: 200.0}
	acountList[2] = mu.User{UserName: "Дима", Account: 300.0}
	acountList[3] = mu.User{UserName: "Петр", Account: 125.0}
	fmt.Println(acountList)

	// Список счетов - история покупок
	//         ID accountList --> ID Order --> Сумма заказа
	billList := map[int]map[int]float32{}
	billList[0] = map[int]float32{0: 0.0}
	billList[1] = map[int]float32{0: 0.0}
	billList[2] = map[int]float32{0: 0.0}
	billList[3] = map[int]float32{0: 0.0}
	// первоначально нулевая история
	fmt.Println(billList)

	fmt.Println("----- 7 -----")

	ordersPrice := []mu.Order{} // список заказов с посчитанной общей ценой

	fmt.Println(seveListwithCost(ordersPrice, itemsPrice, mu.Order{[]string{"Хлеб", "Сосиски"}, 0}))
	fmt.Println(seveListwithCost(ordersPrice, itemsPrice, mu.Order{[]string{"Хлеб", "Сыр"}, 0}))
	fmt.Println(seveListwithCost(ordersPrice, itemsPrice, mu.Order{[]string{"Хлеб", "Рыба"}, 0}))
	fmt.Println(seveListwithCost(ordersPrice, itemsPrice, mu.Order{[]string{"Хлеб", "Рыба"}, 0}))
	fmt.Println(seveListwithCost(ordersPrice, itemsPrice, mu.Order{[]string{"Хлеб", "Рыба", "Ветчина"}, 0}))

}
