package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := []string {"an","lee"}
	for _,person := range people {
		go isSexy(person,c)
	}
	for range people {
		fmt.Println(<-c)
	}
}

func sexyCount(person string) {
	for i := 0; i< 10; i++ {
		fmt.Println(person, i)
		time.Sleep(time.Second)
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 10)
	fmt.Println(person);
	c <- person + "is Sexy"
}