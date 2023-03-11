package study

import "fmt"

/*
Dev
*/
func Aboutif (num1, num2 int) {
	if num1 >= num2 {
		fmt.Println("num1 이 num2 보다 큽니다.")
	} else {
		fmt.Println("num1 이 num2 보다 작습니다.")
	}
}

/*
Dev
*/
func DefineVariableBeforeIfWork (num int) {
	if newVar := float64(num) / 2; newVar >= float64(num) {
		fmt.Println("새로 생성된 값 newVar이 더 큽니다.",newVar)
	} else {
		fmt.Println("새로 생성된 값 newVar이 더 작습니다.", newVar)
	}
}