package study

import "fmt"

/*
Dev
*/
func Pointer (str string) {
	a := str
	b := a
	c := &a

	fmt.Printf("입력받은 값 %v\n", a);
	fmt.Printf("입력 받은 값의 값을 저장한 변수 %v\n", b)
	fmt.Printf("입력 받은 값의 주소를 저장한 변수 %d\n", c)
	fmt.Printf("입력 받은 값의 주소를 저장한 변수가 가리키는 값 %v\n", *c)
	fmt.Printf("===========입력 받은 값에 변화를 준다면========\n")
	a = "Bye World"
	fmt.Printf("입력받은 값 %v\n", a);
	fmt.Printf("입력 받은 값의 값을 저장한 변수 %v\n", b)
	fmt.Printf("입력 받은 값의 주소를 저장한 변수 %d\n", c)
	fmt.Printf("입력 받은 값의 주소를 저장한 변수가 가리키는 값 %v\n", *c)
	fmt.Printf("===========입력 받은 값의 주소를 저장한 변수가 가리키는 값을 변경하면 ========\n")
	*c = "Change World"
	fmt.Printf("입력받은 값 %v\n", a)
}