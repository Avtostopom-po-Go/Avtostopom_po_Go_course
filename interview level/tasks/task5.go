//Добавить код, который выведет тип переменной whoami

func printType(whoami interface{}) {

}

func main() {
	printType(42)
	printType("im string")
	printType(true)
}


//Ответ:

package main

import "fmt"

func printType(whoami interface{}) {
	fmt.Printf("Type of whoami: %T\n", whoami)
	//fmt.Printf("Type of whoami: %v\n", reflect.TypeOf(whoami))
}

func main() {
	printType(42)
	printType("im string")
	printType(true)
}