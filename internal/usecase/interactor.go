package usecase

import (
	"context"
	"errors"

	"github.com/qjawko/currency-convertor/internal/domain"
	"github.com/qjawko/currency-convertor/internal/entity"
)

type CurrencyDataProvider interface {
	GetConversionRate(ctx context.Context, amount float64, fromCurrencyID, toCurrencyID string) (entity.CurrencyData, error)
}

type ConversionInteractor struct {
	DataProvider CurrencyDataProvider
}

func NewConversionInteractor(provider CurrencyDataProvider) *ConversionInteractor {
	return &ConversionInteractor{
		DataProvider: provider,
	}
}

func (c *ConversionInteractor) Convert(ctx context.Context, request domain.ConversionRequest) (domain.ConversionResult, error) {
	// Get conversion rate from the external data provider
	apiData, err := c.DataProvider.GetConversionRate(ctx, request.Amount, request.FromCurrency, request.ToCurrency)
	if err != nil {
		return domain.ConversionResult{}, err
	}

	// Extract the conversion rate from the API data
	rate, ok := apiData.Quote[request.ToCurrency]
	if !ok {
		return domain.ConversionResult{}, errors.New("failed to extract conversion rate from API response")
	}

	convertedAmount := request.Amount * rate.Price

	// Create and return the conversion result
	return domain.NewConversionResult(
		request.FromCurrency,
		request.ToCurrency,
		request.Amount,
		convertedAmount,
		rate.Price,
		rate.LastUpdated,
	), nil
}
