package main

import "fmt"

func main() {
	i := Plus(1, 2)
	fmt.Println(i)

	i2, i3 := Div(9, 4)
	fmt.Println(i2, i3)

	i4 := Double(1000)
	fmt.Println(i4)

	// 無名関数
	f := func(x, y int) int {
		return x + y
	}
	i5 := f(1, 2)
	fmt.Println(i5)

	CallFunction(func() {
		fmt.Println("im a function")
	})

	f3 := Later()
	fmt.Println(f3("hello"))
	fmt.Println(f3("my"))
	fmt.Println(f3("name"))
	fmt.Println(f3("is go"))

	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Plus(x int, y int) int {
	return x + y
}

func Double(price int) (result int) {
	result = price * 2
	return
}

// 関数を引数にとる関数
func CallFunction(f func()) {
	f()
}

// クロージャーの実装
func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

// ジェネレーターの実装
func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
