//Что выведет код?

type impl struct{}

type I interface {
	C()
}

func (*impl) C() {}

func A() I {
	return nil
}

func B() I {
	var ret *impl
	return ret
}

func main() {
	a := A()
	b := B()
	fmt.Println(a == b)
}


//Ответ:

type impl struct{}

type I interface {
	C()
}

func (*impl) C() {}

func A() I {
	return nil
}

func B() I {
	var ret *impl
	return ret
}

func main() {
	a := A() // true
	b := B() // true
	fmt.Println(a == b) // false
}

//Когда мы сравниваем "a" и "b" с помощью оператора "==",
//мы сравниваем значения интерфейсов. В данном случае "a"
//содержит nil, а "b" содержит указатель на нулевую структуру.