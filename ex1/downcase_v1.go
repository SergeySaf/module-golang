package downcase

func isUpperAlpha(ch byte) bool {
	return ch >= 'A' && ch <= 'Z'
}

func toDown(ch byte) byte {
	return ch + 32
}

func Downcase(input string) (string, error) {
	res := ""

	for _, ch := range []byte(input) {
		if isUpperAlpha(ch) {
			res += string(toDown(ch))
		} else {
			res += string(ch)
		}
	}

	return res, nil 
}

