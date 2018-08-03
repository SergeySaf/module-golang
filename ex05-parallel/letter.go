package letter

import "fmt"

//Frequency ...
func Frequency(str string) map[rune]int {
	myMap := make(map[rune]int)

	for _, value := range str {
		myMap[value]++
	}
	fmt.Println(myMap)
	return myMap
}

func parrFrequency(str string, c chan map[rune]int) {
	c <- Frequency(str)
}

//ConcurrentFrequency ...
func ConcurrentFrequency(arrStr []string) map[rune]int {
	myMap := make(map[rune]int)
	c := make(chan map[rune]int)
	for i := range arrStr {
		go parrFrequency(arrStr[i], c)
	}
	for range arrStr {
		for letter, count := range <-c {
			myMap[letter] += count
		}
	}
	fmt.Println(myMap)
	return myMap
}
