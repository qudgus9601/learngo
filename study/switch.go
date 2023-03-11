package study

import "fmt"

/*
Dev
*/
func AboutSwitch(age int) {
	switch {
	case age <= 13:
		fmt.Println("초딩")
	case age <= 16 && age > 13:
		fmt.Println("중딩")
	case age <= 19 && age > 16:
		fmt.Println("고딩")
	case age > 19 :
		fmt.Println("대딩")
	}
}

/*
Dev
*/
func DefineVariableBeforeSwitchWork(age int) {
	switch manAge := age - 1;{
	case manAge >= 19:
		fmt.Println("술 마셔도 괜찮다")
	case manAge < 19:
		fmt.Println("으디 어린느무시끼가")
	}
}

/*
Dev
*/
func OnlyEqualSwitch(age int) {
	switch age{
	case 14:
		fmt.Println("초졸 개추")
	case 17:
		fmt.Println("중졸 개추")
	case 20:
		fmt.Println("고졸 개추")
	default :
		fmt.Println("평범쓰")
	}
}