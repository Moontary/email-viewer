package handlers

import (
	"backViewer/pkg/entity"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
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

// NewEmailHandler
// Adds handlers app routes
func NewEmailHandler(repo EmailRepo) http.Handler {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	eh := EmailHandler{repo: repo}

	//router.Use(middleware.Logger)
	r.Get("/", eh.getEmails)
	r.Post("/", eh.createEmail)
	r.Delete("/{id}", eh.deleteEmail)
	return r
}

func (eh EmailHandler) getEmails(w http.ResponseWriter, r *http.Request) {
	emails, _ := eh.repo.GetAll(r.Context())
	if err := json.NewEncoder(w).Encode(emails); err != nil {
		return
	}

}

func (eh EmailHandler) createEmail(w http.ResponseWriter, r *http.Request) {

	email := entity.Email{}
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&email); err != nil {
		http.Error(w, "invalid user input", http.StatusBadRequest)
		return
	}
	validate := validator.New()
	if err := validate.Struct(email); err != nil {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// figure out
	_, err := eh.repo.GetEmailByAddress(ctx, email.Address)
	if err == nil {
		http.Error(w, fmt.Sprintf("Email with address: %s already exist", email.Address), http.StatusBadRequest)
		return
	}

	_, err = eh.repo.AddOne(ctx, &email)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusBadRequest)
		return
	}

	w.Write([]byte("Email created successfully"))
	w.WriteHeader(201)
}

func (eh EmailHandler) deleteEmail(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Email not found", http.StatusNotFound)
		return
	}
	_, err := eh.repo.GetEmailByID(ctx, id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Email with id: %s does not exist", id), http.StatusNotFound)
		return
	}

	err = eh.repo.RemoveOne(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte("Email deleted"))
	w.WriteHeader(200)
}
