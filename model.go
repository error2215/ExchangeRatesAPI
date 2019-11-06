package ExchangeRatesAPI

type ExchangeRatesAPI struct {
	dateFrom       string
	dateTo         string
	baseCurrency   string
	rates          []string
	apiURL         string
	dateRegExp     string
	currencyRegExp string
}

func New() *ExchangeRatesAPI {
	return &ExchangeRatesAPI{
		dateFrom:       "",
		dateTo:         "",
		baseCurrency:   "EUR",
		rates:          nil,
		apiURL:         "https://api.exchangeratesapi.io/",
		dateRegExp:     "",
		currencyRegExp: "",
	}
}

func (a *ExchangeRatesAPI) BuildBody() map[string]string {
	return map[string]string{}
}
