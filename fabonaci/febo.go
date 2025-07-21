package main

func printFebo(first chan int, second chan int, done chan bool) {
	defer func() { done <- true }()

	for range 10 {
		f := <-first
		s := <-second
		print(s, " ")
		first <- s
		second <- f + s
	}
}

func main() {
	first := make(chan int, 1)
	second := make(chan int, 1)
	done := make(chan bool)

	go printFebo(first, second, done)
	first <- 0
	second <- 1

	<-done
}
