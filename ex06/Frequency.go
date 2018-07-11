package letter

type Map_C map[rune]int

func Frequency(str string) Map_C {
	freqC := Map_C{}
	for i := 0; i < len(str); i++ {
		if freqC[rune(str[i])] >= 0 {
			freqC[rune(str[i])]++
		}
	}
	return freqC
}

func ConcurrentFrequency(astr []string) Map_C {
	freqC := Map_C{}
	freq := make(chan Map_C, len(astr))
	for _, i := range astr {
		go func(str string) {
			freq <- Frequency(str)
		}(i)
	}
	for range astr {
		for char, n := range <-freq {
			freqC[char] += n
		}
	}
	return freqC
}
