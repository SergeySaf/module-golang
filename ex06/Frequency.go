package letter

//package main

//import "fmt"

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
		//		freqC = <-freq
	}
	return freqC
}

/*func add(str string) Map_C {
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
	freq := make(chan Map_C, 3)
	for _, i := range astr {
		go func(s string) {
			freq <- add(s)
		}(i)
	}
	//	}(astr[1])
	for range astr {
		for char, n := range <-freq {
			freqC[char] += n
			fmt.Println("freq")
			fmt.Println(char)
			fmt.Println(n)
		}
		//				freqC
		fmt.Println("astr")
	}
	//	freqC = <-freq
	return freqC
}

func main() {
	euro := "annnn"
	dutch := "bb"
	us := "cc"
	a := []string{euro, dutch, us}
	fmt.Println(a[0])
	b := ConcurrentFrequency(a)
	fmt.Println(b[rune(dutch[0])])

}*/
