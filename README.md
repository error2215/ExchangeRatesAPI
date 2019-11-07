# exchange-rates-api - Currency Exchange Rates API SDK

This is an unofficial wrapper for the awesome, free [ExchangeRatesAPI](https://exchangeratesapi.io/), which provides exchange rate lookups courtesy of the [Central European Bank](https://www.ecb.europa.eu/stats/policy_and_exchange_rates/euro_reference_exchange_rates/html/index.en.html).

### 1. Installation:

Use go get to retrieve the SDK to add it to your GOPATH workspace, or project's Go module dependencies.

`go get github.com/error2215/exchange-rates-api`

### 2. Getting Started:

Since the CurrencyExchangeAPI does not require API keys or authentication in order to access and interrogate its API, getting started with this library is easy. The following examples show how to achieve various functions using the library.


**Basic usage:**<br />
Fetch the latest exchange rates from the European Central Bank:

```
import (
  "github.com/sirupsen/logrus"
	api "github.com/error2215/exchange-rates-api"
)

func main() {
	json, err := api.New().Get()
	if err != nil{
		logrus.Error(err)
	}else{
		logrus.Info(json)
	}
}
```

**Historical data:**<br />
Get historical rates for any day since 1999:

```
json, err := api.New().AddDateFrom("2015-01-20").Get()
	if err != nil {
		logrus.Error(err)
	} else {
		logrus.Info(json)
	}
```

**Set the base currency:**<br />
By default, the base currency is set to Euro (EUR), but it can be changed: 

```
json, err := api.New(true).SetBaseCurrency("USD").Get()
	if err != nil {
		logrus.Error(err)
	} else {
		logrus.Info(json)
	}
```

**Fetch specific rates:**<br />
If you do not want all current rates, it's possible to specify only the currencies you want using `AddRate()`. The following code fetches only the exchange rate between USD and EUR:

```
json, err := api.New(true).AddRate("USD").Get()
	if err != nil {
		logrus.Error(err)
	} else {
		logrus.Info(json)
	}
```

Please refer to the [API website](https://exchangeratesapi.io/) for further information and full API docs.

### 3. API Reference:

The following API reference lists the publicly-available methods for the 

#### `exchange-rates-api` Reference:

`AddDateFrom( from string )`:<br />
Set the date from which to retrieve historic rates. `from` should be an ISO 8601 date.

`GetDateFrom()`:<br />
Returns the specified date from which to retrieve historic rates. Returns `""` if none is specified.

`RemoveDateFrom()`:<br />
Removes the specified start date for the retrieval of historic rates.

`AddDateTo( to string)`:<br />
Set the end date for the retrieval of historic rates. `to` should be an ISO 8601 date.

`GetDateTo()`:<br />
Returns the specified end date for the retrieval of historic rates. Returns `""` if none is specified.

`RemoveDateTo()`:<br />
Removes the specified end date for the retrieval of historic rates.

`ValidateCurrency( currency string )`:<br />
Checks if a specific currency code is supported. `currency` should be passed as an ISO 4217 code (e.g. `EUR`).<br />
Returns nil if supported, or error if not.

`ValidateDateFormat( date string )`:<br />
Checks if a specific date is supported. `date` should be passed as an ISO 8601 notation (e.g. YYYY-MM-DD).<br />
Returns nil if supported, or error if not.

`SetBaseCurrency( currency string)`:<br />
Set the base currency to be used for exchange rates. `currency` should be passed an ISO 4217 code (e.g. `EUR`).<br />
`currency` must be one of the [supported currency codes](#5-supported-currencies).

`GetBaseCurrency()`:<br />
Returns the currently specified base currency. If `SetBaseCurrency()` hasn't been called, this will return the default base currency `EUR`.

`AddRate( currency string )`:<br />
Adds a new currency to be retrieved. `currency` should be passed an ISO 4217 code (e.g. `EUR`).<br />
`currency` must be one of the [supported currency codes](#5-supported-currencies).<br />
If no rates are added, **all** rates will be returned.

`RemoveRate( currency string )`:<br />
Removes a currency that has already been added to the retrieval list.  `currency` should be passed an ISO 4217 code (e.g. `EUR`).<br />
`currency` must be one of the [supported currency codes](#5-supported-currencies).

`Get()`:<br />
Send off the request to the API and return a raw JSON response. 

`GetSupportedCurrencies()`:<br />
Returns a slice of supported currency codes.

`GetSupportedCurrenciesInString(string delimiter)`:<br />
Returns a string of supported currency codes separated by delimiter.

`GetRates()`:<br />
Returns a slice of the currently specified rates to retrieve.

`GetRatesInString(string delimiter)`:<br />
Returns a string of currently specified rates to retrieve separated by delimiter.

`GetBaseCurrency()`:<br />
Returns the base currency of the request. If not base currency was specified using `SetBaseCurrency` on the request, this will return the default (`EUR`).


### 4. Supported Currencies:

The library supports any currency currently available on the European Central Bank's web service, which at the time of the latest release are as follows:

![](https://www.ecb.europa.eu/shared/img/flags/AUD.gif) Australian Dollar (AUD)<br />
![](https://www.ecb.europa.eu/shared/img/flags/BRL.gif) Brazilian Real (BRL)<br />
![](https://www.ecb.europa.eu/shared/img/flags/GBP.gif) British Pound Sterline (GBP)<br />
![](https://www.ecb.europa.eu/shared/img/flags/BGN.gif) Bulgarian Lev (BGN)<br />
![](https://www.ecb.europa.eu/shared/img/flags/CAD.gif) Canadian Dollar (CAD)<br />
![](https://www.ecb.europa.eu/shared/img/flags/CNY.gif) Chinese Yuan Renminbi (CNY)<br />
![](https://www.ecb.europa.eu/shared/img/flags/HRK.gif) Croatian Kuna (HRK)<br />
![](https://www.ecb.europa.eu/shared/img/flags/CZK.gif) Czech Koruna (CZK)<br />
![](https://www.ecb.europa.eu/shared/img/flags/DKK.gif) Danish Krone (DKK)<br />
![](https://www.ecb.europa.eu/shared/img/flags/EUR.gif) Euro (EUR)<br />
![](https://www.ecb.europa.eu/shared/img/flags/HKD.gif) Hong Kong Dollar (HKD)<br />
![](https://www.ecb.europa.eu/shared/img/flags/HUF.gif) Hungarian Forint (HUF)<br />
![](https://www.ecb.europa.eu/shared/img/flags/ISK.gif) Icelandic Krona (ISK)<br />
![](https://www.ecb.europa.eu/shared/img/flags/IDR.gif) Indonesian Rupiah (IDR)<br />
![](https://www.ecb.europa.eu/shared/img/flags/INR.gif) Indian Rupee (INR)<br />
![](https://www.ecb.europa.eu/shared/img/flags/ILS.gif) Israeli Shekel (ILS)<br />
![](https://www.ecb.europa.eu/shared/img/flags/JPY.gif) Japanese Yen (JPY)<br />
![](https://www.ecb.europa.eu/shared/img/flags/MYR.gif) Malaysian Ringgit (MYR)<br />
![](https://www.ecb.europa.eu/shared/img/flags/MXN.gif) Mexican Peso (MXN)<br />
![](https://www.ecb.europa.eu/shared/img/flags/NZD.gif) New Zealand Dollar (NZD)<br />
![](https://www.ecb.europa.eu/shared/img/flags/NOK.gif) Norwegian Krone (NOK)<br />
![](https://www.ecb.europa.eu/shared/img/flags/PHP.gif) Philippine Peso (PHP)<br />
![](https://www.ecb.europa.eu/shared/img/flags/PLN.gif) Polish Zloty (PLN)<br />
![](https://www.ecb.europa.eu/shared/img/flags/RON.gif) Romanian Leu (RON)<br />
![](https://www.ecb.europa.eu/shared/img/flags/RUB.gif) Russian Rouble (RUB)<br />
![](https://www.ecb.europa.eu/shared/img/flags/SGD.gif) Singapore Dollar (SGD)<br />
![](https://www.ecb.europa.eu/shared/img/flags/ZAR.gif) South African Rand (ZAR)<br />
![](https://www.ecb.europa.eu/shared/img/flags/KRW.gif) South Korean Won (KRW)<br />
![](https://www.ecb.europa.eu/shared/img/flags/SEK.gif) Swedish Krona (SEK)<br />
![](https://www.ecb.europa.eu/shared/img/flags/CHF.gif) Swiss Franc (CHF)<br />
![](https://www.ecb.europa.eu/shared/img/flags/THB.gif) Thai Baht (THB)<br />
![](https://www.ecb.europa.eu/shared/img/flags/TRY.gif) Turkish Lira (TRY)<br />
![](https://www.ecb.europa.eu/shared/img/flags/USD.gif) US Dollar (USD)<br />


### 5. Bugs & Features:

If you have spotted any bugs, or would like to request additional features from the library, please file an issue via the Issue Tracker on the project's Github page: [https://github.com/error2215/exchange-rates-api/issues](https://github.com/error2215/exchange-rates-api/issues).
