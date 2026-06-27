package main

import (
	"fmt"

	"go-digital-library/internal/books"
	"go-digital-library/internal/config"
	"go-digital-library/internal/console"
	"go-digital-library/internal/database"
	"go-digital-library/internal/loans"
	"go-digital-library/internal/users"
)

func main() {
	dbConfig := config.LoadDatabaseConfig()

	db, err := database.Connect(dbConfig)
	if err != nil {
		fmt.Println("Error al conectar con PostgreSQL:", err)
		return
	}
	defer db.Close()

	fmt.Println("Conexion a PostgreSQL exitosa.")

	bookService := books.NewBookService()
	userService := users.NewUserService()
	loanService := loans.NewLoanService(bookService, userService)

	menu := console.NewMenu(bookService, userService, loanService)
	menu.Ejecutar()
}
