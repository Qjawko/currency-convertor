package conversion

import "context"

// Response represents the result of a currency conversion.
type Response struct {
	ConvertedAmount float64
	FromCurrency    string
	ToCurrency      string
	Rate            float64
	LastUpdated     string
}

// CurrencyDataProvider defines methods to get currency data.
type CurrencyDataProvider interface {
	GetConversionRate(ctx context.Context, amount float64, fromCurrencyID, toCurrencyID string) (APIData, error)
}

// APIData represents the main data structure from the API response.
type APIData struct {
	Data   []Data `json:"data"`
	Status Status `json:"status"`
}

// Data represents the nested data structure in the API response.
type Data struct {
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

// Status represents the status of the API response.
type Status struct {
	Timestamp    string `json:"timestamp"`
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Elapsed      int    `json:"elapsed"`
	CreditCount  int    `json:"credit_count"`
	Notice       string `json:"notice"`
}
