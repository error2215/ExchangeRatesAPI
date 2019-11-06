package ExchangeRatesAPI

import (
	"io/ioutil"
	"net/http"
)

func (a *ExchangeRatesAPI) Get() ([]byte, error) {
	resp, err := http.Get(a.apiURL + a.buildQuery())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
