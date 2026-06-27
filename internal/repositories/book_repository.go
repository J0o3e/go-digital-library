package repositories

import (
	"database/sql"
	"errors"

	"go-digital-library/internal/models"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (r *BookRepository) Crear(titulo string, autor string, categoria string, anio int) (*models.Libro, error) {
	query := `
		INSERT INTO libros (titulo, autor, categoria, anio, disponible)
		VALUES ($1, $2, $3, $4, true)
		RETURNING id, titulo, autor, categoria, anio, disponible
	`

	var id int
	var tituloDB string
	var autorDB string
	var categoriaDB string
	var anioDB int
	var disponibleDB bool

	err := r.db.QueryRow(query, titulo, autor, categoria, anio).Scan(
		&id,
		&tituloDB,
		&autorDB,
		&categoriaDB,
		&anioDB,
		&disponibleDB,
	)

	if err != nil {
		return nil, err
	}

	libro, err := models.NuevoLibro(id, tituloDB, autorDB, categoriaDB, anioDB)
	if err != nil {
		return nil, err
	}

	if !disponibleDB {
		err = libro.Prestar()
		if err != nil {
			return nil, err
		}
	}

	return libro, nil
}

func (r *BookRepository) ListarTodos() ([]*models.Libro, error) {
	query := `
		SELECT id, titulo, autor, categoria, anio, disponible
		FROM libros
		ORDER BY id
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	libros := []*models.Libro{}

	for rows.Next() {
		libro, err := escanearLibro(rows)
		if err != nil {
			return nil, err
		}

		libros = append(libros, libro)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return libros, nil
}

func (r *BookRepository) BuscarPorID(id int) (*models.Libro, error) {
	query := `
		SELECT id, titulo, autor, categoria, anio, disponible
		FROM libros
		WHERE id = $1
	`

	row := r.db.QueryRow(query, id)

	libro, err := escanearLibro(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("Libro no encntrado")
		}

		return nil, err
	}

	return libro, nil
}

func (r *BookRepository) BuscarPorTexto(texto string) ([]*models.Libro, error) {
	query := `
		SELECT id, titulo, autor, categoria, anio, disponible
		FROM libros
		WHERE LOWER(titulo) LIKE LOWER($1)
		   OR LOWER(autor) LIKE LOWER($1)
		   OR LOWER(categoria) LIKE LOWER($1)
		ORDER BY id
	`

	patron := "%" + texto + "%"

	rows, err := r.db.Query(query, patron)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	libros := []*models.Libro{}

	for rows.Next() {
		libro, err := escanearLibro(rows)
		if err != nil {
			return nil, err
		}

		libros = append(libros, libro)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return libros, nil
}

func (r *BookRepository) ActualizarDisponibilidad(id int, disponible bool) error {
	query := `
		UPDATE libros
		SET disponible = $1
		WHERE id = $2
	`

	resultado, err := r.db.Exec(query, disponible, id)
	if err != nil {
		return err
	}

	filasAfectadas, err := resultado.RowsAffected()
	if err != nil {
		return err
	}

	if filasAfectadas == 0 {
		return errors.New("Libro no encontrado")
	}

	return nil
}

type scannerLibro interface {
	Scan(dest ...any) error
}

func escanearLibro(scanner scannerLibro) (*models.Libro, error) {
	var id int
	var titulo string
	var autor string
	var categoria string
	var anio int
	var disponible bool

	err := scanner.Scan(
		&id,
		&titulo,
		&autor,
		&categoria,
		&anio,
		&disponible,
	)
	if err != nil {
		return nil, err
	}

	libro, err := models.NuevoLibro(id, titulo, autor, categoria, anio)
	if err != nil {
		return nil, err
	}

	if !disponible {
		err = libro.Prestar()
		if err != nil {
			return nil, err
		}
	}

	return libro, nil
}
