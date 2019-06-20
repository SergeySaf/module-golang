package goroutines

func Change_str(input, output chan string, close_sig chan bool) {
	var buff []byte
	str := <-input
	buff = append(buff, '(')
	for i := 0; i < len(str); i++ {
		buff = append(buff, str[i])
	}
	buff = append(buff, ')')
	str2 := string(buff[:])
	output <- str2
	close_sig <- true
}

func Chan_close(close_sig chan bool, output chan string) {
	s := <-close_sig
	if s == true {
		close(output)
	}
}

func Process(input chan string) chan string {
	close_sig := make(chan bool)
	output := make(chan string)
	go Change_str(input, output, close_sig)
	go Chan_close(close_sig, output)
	return output
}
