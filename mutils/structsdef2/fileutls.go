package structsdef1

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	openFileError   = errors.New("Open file error")
	collectionEmpty = errors.New("Empty collection")
	temDirs         = `C:\MyGo\tmp\`
)

func gprep(colname string, gmode int64) *os.File {
	vcolname := fmt.Sprintf("%s.txt", colname)
	var f *os.File = nil
	var err error
	if gmode == 0 {
		f, err = os.OpenFile(fmt.Sprintf("%s%s", temDirs, vcolname),
			os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0777)
	} else {
		f, err = os.OpenFile(fmt.Sprintf("%s%s", temDirs, vcolname),
			os.O_RDONLY, 0777)
	}

	if err != nil {
		f = nil
		panic(openFileError)
	}
	return f
}

func gdef(f *os.File, colname string, gmode int64) {
	if r := recover(); r != nil {
		switch r {
		case openFileError:
			{
				fmt.Println(openFileError)
				f.Close()
			}
		case collectionEmpty:
			{
				fmt.Println(collectionEmpty)
			}
		default:
			{
				if f != nil {
					fmt.Println("неизвестная ошибка")
					f.Close()
				}
			}
		}
	} else {
		f.Close()
		switch gmode {
		case 0:
			{
				fmt.Printf("%s Saved !!!\n", colname)
			}
		case 1:
			{
				fmt.Printf("%s Loaded !!!\n", colname)
			}
		}

	}
}

//SaveAcountList - схраним пользователей
func SaveAcountList(vacountList map[string]*User) {
	if len(vacountList) == 0 {
		panic(collectionEmpty)
	}
	f := gprep("acountList", 0)
	for id, item := range vacountList {
		str := fmt.Sprintf("%s:%s:%f:%d\n",
			id, item.Email, item.Account, item.UserType)
		_, err := f.WriteString(str)
		if err != nil {
			fmt.Println("Строка %s пропущена с ошибкой %v ", str, err)
			continue
		}
	}
	defer gdef(f, "acountList", 0)
}

//LoadAcountList - загрузим пользователей
func LoadAcountList() map[string]*User {
	vacountList := map[string]*User{}
	f := gprep("acountList", 1)

	defer gdef(f, "acountList", 1)

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(errors.New(fmt.Sprintf("error reading file %s", err)))
		}
		line = strings.TrimSuffix(line, "\n") // убрать символ перевода строки
		strList := strings.Split(line, ":")
		fAccount, _ := strconv.ParseFloat(strList[2], 64)
		fAccount32 := float32(fAccount)
		fUserType, _ := strconv.ParseInt(strList[3], 10, 64)
		fUserType32 := int(fUserType)
		vuser := User{Email: strList[1],
			Account:  fAccount32,
			UserType: fUserType32}
		vacountList[strList[0]] = &vuser
	}
	return vacountList
}

//SaveBillList - схраним историю счтов
func SaveBillList(vbillList map[string]map[int]float32) {
	if len(vbillList) == 0 {
		panic(collectionEmpty)
	}
	f := gprep("billList", 0)
	for id, bills := range vbillList {
		strb := ""
		for idx, bill := range bills {
			strb = fmt.Sprintf("%s,%d:%f", strb, idx, bill)
		}
		str := fmt.Sprintf("%s#%s\n", id, strb[1:]) // выкинул лишнюю запятую
		_, err := f.WriteString(str)
		if err != nil {
			fmt.Println("Строка %s пропущена с ошибкой %v ", str, err)
			continue
		}
	}
	defer gdef(f, "billList", 0)
}

//LoadBillList - загрузим историю счтов
func LoadBillList() map[string]map[int]float32 {
	vbillList := map[string]map[int]float32{}
	f := gprep("billList", 1)

	defer gdef(f, "billList", 1)

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(errors.New(fmt.Sprintf("error reading file %s", err)))
		}

		line = strings.TrimSuffix(line, "\n") // убрать символ перевода строки

		strList := strings.Split(line, "#")
		uname := strList[0]
		vmap := map[int]float32{}
		blist := strings.Split(strList[1], ",")
		for _, item := range blist {
			myitem := strings.Split(item, ":")
			fb, _ := strconv.ParseInt(myitem[0], 10, 64)
			fb32 := int(fb)
			fa, _ := strconv.ParseFloat(myitem[1], 64)
			fa32 := float32(fa)
			vmap[fb32] = fa32
		}
		vbillList[uname] = vmap
	}
	return vbillList
}

//SaveItemsPrice - схраним каталог товаров
func SaveItemsPrice(vitemsPrice map[string]*ItemPrice) {
	if len(vitemsPrice) == 0 {
		panic(collectionEmpty)
	}
	f := gprep("itemsPrice", 0)
	for name, items := range vitemsPrice {
		strb := fmt.Sprintf("#ItemPrice:%f,ItemType:%d",
			items.ItemPrice, items.ItemType)
		str := fmt.Sprintf("%s%s\n", name, strb)
		_, err := f.WriteString(str)
		if err != nil {
			fmt.Println("Строка %s пропущена с ошибкой %v ", str, err)
			continue
		}
	}
	defer gdef(f, "itemsPrice", 0)
}

//LoadItemsPrice - загрузим каталог товаров
func LoadItemsPrice() map[string]*ItemPrice {
	vitemsPrice := map[string]*ItemPrice{}
	f := gprep("itemsPrice", 1)

	defer gdef(f, "itemsPrice", 1)

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(errors.New(fmt.Sprintf("error reading file %s", err)))
		}
		line = strings.TrimSuffix(line, "\n") // убрать символ перевода строки
		strList := strings.Split(line, "#")
		uname := strList[0]
		itemarr := strings.Split(strList[1], ",")
		var fa float64
		var fa32 float32
		var fb int64
		var fb32 int

		for _, items := range itemarr {
			aitems := strings.Split(items, ":")
			if aitems[0] == "ItemPrice" {
				fa, _ = strconv.ParseFloat(aitems[1], 64)
				fa32 = float32(fa)
			}

			if aitems[0] == "ItemType" {
				fb, _ = strconv.ParseInt(aitems[1], 10, 64)
				fb32 = int(fb)
			}
		}
		vitemsPrice[uname] = &ItemPrice{ItemPrice: fa32, ItemType: fb32}
	}
	return vitemsPrice
}

//SaveOrdersPrice - схраним заказы с ценами
func SaveOrdersPrice(vordersPrice []Order) {
	if len(vordersPrice) == 0 {
		panic(collectionEmpty)
	}
	f := gprep("ordersPrice", 0)
	for _, elem := range vordersPrice {
		str := strings.Join(elem.Items, ",")
		str = fmt.Sprintf("%s#%f,%d\n", str, elem.TotalSum, elem.OrderType)
		_, err := f.WriteString(str)
		if err != nil {
			fmt.Println("Строка %s пропущена с ошибкой %v ", str, err)
			continue
		}
	}
	defer gdef(f, "ordersPrice", 0)
}

//LoadOrdersPrice - загрузим заказы с ценами
func LoadOrdersPrice() []Order {
	vordersPrice := []Order{}
	f := gprep("ordersPrice", 1)

	defer gdef(f, "ordersPrice", 1)

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(errors.New(fmt.Sprintf("error reading file %s", err)))
		}
		line = strings.TrimSuffix(line, "\n") // убрать символ перевода строки
		strList := strings.Split(line, "#")
		items := strings.Split(strList[0], ",")
		itemarr := strings.Split(strList[1], ",")

		fa, _ := strconv.ParseFloat(itemarr[0], 64)
		fa32 := float32(fa)
		fb, _ := strconv.ParseInt(itemarr[1], 10, 64)
		fb32 := int(fb)
		vOrder := Order{Items: items, TotalSum: fa32, OrderType: fb32}
		vordersPrice = append(vordersPrice, vOrder)
	}
	return vordersPrice
}
