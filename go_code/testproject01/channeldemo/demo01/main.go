package main

import (
	"fmt"
)

func PutNum(intChan chan int) {
	for i := 2; i <= 8000; i++ {
		intChan <- i
	}

	defer close(intChan)
}

// 取出来并求是否为素数，如果是就放入 primeChan
func IsPrime(intChan chan int, primeChan chan int, resChan chan bool) {
	var flag bool
	for {
		//time.Sleep(time.Microsecond * 10)
		num1, ok := <-intChan
		if !ok {
			break
		}
		flag = true //假设是素数
		//判断是否为素数
		for i := 2; i < num1; i++ {
			if num1%i == 0 {
				flag = false
				break
			}
		}

		if flag {
			//将这个数放入primeChan
			primeChan <- num1
		}
	}

	fmt.Println("这个协程结束，退出")

	resChan <- true
}

func main() {
	intChan := make(chan int, 2000)
	primeChan := make(chan int, 2000)
	resChan := make(chan bool, 4)

	go PutNum(intChan)

	for i := 0; i < 4; i++ {
		go IsPrime(intChan, primeChan, resChan)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-resChan
		}
		close(primeChan)
	}()

	//遍历取出primeChan

	for {
		res, ok := <-primeChan
		if !ok {
			break
		}

		fmt.Printf("素数=%v \n", res)
	}

	fmt.Println("主线程退出")
}
