package brackets

func Bracket(str string) (bool, error) {
	openn := "[{("
	closse := "]})"
	var st [100]byte
	n := 0
	lens := len(str)
	if lens == 0 {
		return true, nil
	}
	for i := 0; i < lens; i++ {
		for k := 0; k < 3; k++ {
			if str[i] == openn[k] {
				st[n] = str[i]
				n++
			} else if str[i] == closse[k] {
				if (n - 1) < 0 {
					return false, nil
				}
				if st[n-1] == openn[k] {
					n--
				} else {
					return false, nil
				}
			}
		}
	}
	if n == 0 {
		return true, nil
	}
	return false, nil
}
