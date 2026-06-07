package main

import (
	"fmt"
	"go-digital-library/internal/models"
)

func main() {
	libro, err := models.NuevoLibro(1, "Clean Code", "Robert Martin", "Educativo", 2008)
	if err != nil {
		fmt.Println("Error al crear el libro:", err)
		return
	}

	usuario, err := models.NuevoUsuario(1, "Jose", "jose@email.com")
	if err != nil {
		fmt.Println("Error al crear el usuario:", err)
		return
	}

	prestamo, err := models.NuevoPrestamo(1, libro.ID(), usuario.ID())
	if err != nil {
		fmt.Println("Error al crear prestamo:", err)
		return
	}

	err = libro.Prestar()
	if err != nil {
		fmt.Println("Error al prestar libro:", err)
		return
	}

	fmt.Println("Sistema Go Digital Library")
	fmt.Println("------------------------")
	fmt.Println("Libro:", libro.Titulo())
	fmt.Println("Autor:", libro.Autor())
	fmt.Println("Categoria:", libro.Categoria())
	fmt.Println("Anio:", libro.Anio())
	fmt.Println("Usuario:", usuario.Nombre())
	fmt.Println("Correo:", usuario.Correo())
	fmt.Println("Prestamo ID:", prestamo.ID())
	fmt.Println("Libro disponible:", libro.Disponible())
	fmt.Println("Prestamo activo:", prestamo.Activo())
}
