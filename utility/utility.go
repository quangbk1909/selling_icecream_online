package utility

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func DecodeDataFromJsonFile(f *os.File, data interface{}) error {
	byteValue, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return err
	}

	return nil
}
