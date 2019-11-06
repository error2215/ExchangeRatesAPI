package ExchangeRatesAPI

import (
	"io/ioutil"
	"net/http"
)

func (a *ExchangeRatesAPI) Get() (string, error) {
	resp, err := http.Get(a.apiURL + a.buildQuery())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
