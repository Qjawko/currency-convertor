# Currency Convertor CLI

Currency Convertor is a command-line interface (CLI) tool that allows users to convert between different currencies using the latest exchange rates provided by the CoinMarketCap API.

## Features

- Convert between different currencies.
- Utilizes real-time exchange rates from the CoinMarketCap API.
- Graceful shutdown to handle interruptions.

## Prerequisites

- Go (at least version 1.19)
- A valid CoinMarketCap API key

## Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/qjawko/currency-convertor.git
   ```

2. Navigate to the project directory:
   ```bash
   cd currency-convertor
   ```

3. Build the project:
   ```bash
   make
   ```

4. The binary will be available in `./build/bin`.

## Configuration

Before you can use the Currency Convertor CLI, you need to provide your CoinMarketCap API key.

You can configure it using an environment variable:

- Environment variable:
  ```bash
  export CURRENCY_CONVERTOR_API_KEY=YOUR_COINMARKETCAP_API_KEY
  ```

## Usage

```bash
./currency-convertor <amount> <from_currency> <to_currency>
```

- `<amount>`: The amount of currency you want to convert.
- `<from_currency>`: The currency code of the currency you want to convert from.
- `<to_currency>`: The currency code of the currency you want to convert to.

For help:
```bash
./currency-convertor help
```

## Examples

To convert 100 US Dollars to Euros:
```bash
./currency-convertor 100 USD EUR
```

To convert 1500 Japanese Yen to US Dollars:
```bash
./currency-convertor 1500 JPY USD
```

## Contributing

If you'd like to contribute, please fork the repository and make changes as you'd like. Pull requests are warmly welcome.

## License

This project is open-source and available under the [MIT License](LICENSE).

## Acknowledgments

- Thanks to CoinMarketCap for providing the currency conversion API.