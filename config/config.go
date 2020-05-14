package config

import (
	"encoding/json"
	"io/ioutil"
)

type config interface {
	Unmarshal(path string) error
}

//Unmarshal config for bill items
func (b BillRenameItemList) Unmarshal(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	slice := BillRenameItemList{}
	err = json.Unmarshal(data, &slice)
	if err != nil {
		return err
	}
	b = slice
	return nil
}
