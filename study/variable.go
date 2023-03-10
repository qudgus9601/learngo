package study

import "fmt"

/*
Dev
	변수의 내용을 출력합니다.
*/
func Variable() {
	var variableVariable string = "var 로 선언된 변수는 가변성입니다."
	variableVariable2 := "var 는 두가지 방식으로 선언 가능합니다."
	fmt.Println(variableVariable, variableVariable2)
}

/*
Dev
	상수에 대해 출력합니다.
*/
func Constant() {
	const constantVariable string = "const 로 선언된 변수는 불변입니다."
	fmt.Println(constantVariable)
}
