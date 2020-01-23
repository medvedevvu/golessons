package shopcompetition

import (
	"testing"
)

func TetsRegister(t *testing.T) {
	accList := new(AccountList)
	err := accList.Register("Vasily")
	if err != nil {
		t.Errorf(" не смогли зарегестрировать ")
	}
	t.Log(accList)
}
