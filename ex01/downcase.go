package downcase

func ByteToString(c []byte) string {
	var n int
	for i, b := range c {
		if b == 0 {
			break
		}
		n = i
	}
	return string(c[:n+1])
}

func Downcase(s string) (string, error) {
	up := "QWERTYUIOPASDFGHJKLZXCVBNM"
	lenn := len(s)
	var barr [100]byte
	for i := 0; i < lenn; i++ {
		flag := 0
		for k := 0; k < 26; k++ {
			if s[i] == up[k] {
				barr[i] = s[i] + 32
				flag = 1
			}
			if flag == 0 {
				barr[i] = s[i]
			}
		}
	}
	res := ByteToString(barr[:])
	return res, nil
}
