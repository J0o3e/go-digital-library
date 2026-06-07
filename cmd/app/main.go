package main

import (
	"fmt"
	"go-digital-library/internal/users"
)

func main() {
	userService := users.NewUserService()

	usuario1, err := userService.RegistrarUsuario("Jose", "jose@email.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	usuario2, err := userService.RegistrarUsuario("Ana", "ana@email.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Usuarios registrados:")
	fmt.Println(usuario1.ID(), "-", usuario1.Nombre(), "-", usuario1.Correo())
	fmt.Println(usuario2.ID(), "-", usuario2.Nombre(), "-", usuario2.Correo())

	fmt.Println("\nListado completo:")
	for _, usuario := range userService.ListarUsuarios() {
		fmt.Println(usuario.ID(), usuario.Nombre(), "-", usuario.Correo(), "- Activo:", usuario.Activo())
	}

	fmt.Println("\nBusqueda por correo:")
	usuarioBuscado, err := userService.BuscarPorCorreo("jose@email.com")
	if err != nil {
		fmt.Println("Error, err")
		return
	}
	fmt.Println("Encontrado:", usuarioBuscado.Nombre())

	fmt.Println("\nBusqueda por texto: ana")
	resultados := userService.BuscarPorTexto("ana")
	for _, usuario := range resultados {
		fmt.Println("Encontrado:", usuario.Nombre(), "-", usuario.Correo())
	}

	err = userService.DesactivarUsuario(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	usuarioDesactivado, err := userService.BuscarPorID(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nUsuario desactivado:", usuarioDesactivado.Nombre())
	fmt.Println("Activo:", usuarioDesactivado.Activo())

}
