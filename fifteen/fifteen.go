// задание 15 
package main

import (
	"fmt" 
	"sort"
)

func main() {
   bslice := []string{"az", "ax" , "cq" , "wa" , "fs" , "sd"}
   fmt.Println( bslice)   
   // --- обычная сортировка 
   sort.Strings( bslice )	
   fmt.Println( bslice , "----обычная") 
   // --- реверсивная сортировка
   sort.Sort( sort.Reverse( sort.StringSlice(bslice) ) )
   fmt.Println( bslice, "----реверсивная") 
   // --- лексикографическая
   /*
	   Когда писал сотрел вот это https://www.youtube.com/watch?v=4MkLnYxpflI
   */
   sort.SliceStable( bslice, 
	   func(i, j int) bool { 
		   return (bslice[i][0]< bslice[j][0]) && 
				  (bslice[i][len(bslice[i])-1]< 
				   bslice[j][len(bslice[j])-1])
	   })
	fmt.Println( bslice, "----лексикографическая")  
}
