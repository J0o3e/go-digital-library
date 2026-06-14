package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"go-digital-library/internal/books"
	"go-digital-library/internal/loans"
	"go-digital-library/internal/users"
)

func main() {
	bookService := books.NewBookService()
	userService := users.NewUserService()
	loanService := loans.NewLoanService(bookService, userService)

	reader := bufio.NewReader(os.Stdin)

	for {
		mostrarMenu()

		opcionTexto := leerTexto(reader, "Seleccione una opcion: ")
		opcion, err := strconv.Atoi(opcionTexto)
		if err != nil {
			fmt.Println("Opcion invalida. Debe ingresar un numero.")
			continue
		}

		switch opcion {
		case 1:
			registrarLibro(reader, bookService)
		case 2:
			listarLibros(bookService)
		case 3:
			registrarUsuario(reader, userService)
		case 4:
			listarUsuarios(userService)
		case 5:
			registrarPrestamo(reader, loanService)
		case 6:
			listarPrestamosActivos(loanService)
		case 7:
			devolverPrestamo(reader, loanService)
		case 8:
			buscarLibros(reader, bookService)
		case 9:
			buscarUsuarios(reader, userService)
		case 10:
			fmt.Println("Saliendo del sistema...")
			return
		default:
			fmt.Println("Opcion no valida.")
		}
	}
}

func mostrarMenu() {
	fmt.Println("\n===== GO DIGITAL LIBRARY =====")
	fmt.Println("1. Registrar libro")
	fmt.Println("2. Listar libros")
	fmt.Println("3. Registrar usuario")
	fmt.Println("4. Listar usuarios")
	fmt.Println("5. Registar prestamo")
	fmt.Println("6. Listar prestamos activos")
	fmt.Println("7. Devolver prestamo")
	fmt.Println("8. Buscar libros")
	fmt.Println("9. Buscar usuarios")
	fmt.Println("10. Salir")
}

func leerTexto(reader *bufio.Reader, mensaje string) string {
	fmt.Println(mensaje)
	texto, _ := reader.ReadString('\n')
	return strings.TrimSpace(texto)
}

func leerEntero(reader *bufio.Reader, mensaje string) (int, error) {
	texto := leerTexto(reader, mensaje)
	numero, err := strconv.Atoi(texto)
	if err != nil {
		return 0, err
	}

	return numero, nil
}

func registrarLibro(reader *bufio.Reader, bookService *books.BookService) {
	fmt.Println("\n--- Registrar libro ---")

	titulo := leerTexto(reader, "Titulo: ")
	autor := leerTexto(reader, "Autor: ")
	categoria := leerTexto(reader, "Categoria: ")

	anio, err := leerEntero(reader, "Anio: ")
	if err != nil {
		fmt.Println("Error: el anio debe ser un numero.")
		return
	}

	libro, err := bookService.RegistrarLibro(titulo, autor, categoria, anio)
	if err != nil {
		fmt.Println("Error al registrar libro:", err)
		return
	}

	fmt.Println("Libro registrado correctamente.")
	fmt.Println("ID:", libro.ID(), "- Titulo:", libro.Titulo())
}

func listarLibros(bookService *books.BookService) {
	fmt.Println("\n--- Lista de libros ---")

	libros := bookService.ListarLibros()
	if len(libros) == 0 {
		fmt.Println("No hay libros registrados.")
		return
	}

	for _, libro := range libros {
		estado := "Disponible"
		if !libro.Disponible() {
			estado = "Prestado"
		}

		fmt.Println(
			"ID:", libro.ID(),
			"| Titulo:", libro.Titulo(),
			"| Autor:", libro.Autor(),
			"| Categoria:", libro.Categoria(),
			"| Anio:", libro.Anio(),
			"| Estado:", estado,
		)
	}
}

func buscarLibros(reader *bufio.Reader, bookService *books.BookService) {
	fmt.Println("\n--- Buscar libros ---")

	texto := leerTexto(reader, "Ingresar titulo, autor o categoria: ")
	resultados := bookService.BuscarPorTexto(texto)

	if len(resultados) == 0 {
		fmt.Println("No se encontraron libros.")
		return
	}

	for _, libro := range resultados {
		estado := "Disponible"
		if !libro.Disponible() {
			estado = "Prestado"
		}

		fmt.Println(
			"ID:", libro.ID(),
			"| Titulo:", libro.Titulo(),
			"| Autor:", libro.Autor(),
			"| Categoria:", libro.Categoria(),
			"| Anio:", libro.Anio(),
			"| Estado:", estado,
		)
	}
}

