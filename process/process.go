package process

import "fmt"

func CheckEven(number int) string {
	if number%2 == 0 {
		return "YES"
	}
	return "NO"
}

func main() {
	fmt.Println("Test 6 : ", CheckEven(6))
	fmt.Println("Test 11 : ", CheckEven(11))
}
