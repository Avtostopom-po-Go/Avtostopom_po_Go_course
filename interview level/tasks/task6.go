//Исправить функцию, чтобы она работала.
//Сигнатуру менять нельзя

func printNumber(ptrToNumber interface{}) {
	if ptrToNumber != nil {
		fmt.Println(*ptrToNumber.(*int))
	} else {
		fmt.Println("nil")
	}
}

func main() {

	v := 10
	printNumber(&v)
	var pv *int

	printNumber(pv)
	pv = &v
	printNumber(pv)
}


//Ответ:

func printNumber(ptrToNumber interface{}) {

	if ptrToNumber == nil {
		fmt.Println("nil")
		return
	}

	num, ok := ptrToNumber.(*int)

	if !ok || num == nil {
		fmt.Println("nil")
		return
	}

	fmt.Println(*num)
}
