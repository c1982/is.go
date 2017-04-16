package Is

const(
	isbn13Regex         = `(?:[\d]-?){12}[\dxX]`
	isbn10Regex         = `(?:[\d]-?){9}[\dxX]`
)

func ISBN13(input string) bool {
	return isMatch(input, isbn13Regex)
}

func ISBN10(input string) bool {
	return isMatch(input, isbn10Regex)
}
