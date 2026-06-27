package repositories

import (
	"database/sql"
	"errors"
	"time"

	"go-digital-library/internal/models"
)

type LoanRepository struct {
	db *sql.DB
}

func NewLoanRepository(db *sql.DB) *LoanRepository {
	return &LoanRepository{
		db: db,
	}
}

func (r *LoanRepository) Crear(libroID int, usuarioID int) (*models.Prestamo, error) {
	query := `
		INSERT INTO prestamos (libro_id, usuario_id, fecha_prestamo, fecha_devolucion, activo)
		VALUES ($1, $2, CURRENT_TIMESTAMP, NULL, true)
		RETURNING id, libro_id, usuario_id, fecha_prestamo, fecha_devolucion, activo
	`

	row := r.db.QueryRow(query, libroID, usuarioID)

	prestamo, err := escanearPrestamo(row)
	if err != nil {
		return nil, err
	}

	return prestamo, nil
}

func (r *LoanRepository) ListarTodos() ([]*models.Prestamo, error) {
	query := `
		SELECT id, libro_id, usuario_id, fecha_prestamo, fecha_devolucion, activo
		FROM prestamos
		ORDER BY id
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	prestamos := []*models.Prestamo{}

	for rows.Next() {
		prestamo, err := escanearPrestamo(rows)
		if err != nil {
			return nil, err
		}

		prestamos = append(prestamos, prestamo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return prestamos, nil
}

func (r *LoanRepository) ListarActivos() ([]*models.Prestamo, error) {
	query := `
		SELECT id, libro_id, usuario_id, fecha_prestamo, fecha_devolucion, activo
		FROM prestamos
		WHERE activo = true
		ORDER BY id
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	prestamos := []*models.Prestamo{}

	for rows.Next() {
		prestamo, err := escanearPrestamo(rows)
		if err != nil {
			return nil, err
		}

		prestamos = append(prestamos, prestamo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return prestamos, nil
}

func (r *LoanRepository) BuscarPorID(id int) (*models.Prestamo, error) {
	query := `
		SELECT id, libro_id, usuario_id, fecha_prestamo, fecha_devolucion, activo
		FROM prestamos
		WHERE id = $1
	`

	row := r.db.QueryRow(query, id)

	prestamo, err := escanearPrestamo(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("Prestamo no encontrado")
		}

		return nil, err
	}

	return prestamo, nil
}

func (r *LoanRepository) BuscarActivoPorLibro(LibroID int) (*models.Prestamo, error) {
	query := `
		SELECT id, libro_id, usuario_id, fecha_prestamo, fecha_devolucion, activo
		FROM prestamos
		WHERE libro_id = $1
		  AND activo = true
		LIMIT 1
	`

	row := r.db.QueryRow(query, LibroID)

	prestamo, err := escanearPrestamo(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("No existe un prestamo activo para ese libro")
		}

		return nil, err
	}

	return prestamo, nil
}

func (r *LoanRepository) MarcarDevuelto(id int) error {
	query := `
		UPDATE prestamos
		SET activo = false,
		    fecha_devolucion = CURRENT_TIMESTAMP
		WHERE id = $1
		  AND activo = true
	`

	resultado, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	filasAfectadas, err := resultado.RowsAffected()
	if err != nil {
		return err
	}

	if filasAfectadas == 0 {
		return errors.New("Prestamo no encontrado o ya fue devuelto")
	}

	return nil
}

type scannerPrestamo interface {
	Scan(dest ...any) error
}

func escanearPrestamo(scanner scannerPrestamo) (*models.Prestamo, error) {
	var id int
	var libroID int
	var usuarioID int
	var fechaPrestamo time.Time
	var fechaDevolucion sql.NullTime
	var activo bool

	err := scanner.Scan(
		&id,
		&libroID,
		&usuarioID,
		&fechaPrestamo,
		&fechaDevolucion,
		&activo,
	)
	if err != nil {
		return nil, err
	}

	var fechaDevolucionPtr *time.Time
	if fechaDevolucion.Valid {
		fechaDevolucionPtr = &fechaDevolucion.Time
	}

	prestamo, err := models.NuevoPrestamoDesdeBD(
		id,
		libroID,
		usuarioID,
		fechaPrestamo,
		fechaDevolucionPtr,
		activo,
	)
	if err != nil {
		return nil, err
	}

	return prestamo, nil
}
