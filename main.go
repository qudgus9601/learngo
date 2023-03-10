package main

import (
	"fmt"
	"learngo/study"
)

func main() {
	// 변수
	study.Variable()
	study.Constant()

	// 함수
	study.ManyParamFunc("h","e","l","l","o")
	study.ManyReturnFunc("str1", "str2")
	study.UnCompileSomeVariableFunc("str1","str2")
	fmt.Println(study.NakedReturnFunc("hello world"));
	study.DeferFunc();
	// 
}