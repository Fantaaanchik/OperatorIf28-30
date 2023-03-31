package Input

import "fmt"

var (
	C           int16
	numberAbout int16
)

func AboutNum() {

	fmt.Printf("If30.\nДано целое число, лежащее в диапазоне 1–999. Вывести его строку-описание вида " +
		"«четное двузначное число», «нечетное трехзначное число» и т. д.\n")
	fmt.Printf("Введите число в диапазоне 1-999:\n")

	_, err := fmt.Scan(&numberAbout)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	C = numberAbout

	return
}
