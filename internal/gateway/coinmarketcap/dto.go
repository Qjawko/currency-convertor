package coinmarketcap

// ConvertRequestDTO represents the data required for a currency conversion request.
type ConvertRequestDTO struct {
	Amount       float64 `json:"amount"`
	FromCurrency string  `json:"from_currency"`
	ToCurrency   string  `json:"to_currency"`
}

// ConvertResponseDTO represents the result of a currency conversion.
type ConvertResponseDTO struct {
	FromCurrency    string  `json:"from_currency"`
	ToCurrency      string  `json:"to_currency"`
	OriginalAmount  float64 `json:"original_amount"`
	ConvertedAmount float64 `json:"converted_amount"`
	ConversionRate  float64 `json:"conversion_rate"`
	LastUpdated     string  `json:"last_updated"`
}

// APIErrorDTO represents an error message returned from an external API.
type APIErrorDTO struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
