package api

import (
	"encoding/json"
	"net/http"
)

type loanResponse struct {
	ID              int    `json:"id"`
	LibroID         int    `json:"libro_id"`
	LibroTitulo     string `json:"libro_titulo"`
	UsuarioID       int    `json:"usuario_id"`
	UsuarioNombre   string `json:"usuario_nombre"`
	FechaPrestamo   string `json:"fecha_prestamo"`
	FechaDevolucion string `json:"fecha_devolucion"`
	Activo          bool   `json:"activo"`
}

type createLoanRequest struct {
	LibroID   int `json:"libro_id"`
	UsuarioID int `json:"usuario_id"`
}

type returnLoanRequest struct {
	PrestamoID int `json:"prestamo_id"`
}

func (s *Server) registerLoanRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/loans", s.handleLoans)
	mux.HandleFunc("/api/loans/active", s.handleActiveLoans)
	mux.HandleFunc("/api/loans/history", s.handleLoanHistory)
	mux.HandleFunc("/api/loans/return", s.handleReturnLoan)
}

func (s *Server) handleLoans(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.handleCreateLoan(w, r)
	default:
		writeError(w, http.StatusMethodNotAllowed, "Metodo no permitido")
	}
}

func (s *Server) handleCreateLoan(w http.ResponseWriter, r *http.Request) {
	var request createLoanRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeError(w, http.StatusBadRequest, "JSON invalido")
		return
	}

	prestamo, err := s.loanService.RegistrarPrestamo(request.LibroID, request.UsuarioID)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	response, err := s.buildLoanResponse(prestamo.ID())
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, response)
}

func (s *Server) handleActiveLoans(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "Metodo no permitido")
		return
	}

	prestamos := s.loanService.ListarPrestamosActivos()

	response := []loanResponse{}

	for _, prestamo := range prestamos {
		item, err := s.buildLoanResponse(prestamo.ID())
		if err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}

		response = append(response, item)
	}

	writeJSON(w, http.StatusOK, response)
}

func (s *Server) handleLoanHistory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "Metodo no permitido")
		return
	}

	prestamos := s.loanService.ListarPrestamos()

	response := []loanResponse{}

	for _, prestamo := range prestamos {
		item, err := s.buildLoanResponse(prestamo.ID())
		if err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}

		response = append(response, item)
	}

	writeJSON(w, http.StatusOK, response)
}

func (s *Server) handleReturnLoan(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "Metodo no permitido")
		return
	}

	var request returnLoanRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeError(w, http.StatusBadRequest, "JSON invalido")
		return
	}

	err = s.loanService.DevolverPrestamo(request.PrestamoID)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	response := map[string]string{
		"message": "Prestamo devuelto correctamente",
	}

	writeJSON(w, http.StatusOK, response)
}

func (s *Server) buildLoanResponse(prestamoID int) (loanResponse, error) {
	prestamo, err := s.loanService.BuscarPorID(prestamoID)
	if err != nil {
		return loanResponse{}, err
	}

	libro, err := s.bookService.BuscarPorID(prestamo.LibroID())
	if err != nil {
		return loanResponse{}, err
	}

	usuario, err := s.userService.BuscarPorID(prestamo.UsuarioID())
	if err != nil {
		return loanResponse{}, err
	}

	fechaPrestamo := prestamo.FechaPrestamo().Format("2006-01-02 15:04")

	fechaDevolucion := "Pendiente"
	if prestamo.FechaDevolucion() != nil {
		fechaDevolucion = prestamo.FechaDevolucion().Format("2006-01-02 15:04")
	}

	return loanResponse{
		ID:              prestamo.ID(),
		LibroID:         prestamo.LibroID(),
		LibroTitulo:     libro.Titulo(),
		UsuarioID:       prestamo.UsuarioID(),
		UsuarioNombre:   usuario.Nombre(),
		FechaPrestamo:   fechaPrestamo,
		FechaDevolucion: fechaDevolucion,
		Activo:          prestamo.Activo(),
	}, nil
}


