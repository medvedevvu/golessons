//задание 2, 3
package main

import (
	"flag"
	"fmt"
	mu "mutils/structsdef2"
	"sort"
)

/* Добавление нового товра в справочник - если есть измениться цена и тип */

func addItemsPrice(itemsPrice map[string]*mu.ItemPrice,
	itemName string,
	item *mu.ItemPrice) string {
	/* проверим наличие в каталоге */
	vItemPrice, ok := itemsPrice[itemName]
	msg := itemName
	if !ok { // не нашли добавляем в каталог
		itemsPrice[itemName] = item
		msg += "--новый--товар--по цене--:" + fmt.Sprintf("%.2f", item.ItemPrice)
	} else { // нашли - обноавляем цену и тип
		msg += "--обновиление--цены--старая:" + fmt.Sprintf("%.2f", vItemPrice.ItemPrice) +
			"-- старый--тип:" + fmt.Sprintf("%d", vItemPrice.ItemType) +
			" новая:" + fmt.Sprintf("%.2f", item.ItemPrice) +
			" новый тип:" + fmt.Sprintf("%d", item.ItemType)
		itemsPrice[itemName] = item
	}
	return msg
}

// PrintCatalog печать каталога
func PrintCatalog(itemsPrice map[string]*mu.ItemPrice) {
	for name, item := range itemsPrice {
		fmt.Printf("Name: %s Price: %.2f  Type: %d \n", name, item.ItemPrice,
			item.ItemType)
	}
}

// PrintUsers печать каталога
func PrintUsers(acountList map[string]*mu.User) {
	for name, item := range acountList {
		fmt.Printf("Name: %s  Email: %s Account: %.2f  Type: %d\n", name, item.Email, item.Account, item.UserType)
	}
}

func createOrder(itemsList map[string]*mu.ItemPrice, orderItems []string) mu.Order {
	tovar := []string{}
	probnik := []string{}
	itog := []string{}
	/* определимся сколько товаров и сколько пробников
	   пробники можем добавлять только к комплекту = 2 товара + пробник
	*/
	for _, name := range orderItems {
		vitemsList, ok := itemsList[name]
		if ok {
			switch vitemsList.ItemType {
			case 2:
				probnik = append(probnik, name)
			default:
				tovar = append(tovar, name)
			}
		}

	}
	/*
	   два товра + пробник == 3
	   если выбрал только пробники - вернуть пустой набор nil
	   если больше 1,2,...,4....> товаров - то вернуть список без пробника
	*/
	switch {
	case (len(probnik) == 1 && len(tovar) == 2):
		{
			itog = append(itog, probnik...)
			itog = append(itog, tovar...)
			return mu.Order{Items: itog, OrderType: 0, TotalSum: 0}
		}
	case (len(probnik) != 0 && len(tovar) == 0):
		{
			return mu.Order{}
		}
	case (len(probnik) == 0 && len(tovar) == 0):
		{
			return mu.Order{}
		}
	default:
		{
			itog = append(itog, tovar...)
			return mu.Order{Items: itog, OrderType: 0, TotalSum: 0}
		}
	}
}

/* Получить цену заказа по списку товаров - если товара
   нет в справочнике - сообщить об этом пользователю
   вернуть заказ с посчитанной ценой
*/
func getOrderCost(itemsList map[string]*mu.ItemPrice, shopList mu.Order) float32 {
	var ordrCost float32 = 0
	for _, shopName := range shopList.Items { // бегу по списку товаров в заказе
		vitemList, ok := itemsList[shopName]
		if ok { // нашли
			if vitemList.ItemType != 2 {
				ordrCost += vitemList.ItemPrice
			} else {
				fmt.Println(" товар >>" + shopName + "<< пробник с нулевой ценой ")
			}

		} else {
			fmt.Println(" товара >>" + shopName + "<< нет в каталоге")
		}
	}
	return ordrCost
}

/* Получить тип заказа по списку товаров -
 */
func getOrderType(itemsList map[string]*mu.ItemPrice,
	shopList mu.Order) int {
	/*
	   0 - товар
	   1 - набор (всегда 2 товара )
	   2 - (набор + пробник или товар + пробник )
	*/
	filteredItems := []string{}
	for _, name := range shopList.Items {
		_, ok := itemsList[name]
		if ok {
			filteredItems = append(filteredItems, name)
		}
	}
	if len(filteredItems) == 3 { // набор
		isNaborPlusProbnik := false
		countProbniki := 0
		for _, name := range filteredItems {
			vitemsList, ok := itemsList[name]
			if ok {
				if vitemsList.ItemType == 2 {
					isNaborPlusProbnik = true
					countProbniki++
				}
			}
		}
		if isNaborPlusProbnik && (countProbniki == 1) {
			return 2
		} else {
			return 1
		}
	} else {
		switch {
		case len(filteredItems) == 2:
			{
				return 1
			}
		case len(filteredItems) > 0:
			{
				return 0
			}
		default:
			{
				return -1
			}
		}

	}
}

