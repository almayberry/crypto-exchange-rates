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

func main() {

	amount, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("value must be a number (type float)")
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

	var result responseInfo
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	cryptoOneValue, err := strconv.ParseFloat(result.Data.Rates[cryptoOne], 64)
	if err != nil {
		fmt.Println("Crypto One must be a valued crypto type.")
		os.Exit(1)
	}

	cryptoTwoValue, err := strconv.ParseFloat(result.Data.Rates[cryptoTwo], 64)
	if err != nil {
		fmt.Println("Crypto Two must be a valued crypto type.")
		os.Exit(1)
	}

	fmt.Print("$")
	fmt.Print(amount * 0.7)
	fmt.Print(" => ")
	fmt.Print((amount * 0.7) * cryptoOneValue)
	fmt.Println(" " + cryptoOne)

	fmt.Print("$")
	fmt.Print(amount * 0.3)
	fmt.Print(" => ")
	fmt.Print((amount * 0.3) * cryptoTwoValue)
	fmt.Print(" " + cryptoTwo)

}
