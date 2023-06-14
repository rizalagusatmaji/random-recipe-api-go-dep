package apis

import (
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetCurrencyExchange() (float64, error) {
	url := "https://currency-exchange.p.rapidapi.com/exchange?from=USD&to=IDR&q=1.0"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return 0.0, nil
	}
	req.Header.Add("X-RapidAPI-Host", "currency-exchange.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", "e9b079b792msh3e1aa8608df520dp15d272jsn0831149d10d6")

	res, err := client.Do(req)
	if err != nil {
		return 0.0, nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0.0, err
	}

	priceStr := string(body)
	priceFloat, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return 0, err
	}

	return priceFloat, nil
}
