package shopcompetition

import (
	"fmt"
	"testing"
)

/*func TestRegister(t *testing.T) {
	accList := new(AccountList)
	nameList := []string{"Vasily", "Vasily", "Petr", "Oleg"}
	for _, name := range nameList {
		err := accList.Register(name)
		if err != nil {
			t.Errorf(" не смогли зарегестрировать такой пользователь %s уже есть ", "Vasily")
		}
	}
	t.Log(accList)
}
*/
func TestAddBalance(t *testing.T) {

	accList := new(AccountList)
	nameList := []string{"Vasily", "Petr", "Oleg"}
	for _, name := range nameList {
		err := accList.Register(name)
		if err != nil {
			t.Errorf(" не смогли зарегестрировать такой пользователь %s уже есть ", "Vasily")
		}
	}
	for _, name := range nameList {
		err := accList.AddBalance(name, 999.23)
		if err != nil {
			t.Errorf(" не смогли установить баланс пользователю %s", name)
		}
	}
	fmt.Println(accList)
}
