package cipher


func Downcase(input string) string {
	size := len(input)
	count := 0
	for i := 0; i < size; i++ {
		if (input[i] >= 65 && input[i] <= 90) || (input[i] >= 97 && input[i] <= 122) {
			count++
		}
	}
	b := make([]byte, count)
	count = 0
	for i := 0; i < size; i++ {
		char := input[i]
		if (input[i] >= 65 && input[i] <= 90) || (input[i] >= 97 && input[i] <= 122) {
			if char >= 'A' && char <= 'Z' {
				char += 32
			}
			b[count] = char
			count++
		}
	}
	return string(b)
}

func key_is_valid (key string) bool {
	flag := false
	if len(key) < 1 {
		return false
	}
	for _, value := range key {
		if value != 'a' {
			flag = true
		}
		if value < 'a' || value > 'z' {
			return false
		}
	}
	if flag {
		return true
	}
	return false
}

func endless_str(input string,size int) string {
	buff := make([]byte, size)
	j := 0
	for i := range buff {
		if j == len(input) {
			j = 0
		}
		buff[i] = input[j]
		j++
	}
	return string(buff)
}

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type MyCaesar struct{}

func (c MyCaesar) Encode(str string) string {
	new_str := Downcase(str)
	buff := make([]byte, len(new_str))
	for i := range buff {
		char := new_str[i]
		if char >= 'x' {
			char -= 23
		} else {
			char += 3
		}
		buff[i] = char
	}
	return string(buff)
}

func (c MyCaesar) Decode(str string) string {
	buff := make([]byte, len(str))
	for i := range buff {
		char := str[i]
		if char < byte(100) {
			char += byte(23)
		} else {
			char -= byte(3)
		}
		buff[i] = char
	}
	return string(buff)
}

func NewCaesar() Cipher {
	return MyCaesar{}
}

type MyShift struct {
	shift int
}

func (c MyShift) Encode(str string) string {
	new_str := Downcase(str)
	buff := make([]byte, len(new_str))
	for i := range buff {
		char := new_str[i] + byte(c.shift)
		if char > 122 {
			char -= 26
		} else if char < 97 {
			char += 26
		}
		buff[i] = char
	}
	return string(buff)
}

func (c MyShift) Decode(str string) string {
	new_str := Downcase(str)
	buff := make([]byte, len(new_str))
	for i := range buff {
		char := new_str[i] - byte(c.shift)
		if char > 122 {
			char -= 26
		} else if char < 97 {
			char += 26
		}
		buff[i] = char
	}
	return string(buff)
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

func (c MyVigenere) Encode(str string) string {
	new_str := Downcase(str)
	buff := make([]byte, len(new_str))
	new_key := endless_str(c.key, len(new_str))
	for i := range(buff) {
		shift := new_key[i] - 97
		if shift != 0 {
			char := new_str[i] + shift
			if char > 122 {
				char -=26
			} else if char < 97 {
				char += 26
			} 
			buff[i] = char
		} else {
			buff[i] = new_str[i] 
		}
	}	
	return string(buff)
}

func (c MyVigenere) Decode(str string) string {
	new_str := Downcase(str)
	buff := make([]byte, len(str))
	new_key := endless_str(c.key, len(str))
	for i := range(buff) {
		shift := new_key[i] - 97
		if shift != 0 {
			char := new_str[i] - shift
			if char > 122 {
				char -=26
			} else if char < 97 {
				char += 26
			}
			buff[i] = char
		} else {
			buff[i] = new_str[i]
		}
	}
	return string(buff)
}

func NewVigenere(key string) Cipher {
	if  key_is_valid(key) {
		return MyVigenere{key}
	}
	return nil
}
