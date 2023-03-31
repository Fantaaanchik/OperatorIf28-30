package funcs

import (
	"fmt"
	"yearCounter/Input"
)

func DataCheckerIf28() {

	Input.InputData()

	if Input.A%4 == 0 && Input.A%100 == 0 && Input.A%400 == 0 {
		fmt.Printf("%v год является высококосным годом", Input.A)
	} else {
		fmt.Printf("%v год является обычным годом", Input.A)
	}

	return
}