/*
 используется для сравнения списка закза и сохраненные заказы с ценами 203 строка
  в функции seveListwithCostAndOrderType
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

/*
сохраним список товаров с ценой во время запроса от пользователя
*/

func seveListwithCostAndOrderType(
	ordersPrice *[]mu.Order, // списки товаров с ценами
	itemsPrice map[string]*mu.ItemPrice, // справочник товаров
	itemsList mu.Order) mu.Order { // список товаров заказа
	// посмотрим , есть ли такая запись в справочнике список товаров - с ценой
	if len(*ordersPrice) == 0 { // список пустой
		itemsList.TotalSum = getOrderCost(itemsPrice, itemsList)
		itemsList.OrderType = getOrderType(itemsPrice, itemsList)
		*ordersPrice = append(*ordersPrice, itemsList)
		return itemsList
	}

	exists := false
	//var tmpordersPrice float32 = 0
	for _, oItem := range *ordersPrice {
		if exists = compStrArr(oItem.Items, itemsList.Items); exists && (oItem.OrderType == itemsList.OrderType) { // такой список товаров уже есть
			break
		}
	}
	itemsList.TotalSum = getOrderCost(itemsPrice, itemsList) // счет в любом случае
	itemsList.OrderType = getOrderType(itemsPrice, itemsList)
	if !exists {
		*ordersPrice = append(*ordersPrice, itemsList)
		return itemsList
	} else {
		fmt.Printf("Список %s с типом %d уже есть \n", itemsList.Items, itemsList.OrderType)
		return itemsList
	}
}

/*
   Регистрация заказа с корректировкой остатка у пользователя
*/

func orderRegister(acountList map[string]*mu.User, // список пользователей
	ordersPrice *[]mu.Order, // списки товаров с ценами
	itemsPrice map[string]*mu.ItemPrice, // справочник товаров
	billList map[string]map[int]float32, // список счетов
	UserName string, // имя пользователя
	itemsList mu.Order,
	// корректировочная скидка
	addDisc float32) { // заказ
	// проверим пользователя
	vacountList, ok := acountList[UserName]
	if !ok {
		fmt.Printf("Пользователь %s не регистрирован !\n", UserName)
		return
	}
	// проверим кредитоспособность
	if vacountList.Account <= 0 {
		fmt.Printf("У пользователя %s нет средств на счету %.2f !\n",
			UserName, vacountList.Account)
	}
	// добавить ветку просмотра
	tmp := seveListwithCostAndOrderType(ordersPrice, itemsPrice, itemsList)
	/*
	  Воспользуемся SetDiscount - сделаем скидку
	*/
	vtnoDiscount,
		vtnoDiscountByComplect,
		vorderTotalSum := tmp.SetDiscount(UserName, acountList, itemsPrice, addDisc)

	if (vorderTotalSum != vtnoDiscount) ||
		(vorderTotalSum != vtnoDiscountByComplect) {
		fmt.Printf("СКИДКА ЕСТЬ :  без скидка  %.2f --\n "+
			"(скидка пользователь и комплект ) %.2f --\n "+
			"(скидка пользователь и комплект и скидочный коэф. )%.2f\n",
			vtnoDiscount, vtnoDiscountByComplect, vorderTotalSum,
		)
	} else {
		fmt.Printf("СКИДКИ НЕТ :  без скидка  %.2f --\n "+
			"(скидка пользователь и комплект ) %.2f --\n "+
			"(скидка пользователь и комплект и скидочный коэф. )%.2f\n",
			vtnoDiscount, vtnoDiscountByComplect, vorderTotalSum,
		)
	}

	saldo := vacountList.Account - tmp.TotalSum
	ostatok := vacountList.Account
	if saldo >= 0 {
		vacountList.Account = saldo
		// сохраним успешный вариант
		// сохраним списание
		billList[UserName][len(billList[UserName])] = tmp.TotalSum

		fmt.Printf("Списание выполнено , пользователь %s остаток: %.2f списание: %.2f сальдо: %.2f  !\n",
			UserName, ostatok, tmp.TotalSum, saldo)

	} else {
		fmt.Printf("У пользователя %s остаток: %.2f списание: %.2f сальдо: %.2f - не достаточно средств !\n",
			UserName, ostatok, tmp.TotalSum, saldo)
	}

}

