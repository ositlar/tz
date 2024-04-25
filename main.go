package main

import (
	"fmt"
	"strconv"
)

func main() {
	//testFirstTask()
	//testSecondTask()
	//testThirdTask()
	testFourthTask()
}

func testFirstTask() {
	arr := []int{0, 1, 2, 5, 8, 10, 11, 15, 12, 21, 41, 25, 1048}
	for _, i := range arr {
		fmt.Println(firstTask(i))
	}
}

func testSecondTask() {
	arr := []int{42, 12, 18}
	fmt.Println(secondTask(arr))
}

func firstTask(n int) string {
	result := strconv.Itoa(n) + " компьютер"
	if n > 10 && n < 20 {
		return result + "ов"
	}
	lastDigit := n % 10
	if lastDigit == 1 {
		return result
	}
	if lastDigit > 1 && lastDigit < 5 {
		return result + "а"
	}
	return result + "ов"
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func secondTask(numbers []int) []int {
	if len(numbers) == 0 {
		return nil
	}

	commonDivisors := make(map[int]bool)
	firstNumber := numbers[0]

	for _, num := range numbers[1:] {
		gcdVal := gcd(firstNumber, num)

		for divisor := 1; divisor <= gcdVal; divisor++ {
			if gcdVal%divisor == 0 {
				commonDivisors[divisor] = true
			}
		}
	}
	delete(commonDivisors, 1)
	result := make([]int, 0)
	for divisor := range commonDivisors {
		result = append(result, divisor)
	}

	return result
}

func testThirdTask() {
	min := 10
	max := 40
	fmt.Println(thirdTask(min, max))
}

func thirdTask(min, max int) []int {
	result := make([]int, 0)

	k := 0
	for i := min; i <= max; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				k++
			}
		}
		if k == 0 {
			result = append(result, i)
		} else {
			k = 0
		}
	}
	return result
}

func testFourthTask() {
	max := 15
	fourthTask(max)
}

func fourthTask(max int) {
	for i := 0; i <= max; i++ {
		if i == 0 {
			fmt.Print("\t")
		} else {
			fmt.Print(i, "\t")
		}
	}
	fmt.Println()
	for i := 1; i <= max; i++ {
		for j := 0; j <= max; j++ {
			if i*j == 0 {
				fmt.Print(i, "\t")
			} else {
				fmt.Print(i*j, "\t")
			}
		}
		fmt.Println()
	}
}
