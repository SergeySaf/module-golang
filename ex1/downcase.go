package downcase


func Downcase(input string) (string, error) {
	res := ""
	alph := map[string]string {
		"A": "a", "B": "b", "C": "c", "D": "d", "E": "e", "F": "f", "G": "g", "H": "h",
		"I": "i", "J": "j", "K": "k", "L": "l", "M": "m", "N": "n", "O": "o", "P": "p",
		"Q": "q", "R": "r", "S": "s", "T": "t", "U": "u", "V": "v", "W": "w","X": "x",
		"Y": "y", "Z": "z",
	}
	for i := 0; i < len(input); i++ {
		value, ok := alph[string(input[i])]
		if ok {
			res += value		
		} else {
			res += string(input[i])
		}
	}
	return res, nil 
}
package downcase
