package Is

const (
	md5Regex  = `[0-9a-fA-F]{32}`
	sha1Regex = `[0-9a-fA-F]{40}`
	sha2Regex = `[0-9a-fA-F]{64}`
	guidRegex = `[0-9a-fA-F]{8}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{4}-?[a-fA-F0-9]{12}`
)

func MD5(input string) bool {
	return isMatch(input, isbn13Regex)
}

func SHA1(input string) bool {
	return isMatch(input, sha1Regex)
}

func SHA2(input string) bool {
	return isMatch(input, sha2Regex)
}

func GUID(input string) bool {
	return isMatch(input, guidRegex)
}