package ExchangeRatesAPI

import (
	"errors"
	"github.com/sirupsen/logrus"
	"net/url"
	"regexp"
	"strings"
)

type ExchangeRatesAPI struct {
	dateFrom       string
	dateTo         string
	baseCurrency   string
	symbols        []string
	apiURL         string
	dateRegExp     string
	currencyRegExp string
}

var supportedCurrencies = []string{
	"USD", "GBP", "EUR", "JPY", "BGN", "CZK", "DKK", "HUF", "PLN", "RON",
	"SEK", "CHF", "ISK", "NOK", "HRK", "RUB", "TRY", "AUD", "BRL", "CAD",
	"CNY", "HKD", "IDR", "ILS", "INR", "KRW", "MXN", "MYR", "NZD", "BHP",
	"SGD", "THB", "ZAR",
}

func New() *ExchangeRatesAPI {
	return &ExchangeRatesAPI{
		dateFrom:     "",
		dateTo:       "",
		baseCurrency: "EUR",
		symbols:      []string{},
		apiURL:       "https://api.exchangeratesapi.io",
	}
}

func (a *ExchangeRatesAPI) GetDateFrom() string {
	return a.dateFrom
}

func (a *ExchangeRatesAPI) GetDateTo() string {
	return a.dateTo
}

func (a *ExchangeRatesAPI) GetSupportedCurrencies() []string {
	return supportedCurrencies
}

func (a *ExchangeRatesAPI) GetSupportedCurrenciesInString(delimiter string) string {
	return strings.Join(supportedCurrencies, delimiter)
}

func (a *ExchangeRatesAPI) GetSymbols() []string {
	return a.symbols
}

func (a *ExchangeRatesAPI) GetSymbolsInString(delimiter string) string {
	return strings.Join(a.symbols, delimiter)
}

func (a *ExchangeRatesAPI) GetBaseCurrency() string {
	return a.baseCurrency
}

func (a *ExchangeRatesAPI) AddDateFrom(from string) *ExchangeRatesAPI {
	if a.validateDateFormat(from) == nil {
		a.dateFrom = from
	} else {
		logrus.Error(a.validateDateFormat(from))
	}
	return a
}

func (a *ExchangeRatesAPI) RemoveDateFrom() *ExchangeRatesAPI {
	a.dateFrom = ""
	return a
}

func (a *ExchangeRatesAPI) AddDateTo(to string) *ExchangeRatesAPI {
	if a.validateDateFormat(to) == nil {
		a.dateTo = to
	} else {
		logrus.Error(a.validateDateFormat(to))
	}
	return a
}

func (a *ExchangeRatesAPI) RemoveDateTo() *ExchangeRatesAPI {
	a.dateTo = ""
	return a
}

func (a *ExchangeRatesAPI) SetBaseCurrency(currency string) *ExchangeRatesAPI {
	if a.validateCurrency(currency) == nil {
		a.baseCurrency = currency
	} else {
		logrus.Error(a.validateCurrency(currency))
	}
	return a
}

func (a *ExchangeRatesAPI) AddRate(currency string) *ExchangeRatesAPI {
	if a.validateCurrency(currency) == nil {
		a.symbols = append(a.symbols, currency)
	} else {
		logrus.Error(a.validateCurrency(currency))
	}
	return a
}

func (a *ExchangeRatesAPI) RemoveRate(currency string) *ExchangeRatesAPI {
	for num, rate := range a.symbols {
		if rate == currency {
			a.symbols = append(a.symbols[:num], a.symbols[num+1:]...)
			break
		}
	}
	return a
}

func (a *ExchangeRatesAPI) buildQuery() string {
	query := ""
	values := url.Values{}
	if a.dateFrom != "" && a.dateTo != "" {
		query = "/history"
		values.Set("start_at", a.GetDateFrom())
		if a.GetDateTo() != "" {
			values.Set("end_at", a.GetDateTo())
		}
		if len(a.GetSymbols()) > 0 {
			values.Set("symbols", a.GetSymbolsInString(","))
		}
		if a.GetBaseCurrency() != "EUR" {
			values.Set("base", a.GetBaseCurrency())
		}
	} else if a.dateFrom != "" {
		query = "/" + a.dateFrom
		if len(a.GetSymbols()) > 0 {
			values.Set("symbols", a.GetSymbolsInString(","))
		}
		if a.GetBaseCurrency() != "EUR" {
			values.Set("base", a.GetBaseCurrency())
		}
		return query + values.Encode()
	} else {
		query = "/latest"
		if len(a.GetSymbols()) > 0 {
			values.Set("symbols", a.GetSymbolsInString(","))
		}
		if a.GetBaseCurrency() != "EUR" {
			values.Set("base", a.GetBaseCurrency())
		}
		return query + values.Encode()
	}
	return "/latest"
}

func (a *ExchangeRatesAPI) validateDateFormat(date string) error {
	re := regexp.MustCompile("((19|20)\\d\\d)-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])")
	if !re.MatchString(date) {
		return errors.New("The specified date is invalid. Please use ISO 8601 notation (e.g. YYYY-MM-DD) ")
	}
	logrus.Info("validated")
	return nil
}

func (a *ExchangeRatesAPI) validateCurrency(currency string) error {
	for _, supported := range supportedCurrencies {
		if supported == currency {
			return nil
		}
	}
	return errors.New("The specified currency code is not currently supported ")
}
