# Run Crypto Exchange Rate
In command line under crypto-exchange-rates run command `go run crypto_70_30.go [dollar amount] [crypto name one] [crypto name two]`

This will out put how much crypto type one you can get with 70% of the dollar amount you input and how much crypto name two you can get with the reamining 30%
#### Example:
$ go run crypto_70_30.go 100 BTC 1INCH

$70 => 0.001125888758491 BTC

$30 => 55.04587155963302 1INCH

# Run tests

# Assignment Prompt

This endpoint provides up-to-the-minute crypto exchange rates relative
to US dollars: https://api.coinbase.com/v2/exchange-rates?currency=USD

That is: each rate is how much of that crypto currency you would get
for 1 dollar. So if you received a value for 0.091 for BTC, that's
saying it's 0.091 per 1 USD.

Your Task:

You are to make a cli that takes in a USD amount as holdings, and
calculates the 70/30 split for 2 given crypto currencies. Stated
simply: I have $X I want to keep in BTC and ETH, 70/30 split. How many
of each should I buy? An example usage would look like:

binary_name 100 BTC ETH

This output should be in the following format

$70.00 => 0.0025 BTC
$30.00 => 0.0160 ETH

This output tells us: Of our 100$ holdings, 70% of that is 70$, which
buys 0.0025 BTC, and 30% of our holdings is 30$, which buys 0.016 ETH.

Make sure to:

Handle possible error cases
Follow Go's formatting (go fmt)
Give best effort: while just a takehome test, give a sample of what
sort of code you like to deliver.

