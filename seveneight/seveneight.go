package main

import "fmt"

func seven() {
	fmt.Println("--- 7 задание ----")
	sl1 := []int{1, 2, 3, 4}
	sl2 := []int{5, 6, 7, 8}
	fmt.Println(sl1)
	fmt.Println(sl2)
	sl1 = append(sl1, sl2...)
	fmt.Println(sl1)
}

func eight() {
	fmt.Println("--- 8 задание ----")
	sl1 := []int{1, 6, 7, 6,9,4}
	sl2 := []int{5, 6, 7, 8,4,4}
	fmt.Println(sl1)
	fmt.Println(sl2)
	sl3 := []int{}
	for _, e := range sl1 {
		for _, e2 := range sl2 {
			if e == e2 {
				exist := false
				for _ , e3 := range sl3 {
					exist = ( e==e3 )
					if exist {
						break 
					}
				}
				if exist == false {
					sl3 = append( sl3 , e)
				}
			}
		}
	}
	sl4 := []int{}
	for _, e := range sl1 {
		exist := false
		for _,e1 := range sl3 {
			exist = ( e==e1 )
			if exist {
				break
			} 
		}
		if exist==false {
			sl4 = append( sl4 , e)
		}
	}
	fmt.Print(sl4)
}

func main() {
	seven()
	eight()
}
