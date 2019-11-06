package ExchangeRatesAPI

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func (a *ExchangeRatesAPI) Get() map[string]interface{} {
	res := make(map[string]interface{})

	resp, err := http.Get(a.apiURL + "/latest")
	if err != nil {
		logrus.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error(err)
	}

	logrus.Info(string(body))
	return res
}
