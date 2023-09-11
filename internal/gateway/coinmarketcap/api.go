package coinmarketcap

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/qjawko/currency-convertor/internal/usecase/conversion"
)

const baseURL = "https://pro-api.coinmarketcap.com/v2/tools/price-conversion"

type APIClient struct {
	APIKey     string
	httpClient *http.Client
}

// NewAPIClient creates a new CoinMarketCap API client.
func NewAPIClient(apiKey string, client *http.Client) *APIClient {
	if client == nil {
		client = &http.Client{}
	}

	return &APIClient{
		APIKey:     apiKey,
		httpClient: client,
	}
}

func (client *APIClient) GetConversionRate(ctx context.Context, amount float64, fromCurrency, toCurrency string) (conversion.APIData, error) {
	var apiData conversion.APIData

	url := fmt.Sprintf("%s?amount=%f&symbol=%s&convert=%s", baseURL, amount, fromCurrency, toCurrency)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return apiData, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-CMC_PRO_API_KEY", client.APIKey)

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return apiData, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return apiData, fmt.Errorf("API returned non-200 status code: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&apiData)
	if err != nil {
		return apiData, err
	}

	if apiData.Status.ErrorCode != 0 {
		return apiData, errors.New(apiData.Status.ErrorMessage)
	}

	return apiData, nil
}
