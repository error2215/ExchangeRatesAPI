package ExchangeRatesAPI

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func (a *ExchangeRatesAPI) Get(debug bool) (string, error) {
	if debug {
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