func registrarUsuario(reader *bufio.Reader, userService *users.UserService) {
	fmt.Println("\n--- Registrar usuario ---")

	nombre := leerTexto(reader, "Nombre: ")
	correo := leerTexto(reader, "Correo: ")

	usuario, err := userService.RegistrarUsuario(nombre, correo)
	if err != nil {
		fmt.Println("Error al registrar usuario:", err)
		return
	}

	fmt.Println("Usuario registrado correctamente.")
	fmt.Println("ID:", usuario.ID(), "- Nombre:", usuario.Nombre())
}

func listarUsuarios(userService *users.UserService) {
	fmt.Println("\n--- Lista de usuarios ---")

	usuarios := userService.ListarUsuarios()
	if len(usuarios) == 0 {
		fmt.Println("No hay usuarios registrados.")
		return
	}

	for _, usuario := range usuarios {
		estado := "Activo"
		if !usuario.Activo() {
			estado = "Inactivo"
		}

		fmt.Println(
			"ID:", usuario.ID(),
			"| Nombre:", usuario.Nombre(),
			"| Correo:", usuario.Correo(),
			"| Estado:", estado,
		)
	}
}

func buscarUsuarios(reader *bufio.Reader, userService *users.UserService) {
	fmt.Println("\n--- Buscar usuarios ---")

	texto := leerTexto(reader, "Ingrese nombre o correo: ")
	resultados := userService.BuscarPorTexto(texto)

	if len(resultados) == 0 {
		fmt.Println("No se encontraron usuarios.")
		return
	}

	for _, usuario := range resultados {
		estado := "Activo"
		if !usuario.Activo() {
			estado = "Inactivo"
		}

		fmt.Println(
			"ID:", usuario.ID(),
			"| Nombre:", usuario.Nombre(),
			"| Correo:", usuario.Correo(),
			"| Estado:", estado,
		)
	}
}

func registrarPrestamo(reader *bufio.Reader, loanService *loans.LoanService) {
	fmt.Println("\n--- Registrar prestamo ---")

	libroID, err := leerEntero(reader, "ID del libro: ")
	if err != nil {
		fmt.Println("Error: el ID del libro debe ser un numero.")
		return
	}

	usuarioID, err := leerEntero(reader, "ID del usuario: ")
	if err != nil {
		fmt.Println("Error: el ID del usuario deber ser un numero")
		return
	}

	prestamo, err := loanService.RegistrarPrestamo(libroID, usuarioID)
	if err != nil {
		fmt.Println("Error al registrar prestamo:", err)
		return
	}

	fmt.Println("Prestamo registrado correctamente.")
	fmt.Println("Prestamo ID:", prestamo.ID())
}

func listarPrestamosActivos(loanService *loans.LoanService) {
	fmt.Println("\n--- Prestamos activos ---")

	prestamos := loanService.ListarPrestamosActivos()
	if len(prestamos) == 0 {
		fmt.Println("No hay prestamos activos.")
		return
	}

	for _, prestamo := range prestamos {
		fmt.Println(
			"ID:", prestamo.ID(),
			"| Libro ID:", prestamo.LibroID(),
			"| Usuario ID:", prestamo.UsuarioID(),
			"| Activo:", prestamo.Activo(),
		)
	}
}

func devolverPrestamo(reader *bufio.Reader, loanService *loans.LoanService) {
	fmt.Println("\n--- Devolver prestamo ---")

	prestamoID, err := leerEntero(reader, "ID del prestamo: ")
	if err != nil {
		fmt.Println("Error: el ID del prestamo debe ser un numero.")
		return
	}

	err = loanService.DevolverPrestamo(prestamoID)
	if err != nil {
		fmt.Println("Error al devoler prestamo:", err)
		return
	}

	fmt.Println("Prestamo devuelto correctamente.")
}
