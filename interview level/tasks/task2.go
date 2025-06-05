//Что выведет? Как исправить?

type Person struct {
	Name string
} 

func changeName(person *Person) {
	person = &Person{
		Name: "Alice",
	}
}

func main() {

	person := &Person{
		Name: "Bob",
	}

	fmt.Println(person.Name) // Выведет "Bob"
	changeName(person)
	fmt.Println(person.Name) // Выведет "Bob"
}


//Ответ:

type Person struct {
	Name string
}

func changeName(person **Person) {
	*person = &Person{
		Name: "Alice",
	}
}

func main() {
	person := &Person{
		Name: "Bob",
	}

	fmt.Println(person.Name)
	changeName(&person)
	fmt.Println(person.Name)
}