func showAccount(acountList map[string]*mu.User, p int) {
	//keys := reflect.ValueOf(acountList).MapKeys()
	keys := []string{}
	for k := range acountList {
		keys = append(keys, k)
	}
	accs := []float64{}
	for _, k := range acountList {
		accs = append(accs, float64(k.Account))
	}

	switch p {
	case 0:
		{
			sort.Strings(keys)
		}
	case 1:
		{
			sort.Sort(sort.Reverse(sort.StringSlice(keys)))
		}
	case 2:
		{
			sort.Float64s(accs)
		}
	case 3:
		{
			sort.Sort(sort.Reverse(sort.Float64Slice(accs)))
		}
	default:
		fmt.Println("--- такой опции нет ---")
	}

	/*
	   Когда сделал вывод вот так перестала слетать сортировка при печати,
	   не знаю на сколько такой способ правильный , но
	   когда делаешь через range - сортировка слетает при выводе на печать.
	*/
	//keys := reflect.ValueOf(acountList).MapKeys()

	if p == 0 || p == 1 {
		for i := 0; i < len(keys); i++ {
			fmt.Printf("Name: %s Price: %.2f \n", keys[i],
				acountList[keys[i]].Account)
		}
	}
	if p == 2 || p == 3 {
		for i := 0; i < len(accs); i++ {
			for name, j := range acountList {
				if float64(j.Account) == accs[i] {
					fmt.Printf("Name: %s Price: %.2f \n", name,
						j.Account)
				}
			}
		}
	}
}

