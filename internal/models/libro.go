package models

import (
	"errors"
	"strings"
)

type Libro struct {
	id         int
	titulo     string
	autor      string
	categoria  string
	anio       int
	disponible bool
}

func NuevoLibro(id int, titulo string, autor string, categoria string, anio int) (*Libro, error) {
	titulo = strings.TrimSpace(titulo)
	autor = strings.TrimSpace(autor)
	categoria = strings.TrimSpace(categoria)

	if id <= 0 {
		return nil, errors.New("El ID del libro debe ser mayor a cero")
	}
	if titulo == "" {
		return nil, errors.New("El titulo del libro no puede esatr vacio")
	}
	if autor == "" {
		return nil, errors.New("El autor del libro no puede estar vacio")
	}
	if categoria == "" {
		return nil, errors.New("La categoriad del libro no puede estar vacia")
	}
	if anio <= 0 {
		return nil, errors.New("El anio del libro debe ser valido")
	}

	libro := &Libro{
		id:         id,
		titulo:     titulo,
		autor:      autor,
		categoria:  categoria,
		anio:       anio,
		disponible: true,
	}

	return libro, nil
}

func (l *Libro) ID() int {
	return l.id
}

func (l *Libro) Titulo() string {
	return l.titulo
}

func (l *Libro) Autor() string {
	return l.autor
}

func (l *Libro) Categoria() string {
	return l.categoria
}

func (l *Libro) Anio() int {
	return l.anio
}

func (l *Libro) Disponible() bool {
	return l.disponible
}

func (l *Libro) Prestar() error {
	if !l.disponible {
		return errors.New("El libro no esta disponible")
	}

	l.disponible = false
	return nil
}

func (l *Libro) Devolver() {
	l.disponible = true
}

func (l *Libro) CambiarTitulo(titulo string) error {
	titulo = strings.TrimSpace(titulo)

	if titulo == "" {
		return errors.New("El titulo no puede estar vacio")
	}

	l.titulo = titulo
	return nil
}

func (l *Libro) CambiarAutor(autor string) error {
	autor = strings.TrimSpace(autor)

	if autor == "" {
		return errors.New("El autor no puede estar vacio")
	}

	l.autor = autor
	return nil
}

func (l *Libro) CambiarCategoria(categoria string) error {
	categoria = strings.TrimSpace(categoria)

	if categoria == "" {
		return errors.New("La categoria no puede estar vacia")
	}

	l.categoria = categoria
	return nil
}
