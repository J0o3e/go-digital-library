package main

import (
	"go-digital-library/internal/books"
	"go-digital-library/internal/console"
	"go-digital-library/internal/loans"
	"go-digital-library/internal/users"
)

func main() {
	bookService := books.NewBookService()
	userService := users.NewUserService()
	loanService := loans.NewLoanService(bookService, userService)

	menu := console.NewMenu(bookService, userService, loanService)
	menu.Ejecutar()
}