func main() {

	pMode := flag.String("pmod", "n", "s - save / l - load ")
	flag.Parse()
	valMode := *pMode

	acountList := map[string]*mu.User{} // каталог пользователей
	// --- положим немного данных о пользователях
	if valMode == "l" { // грузим из файла
		acountList = mu.LoadAcountList()
	} else {
		acountList["Вася"] = &mu.User{Email: "vasya@vlg.ru", Account: 800.0, UserType: 1}
		acountList["Коля"] = &mu.User{Email: "kolya@volgatel.ru", Account: 200.0, UserType: 0}
		acountList["Дима"] = &mu.User{Email: "dima@mail.ru", Account: 300.0, UserType: 1}
		acountList["Петр"] = &mu.User{Email: "petr@onix.ru", Account: 125.0, UserType: 0}
	}
	PrintUsers(acountList)

	// Список счетов - история покупок
	//         UserNAme accountList --> ID Order --> Сумма заказа
	billList := map[string]map[int]float32{}

	if valMode == "l" { // грузим из файла
		billList = mu.LoadBillList()
	} else {
		billList["Вася"] = map[int]float32{0: 0.0}
		billList["Коля"] = map[int]float32{0: 0.0}
		billList["Дима"] = map[int]float32{0: 0.0}
		billList["Петр"] = map[int]float32{0: 0.0}
	}
	// первоначально нулевая история
	fmt.Println(billList)

	//                ItemName
	itemsPrice := map[string]*mu.ItemPrice{} // каталог товаров
	// --- положим немного данных в каталог
	if valMode == "l" { // грузим из файла
		itemsPrice = mu.LoadItemsPrice()
	} else {
		itemsPrice["Спички"] = &mu.ItemPrice{ItemPrice: 1.2, ItemType: 0}
		itemsPrice["Хлеб"] = &mu.ItemPrice{ItemPrice: 20.15, ItemType: 0}
		itemsPrice["Сыр"] = &mu.ItemPrice{ItemPrice: 200.05, ItemType: 0}
		itemsPrice["Рыба"] = &mu.ItemPrice{ItemPrice: 150.45, ItemType: 1}
		itemsPrice["Сосиски"] = &mu.ItemPrice{ItemPrice: 300.45, ItemType: 0}
		itemsPrice["Зубочистки"] = &mu.ItemPrice{ItemPrice: 0, ItemType: 2}

		fmt.Println("----- добавление товара в каталог -----")
		PrintCatalog(itemsPrice)

		fmt.Println(addItemsPrice(itemsPrice, "Сосиски",
			&mu.ItemPrice{ItemPrice: 255.41, ItemType: 1}))
		fmt.Println(addItemsPrice(itemsPrice, "Ветчина",
			&mu.ItemPrice{ItemPrice: 600.32, ItemType: 1}))
	}
	PrintCatalog(itemsPrice)

	fmt.Println("----- сформировать заказ -----")
	v1Order := createOrder(itemsPrice,
		[]string{"Хлеб", "Сосиски", "Рыба", "Зубочистки"})
	fmt.Println(v1Order)
	v2Order := createOrder(itemsPrice,
		[]string{"Хлеб", "Сосиски", "Зубочистки"})
	fmt.Println(v2Order)

	fmt.Println("----- получить цену и тип заказа -----")
	vTotalSum := getOrderCost(itemsPrice, v1Order)  // цена заказа
	vOrderType := getOrderType(itemsPrice, v1Order) // тип заказа
	fmt.Printf("Цена заказа %.2f Тип заказа: %d \n ", vTotalSum, vOrderType)

	vTotalSum = getOrderCost(itemsPrice, v2Order)  // цена заказа
	vOrderType = getOrderType(itemsPrice, v2Order) // тип заказа
	fmt.Printf("Цена заказа %.2f Тип заказа: %d \n ", vTotalSum, vOrderType)

	//PrintCatalog(itemsPrice)

	fmt.Println("----- 7 -----")

	// список заказов с посчитанной общей ценой и типом заказа

	ordersPrice := []mu.Order{}
	if valMode == "l" { // грузим из файла
		ordersPrice = mu.LoadOrdersPrice()
	} else {

		fmt.Println(seveListwithCostAndOrderType(&ordersPrice,
			itemsPrice,
			mu.Order{[]string{"Хлеб", "Сосиски", "Зубочистки"}, 0, 0}))
		fmt.Println(seveListwithCostAndOrderType(&ordersPrice,
			itemsPrice,
			mu.Order{[]string{"Хлеб", "Сыр"}, 0, 0}))
		fmt.Println(seveListwithCostAndOrderType(&ordersPrice,
			itemsPrice,
			mu.Order{[]string{"Хлеб", "Рыба"}, 0, 0}))
		fmt.Println(seveListwithCostAndOrderType(&ordersPrice,
			itemsPrice,
			mu.Order{[]string{"Хлеб", "Рыба"}, 0, 0}))
		fmt.Println(seveListwithCostAndOrderType(&ordersPrice,
			itemsPrice,
			mu.Order{[]string{"Хлеб", "Рыба", "Ветчина"}, 0, 0}))
	}
	fmt.Println(ordersPrice)

	fmt.Println("----- 8 -----")
	PrintUsers(acountList)
	fmt.Println("---------------------------")
	orderRegister(acountList, // списки пользователь
		&ordersPrice, // списки товаров с ценами
		itemsPrice,   // справочник товаров
		billList,     // список счетов
		"Вася",       // пользователь
		mu.Order{[]string{"Хлеб", "Рыба", "Ветчина"}, 0, 0}, 0.8) // список товаров

	orderRegister(acountList, // списки пользователь
		&ordersPrice, // списки товаров с ценами
		itemsPrice,   // справочник товаров
		billList,     // список счетов
		"Вася",       // пользователь
		mu.Order{[]string{"Хлеб", "Рыба", "Ветчина"}, 0, 0}, 0.8) // список товаров

	orderRegister(acountList, // списки пользователь
		&ordersPrice, // списки товаров с ценами
		itemsPrice,   // справочник товаров
		billList,     // список счетов
		"Дима",       // пользователь
		mu.Order{[]string{"Хлеб", "Сосиски"}, 0, 0}, 0.8) // список товаров

	orderRegister(acountList, // списки пользователь
		&ordersPrice, // списки товаров с ценами
		itemsPrice,   // справочник товаров
		billList,     // список счетов
		"Дима",       // пользователь
		mu.Order{[]string{"Хлеб", "Сосиски"}, 0, 0}, 0.8) // список товаров

	fmt.Println("---------------------------")
	//PrintUsers(acountList)
	fmt.Println(billList)

	fmt.Println("----- 9 -----")
	fmt.Println("----- по имени        -----")
	showAccount(acountList, 0)

	fmt.Println("----- по имени реверс -----")
	showAccount(acountList, 1)
	fmt.Println("----- по деньгам      -----")
	showAccount(acountList, 2)
	fmt.Println("----- по деньгам инверсия---")
	showAccount(acountList, 3)

	fmt.Println("----- сохранить справочники в файл ---")
	if valMode == "s" || valMode == "ls" {
		mu.SaveAcountList(acountList)
		mu.SaveBillList(billList)
		mu.SaveItemsPrice(itemsPrice)
		mu.SaveOrdersPrice(ordersPrice)
	}

}
