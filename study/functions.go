package study

import "fmt"

/*
Dev
*/
func ManyParamFunc(param ...string) {
	fmt.Println(param)
}

/*
Dev
*/
func ManyReturnFunc(str1,str2 string) (string,string) {
	fmt.Println(str1,str2)
	return str1,str2
}

/*
Dev
*/
func UnCompileSomeVariableFunc(str1,str2 string) {
	compiledVariable, _ := str1, str2
	fmt.Println(compiledVariable);
}

/*
Dev
*/
func NakedReturnFunc(str1 string) (length int, str string) {
	length = len(str1);
	str = str1
	return;
}

/*
Dev
*/
func DeferFunc() {
	defer fmt.Println("defer done when DeferFunc Ended")
	fmt.Println("DeferFunc Called")
}