package usecase

import (
	"context"
	"errors"

	"github.com/qjawko/currency-convertor/internal/entity"
	"github.com/qjawko/currency-convertor/internal/usecase/conversion"
)

type ConversionInteractor struct {
	DataProvider conversion.CurrencyDataProvider
}

func NewConversionInteractor(provider conversion.CurrencyDataProvider) *ConversionInteractor {
	return &ConversionInteractor{
		DataProvider: provider,
	}
}

func (c *ConversionInteractor) Convert(ctx context.Context, request entity.ConversionRequest) (entity.ConversionResult, error) {
	// Get conversion rate from the external data provider
	apiData, err := c.DataProvider.GetConversionRate(ctx, request.Amount, request.FromCurrency, request.ToCurrency)
	if err != nil {
		return entity.ConversionResult{}, err
	}

	if len(apiData.Data) < 1 {
		return entity.ConversionResult{}, errors.New("unexpected amount of data")
	}

	// Extract the conversion rate from the API data
	rate, ok := apiData.Data[0].Quote[request.ToCurrency]
	if !ok {
		return entity.ConversionResult{}, errors.New("failed to extract conversion rate from API response")
	}

	convertedAmount := request.Amount * rate.Price

	// Create and return the conversion result
	return entity.NewConversionResult(
		request.FromCurrency,
		request.ToCurrency,
		request.Amount,
		convertedAmount,
		rate.Price,
		rate.LastUpdated,
	), nil
}
