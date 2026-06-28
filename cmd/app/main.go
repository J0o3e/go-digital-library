package main

import (
	"fmt"

	"go-digital-library/internal/api"
	"go-digital-library/internal/books"
	"go-digital-library/internal/config"
	"go-digital-library/internal/database"
	"go-digital-library/internal/loans"
	"go-digital-library/internal/repositories"
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

	bookRepository := repositories.NewBookRepository(db)
	userRepository := repositories.NewUserRepository(db)
	loanRepository := repositories.NewLoanRepository(db)

	bookService := books.NewBookDBService(bookRepository)
	userService := users.NewUserDBService(userRepository)
	loanService := loans.NewLoanDBService(bookService, userService, loanRepository)

	server := api.NewServer(bookService, userService, loanService)

	err = server.Start()
	if err != nil {
		fmt.Println("Error al iniciar API:", err)
		return
	}
}
