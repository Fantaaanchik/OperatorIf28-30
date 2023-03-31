package funcs

import (
	"fmt"
	"yearCounter/Input"
)

func DataChecker30() {
	Input.AboutNum()
	if Input.C%2 == 0 && Input.C < 10 {
		fmt.Printf("Число %v является четным одназначным", Input.C)
	} else if Input.C%2 != 0 && Input.C < 10 {
		fmt.Printf("Число %v является нечетным одназначным", Input.C)
	} else if Input.C%2 != 0 && Input.C >= 10 && Input.C < 100 {
		fmt.Printf("Число %v является нечетным двузначным", Input.C)
	} else if Input.C%2 == 0 && Input.C >= 10 && Input.C < 100 {
		fmt.Printf("Число %v является четным двузначным", Input.C)
	} else if Input.C%2 != 0 && Input.C >= 100 && Input.C < 1000 {
		fmt.Printf("Число %v является нечетным трехзначным", Input.C)
	} else if Input.C%2 == 0 && Input.C >= 100 && Input.C < 1000 {
		fmt.Printf("Число %v является четным трехзначным", Input.C)
	} else {
		fmt.Printf("Число %v выходит за пределы диапазона 1-999", Input.C)
	}

	return
}
