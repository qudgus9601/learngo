package main

import (
	"fmt"
	"learngo/banking"
	"learngo/dict"
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

	// bank account
	account := banking.NewAccount("안병현",1000);
	fmt.Println(account);
	account.Deposit(100)
	fmt.Println(account);
	err := account.Withdraw(2000)
	if err != nil {
		// log.Fatalln(err) 프로그램을 종료시킨다.
		fmt.Println(err)
	}
	fmt.Println(account);

	// dictionary
	dictionary := dict.Dictionary{}
	dictionary.Add("hello","world")
	definition, err := dictionary.Search("hello")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
	result, err := dictionary.Update("hello","bye");
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	res,err := dictionary.Search("hello");
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	dictionary.Delete("hello");
	data,err := dictionary.Search("hello");
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}
}