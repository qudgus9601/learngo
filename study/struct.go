package study

import "fmt"

type developer struct {
	name  string
	age   int
	stack []string
}

func Struct() {
	stack := []string{"Javascript", "Typescript", "Go"}
	an := developer{"안병현", 28, stack}
	fmt.Println(an)
}

func StructAccurate() {
	stack := []string{"Javascript", "Typescript", "Go"}
	an := developer{name:"안병현", age:28, stack:stack}
	fmt.Println(an)
}