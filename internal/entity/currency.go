package entity

// ConversionResult represents the result of converting one currency to another.
type ConversionResult struct {
	FromCurrency    string  `json:"from_currency"`
	ToCurrency      string  `json:"to_currency"`
	OriginalAmount  float64 `json:"original_amount"`
	ConvertedAmount float64 `json:"converted_amount"`
	ConversionRate  float64 `json:"conversion_rate"`
	LastUpdated     string  `json:"last_updated"`
}

// NewConversionResult initializes a new ConversionResult entity.
func NewConversionResult(from, to string, originalAmount, convertedAmount, rate float64, lastUpdated string) ConversionResult {
	return ConversionResult{
		FromCurrency:    from,
		ToCurrency:      to,
		OriginalAmount:  originalAmount,
		ConvertedAmount: convertedAmount,
		ConversionRate:  rate,
		LastUpdated:     lastUpdated,
	}
}

// ConversionRequest represents the parameters required for a conversion request.
type ConversionRequest struct {
	Amount       float64
	FromCurrency string
	ToCurrency   string
}
