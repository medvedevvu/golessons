package shopcompetition

/*type Shop interface {
	ProductModifier
	AccountManager
	OrderManager
	Importer
	Exporter
}*/

//NewShop  - constructor
func NewShop(acc AccountList, prdlist ProductList) *Shop {
	return &Shop{acc, prdlist}
}

//AddProduct
func (s *Shop) AddProduct(pr Product) PCSError {
	return s.ProductList.AddProduct(pr)
}

//ModifyProduct
func (s *Shop) ModifyProduct(pr Product) PCSError {
	return s.ProductList.ModifyProduct(pr)
}

// RemoveProduct
func (s *Shop) RemoveProduct(name string) PCSError {
	return s.ProductList.RemoveProduct(name)
}

// Register
func (s *Shop) Register(username string) error {
	return s.AccountList.Register(username)
}

//AddBalance
func (s *Shop) AddBalance(username string, sum float32) error {
	return s.AccountList.AddBalance(username, sum)
}

// Balance
func (s *Shop) Balance(username string) (float32, error) {
	bls, err := s.AccountList.Balance(username)
	return bls, err
}

//GetAccounts
func (s *Shop) GetAccounts(sort AccountSortType) AccountList {
	return s.AccountList.GetAccounts(sort)
}

// Export
func (s *Shop) Export() (dataAcc []byte, dataPrd []byte, errAcc error, errPrd error) {
	dataAcc, errAcc = s.AccountList.Export()
	dataPrd, errPrd = s.ProductList.Export()
	return dataAcc, dataPrd, errAcc, errPrd
}

// Export
func (s *Shop) Import(dataAcc, dataPrd []byte) (errAcc error, errPrd error) {
	errAcc = s.AccountList.Import(dataAcc)
	errPrd = s.ProductList.Import(dataPrd)
	return errAcc, errPrd
}
