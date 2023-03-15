package study

import (
	"fmt"
	"time"
)

func sexyCount(person string) {
	for i := 0; i< 10; i++ {
		fmt.Println(person, i)
		time.Sleep(time.Second)
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 3)
	c <- person + " is Sexy"
}

func goroutine() {
	c := make(chan string)
	people := []string{"an", "lee"}
	for _, person := range people {
		go isSexy(person, c)
	}
	fmt.Println("응답 대기중...")
	for range people {
		fmt.Println("들어온 메세지 : " + <-c)
	}
}