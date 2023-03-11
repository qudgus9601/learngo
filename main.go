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
	
	// 반복문
	study.ForUsage(1,2);
	study.ForUsage1(1,2);
	study.ForUsage2(1,2,3);
	study.ForUsage3(1,2,3);

	// 조건문
	study.Aboutif(3,2)
	study.DefineVariableBeforeIfWork(-1);
}