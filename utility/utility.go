package utility

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
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

func StringSearchText(text string) string {
	// sliceText := strings.Split(text, " ")
	// searchText := "+" + strings.Join(sliceText, " +")
	text = strings.ReplaceAll(text, "kem", "")
	text = strings.ReplaceAll(text, "Kem", "")
	return text
}
