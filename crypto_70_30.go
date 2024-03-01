package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type responseInfo struct {
	Data Data //`json:"data"`
}

type Data struct {
	Currency string
	Rates    map[string]string
}

var result responseInfo

func main() {

	amount, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("value: " + os.Args[1] + " must be a number (type float)")
		os.Exit(1)
	}

	cryptoOne := os.Args[2]
	cryptoTwo := os.Args[3]

	statusCode := getCyrptoExchangeRates()
	if statusCode != http.StatusOK {
		fmt.Println("Request failed with code: " + strconv.FormatInt(int64(statusCode), 64))
	}

	crypto70, err := calculateCryptoInfo(amount, cryptoOne, 0.7)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	crypto30, err := calculateCryptoInfo(amount, cryptoTwo, 0.3)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(crypto70)
	fmt.Println(crypto30)

}

func getCyrptoExchangeRates() int {
	response, err := http.Get("https://api.coinbase.com/v2/exchange-rates?currency=USD")
	if err != nil {
		fmt.Println("No response from request")
		os.Exit(1)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	return response.StatusCode
}

func calculateCryptoInfo(amount float64, cryptoName string, percent float64) (string, error) {
	dollarAmount := amount * percent
	cryptoValue, err := strconv.ParseFloat(result.Data.Rates[cryptoName], 64)
	if err != nil {
		return "", errors.New("Crypto name: " + cryptoName + " must be a valued crypto type.")
	}

	return "$" + strconv.FormatFloat(dollarAmount, 'f', -1, 64) + " => " + strconv.FormatFloat((dollarAmount*cryptoValue), 'f', -1, 64) + " " + cryptoName, nil

}
