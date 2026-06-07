package models

import (
	"errors"
	"strings"
)

type Usuario struct {
	id     int
	nombre string
	correo string
	activo bool
}

func NuevoUsuario(id int, nombre string, correo string) (*Usuario, error) {
	nombre = strings.TrimSpace(nombre)
	correo = strings.TrimSpace(correo)

	if id <= 0 {
		return nil, errors.New("El ID del usuario debe ser amyor a cero")
	}

	if nombre == "" {
		return nil, errors.New("El nombre del usuario no puede estar vacio")
	}

	if correo == "" {
		return nil, errors.New("El correo del usuario no puede estar vacio")
	}

	if !strings.Contains(correo, "@") {
		return nil, errors.New("El correo no es valido")
	}

	usuario := &Usuario{
		id:     id,
		nombre: nombre,
		correo: correo,
		activo: true,
	}

	return usuario, nil
}

func (u *Usuario) ID() int {
	return u.id
}

func (u *Usuario) Nombre() string {
	return u.nombre
}

func (u *Usuario) Correo() string {
	return u.correo
}

func (u *Usuario) Activo() bool {
	return u.activo
}

func (u *Usuario) Desactivar() {
	u.activo = false
}

func (u *Usuario) Activar() {
	u.activo = true
}

func (u *Usuario) CambiarNombre(nombre string) error {
	nombre = strings.TrimSpace(nombre)

	if nombre == "" {
		return errors.New("El nombre no puede estar vacio")
	}

	u.nombre = nombre
	return nil
}

func (u *Usuario) CambiarCorreo(correo string) error {
	correo = strings.TrimSpace(correo)

	if correo == "" {
		return errors.New("El correo no puede estar vacio")
	}

	if !strings.Contains(correo, "@") {
		return errors.New("El correo no es valido")
	}

	u.correo = correo
	return nil
}
