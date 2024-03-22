package main

import "fmt"

func main() {

	m := map[string]int{} // тип ключа - тип значения

	m["a"] = 1
	m["b"] = 2

	// проверка на существование
	value, ok := m["c"]
	fmt.Println(ok)
	fmt.Println(value)

	value = m["c"]
	fmt.Println(value)

	// удалить
	delete(m, "a")

	value, ok = m["a"]
	fmt.Println(ok)
	fmt.Println(value)

	value = m["a"]
	fmt.Println(value)

	// пройтись циклом
	for key, value := range m {
		fmt.Printf("%s -- %d\n", key, value)
	}

}
