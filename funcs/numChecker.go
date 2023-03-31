package funcs

import (
	"fmt"
	"yearCounter/Input"
)

func DataCheckerIf29() {

	Input.NumInput()

	if Input.B < 0 && Input.B%2 == 0 {
		fmt.Printf("Число %v это отрицательное четное число", Input.B)
	} else if Input.B < 0 && Input.B%2 != 0 {
		fmt.Printf("Число %v это отрицательное нечетное число", Input.B)
	} else if Input.B > 0 && Input.B%2 == 0 {
		fmt.Printf("Число %v это положительное четное число", Input.B)
	} else if Input.B > 0 && Input.B%2 != 0 {
		fmt.Printf("Число %v это положительное нечетное число", Input.B)
	} else {
		fmt.Printf("Число %v это нулевое число", Input.B)
	}

	return
}
