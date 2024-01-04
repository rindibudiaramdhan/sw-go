package helper

// Return substring of the string starting with the character number (start) and the length of the returned substring.
func Substr(str string, start int, length int) string {
	return str[start : start+length]
}
