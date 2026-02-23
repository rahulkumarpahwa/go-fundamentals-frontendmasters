package api

import (
	"encoding/json"
	"fmt"
	"io"
	"mod16/datatypes"
	"net/http"
	"strings"
)

const apiURl = "https://open.er-api.com/v6/latest/%s"

func GetPrice(currency string) (*datatypes.Rate, error) {
	if len(currency) < 3 {
		return nil, fmt.Errorf("Required length 3, given : %d", len(currency))
	}
	currencyUpperCase := strings.ToUpper(currency)
	resp, err := http.Get(fmt.Sprintf(apiURl, currencyUpperCase))
	if err != nil {
		return nil, err
	}
	var currencyApiResponse CurrencyApiResponse
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodyBytes, &currencyApiResponse)
		if err != nil {
			return nil, err
		}

	} else {
		return nil, fmt.Errorf("The Status Code is Other than Ok, %v", resp.StatusCode)
	}

	rate := datatypes.Rate{Currency: currencyApiResponse.BaseCode, Price: currencyApiResponse.Rates[currencyUpperCase], List: currencyApiResponse.Rates}

	return &rate, err

}
