package users

import (
	"errors"
	"strings"

	"go-digital-library/internal/models"
	"go-digital-library/internal/repositories"
)

type UserService struct {
	usuarios        []*models.Usuario
	indicePorID     map[int]*models.Usuario
	indicePorCorreo map[string]*models.Usuario
	siguienteID     int
	userRepository  *repositories.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		usuarios:        []*models.Usuario{},
		indicePorID:     make(map[int]*models.Usuario),
		indicePorCorreo: make(map[string]*models.Usuario),
		siguienteID:     1,
		userRepository:  nil,
	}
}

func NewUserDBService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{
		usuarios:        []*models.Usuario{},
		indicePorID:     make(map[int]*models.Usuario),
		indicePorCorreo: make(map[string]*models.Usuario),
		siguienteID:     1,
		userRepository:  userRepository,
	}
}

func (s *UserService) RegistrarUsuario(nombre string, correo string) (*models.Usuario, error) {
	correo = strings.ToLower(strings.TrimSpace(correo))

	if s.userRepository != nil {
		_, err := s.userRepository.BuscarPorCorreo(correo)
		if err == nil {
			return nil, errors.New("Ya existe un usuario con este correo")
		}

		return s.userRepository.Crear(nombre, correo)
	}

	if _, existe := s.indicePorCorreo[correo]; existe {
		return nil, errors.New("Ya existe un usuario con este correo")
	}

	usuario, err := models.NuevoUsuario(s.siguienteID, nombre, correo)
	if err != nil {
		return nil, err
	}

	s.usuarios = append(s.usuarios, usuario)
	s.indicePorID[usuario.ID()] = usuario
	s.indicePorCorreo[usuario.Correo()] = usuario
	s.siguienteID++

	return usuario, nil
}

func (s *UserService) ListarUsuarios() []*models.Usuario {
	if s.userRepository != nil {
		usuarios, err := s.userRepository.ListarTodos()
		if err != nil {
			return []*models.Usuario{}
		}

		return usuarios
	}

	return s.usuarios
}

func (s *UserService) BuscarPorID(id int) (*models.Usuario, error) {
	if s.userRepository != nil {
		return s.userRepository.BuscarPorID(id)
	}

	usuario, existe := s.indicePorID[id]
	if !existe {
		return nil, errors.New("Usuario no encontrado")
	}

	return usuario, nil
}

func (s *UserService) BuscarPorCorreo(correo string) (*models.Usuario, error) {
	correo = strings.ToLower(strings.TrimSpace(correo))

	if s.userRepository != nil {
		return s.userRepository.BuscarPorCorreo(correo)
	}

	usuario, existe := s.indicePorCorreo[correo]
	if !existe {
		return nil, errors.New("Usuario no encontrado")
	}

	return usuario, nil
}

func (s *UserService) BuscarPorTexto(texto string) []*models.Usuario {
	if s.userRepository != nil {
		usuarios, err := s.userRepository.BuscarPorTexto(texto)
		if err != nil {
			return []*models.Usuario{}
		}

		return usuarios
	}

	texto = strings.ToLower(strings.TrimSpace(texto))
	resultados := []*models.Usuario{}

	for _, usuario := range s.usuarios {
		nombre := strings.ToLower(usuario.Nombre())
		correo := strings.ToLower(usuario.Correo())

		if strings.Contains(nombre, texto) || strings.Contains(correo, texto) {
			resultados = append(resultados, usuario)
		}
	}

	return resultados
}

func (s *UserService) DesactivarUsuario(id int) error {
	if s.userRepository != nil {
		return s.userRepository.ActualizarEstado(id, false)
	}

	usuario, err := s.BuscarPorID(id)
	if err != nil {
		return err
	}

	usuario.Desactivar()
	return nil
}

func (s *UserService) ActivarUsuario(id int) error {
	if s.userRepository != nil {
		return s.userRepository.ActualizarEstado(id, true)
	}

	usuario, err := s.BuscarPorID(id)
	if err != nil {
		return err
	}

	usuario.Activar()
	return nil
}
