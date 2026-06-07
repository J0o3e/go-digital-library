package models

import (
	"errors"
	"time"
)

type Prestamo struct {
	id              int
	libroID         int
	usuarioID       int
	fechaPrestamo   time.Time
	fechaDevolucion *time.Time
	activo          bool
}

func NuevoPrestamo(id int, libroID int, usuarioID int) (*Prestamo, error) {
	if id <= 0 {
		return nil, errors.New("El ID del prestamo debe ser mayor a cero")
	}

	if libroID <= 0 {
		return nil, errors.New("El ID del libro debe ser mayor a cero")
	}

	if usuarioID <= 0 {
		return nil, errors.New("El ID del usuario deber ser mayor a cero")
	}

	prestamo := &Prestamo{
		id:              id,
		libroID:         libroID,
		usuarioID:       usuarioID,
		fechaPrestamo:   time.Now(),
		fechaDevolucion: nil,
		activo:          true,
	}

	return prestamo, nil
}

func (p *Prestamo) ID() int {
	return p.id
}

func (p *Prestamo) LibroID() int {
	return p.libroID
}

func (p *Prestamo) UsuarioID() int {
	return p.usuarioID
}

func (p *Prestamo) FechaPrestamo() time.Time {
	return p.fechaPrestamo
}

func (p *Prestamo) FechaDevolucion() *time.Time {
	return p.fechaDevolucion
}

func (p *Prestamo) Activo() bool {
	return p.activo
}

func (p *Prestamo) Devolver() error {
	if !p.activo {
		return errors.New("El prestamo ya fue devuelto")
	}

	fechaActual := time.Now()
	p.fechaDevolucion = &fechaActual
	p.activo = false

	return nil
}
