package books

import (
	"errors"
	"go-digital-library/internal/models"
	"strings"
)

type BookService struct {
	libros      []*models.Libro
	indicePorID map[int]*models.Libro
	siguienteID int
}

func NewBookService() *BookService {
	return &BookService{
		libros:      []*models.Libro{},
		indicePorID: make(map[int]*models.Libro),
		siguienteID: 1,
	}
}

func (s *BookService) RegistrarLibro(titulo string, autor string, categoria string, anio int) (*models.Libro, error) {
	libro, err := models.NuevoLibro(s.siguienteID, titulo, autor, categoria, anio)
	if err != nil {
		return nil, err
	}

	s.libros = append(s.libros, libro)
	s.indicePorID[libro.ID()] = libro
	s.siguienteID++

	return libro, nil
}

func (s *BookService) ListarLibros() []*models.Libro {
	return s.libros
}

func (s *BookService) BuscarPorID(id int) (*models.Libro, error) {
	libro, existe := s.indicePorID[id]
	if !existe {
		return nil, errors.New("Libro no encontrado")
	}

	return libro, nil
}

func (s *BookService) BuscarPorTexto(texto string) []*models.Libro {
	texto = strings.ToLower(strings.TrimSpace(texto))
	resultados := []*models.Libro{}

	for _, libro := range s.libros {
		titulo := strings.ToLower(libro.Titulo())
		autor := strings.ToLower(libro.Autor())
		categoria := strings.ToLower(libro.Categoria())

		if strings.Contains(titulo, texto) ||
			strings.Contains(autor, texto) ||
			strings.Contains(categoria, texto) {
			resultados = append(resultados, libro)
		}
	}

	return resultados
}

func (s *BookService) PrestarLibro(id int) error {
	libro, err := s.BuscarPorID(id)
	if err != nil {
		return err
	}

	return libro.Prestar()
}

func (s *BookService) DevolverLibro(id int) error {
	libro, err := s.BuscarPorID(id)
	if err != nil {
		return err
	}

	libro.Devolver()
	return nil
}
