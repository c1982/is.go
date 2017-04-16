package Is

const (
	zipCodeRegex  = `^[0-9]{5}(?:-[0-9]{4})?$`
	postCodeRegex = `^[A-Z]{1,2}[0-9RCHNQ][0-9A-Z]?\s?[0-9][ABD-HJLNP-UW-Z]{2}$|^[A-Z]{2}-?[0-9]{4}$`
	postBoxRegex  = `(?i)P\.? ?O\.? Box \d+`
	phoneRegex    = `(?:(?:\+?\d{1,3}[-.\s*]?)?(?:\(?\d{3}\)?[-.\s*]?)?\d{3}[-.\s*]?\d{4,6})|(?:(?:(?:\(\+?\d{2}\))|(?:\+?\d{2}))\s*\d{2}\s*\d{3}\s*\d{4})`
)

func UsZipCode(input string) bool {
	return isMatch(input, zipCodeRegex)
}

func UkPostCode(input string) bool {
	return isMatch(input, postCodeRegex)
}

func PostBox(input string) bool {
	return isMatch(input, postBoxRegex)
}

func Phone(input string) bool {
	return isMatch(input, phoneRegex)
}

