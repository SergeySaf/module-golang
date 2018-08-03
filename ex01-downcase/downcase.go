package downcase

func Downcase(input string) (string, error) {
	res := ""

	for _, letter := range input {
		if letter >= 'A' && letter <= 'Z' {
			res += string(letter + 32)
		} else {
			res += string(letter)
		}
	}

	return res, nil
}
