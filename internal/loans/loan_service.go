package loans

import (
	"errors"

	"go-digital-library/internal/books"
	"go-digital-library/internal/models"
	"go-digital-library/internal/users"
)

type LoanService struct {
	prestamos              []*models.Prestamo
	indicePorID            map[int]*models.Prestamo
	prestamoActivoPorLibro map[int]*models.Prestamo
	siguienteID            int
	bookService            *books.BookService
	userService            *users.UserService
}

func NewLoanService(bookService *books.BookService, userService *users.UserService) *LoanService {
	return &LoanService{
		prestamos:              []*models.Prestamo{},
		indicePorID:            make(map[int]*models.Prestamo),
		prestamoActivoPorLibro: make(map[int]*models.Prestamo),
		siguienteID:            1,
		bookService:            bookService,
		userService:            userService,
	}
}

func (s *LoanService) RegistrarPrestamo(libroID int, usuarioID int) (*models.Prestamo, error) {
	libro, err := s.bookService.BuscarPorID(libroID)
	if err != nil {
		return nil, err
	}

	usuario, err := s.userService.BuscarPorID(usuarioID)
	if err != nil {
		return nil, err
	}

	if !usuario.Activo() {
		return nil, errors.New("El usuario no esta activo")
	}

	if !libro.Disponible() {
		return nil, errors.New("El libro no esta disponible")
	}

	prestamo, err := models.NuevoPrestamo(s.siguienteID, libro.ID(), usuario.ID())
	if err != nil {
		return nil, err
	}

	err = s.bookService.PrestarLibro(libro.ID())
	if err != nil {
		return nil, err
	}

	s.prestamos = append(s.prestamos, prestamo)
	s.indicePorID[prestamo.ID()] = prestamo
	s.prestamoActivoPorLibro[libro.ID()] = prestamo
	s.siguienteID++

	return prestamo, nil

}

func (s *LoanService) ListarPrestamos() []*models.Prestamo {
	return s.prestamos
}

func (s *LoanService) ListarPrestamosActivos() []*models.Prestamo {
	activos := []*models.Prestamo{}

	for _, prestamo := range s.prestamos {
		if prestamo.Activo() {
			activos = append(activos, prestamo)
		}
	}

	return activos
}

func (s *LoanService) BuscarPorID(id int) (*models.Prestamo, error) {
	prestamo, existe := s.indicePorID[id]
	if !existe {
		return nil, errors.New("Prestamo no encontrado")
	}

	return prestamo, nil
}

func (s *LoanService) DevolverPrestamo(prestamoID int) error {
	prestamo, err := s.BuscarPorID(prestamoID)
	if err != nil {
		return err
	}

	if !prestamo.Activo() {
		return errors.New("El prestamo ya fue devuelto")
	}

	err = prestamo.Devolver()
	if err != nil {
		return err
	}

	err = s.bookService.DevolverLibro(prestamo.LibroID())
	if err != nil {
		return err
	}

	delete(s.prestamoActivoPorLibro, prestamo.LibroID())

	return nil
}

func (s *LoanService) BuscarPrestamoActivoPorLibro(libroID int) (*models.Prestamo, error) {
	prestamo, existe := s.prestamoActivoPorLibro[libroID]
	if !existe {
		return nil, errors.New("No existe un prestamo activo para este libro")
	}

	return prestamo, nil
}
