package Is

const (
	ccRegex      = `(?:(?:(?:\d{4}[- ]?){3}\d{4}|\d{15,16}))`
	visaRegex    = `4\d{3}[\s-]?\d{4}[\s-]?\d{4}[\s-]?\d{4}`
	masterRegex  = `5[1-5]\d{2}[\s-]?\d{4}[\s-]?\d{4}[\s-]?\d{4}`
	bitCoinRegex = `[13][a-km-zA-HJ-NP-Z1-9]{25,34}`
	ibanRegex    = `[A-Z]{2}\d{2}[A-Z0-9]{4}\d{7}([A-Z\d]?){0,16}`
)

func CreditCard(input string) bool {
	return isMatch(input, ccRegex)
}

func Visa(input string) bool {
	return isMatch(input, visaRegex)
}

func Master(input string) bool {
	return isMatch(input, masterRegex)
}

func BitCoin(input string) bool {
	return isMatch(input, bitCoinRegex)
}

func IBAN(input string) bool {
	return isMatch(input, ibanRegex)
}

