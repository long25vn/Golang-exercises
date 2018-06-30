// concurrency-channel-1 project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	e := make(chan int, 0)
	o := make(chan int, 0)
	brk := make(chan bool, 0)
	var even, odd []int
	go func() {
		for i := 0; i < 100; i++ {
			{
				if i%2 == 1 {
					o <- i
					//	odd = append(odd, i)
				} else {
					e <- i
					//	fmt.Println("even value")
					//even = append(even, i)
				}

			}
			if i == 99 {
				brk <- true
			}
		}
	}()

	go func() {
		for {
			select {
			case evn := <-e:
				fmt.Println("chann even", evn)
				even = append(even, evn)
			case od := <-o:
				fmt.Println("chann odd", od)
				odd = append(odd, od)
			case <-brk:
				break
			}
		}
	}()
	_, _ = fmt.Scanln()
	fmt.Println(even, odd)

}
