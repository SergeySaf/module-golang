package brackets

import "strings"

func clear_brackets(str string) string {
	buff := "" 
	for _, value := range str {
		if (value == '}') || (value == '{') || (value == ')') || (value == '(') || (value == ']') || (value == '[') {
			buff += string(value)
		} else {
			continue
		}
	}
	return string(buff)
}

func Bracket(new_str string) (bool, error) {
	str := clear_brackets(new_str)
	flag := true
	brace1 := "()"
	brace2 := "{}"
	brace3 := "[]"
	for flag {
		if strings.Contains(str, brace1) || strings.Contains(str, brace2) || strings.Contains(str, brace3) {
			str = strings.Replace(str, brace1, "", -1)
			str = strings.Replace(str, brace2, "", -1)
			str = strings.Replace(str, brace3, "", -1)
		} else if str == "" {
			return true, nil
		} else {
			flag = false
		}
	}
	return flag, nil
}

