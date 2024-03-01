package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type responseInfo struct {
	Data struct {
		Currency string `json:"currency"`
		Rates    map[string]string
	} `json:"data"`
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

	response, err := http.Get("https://api.coinbase.com/v2/exchange-rates?currency=USD")
	if err != nil {
		fmt.Println("No response from request")
		os.Exit(1)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body) // response body is []byte

	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	displayCryptoInfo(amount, cryptoOne, 0.7)
	displayCryptoInfo(amount, cryptoTwo, 0.3)

}

func displayCryptoInfo(amount float64, cryptoName string, percent float64) {
	dollarAmount := amount * percent
	cryptoValue, err := strconv.ParseFloat(result.Data.Rates[cryptoName], 64)
	if err != nil {
		fmt.Println("Crypto name: " + cryptoName + " must be a valued crypto type.")
		os.Exit(1)
	}
	fmt.Print("$")
	fmt.Print(dollarAmount)
	fmt.Print(" => ")
	fmt.Print((dollarAmount) * cryptoValue)
	fmt.Println(" " + cryptoName)

}
