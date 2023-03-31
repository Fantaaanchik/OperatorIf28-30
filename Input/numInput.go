package Input

import "fmt"

var (
	B int16
)

func NumInput() {
	var number int16
	fmt.Printf("If29.\nДано целое число. Вывести его строку-описание вида «отрицательное четное число», «нулевое " +
		"число», «положительное нечетное число» и т. д.\n")
	fmt.Printf("Введите число:\n")
	_, er := fmt.Scan(&number)
	if er != nil {
		fmt.Println(er.Error())
		return
	}
	B = number

	return
}
