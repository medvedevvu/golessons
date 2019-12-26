// задания 9 , 10 , 11 , 12 , 13
package main

import "fmt"

func main() {
	temp := []int {1,2,3,4,5,6,7}
	fmt.Println( "--- 9 ---"  )
    fmt.Println( temp     )
	fmt.Println( f9(temp) )

	fmt.Println( "--- 10 ---"  )
    fmt.Println( temp     )
	fmt.Println( f10(temp, 3 ) )

	fmt.Println( "--- 11 ---"  )
    fmt.Println( temp     )
	fmt.Println( f11(temp) )

	fmt.Println( "--- 12 ---"  )
    fmt.Println( temp     )
	fmt.Println( f12(temp, 3 ) )

	fmt.Println( "--- 13 ---"  )
	test := f13(temp)
	test[0] = 77777 
	fmt.Println( test )
    fmt.Println( temp )	
}

func f9(ltemp []int ) []int {
	return append( ltemp[1:] , ltemp[0:1]... )
}

func f10(ltemp []int , delta int ) []int {
	for i:= 0 ; i < delta ; i++ {
		ltemp= append( ltemp[1:] , ltemp[0:1]... )
	}
	return ltemp
}

func f11(ltemp []int ) []int {
	start := len(ltemp)-1
	end   := len(ltemp)
	return append( ltemp[start:end] , ltemp[0:start]...)
}

func f12(ltemp []int , delta int ) []int {
	for i:= 0 ; i < delta ; i++ {
	 start := len(ltemp)-1
	 end   := len(ltemp)
	 ltemp = append( ltemp[start:end] , ltemp[0:start]...)
	}
	return ltemp
}

func f13(ltemp []int ) []int {
	return append( []int{} , ltemp... )
}
