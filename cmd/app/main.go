package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/qjawko/currency-convertor/internal/domain"
	"github.com/qjawko/currency-convertor/internal/gateway/coinmarketcap"
	"github.com/qjawko/currency-convertor/internal/usecase"
	"github.com/qjawko/currency-convertor/pkg/config"
)

// Converter provides methods to handle currency conversion.
type Converter interface {
	Convert(ctx context.Context, request domain.ConversionRequest) (domain.ConversionResult, error)
}

type App struct {
	converter Converter
}

func NewApp(converter Converter) *App {
	return &App{
		converter: converter,
	}
}

func (app *App) Run(ctx context.Context, args []string) error {
	if len(args) < 2 {
		app.printHelp()
		return nil
	}

	// Check for the "help" argument
	if args[1] == "help" {
		app.printHelp()
		return nil
	}

	if len(args) < 4 {
		return errors.New("usage: app <amount> <from_currency> <to_currency>")
	}

	amount, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return fmt.Errorf("invalid amount: %w", err)
	}
	fromCurrency := args[2]
	toCurrency := args[3]

	result, err := app.converter.Convert(ctx, domain.ConversionRequest{
		Amount:       amount,
		FromCurrency: fromCurrency,
		ToCurrency:   toCurrency,
	})
	if err != nil {
		return fmt.Errorf("failed to convert currency: %w", err)
	}

	fmt.Printf("%.2f %s is equivalent to %.2f %s\n", amount, fromCurrency, result.ConvertedAmount, toCurrency)
	return nil
}

func (app *App) printHelp() {
	fmt.Println(helpMessage)
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-stop
		cancel()
	}()

	apiClient := coinmarketcap.NewAPIClient(cfg.APIKey, NewHttpClient())
	converter := usecase.NewConversionInteractor(apiClient)

	app := NewApp(converter)

	if err := app.Run(ctx, os.Args); err != nil {
		if errors.Is(err, context.Canceled) {
			fmt.Println("request canceled")
			return
		}

		log.Fatalf("App terminated with error: %v", err)
	}
}

func NewHttpClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 10, // 10-second timeout for requests
	}
}
