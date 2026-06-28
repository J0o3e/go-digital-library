package api

import (
	"encoding/json"
	"net/http"
)

type userResponse struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Correo string `json:"correo"`
	Activo bool   `json:"activo"`
}

type createUserRequest struct {
	Nombre string `json:"nombre"`
	Correo string `json:"correo"`
}

func (s *Server) registerUserRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/users", s.handleUsers)
	mux.HandleFunc("/api/users/search", s.handleSearchUsers)
}

func (s *Server) handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleListUsers(w, r)
	case http.MethodPost:
		s.handleCreateUser(w, r)
	default:
		writeError(w, http.StatusMethodNotAllowed, "Metodo no permitido")
	}
}

func (s *Server) handleListUsers(w http.ResponseWriter, r *http.Request) {
	usuarios := s.userService.ListarUsuarios()

	response := []userResponse{}

	for _, usuario := range usuarios {
		response = append(response, userResponse{
			ID:     usuario.ID(),
			Nombre: usuario.Nombre(),
			Correo: usuario.Correo(),
			Activo: usuario.Activo(),
		})
	}

	writeJSON(w, http.StatusOK, response)
}

func (s *Server) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var request createUserRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeError(w, http.StatusBadRequest, "JSON invalido")
		return
	}

	usuario, err := s.userService.RegistrarUsuario(request.Nombre, request.Correo)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	response := userResponse{
		ID:     usuario.ID(),
		Nombre: usuario.Nombre(),
		Correo: usuario.Correo(),
		Activo: usuario.Activo(),
	}

	writeJSON(w, http.StatusCreated, response)
}

func (s *Server) handleSearchUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "Metodo no permitido")
		return
	}

	query := r.URL.Query().Get("q")
	if query == "" {
		writeError(w, http.StatusBadRequest, "Debe enviar el parametro q")
		return
	}

	usuarios := s.userService.BuscarPorTexto(query)

	response := []userResponse{}

	for _, usuario := range usuarios {
		response = append(response, userResponse{
			ID:     usuario.ID(),
			Nombre: usuario.Nombre(),
			Correo: usuario.Correo(),
			Activo: usuario.Activo(),
		})
	}

	writeJSON(w, http.StatusOK, response)
}
