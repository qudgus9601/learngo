package study

import "fmt"

/*
Dev
*/
func Map(key, value string) {
	newMap := map[string]string{key: value}

	fmt.Println(newMap)
}

/*
Dev
*/
func MapRange(key,value string) {
	fmt.Println("map에도 for range 를 사용하여 모든 키 값을 출력 할 수 있습니다.")
	newMap := map[string]string{key:value}

	for key, value := range newMap {
		fmt.Printf("key: %v, value: %v",key,value)
	}
}