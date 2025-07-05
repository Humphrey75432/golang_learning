package main

import "time"

const (
	idCheckTmCost   = 60
	bodyCheckTmCost = 120
	xRayCheckTmCost = 180
)

func idCheck(id string) int {
	time.Sleep(time.Millisecond * time.Duration(idCheckTmCost))
	println("\tgoroutine-", id, ": idCheck ok")
	return idCheckTmCost
}

func bodyCheck(id string) int {
	time.Sleep(time.Millisecond * time.Duration(bodyCheckTmCost))
	println("\tgoroutine-", id, ": bodyCheck ok")
	return bodyCheckTmCost
}

func xRayCheck(id string) int {
	time.Sleep(time.Millisecond * time.Duration(xRayCheckTmCost))
	println("\tgoroutine-", id, ": xRayCheck ok")
	return xRayCheckTmCost
}

func start(id string, f func(string) int, next chan<- struct{}) (chan<- struct{}, chan<- struct{}, <-chan int) {
	queue := make(chan struct{}, 10)
	quit := make(chan struct{})
	result := make(chan int)

	go func() {
		total := 0
		for {
			select {
			case <-quit:
				result <- total
				return
			case v := <-queue:
				total += f(id)
				if next != nil {
					next <- v
				}
			}
		}
	}()

	return queue, quit, result
}

func newAirportSecurityCheckChannel(id string, queue <-chan struct{}) {
	go func(id string) {
		println("goroutine-", id, ": airportSecurityCheck start")
		queue3, quit3, result3 := start(id, xRayCheck, nil)
		queue2, quit2, result2 := start(id, bodyCheck, queue3)
		queue1, quit1, result1 := start(id, idCheck, queue2)

		for {
			select {
			case v, ok := <-queue:
				if !ok {
					close(quit1)
					close(quit2)
					close(quit3)
					total := max(<-result1, <-result2, <-result3)
					println("goroutine-", id, ": total time cost:", total)
					println("goroutine-", id, ": airportSecurityCheck closed")
					return
				}
				queue1 <- v
			}
		}
	}(id)
}

func main() {
	passenger := 30
	queue := make(chan struct{}, passenger)
	newAirportSecurityCheckChannel("channel1", queue)
	newAirportSecurityCheckChannel("channel2", queue)
	newAirportSecurityCheckChannel("channel3", queue)

	time.Sleep(5 * time.Second)
	for i := 0; i < passenger; i++ {
		queue <- struct{}{}
	}
	time.Sleep(5 * time.Second)
	close(queue)
	time.Sleep(1000 * time.Second)
}
