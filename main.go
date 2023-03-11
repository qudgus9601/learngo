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
	// if
	study.Aboutif(3,2)
	study.DefineVariableBeforeIfWork(-1);
	// switch
	study.AboutSwitch(28)
	study.DefineVariableBeforeSwitchWork(25)
	study.OnlyEqualSwitch(3)

	// 포인터
	study.Pointer("Hello World")

	// 배열
	study.ArrayBasic(5)
	study.Slice(5)
	study.Map("name","안병현")
	study.MapRange("name","안병현")

	// 구조체
	study.Struct()
}