package study

import "fmt"

/*
Dev
*/
func ForUsage (nums ...int) {
	fmt.Println("인덱스만을 사용합니다.")
	for idx := range nums {
		fmt.Printf("idx: %d\n",idx);
	}
}

/*
Dev
*/
func ForUsage1 (nums ...int) {
	fmt.Println("값만을 사용합니다.")
	for _, num := range nums {
		fmt.Printf("num: %d\n",num);
	}
}


/*
Dev
*/
func ForUsage2 (nums ...int) {
	fmt.Println("인덱스와 값을 사용합니다.")
	for idx, num := range nums {
		fmt.Printf("idx: %d, num: %d\n",idx, num);
	}
}

/*
Dev
*/
func ForUsage3 (nums ...int) {
	fmt.Println("타 언어와 비슷한 구조를 사용합니다.")
	for i := 0; i < len(nums); i++ {
		fmt.Printf("i: %d\n",i)
	}
}

