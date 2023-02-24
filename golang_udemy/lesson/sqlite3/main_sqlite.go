package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

/*
func IsOne(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}
*/

var Db *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	Db, _ := sql.Open("sqlite3", "./example.sql")
	defer Db.Close()

	/*
		cmd := `CREATE TABLE IF NOT EXISTS persons(
			name STRING,
			age INT
		)`
		_, err := Db.Exec(cmd)
	*/

	/*
		cmd := "INSERT INTO persons (name, age) VALUES (?, ?)"
		_, err := Db.Exec(cmd, "hanako", 19)
	*/

	/*
		cmd := "UPDATE persons SET age = ? WHERE name = ?"
		_, err := Db.Exec(cmd, 25, "tarou")
	*/

	/*
		cmd := "SELECT * FROM persons where age = ?"
		// QueryRow 1レコード取得
		row := Db.QueryRow(cmd, 25)
		var p Person
		err := row.Scan(&p.Name, &p.Age)

		if err != nil {
			if err == sql.ErrNoRows {
				log.Panicln("No row")
			} else {
				log.Panicln(err)
			}
		}
	*/

	cmd := "SELECT * FROM persons"
	// Query 条件にあうもの全て取得
	rows, _ := Db.Query(cmd)
	defer rows.Close()
	var pp []Person
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Panicln(err)
		}
		pp = append(pp, p)
	}

	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}

	/*
		fmt.Println(IsOne(1))
		fmt.Println(IsOne(2))

		s := []int{1, 2, 3, 4, 5}
		fmt.Println(alib.Average(s))
	*/
	/*
			fmt.Println(foo.Max)

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
	*/
}
