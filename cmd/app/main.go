package main

import (
	"fmt"

	"go-digital-library/internal/books"
	"go-digital-library/internal/loans"
	"go-digital-library/internal/users"
)

func main() {
	bookService := books.NewBookService()
	userService := users.NewUserService()
	loanService := loans.NewLoanService(bookService, userService)

	libro1, err := bookService.RegistrarLibro("Clean Code", "Robert Martin", "Programacion", 2008)
	if err != nil {
		fmt.Println("Error al registrar el libro:", err)
		return
	}

	libro2, err := bookService.RegistrarLibro("El principito", "Antoine de Saint-Exupery", "Literatura", 1943)
	if err != nil {
		fmt.Println("Error al registrar el libro", err)
		return
	}

	usuario1, err := userService.RegistrarUsuario("Jose", "jose@email.com")
	if err != nil {
		fmt.Println("Error al registrar el usuario:", err)
		return
	}

	usuario2, err := userService.RegistrarUsuario("Ana", "ana@email.com")
	if err != nil {
		fmt.Println("Error al registrar el usuario:", err)
		return
	}

	fmt.Println("Libros registrados:")
	fmt.Println(libro1.ID(), "-", libro1.Titulo(), "- Disponible:", libro1.Disponible())
	fmt.Println(libro2.ID(), "-", libro2.Titulo(), "- Disponible:", libro2.Disponible())

	fmt.Println("\nUsuarios registrados:")
	fmt.Println(usuario1.ID(), "-", usuario1.Nombre(), "-", usuario1.Correo())
	fmt.Println(usuario2.ID(), "-", usuario2.Nombre(), "-", usuario2.Correo())

	prestamo, err := loanService.RegistrarPrestamo(libro1.ID(), usuario1.ID())
	if err != nil {
		fmt.Println("Error al registrar prestamo:", err)
		return
	}

	fmt.Println("\nPrestamo registrado:")
	fmt.Println("Prestamo ID:", prestamo.ID())
	fmt.Println("Libro ID:", prestamo.LibroID())
	fmt.Println("Usuario ID:", prestamo.UsuarioID())
	fmt.Println("Activo:", prestamo.Activo())

	libroPrestado, err := bookService.BuscarPorID(libro1.ID())
	if err != nil {
		fmt.Println("Error al buscar libro:", err)
		return
	}

	fmt.Println("\nEstado del libro despues del prestamo:")
	fmt.Println("Libro:", libroPrestado.Titulo())
	fmt.Println("Disponible:", libroPrestado.Disponible())

	fmt.Println("\nPrestamos activos:")
	for _, p := range loanService.ListarPrestamosActivos() {
		fmt.Println("Prestamo:", p.ID(), "- Libro:", p.LibroID(), "- Usuario:", p.UsuarioID())
	}

	err = loanService.DevolverPrestamo(prestamo.ID())
	if err != nil {
		fmt.Println("Error al devolver prestamo:", err)
		return
	}

	libroDevuelto, err := bookService.BuscarPorID(libro1.ID())
	if err != nil {
		fmt.Println("Error al buscar libro:", err)
		return
	}

	fmt.Println("\nPrestamo devuelto correctamente")
	fmt.Println("Libro:", libroDevuelto.Titulo())
	fmt.Println("Disponible:", libroDevuelto.Disponible())

	prestamoDevuelto, err := loanService.BuscarPorID(prestamo.ID())
	if err != nil {
		fmt.Println("Error al buscar el prestamo:", err)
		return
	}

	fmt.Println("Prestamo activo:", prestamoDevuelto.Activo())
}
