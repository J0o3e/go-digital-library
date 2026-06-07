package main

import (
	"fmt"
	"go-digital-library/internal/books"
)

func main() {
	bookService := books.NewBookService()

	libro1, err := bookService.RegistrarLibro("Clean Code", "Robert Martin", "Programacion", 2008)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	libro2, err := bookService.RegistrarLibro("El Principito", "Antoine de Saint-Exupery", "Literatura", 1943)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Libros registrados:")
	fmt.Println(libro1.ID(), "-", libro1.Titulo())
	fmt.Println(libro2.ID(), "-", libro2.Titulo())

	fmt.Println("\nListado completo")
	for _, libro := range bookService.ListarLibros() {
		fmt.Println("Encontrado:", libro.Titulo())
	}

	fmt.Println("\nBusqueda por texto: programacion")
	resultados := bookService.BuscarPorTexto("programacion")
	for _, libro := range resultados {
		fmt.Println("Encontrado:", libro.Titulo())
	}

	err = bookService.PrestarLibro(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	libroBuscado, err := bookService.BuscarPorID(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nLibro prestado:", libroBuscado.Titulo())
	fmt.Println("Disponible:", libroBuscado.Disponible())
}
