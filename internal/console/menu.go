package console

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

type Menu struct {
	reader      *bufio.Reader
	bookService *books.BookService
	userService *users.UserService
	loanService *loans.LoanService
}

func NewMenu(
	bookService *books.BookService,
	userService *users.UserService,
	loanService *loans.LoanService,
) *Menu {
	return &Menu{
		reader:      bufio.NewReader(os.Stdin),
		bookService: bookService,
		userService: userService,
		loanService: loanService,
	}
}

func (m *Menu) Ejecutar() {
	for {
		m.mostrarMenu()

		opcionTexto := m.leerTexto("Seleccione una opcion: ")
		opcion, err := strconv.Atoi(opcionTexto)
		if err != nil {
			fmt.Println("Opcion invalida. Debe ingresar un numero.")
			continue
		}

		switch opcion {
		case 1:
			m.registrarLibro()
		case 2:
			m.listarLibros()
		case 3:
			m.registrarUsuario()
		case 4:
			m.listarUsuarios()
		case 5:
			m.registrarPrestamo()
		case 6:
			m.listarPrestamosActivos()
		case 7:
			m.devolverPrestamo()
		case 8:
			m.buscarLibros()
		case 9:
			m.buscarUsuarios()
		case 10:
			fmt.Println("Saliendo del sistema...")
			return
		default:
			fmt.Println("Opcion no valida.")
		}
	}
}

func (m *Menu) mostrarMenu() {
	fmt.Println("\n===== GO DIGITAL LIBRARY =====")
	fmt.Println("1. Registrar libro")
	fmt.Println("2. Listar libros")
	fmt.Println("3. Registrar usuario")
	fmt.Println("4. Listar usuarios")
	fmt.Println("5. Registrar prestamo")
	fmt.Println("6. Listar prestamos activos")
	fmt.Println("7. Devolver prestamo")
	fmt.Println("8. Buscar libros")
	fmt.Println("9. Buscar usuarios")
	fmt.Println("10. Salir")
}

func (m *Menu) leerTexto(mensaje string) string {
	fmt.Print(mensaje)
	texto, _ := m.reader.ReadString('\n')
	return strings.TrimSpace(texto)
}

func (m *Menu) leerEntero(mensaje string) (int, error) {
	texto := m.leerTexto(mensaje)
	numero, err := strconv.Atoi(texto)
	if err != nil {
		return 0, err
	}

	return numero, nil
}

func (m *Menu) registrarLibro() {
	fmt.Println("\n--- Registrar libro ---")

	titulo := m.leerTexto("Titulo: ")
	autor := m.leerTexto("Autor: ")
	categoria := m.leerTexto("Categoria: ")

	anio, err := m.leerEntero("Anio: ")
	if err != nil {
		fmt.Println("Error: el anio debe ser un numero.")
		return
	}

	libro, err := m.bookService.RegistrarLibro(titulo, autor, categoria, anio)
	if err != nil {
		fmt.Println("Error al registrar libro:", err)
		return
	}

	fmt.Println("Libro registrado correctamente.")
	fmt.Println("ID:", libro.ID(), "- Titulo:", libro.Titulo())
}

func (m *Menu) listarLibros() {
	fmt.Println("\n--- Lista de libros ---")

	libros := m.bookService.ListarLibros()
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

func (m *Menu) buscarLibros() {
	fmt.Println("\n--- Buscar libros ---")

	texto := m.leerTexto("Ingrese titulo, autor o categoria: ")
	resultados := m.bookService.BuscarPorTexto(texto)

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

func (m *Menu) registrarUsuario() {
	fmt.Println("\n--- Registrar usuario ---")

	nombre := m.leerTexto("Nombre: ")
	correo := m.leerTexto("Correo: ")

	usuario, err := m.userService.RegistrarUsuario(nombre, correo)
	if err != nil {
		fmt.Println("Error al registrar usuario:", err)
		return
	}

	fmt.Println("Usuario registrado correctamente.")
	fmt.Println("ID:", usuario.ID(), "- Nombre:", usuario.Nombre())
}

func (m *Menu) listarUsuarios() {
	fmt.Println("\n--- Lista de usuarios ---")

	usuarios := m.userService.ListarUsuarios()
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

func (m *Menu) buscarUsuarios() {
	fmt.Println("\n--- Buscar usuarios ---")

	texto := m.leerTexto("Ingrese nombre o correo: ")
	resultados := m.userService.BuscarPorTexto(texto)

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

func (m *Menu) registrarPrestamo() {
	fmt.Println("\n--- Registrar prestamo ---")

	libroID, err := m.leerEntero("ID del libro: ")
	if err != nil {
		fmt.Println("Error: el ID del libro debe ser un numero.")
		return
	}

	usuarioID, err := m.leerEntero("ID del usuario: ")
	if err != nil {
		fmt.Println("Error: el ID del usuario debe ser un numero.")
		return
	}

	prestamo, err := m.loanService.RegistrarPrestamo(libroID, usuarioID)
	if err != nil {
		fmt.Println("Error al registrar prestamo:", err)
		return
	}

	fmt.Println("Prestamo registrado correctamente.")
	fmt.Println("Prestamo ID:", prestamo.ID())
}

func (m *Menu) listarPrestamosActivos() {
	fmt.Println("\n--- Prestamos activos ---")

	prestamos := m.loanService.ListarPrestamosActivos()
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

func (m *Menu) devolverPrestamo() {
	fmt.Println("\n--- Devolver prestamo ---")

	prestamoID, err := m.leerEntero("ID del prestamo: ")
	if err != nil {
		fmt.Println("Error: el ID del prestamo debe ser un numero.")
		return
	}

	err = m.loanService.DevolverPrestamo(prestamoID)
	if err != nil {
		fmt.Println("Error al devolver prestamo:", err)
		return
	}

	fmt.Println("Prestamo devuelto correctamente.")
}
