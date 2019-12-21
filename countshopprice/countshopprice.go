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
*/
func getOrderCost(itemList map[int]mu.ItemPrice, shopList []string) float32 {
	var ordrCost float32 = 0
	for _, shopName := range shopList { // бегу по списку товаров в заказе
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

func main() {
	itemsPrice := map[int]mu.ItemPrice{}
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
	fmt.Printf("Цена заказа %.2f\n", getOrderCost(itemsPrice, []string{"Хлеб", "Сосиски", "Салями"}))

	acountList := map[int]mu.User{}
	// --- положим немного данных о пользователях
	acountList[0] = mu.User{UserName: "Вася", Account: 800.0}
	acountList[1] = mu.User{UserName: "Коля", Account: 200.0}
	acountList[2] = mu.User{UserName: "Дима", Account: 300.0}
	acountList[2] = mu.User{UserName: "Петр", Account: 125.0}
	fmt.Println(acountList)

}
