package cipher

func ByteToString(c []byte) string {
	n := -1
	for i, b := range c {
		if b == 0 {
			break
		}
		n = i
	}
	return string(c[:n+1])
}

func LongKey(key string, size int) string {
	var barr [100]byte
	j := 0
	lenkey := len(key)
	for i := 0; i < size; i++ {
		if j >= lenkey {
			j = 0
		}
		barr[i] = key[j]
		j++
	}
	res := ByteToString(barr[:])
	return res
}

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type MyCaesar struct{}

func (c MyCaesar) Encode(s string) string {
	ex := "azAZ"
	key := "d"
	lens := len(s)
	n := 0
	var barr [100]byte
	for i := 0; i < lens; i++ {
		k := key[0] - ex[0]
		lim := ex[1] - k
		lim2 := ex[3] - k
		if s[i] <= ex[3] && s[i] >= ex[2] {
			if s[i] > lim2 {
				barr[n] = s[i] + k - 26 + 32
			} else {
				barr[n] = s[i] + k + 32
			}
			n++
		}
		if s[i] <= ex[1] && s[i] >= ex[0] {
			if s[i] > lim {
				barr[n] = s[i] + k - 26
			} else {
				barr[n] = s[i] + k
			}
			n++
		}
	}
	res := ByteToString(barr[:])
	return res
}

func (c MyCaesar) Decode(s string) string {
	ex := "aZ"
	lens := len(s)
	key := "d"
	var barr [100]byte
	for i := 0; i < lens; i++ {
		k := key[0] - ex[0]
		lim := ex[0] + k
		if s[i] < lim {
			barr[i] = s[i] - k + 26
		} else {
			barr[i] = s[i] - k
		}
	}
	res := ByteToString(barr[:])
	return res
}

func NewCaesar() Cipher {
	return MyCaesar{}
}

type MyShift struct {
	shift int
}

func (c MyShift) Encode(s string) string {
	ex := "azAZ"
	k := byte(0)
	lens := len(s)
	n := 0
	if c.shift < 0 {
		k = byte(c.shift + 26)
	} else {
		k = byte(c.shift)
	}
	var barr [100]byte
	for i := 0; i < lens; i++ {
		lim := ex[1] - k
		lim2 := ex[3] - k
		if s[i] <= ex[3] && s[i] >= ex[2] {
			if s[i] > lim2 {
				barr[n] = s[i] + k - 26 + 32
			} else {
				barr[n] = s[i] + k + 32
			}
			n++
		}
		if s[i] <= ex[1] && s[i] >= ex[0] {
			if s[i] > lim {
				barr[n] = s[i] + k - 26
			} else {
				barr[n] = s[i] + k
			}
			n++
		}
	}
	res := ByteToString(barr[:])
	return res
}

func (c MyShift) Decode(s string) string {
	ex := "aZ"
	lens := len(s)
	k := byte(0)
	if c.shift < 0 {
		k = byte(c.shift + 26)
	} else {
		k = byte(c.shift)
	}
	var barr [100]byte
	for i := 0; i < lens; i++ {
		lim := ex[0] + k
		if s[i] < lim {
			barr[i] = s[i] - k + 26
		} else {
			barr[i] = s[i] - k
		}
	}
	res := ByteToString(barr[:])
	return res
}

func NewShift(shift int) Cipher {
	if shift >= 1 && shift <= 25 || shift <= -1 && shift >= -25 {
		return MyShift{shift}
	}
	return nil
}

type MyVigenere struct {
	key string
}

func (c MyVigenere) Encode(s string) string {
	ex := "azAZ"
	lens := len(s)
	key := LongKey(c.key, lens)
	n := 0
	var barr [100]byte
	for i := 0; i < lens; i++ {
		k := key[n] - ex[0]
		lim := ex[1] - k
		lim2 := ex[3] - k
		if s[i] <= ex[3] && s[i] >= ex[2] {
			if s[i] > lim2 {
				barr[n] = s[i] + k - 26 + 32
			} else {
				barr[n] = s[i] + k + 32
			}
			n++
		}
		if s[i] <= ex[1] && s[i] >= ex[0] {
			if s[i] > lim {
				barr[n] = s[i] + k - 26
			} else {
				barr[n] = s[i] + k
			}
			n++
		}
	}
	res := ByteToString(barr[:])
	return res
}

func (c MyVigenere) Decode(s string) string {
	ex := "aZ"
	lens := len(s)
	key := LongKey(c.key, lens)
	var barr [100]byte
	for i := 0; i < lens; i++ {
		k := key[i] - ex[0]
		lim := ex[0] + k
		if s[i] < lim {
			barr[i] = s[i] - k + 26
		} else {
			barr[i] = s[i] - k
		}
	}
	res := ByteToString(barr[:])
	return res
}

func Validat(key string) bool {
	ex := "az"
	lenkey := len(key)
	if lenkey < 1 {
		return false
	}
	valid := 0
	for i := 0; i < lenkey; i++ {
		if key[i] > ex[1] || key[i] < ex[0] {
			return false
		} else if key[i] != ex[0] {
			valid = 1
		}
	}
	if valid == 1 {
		return true
	} else {
		return false
	}
}

func NewVigenere(key string) Cipher {
	if Validat(key) {
		return MyVigenere{key}
	}
	return nil
}
