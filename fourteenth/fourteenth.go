// задание 14 
package main

import "fmt"

func main() {
	fmt.Println("--- 14 ---")
	bslice := []int{0,1,2,3,4,5,6,7,8}  // наивно полагаю, что идут по очереди 
	even := []int{}  // четные
	odd  := []int{}  // нечетные
	for _, elem := range bslice {
		if elem % 2 == 0 {
			even = append(even , elem )	
		} else {
			odd = append(odd  , elem )	
		}
	} 
	fmt.Println( bslice ) // исходный 
	fmt.Println( even )   // выбрал все четные
	fmt.Println( odd  )   // выбрал все нечетные
	rez := []int{}


	for i:= 0 ; i < len(odd); i++ {  // клеим через один 
		rez = append( rez, odd[i] )
		rez = append( rez, even[i])
	}
	if len(even) != len(odd) { 
		//  заканчивается на четное - добавляем последний в хвост 
		rez = append( rez, even[len(odd)] )
	}
	fmt.Println( rez  )
}
