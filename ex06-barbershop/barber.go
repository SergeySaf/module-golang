package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	maxWaiting = 3
	nCustomers = 8
)

var (
	lobby = make(chan chan int, maxWaiting)
	wg    = new(sync.WaitGroup)
	names = []string{"Bob", "John", "Pal", "Josh", "Alex", "Igor", "Andrew", "Sergei"}
)

func barber() {
	for {
		select {
		case ch, _ := <-lobby:
			id := <-ch
			fmt.Println("Barber Phill cuts the hair of", names[id])
			time.Sleep(2 * time.Second)
			fmt.Printf("Barber Phill finished cutting %v\n", names[id])
			ch <- 0
		}
	}
}

func customer(id int) {
	defer wg.Done()
	ch := make(chan int)
	fmt.Println(names[id], "enters the barbershop.")
	select {
	case lobby <- ch:
		ch <- id
		<-ch
	default:
		fmt.Println(names[id], "will come after one hour.")
		time.Sleep(3 * time.Second)
		wg.Add(1)
		customer(id)
	}
}

func main() {
	wg.Add(nCustomers)
	go barber()
	for i := 0; i < nCustomers; i++ {
		time.Sleep(1 * time.Second)
		go customer(i)
	}
	wg.Wait()
	fmt.Println("End of the day")
}
