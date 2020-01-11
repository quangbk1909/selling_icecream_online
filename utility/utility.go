package utility

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"vinid_project/model"

	"github.com/jasonwinn/geocoder"
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
	sliceText := strings.Split(text, " ")
	// searchText := "+" + strings.Join(sliceText, " +")
	if len(sliceText) <= 1 {
		return text
	} else {
		text = strings.ReplaceAll(text, "kem", "")
		text = strings.ReplaceAll(text, "Kem", "")
		return text
	}
}

func MakeResponse(statusCode int, message string, data interface{}) model.ResponseForm {
	var response model.ResponseForm
	var metaDataRes model.MetaDataResponse
	metaDataRes.Code = statusCode
	metaDataRes.Message = message

	response.Data = data
	response.Meta = metaDataRes

	return response

}

func GetAddressFromCoordinates(latitude float64, longitude float64) (string, error) {
	address, err := geocoder.ReverseGeocode(latitude, longitude)
	if err != nil {
		return "", err
	}
	return address.Street + ", " + address.City, nil
}
