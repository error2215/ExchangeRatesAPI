package ExchangeRatesAPI

import (
	"errors"
	"net/url"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
)

type exchangeRatesAPI struct {
	dateFrom     string   // Date from which to request historic rates ('start_at' in url query)
	dateTo       string   // Date to which to request historic rates ('end_at' int url query)
	baseCurrency string   // The base currency (default is EUR) ('base' in url query)
	rates        []string // Exchange rates to fetch ('rates' in url query)
	apiURL       string   // The URL of the API ('https://api.exchangeratesapi.io')
	debug        bool     // debug mode (write errors in console if enable)
}

// List of supported currencies
var supportedCurrencies = []string{
	"USD", "GBP", "EUR", "JPY", "BGN", "CZK", "DKK", "HUF", "PLN", "RON",
	"SEK", "CHF", "ISK", "NOK", "HRK", "RUB", "TRY", "AUD", "BRL", "CAD",
	"CNY", "HKD", "IDR", "ILS", "INR", "KRW", "MXN", "MYR", "NZD", "BHP",
	"SGD", "THB", "ZAR",
}

func New(debug bool) *exchangeRatesAPI {
	return &exchangeRatesAPI{
		debug:        debug,
		dateFrom:     "",
		dateTo:       "",
		baseCurrency: "EUR",
		rates:        []string{},
		apiURL:       "https://api.exchangeratesapi.io",
	}
}

// Get the 'from' date
func (a *exchangeRatesAPI) GetDateFrom() string {
	return a.dateFrom
}

// Get the 'to' date
func (a *exchangeRatesAPI) GetDateTo() string {
	return a.dateTo
}

// Get the supported currencies
func (a *exchangeRatesAPI) GetSupportedCurrencies() []string {
	return supportedCurrencies
}

// Get the supported currencies in string with delimiter parameter
func (a *exchangeRatesAPI) GetSupportedCurrenciesInString(delimiter string) string {
	return strings.Join(supportedCurrencies, delimiter)
}

// Get the rates
func (a *exchangeRatesAPI) GetRates() []string {
	return a.rates
}

// Get the rates in string with delimiter parameter
func (a *exchangeRatesAPI) GetRatesInString(delimiter string) string {
	return strings.Join(a.rates, delimiter)
}

// Get the base currency
func (a *exchangeRatesAPI) GetBaseCurrency() string {
	return a.baseCurrency
}

// Add a from date
func (a *exchangeRatesAPI) AddDateFrom(from string) *exchangeRatesAPI {
	if a.ValidateDateFormat(from) == nil {
		a.dateFrom = from
	} else if a.debug {
		logrus.Error(a.ValidateDateFormat(from))
	}
	return a
}

// Remove a from date
func (a *exchangeRatesAPI) RemoveDateFrom() *exchangeRatesAPI {
	a.dateFrom = ""
	return a
}

//Add a to date
func (a *exchangeRatesAPI) AddDateTo(to string) *exchangeRatesAPI {
	if a.ValidateDateFormat(to) == nil {
		a.dateTo = to
	} else if a.debug {
		logrus.Error(a.ValidateDateFormat(to))
	}
	return a
}

//Remove a to date
func (a *exchangeRatesAPI) RemoveDateTo() *exchangeRatesAPI {
	a.dateTo = ""
	return a
}

//Set base currency
func (a *exchangeRatesAPI) SetBaseCurrency(currency string) *exchangeRatesAPI {
	if a.ValidateCurrency(currency) == nil {
		a.baseCurrency = currency
	} else if a.debug {
		logrus.Error(a.ValidateCurrency(currency))
	}
	return a
}

//Add specified currency to the returned rates
func (a *exchangeRatesAPI) AddRate(currency string) *exchangeRatesAPI {
	if a.ValidateCurrency(currency) == nil {
		a.rates = append(a.rates, currency)
	} else if a.debug {
		logrus.Error(a.ValidateCurrency(currency))
	}
	return a
}

//Remove specified currency from returned rates
func (a *exchangeRatesAPI) RemoveRate(currency string) *exchangeRatesAPI {
	for num, rate := range a.rates {
		if rate == currency {
			a.rates = append(a.rates[:num], a.rates[num+1:]...)
			break
		}
	}
	return a
}

// Build a URL query from parameters
func (a *exchangeRatesAPI) buildQuery() string {
	query := ""
	values := url.Values{}
	if a.dateFrom != "" && a.dateTo != "" {
		query = "/history?"
		values.Set("start_at", a.GetDateFrom())
		if a.GetDateTo() != "" {
			values.Set("end_at", a.GetDateTo())
		}
		if len(a.GetRates()) > 0 {
			values.Set("rates", a.GetRatesInString(","))
		}
		if a.GetBaseCurrency() != "EUR" {
			values.Set("base", a.GetBaseCurrency())
		}
		return query + values.Encode()
	} else if a.dateFrom != "" {
		query = "/" + a.dateFrom
		if len(a.GetRates()) > 0 {
			values.Set("rates", a.GetRatesInString(","))
		}
		if a.GetBaseCurrency() != "EUR" {
			values.Set("base", a.GetBaseCurrency())
		}
		return query + values.Encode()
	} else {
		query = "/latest?"
		if len(a.GetRates()) > 0 {
			values.Set("rates", a.GetRatesInString(","))
		}
		if a.GetBaseCurrency() != "EUR" {
			values.Set("base", a.GetBaseCurrency())
		}
		return query + values.Encode()
	}
}

// Validate date format, must be ISO 8601 notation (e.g. YYYY-MM-DD)
func (a *exchangeRatesAPI) ValidateDateFormat(date string) error {
	re := regexp.MustCompile("((19|20)\\d\\d)-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])")
	if !re.MatchString(date) {
		return errors.New("The specified date is invalid. Please use ISO 8601 notation (e.g. YYYY-MM-DD) ")
	}
	return nil
}

// Validate currency code, must be in list of supported currencies and ISO 4217 notation (e.g. EUR).
func (a *exchangeRatesAPI) ValidateCurrency(currency string) error {
	for _, supported := range supportedCurrencies {
		if supported == currency {
			return nil
		}
	}
	return errors.New("The specified currency code is not currently supported or it has bad format, must be ISO 4217 notation (e.g. EUR) ")
}
