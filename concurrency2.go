package main

func sum(ints []int, c chan int) {
	sum := 0
	for _, v := range ints {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	a := []int{1, 2, 3, 4}
	b := []int{2, 4, 6, 8}

	c := make(chan int)
	go sum(a, c)
	go sum(b, c)
	x := <-c // receive from c
	y := <-c

	println(x, y, x+y)
}
