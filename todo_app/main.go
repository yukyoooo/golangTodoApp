package main

import (
	"fmt"
	"todo_app/app/controllers"
	"todo_app/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()

	/*
		fmt.Println(config.Config.Port)
		fmt.Println(config.Config.SQLDriver)
		fmt.Println(config.Config.DbName)
		fmt.Println(config.Config.LogFile)

		log.Println("test")
	*/

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.PassWord = "testtest"
	// u.CreateUser()

	// u, _ := models.GetUser(3)
	// u.CreateTodo("Third Todo")

	// u2, _ := models.GetUser(2)
	// todos, _ := u2.GetTodosByUser()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// fmt.Println(u)

	// u.Name = "Test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.CreateTodo("Second Todo")
	// t, _ := models.GetTodo(1)

	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	t, _ := models.GetTodo(1)
	// t.Content = "update todo"
	// t.UpdateTodo()
	t.DeleteTodo()
}
