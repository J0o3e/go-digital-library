package api

import (
	"fmt"
	"net/http"

	"go-digital-library/internal/books"
	"go-digital-library/internal/loans"
	"go-digital-library/internal/users"
)

type Server struct {
	bookService *books.BookService
	userService *users.UserService
	loanService *loans.LoanService
}

func NewServer(
	bookService *books.BookService,
	userService *users.UserService,
	loanService *loans.LoanService,
) *Server {
	return &Server{
		bookService: bookService,
		userService: userService,
		loanService: loanService,
	}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()

	s.registerBookRoutes(mux)
	s.registerUserRoutes(mux)
	s.registerLoanRoutes(mux)

	fmt.Println("API iniciado en http://localhost:8080")
	fmt.Println("servicios disponibles:")
	fmt.Println("GET http://localhost:8080/api/books")
	fmt.Println("POST http://localhost:8080/api/books")
	fmt.Println("GET http://localhost:8080/api/books/search?q=texto")
	fmt.Println("GET  http://localhost:8080/api/users")
	fmt.Println("POST http://localhost:8080/api/users")
	fmt.Println("GET  http://localhost:8080/api/users/search?q=texto")
	fmt.Println("POST http://localhost:8080/api/loans")
	fmt.Println("GET  http://localhost:8080/api/loans/active")
	fmt.Println("GET  http://localhost:8080/api/loans/history")
	fmt.Println("POST http://localhost:8080/api/loans/return")

	return http.ListenAndServe(":8080", mux)
}
