package goroutines

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(id int, jobs <-chan string) {
	defer wg.Done()
	spawned := false
	for j := range jobs {
		if spawned != true {
			fmt.Printf("worker:%d spawning\n", id)
			spawned = true
		}
		l, _ := time.ParseDuration(j + "s")
		fmt.Printf("worker:%d sleep:%s\n", id, j)
		time.Sleep(l)
	}
	if spawned == true {
		fmt.Printf("worker:%d stopping\n", id)
	}
}

func Run(num int) {
	jobs := make(chan string, num)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	id := 1
	for scanner.Scan() {
		if id <= num {
			wg.Add(1)
			go worker(id, jobs)
			id++
		}
		input := string(scanner.Bytes())
		jobs <- input
	}
	close(jobs)
	wg.Wait()
}
