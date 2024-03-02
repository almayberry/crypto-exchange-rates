package main

import (
	"net/http"
	"strings"
	"testing"
)

func Test_getCryptoExchangeRates(t *testing.T) {
	statusCode := getCryptoExchangeRates()

	if statusCode != http.StatusOK {
		t.Errorf("got statusCode %q expected %q", statusCode, http.StatusOK)
	}
}

func Test_calculateCryptoInfo_70HappyPath(t *testing.T) {
	rates := make(map[string]string)
	rates["test"] = "1.2"
	result = responseInfo{Data: Data{Currency: "USD", Rates: rates}}

	got, err := calculateCryptoInfo(100, "test", 0.7)

	if !strings.Contains(got, "$70 => 84 test") {
		t.Errorf("Did not return correct value. Instead got: %q", got)
	}
	if err != nil {
		t.Errorf("Got unexpected error: %q", err)
	}
}

func Test_calculateCryptoInfo_30HappyPath(t *testing.T) {
	rates := make(map[string]string)
	rates["test"] = ".056"
	result = responseInfo{Data: Data{Currency: "USD", Rates: rates}}

	got, err := calculateCryptoInfo(100, "test", 0.3)

	if !strings.Contains(got, "$30 => 1.68 test") {
		t.Errorf("Did not return correct value. Instead got: %q", got)
	}
	if err != nil {
		t.Errorf("Got unexpected error: %q", err)
	}

}

func Test_calculateCryptoInfo_ErrorCryptoName(t *testing.T) {
	rates := make(map[string]string)
	rates["test"] = ".056"
	result = responseInfo{Data: Data{Currency: "USD", Rates: rates}}

	_, err := calculateCryptoInfo(100, "errorName", 0.3)

	if err == nil {
		t.Errorf("Error was not found when it was expected.")
	}

}

func Test_roundDollarAmount(t *testing.T) {
	type addTest struct {
		arg1, expected float64
	}

	var addTests = []addTest{
		{100.5555, 100.56},
		{100.22252, 100.22},
	}

	for _, test := range addTests {
		if output := roundDollarAmount(test.arg1); output != test.expected {
			t.Errorf("Output %f not equal to expected %f", output, test.expected)
		}
	}
}

func Test_truncateDollarAmount(t *testing.T) {
	type addTest struct {
		arg1, expected float64
	}

	var addTests = []addTest{
		{100.5555, 100.55},
		{100.22222, 100.22},
	}

	for _, test := range addTests {
		if output := truncateDollarAmount(test.arg1); output != test.expected {
			t.Errorf("Output %f not equal to expected %f", output, test.expected)
		}
	}
}
