package main

import "fmt"

func main() {
	fmt.Println("11~100 사이의 두 숫자를 입력해주세요!")
	var first_num int
	var second_num int

	fmt.Scanln(&first_num)
	fmt.Scanln(&second_num)
	var minus_value int
	minus_value = (first_num - second_num)

	fmt.Println("입력한 숫자들의 차감 갑 =", minus_value)

	if (minus_value <= 0) || (minus_value%11 == 0) {
		fmt.Println("차감 값이 11의 배수이거나 0 이하의 값으로 종료함")
		return
	}

	for i := 1; i <= 89; i++ {
		first_num++
		minus_value = (first_num - second_num)
		if minus_value%11 == 0 {
			fmt.Println("11의 배수가 되는 회수 =", i)
			break
		}
	}

}
