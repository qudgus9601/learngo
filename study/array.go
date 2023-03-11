package study

import "fmt"

func ArrayBasic(length int) {
	fmt.Println("make 함수를 이용하여 배열의 길이를 지정할 수 있습니다.")
	fmt.Println("make 함수를 이용하지 않고 배열의 크기를 동적으로 정하려면 slice를 이용해야합니다.")
	arr := make([]int, length)
	for i := 0; i < len(arr); i++ {
		arr[i] = i;
	}
	fmt.Println(arr)
}

func Slice(length int) {
	arr := []int{}
	for i := 0; i<length; i++ {
		arr = append(arr, i);
	}

	fmt.Println(arr)
}