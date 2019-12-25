//  задание 1
package main

import (
	"fmt"
	mu "mutils/structsdef"
)

/* Добавление нового товра в справочник - если есть измениться цена  */
func addItemsPrice(itemsPrice map[int]*mu.ItemPrice, item *mu.ItemPrice) string {
	/* проверим наличие в каталоге */
	fnd := false
	vIdx := 0
	msg := item.ItemName
	var oldPrice float32 = 0
	for cidx, citem := range itemsPrice {
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
		itemsPrice[vIdx] = item
	} else {
		itemsPrice[len(itemsPrice)] = item
		msg += "--новый--товар--по цене--:" + fmt.Sprintf("%.2f", item.ItemPrice)
	}
	return msg
}

func PrintCatalog(itemsPrice map[int]*mu.ItemPrice) {
	for _, item := range itemsPrice {
		fmt.Printf("Name: %s Price: %.2f \n", item.ItemName, item.ItemPrice)
	}
}

func PrintUsers(acountList map[int]*mu.User) {
	for _, item := range acountList {
		fmt.Printf("Name: %s Price: %.2f \n", item.UserName, item.Account)
	}
}

/* Получить цену заказа по списку товаров - если товара
   нет в справочнике - сообщить об этом пользователю
   вернуть заказ с посчитанной ценой
*/
func getOrderCost(itemsList map[int]*mu.ItemPrice, shopList mu.Order) float32 {
	var ordrCost float32 = 0
	for _, shopName := range shopList.Items { // бегу по списку товаров в заказе
		fond := false
		for _, item := range itemsList { // бегу по списку товаров в каталоге
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
	ordersPrice *[]mu.Order, // списки товаров с ценами
	itemsPrice map[int]*mu.ItemPrice, // справочник товаров
	itemsList mu.Order) mu.Order { // список товаров заказа
	// посмотрим , есть ли такая запись в справочнике список товаров - с ценой
	if len(*ordersPrice) == 0 {
		itemsList.TotalSum = getOrderCost(itemsPrice, itemsList)
		*ordersPrice = append(*ordersPrice, itemsList)
		return itemsList
	}

	exists := false
	//var tmpordersPrice float32 = 0
	for _, oItem := range *ordersPrice {
		//tmpordersPrice = oItem.TotalSum
		if exists = compStrArr(oItem.Items, itemsList.Items); exists { // такой список товаров уже есть
			break
		}
	}
	itemsList.TotalSum = getOrderCost(itemsPrice, itemsList) // счет в любом случае
	if !exists {
		*ordersPrice = append(*ordersPrice, itemsList)
		return itemsList
	} else {
		fmt.Printf("Список %s уже есть \n", itemsList.Items)
		return itemsList
	}
}

/*
   Регистрация заказа с корректировкой остатка у пользователя
*/

func orderRegister(acountList map[int]*mu.User, // список пользователей
	ordersPrice *[]mu.Order, // списки товаров с ценами
	itemsPrice map[int]*mu.ItemPrice, // справочник товаров
	billList map[int]map[int]float32, // список счетов
	user mu.User, // пользователь
	itemsList mu.Order) { // заказ
	// проверим пользователя
	var ostatok float32 = 0
	//var totalCost float32 = 0
	var vIxd int = -1
	for idx, iUser := range acountList {
		if iUser.UserName == user.UserName {
			vIxd = idx
			ostatok = iUser.Account
			break
		}
	}
	if vIxd == -1 { // индекс пользователя не найден
		fmt.Printf("Пользователь %s не регистрирован !\n", user.UserName)
		return
	}
	// проверим кредитоспособность
	if ostatok <= 0 {
		fmt.Printf("У пользователя %s нет средств на счету %.2f !\n", user.UserName, ostatok)
	}
	// добавить ветку просмотра

	tmp := seveListwithCost(ordersPrice, itemsPrice, itemsList)
	var saldo float32 = ostatok - tmp.TotalSum
	if saldo >= 0 {
		//var x = (*acountList)[vIxd]
		//x.Account = saldo
		//(*acountList)[vIxd] = x
		acountList[vIxd].Account = saldo
		// сохраним успешный вариант
		// сохраним списание
		billList[vIxd][len(billList[vIxd])] = tmp.TotalSum

		fmt.Printf("Списание выполнено , пользователь %s остаток: %.2f списание: %.2f сальдо: %.2f  !\n",
			user.UserName, ostatok, tmp.TotalSum, saldo)

	} else {
		fmt.Printf("У пользователя %s остаток: %.2f списание: %.2f сальдо: %.2f - не достаточно средств !\n",
			user.UserName, ostatok, tmp.TotalSum, saldo)
	}

}

//func showAccount(acountList map[int]mu.User) {
//
//}

func main() {
	itemsPrice := map[int]*mu.ItemPrice{} // каталог товаров
	// --- положим немного данных в каталог
	itemsPrice[0] = &mu.ItemPrice{ItemName: "Спички", ItemPrice: 1.2}
	itemsPrice[1] = &mu.ItemPrice{ItemName: "Хлеб", ItemPrice: 20.15}
	itemsPrice[2] = &mu.ItemPrice{ItemName: "Сыр", ItemPrice: 200.05}
	itemsPrice[3] = &mu.ItemPrice{ItemName: "Рыба", ItemPrice: 150.45}
	itemsPrice[4] = &mu.ItemPrice{ItemName: "Сосиски", ItemPrice: 300.45}

	fmt.Println("----- добавление товара в каталог -----")
	PrintCatalog(itemsPrice)

	fmt.Println(addItemsPrice(itemsPrice, &mu.ItemPrice{ItemName: "Сосиски", ItemPrice: 255.41}))
	fmt.Println(addItemsPrice(itemsPrice, &mu.ItemPrice{ItemName: "Ветчина", ItemPrice: 600.32}))
	PrintCatalog(itemsPrice)

	fmt.Println("----- получить цену заказа -----")
	vTempOrder := mu.Order{[]string{"Хлеб", "Сосиски", "Салями"}, 0}
	fmt.Printf("Цена заказа %.2f\n", getOrderCost(itemsPrice, vTempOrder))

	PrintCatalog(itemsPrice)

	acountList := map[int]*mu.User{} // каталог пользователей
	// --- положим немного данных о пользователях
	acountList[0] = &mu.User{UserName: "Вася", Account: 800.0}
	acountList[1] = &mu.User{UserName: "Коля", Account: 200.0}
	acountList[2] = &mu.User{UserName: "Дима", Account: 300.0}
	acountList[3] = &mu.User{UserName: "Петр", Account: 125.0}
	PrintUsers(acountList)

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

	fmt.Println(seveListwithCost(&ordersPrice, itemsPrice, mu.Order{[]string{"Хлеб", "Сосиски"}, 0}))
	fmt.Println(seveListwithCost(&ordersPrice, itemsPrice, mu.Order{[]string{"Хлеб", "Сыр"}, 0}))
	fmt.Println(seveListwithCost(&ordersPrice, itemsPrice, mu.Order{[]string{"Хлеб", "Рыба"}, 0}))
	fmt.Println(seveListwithCost(&ordersPrice, itemsPrice, mu.Order{[]string{"Хлеб", "Рыба"}, 0}))
	fmt.Println(seveListwithCost(&ordersPrice, itemsPrice, mu.Order{[]string{"Хлеб", "Рыба", "Ветчина"}, 0}))

	fmt.Println(ordersPrice)
	fmt.Println("----- 8 -----")
	PrintUsers(acountList)
	fmt.Println("---------------------------")
	orderRegister(acountList, // списки пользователь
		&ordersPrice,   // списки товаров с ценами
		itemsPrice,     // справочник товаров
		billList,       // список счетов
		*acountList[0], // пользователь
		mu.Order{[]string{"Хлеб", "Рыба", "Ветчина"}, 0}) // список товаров

	orderRegister(acountList, // списки пользователь
		&ordersPrice,   // списки товаров с ценами
		itemsPrice,     // справочник товаров
		billList,       // список счетов
		*acountList[0], // пользователь
		mu.Order{[]string{"Хлеб", "Рыба", "Ветчина"}, 0}) // список товаров

	orderRegister(acountList, // списки пользователь
		&ordersPrice,   // списки товаров с ценами
		itemsPrice,     // справочник товаров
		billList,       // список счетов
		*acountList[2], // пользователь
		mu.Order{[]string{"Хлеб", "Сосиски"}, 0}) // список товаров

	fmt.Println("---------------------------")
	PrintUsers(acountList)
	fmt.Println(billList)

	//fmt.Println("----- 9 -----")
	//fmt.Println("----- по имени        -----")

	//showAccount(acountList, 0)
	/*fmt.Println("----- по имени реверс -----")
	showAccount(acountList, 1)
	fmt.Println("----- по деньгам      -----")
	showAccount(acountList, 2)
	fmt.Println("----- по деньгам инверсия---")
	showAccount(acountList, 3)*/
}