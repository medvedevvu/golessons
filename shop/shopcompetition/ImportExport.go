package shopcompetition

import "encoding/json"

/**
//Exporter - интерфейс для получения полного состояния магазина.
type Exporter interface {
	Export() ([]byte, error)
}

//Importer - интерфейс для загрузки состояния магазина. Принимает формат который возвращает Exporter.
type Importer interface {
	Import(data []byte) error
}
*/

func (ac AccountList) Export() ([]byte, error) {
	data, err := json.Marshal(ac)
	return data, err
}

func (ac *AccountList) Import(data []byte) error {
	err := json.Unmarshal(data, ac)
	return err
}

func (pl ProductList) Export() ([]byte, error) {
	data, err := json.Marshal(pl)
	return data, err
}

func (pl *ProductList) Import(data []byte) error {
	err := json.Unmarshal(data, pl)
	return err

}
