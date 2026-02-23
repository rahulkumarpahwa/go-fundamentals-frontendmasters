package api

type CurrencyApiResponse struct {
	BaseCode string  `json:"base_code"`
	Rates map[string]float64 `json:"rates"` 
}

// Note : we don't need to write all the properties of the json response body.