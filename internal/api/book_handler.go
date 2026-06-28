package api

import (
	"encoding/json"
	"net/http"
)

type bookResponse struct {
	ID         int    `json:"id"`
	Titulo     string `json:"titulo"`
	Autor      string `json:"autor"`
	Categoria  string `json:"categoria"`
	Anio       int    `json:"anio"`
	Disponible bool   `json:"disponible"`
}

type createBookRequest struct {
	Titulo    string `json:"titulo"`
	Autor     string `json:"autor"`
	Categoria string `json:"categoria"`
	Anio      int    `json:"anio"`
}

func (s *Server) registerBookRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/books", s.handleBooks)
	mux.HandleFunc("/api/books/search", s.handleSearchBooks)
}

func (s *Server) handleBooks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleListBooks(w, r)
	case http.MethodPost:
		s.handleCreateBook(w, r)
	default:
		writeError(w, http.StatusMethodNotAllowed, "Metodo no permitido")
	}
}

func (s *Server) handleListBooks(w http.ResponseWriter, r *http.Request) {
	libros := s.bookService.ListarLibros()

	response := []bookResponse{}

	for _, libro := range libros {
		response = append(response, bookResponse{
			ID:         libro.ID(),
			Titulo:     libro.Titulo(),
			Autor:      libro.Autor(),
			Categoria:  libro.Categoria(),
			Anio:       libro.Anio(),
			Disponible: libro.Disponible(),
		})
	}

	writeJSON(w, http.StatusOK, response)
}

func (s *Server) handleCreateBook(w http.ResponseWriter, r *http.Request) {
	var request createBookRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeError(w, http.StatusBadRequest, "JSON invalido")
		return
	}

	libro, err := s.bookService.RegistrarLibro(
		request.Titulo,
		request.Autor,
		request.Categoria,
		request.Anio,
	)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	response := bookResponse{
		ID:         libro.ID(),
		Titulo:     libro.Titulo(),
		Autor:      libro.Autor(),
		Categoria:  libro.Categoria(),
		Anio:       libro.Anio(),
		Disponible: libro.Disponible(),
	}

	writeJSON(w, http.StatusCreated, response)
}

func (s *Server) handleSearchBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "Metodo no permitido")
		return
	}

	query := r.URL.Query().Get("q")
	if query == "" {
		writeError(w, http.StatusBadRequest, "Debe enviar el parametro q")
		return
	}

	libros := s.bookService.BuscarPorTexto(query)

	response := []bookResponse{}

	for _, libro := range libros {
		response = append(response, bookResponse{
			ID:         libro.ID(),
			Titulo:     libro.Titulo(),
			Autor:      libro.Autor(),
			Categoria:  libro.Categoria(),
			Anio:       libro.Anio(),
			Disponible: libro.Disponible(),
		})
	}

	writeJSON(w, http.StatusOK, response)
}

