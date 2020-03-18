package shop_competition

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestManyGorutins(t *testing.T) {
	groutList := []string{"TestAddProduct", 
						   "TestRemoveProduct", 
						   "TestOptModifyProducts",
						   "TestAddAndRemoveBoundle",
						   "TestSimpleRemoveBoundle" ,
						   "TestAddBalance" ,
						   "TestCircleCheckBalance",
						   "TestAsyncPlaceOrder" ,
						   "TestPlaceOrderAndAddBalanc",
						}
	groutCount := 50
	var wg sync.WaitGroup
	for idx := 1; idx <= groutCount; idx++ {
		value := rand.Int31n(int32(len(groutList)))
		wg.Add(1)
		go func(value int32, t  *testing.T) {
			defer wg.Done()
			TaskGo(value,t)
			time.Sleep(time.Microsecond)
		}(value, t )
	}
	wg.Wait()
	fmt.Println("Ok")
}

func TaskGo(id int32, t *testing.T) {
	switch {
	case id == 0:
		TestAddProduct(t)
	case id == 1:
		TestRemoveProduct(t)
	case id == 2:
		TestOptModifyProducts(t)
	case id == 3:
		TestAddAndRemoveBoundle(t)
	case id == 4:
		TestSimpleRemoveBoundle(t)
	case id == 5:	
		TestAddBalance(t)
	case id == 6:	
		TestAsyncPlaceOrder(t)
	case id == 7:	
		TestPlaceOrderAndAddBalanc(t)
	}
}
