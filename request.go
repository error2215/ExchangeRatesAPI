package ExchangeRatesAPI

import (
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Make API request with set parameters
func (a *exchangeRatesAPI) Get() (string, error) {
	if a.debug {
		logrus.Info("Query:" + a.buildQuery())
	}
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
