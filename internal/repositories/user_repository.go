package repositories

import (
	"database/sql"
	"errors"

	"go-digital-library/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Crear(nombre string, correo string) (*models.Usuario, error) {
	query := `
		INSERT INTO usuarios (nombre, correo, activo)
		VALUES ($1, $2, true)
		RETURNING id, nombre, correo, activo
	`

	var id int
	var nombreDB string
	var correoDB string
	var activoDB bool

	err := r.db.QueryRow(query, nombre, correo).Scan(
		&id,
		&nombreDB,
		&correoDB,
		&activoDB,
	)
	if err != nil {
		return nil, err
	}

	usuario, err := models.NuevoUsuario(id, nombreDB, correoDB)
	if err != nil {
		return nil, err
	}

	if !activoDB {
		usuario.Desactivar()
	}

	return usuario, nil
}

func (r *UserRepository) ListarTodos() ([]*models.Usuario, error) {
	query := `
		SELECT id, nombre, correo, activo
		FROM usuarios
		ORDER BY id
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	usuarios := []*models.Usuario{}

	for rows.Next() {
		usuario, err := escanearUsuario(rows)
		if err != nil {
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return usuarios, nil
}

func (r *UserRepository) BuscarPorID(id int) (*models.Usuario, error) {
	query := `
		SELECT id, nombre, correo, activo
		FROM usuarios
		WHERE id = $1
	`

	row := r.db.QueryRow(query, id)

	usuario, err := escanearUsuario(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("Usuario no encontrado")
		}

		return nil, err
	}

	return usuario, nil
}

func (r *UserRepository) BuscarPorCorreo(correo string) (*models.Usuario, error) {
	query := `
		SELECT id, nombre, correo, activo
		FROM usuarios
		WHERE correo = $1
	`

	row := r.db.QueryRow(query, correo)

	usuario, err := escanearUsuario(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("Usuario no encontrado")
		}

		return nil, err
	}

	return usuario, nil
}

func (r *UserRepository) BuscarPorTexto(texto string) ([]*models.Usuario, error) {
	query := `
		SELECT id, nombre, correo, activo
		FROM usuarios
		WHERE LOWER(nombre) LIKE LOWER($1)
		   OR LOWER(correo) LIKE LOWER($1)
		ORDER BY id
	`

	patron := "%" + texto + "%"

	rows, err := r.db.Query(query, patron)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	usuarios := []*models.Usuario{}

	for rows.Next() {
		usuario, err := escanearUsuario(rows)
		if err != nil {
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return usuarios, nil
}

func (r *UserRepository) ActualizarEstado(id int, activo bool) error {
	query := `
		UPDATE usuarios
		SET activo = $1
		WHERE id = $2
	`

	resultado, err := r.db.Exec(query, activo, id)
	if err != nil {
		return err
	}

	filasAfectadas, err := resultado.RowsAffected()
	if err != nil {
		return err
	}

	if filasAfectadas == 0 {
		return errors.New("Usuario no encontrado")
	}

	return nil
}

type scannerUsuario interface {
	Scan(dest ...any) error
}

func escanearUsuario(scanner scannerUsuario) (*models.Usuario, error) {
	var id int
	var nombre string
	var correo string
	var activo bool

	err := scanner.Scan(
		&id,
		&nombre,
		&correo,
		&activo,
	)
	if err != nil {
		return nil, err
	}

	usuario, err := models.NuevoUsuario(id, nombre, correo)
	if err != nil {
		return nil, err
	}

	if !activo {
		usuario.Desactivar()
	}

	return usuario, nil
}
