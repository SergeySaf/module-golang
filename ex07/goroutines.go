package goroutines

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

var wait_buf sync.WaitGroup
var stop sync.Mutex

func Workers(tasks chan string, number int) {
	defer wait_buf.Done()
	flag := 0
	stop.Unlock()

	for i := range tasks {
		if flag == 0 {
			fmt.Printf("worker:%d spawning\n", number)
		}
		flag++
		timer, _ := time.ParseDuration(i + "s")
		fmt.Printf("worker:%d sleep:%s\n", number, i)
		time.Sleep(timer)
	}

	if flag > 0 {
		fmt.Printf("worker:%d stopping\n", number)
	}
}

func Run(num int) {
	scanner := bufio.NewScanner(os.Stdin)
	tasks := make(chan string, 64)

	for scanner.Scan() {
		input := string(scanner.Bytes())
		tasks <- input
	}

	for i := 1; i <= num; i++ {
		wait_buf.Add(1)
		stop.Lock()
		go Workers(tasks, i)
	}
	close(tasks)
	wait_buf.Wait()
}
