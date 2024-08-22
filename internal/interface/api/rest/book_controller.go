package rest

import (
	"Or/Library/internal/application/book"
	bookDto "Or/Library/internal/interface/api/dtos/book"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type BookController struct {
	bookService *book.BookService
}

func NewBookController(bookService *book.BookService) *BookController {
	return &BookController{bookService}
}

func (c *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	var bookDTO bookDto.BookDTO
	err := json.NewDecoder(r.Body).Decode(&bookDTO)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	book := bookDto.MapToBookDomainFromDTO(&bookDTO)
	err = c.bookService.CreateBook(r.Context(), book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *BookController) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	book, err := c.bookService.GetBook(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	bookDTO := bookDto.MapToBookDTO(book)
	json.NewEncoder(w).Encode(bookDTO)
}

func (c *BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var bookDTO bookDto.BookDTO
	err := json.NewDecoder(r.Body).Decode(&bookDTO)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	b := bookDto.MapToBookDomainFromDTO(&bookDTO)
	b.ID = id // Ensure the ID is set for the update
	err = c.bookService.UpdateBook(r.Context(), b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := c.bookService.DeleteBook(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *BookController) SearchBooks(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")

	books, err := c.bookService.SearchBooks(r.Context(), query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bookDTOs := make([]bookDto.BookDTO, len(books))
	for i, b := range books {
		bookDTOs[i] = bookDto.MapToBookDTO(b)
	}
	json.NewEncoder(w).Encode(bookDTOs)
}
