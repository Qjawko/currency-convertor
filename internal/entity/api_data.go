package entity

// ConversionResult represents the main data structure from the API response.
type ConversionResult struct {
	Data   []CurrencyData `json:"data"`
	Status Status         `json:"status"`
}

// CurrencyData represents the nested data structure in the API response.
type CurrencyData struct {
	Symbol      string           `json:"symbol"`
	ID          int              `json:"id"`
	Name        string           `json:"name"`
	Amount      float64          `json:"amount"`
	LastUpdated string           `json:"last_updated"`
	Quote       map[string]Quote `json:"quote"`
}

// Quote represents the conversion rate and last update time for a specific currency.
type Quote struct {
	Price       float64 `json:"price"`
	LastUpdated string  `json:"last_updated"`
}

// Status represents the status of the response.
type Status struct {
	Timestamp    string `json:"timestamp"`
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Elapsed      int    `json:"elapsed"`
	CreditCount  int    `json:"credit_count"`
	Notice       string `json:"notice"`
}
