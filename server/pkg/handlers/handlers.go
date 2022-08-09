package handlers

import (
	"backViewer/pkg/entity"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type EmailHandler struct {
	repo EmailRepo
}

type EmailRepo interface {
	GetAll(ctx context.Context) ([]entity.Email, error)
	AddOne(ctx context.Context, email *entity.Email) (*entity.Email, error)
	GetEmailByID(ctx context.Context, id string) (*entity.Email, error)
	GetEmailByAddress(ctx context.Context, address string) (*entity.Email, error)
	RemoveOne(ctx context.Context, id string) error
}

// NewEmailHandler Adds handlers app routes
func NewEmailHandler(repo EmailRepo) http.Handler {

	// router init
	r := chi.NewRouter()

	// CORS handling
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	eh := EmailHandler{repo: repo}

	// middleware and handlers init
	r.Use(middleware.Logger)
	r.Get("/", eh.getEmails)
	r.Post("/", eh.createEmail)
	r.Delete("/{id}", eh.deleteEmail)
	return r
}

// getEmails GET handler
func (eh EmailHandler) getEmails(w http.ResponseWriter, r *http.Request) {
	emails, _ := eh.repo.GetAll(r.Context())
	if err := json.NewEncoder(w).Encode(emails); err != nil {
		return
	}

}

// createEmail POST handler
func (eh EmailHandler) createEmail(w http.ResponseWriter, r *http.Request) {

	email := entity.Email{}
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&email); err != nil {
		http.Error(w, "Invalid user input", http.StatusBadRequest)
		return
	}

	// validator init
	validate := validator.New()
	if err := validate.Struct(email); err != nil {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// GetEmailByAddress checks if email with such address exist in repo
	// For duplicate email 400 code drop
	_, err := eh.repo.GetEmailByAddress(ctx, email.Address)
	if err == nil {
		http.Error(w, fmt.Sprintf("Email with address: %s already exist", email.Address), http.StatusBadRequest)
		return
	}

	// AddOne checks if adding goes right
	// Error code 400 dropped for incorrect value
	_, err = eh.repo.AddOne(ctx, &email)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusBadRequest)
		return
	}

	// check for correct encoding in ResponseWriter
	if err := json.NewEncoder(w).Encode(&email); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// deleteEmail DELETE handler
func (eh EmailHandler) deleteEmail(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	// URLParam used to check id for deletion
	// 404 dropped if error occurred
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Email not found", http.StatusNotFound)
		return
	}

	// GetEmailByID checks if email with such id exist
	// 404 dropped if error occurred
	_, err := eh.repo.GetEmailByID(ctx, id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Email with id: %s does not exist", id), http.StatusNotFound)
		return
	}

	// 400 dropped if error occurred
	err = eh.repo.RemoveOne(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte("Email deleted"))
	w.WriteHeader(http.StatusOK)
}
