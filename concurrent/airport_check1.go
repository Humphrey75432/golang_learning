package main

import "time"

const (
	idCheckTmCost   = 60
	bodyCheckTmCost = 120
	xRayCheckTmCost = 180
)

func idCheck(id int) int {
	time.Sleep(time.Millisecond * time.Duration(idCheckTmCost))
	println("\tgoroutine-", id, ": idCheck ok")
	return idCheckTmCost
}

func bodyCheck(id int) int {
	time.Sleep(time.Millisecond * time.Duration(bodyCheckTmCost))
	println("\tgoroutine-", id, ": bodyCheck ok")
	return bodyCheckTmCost
}

func xRayCheck(id int) int {
	time.Sleep(time.Millisecond * time.Duration(xRayCheckTmCost))
	println("\tgoroutine-", id, ": xRayCheck ok")
	return xRayCheckTmCost
}

func airportSecurityCheck(id int) int {
	println("goroutine-", id, ": airportSecurityCheck start")
	total := 0
	total += idCheck(id)
	total += bodyCheck(id)
	total += xRayCheck(id)
	println("goroutine-", id, ": airportSecurityCheck ok")
	return total
}

func start(id int, f func(int) int, queue <-chan struct{}) <-chan int {
	c := make(chan int)
	go func() {
		total := 0
		for {
			_, ok := <-queue
			if !ok {
				c <- total
				return
			}
			total += f(id)
		}
	}()
	return c
}

func main() {
	total := 0
	passenger := 30
	c := make(chan struct{})
	c1 := start(1, airportSecurityCheck, c)
	c2 := start(2, airportSecurityCheck, c)
	c3 := start(3, airportSecurityCheck, c)

	for i := 0; i < passenger; i++ {
		c <- struct{}{}
	}
	close(c)

	total = max(total, <-c1, <-c2, <-c3)
	println("total time cost:", total)
}
