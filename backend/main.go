package main

import (
	"backViewer/internal/model"
	"backViewer/internal/mongo"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"net/url"
)

type Email struct {
	Email string `json:"email,omitempty" validate:"required,email"`
}

var mr *mongo.NewEmailRepo

func main() {

	mongoDBConn := "mongodb://localhost:27017"
	mr = mongo.NewHandler(mongoDBConn)

	r := registerRoutes()
	log.Fatal(http.ListenAndServe(":3000", r))

}

func registerRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/", func(r chi.Router) {
		r.Get("/emails", getAllEmails)
		r.Post("/email", addEmail)
		r.Delete("/{email}", deleteEmail)
	})
	return r
}

func getAllEmails(w http.ResponseWriter, r *http.Request) {
	emails, _ := mr.GetAll(r.Context())
	err := json.NewEncoder(w).Encode(emails)
	if err != nil {
		return
	}

}

func (e *Email) validate() url.Values {
	err := url.Values{}

	if e.Email == "" {
		err.Add("Email", "Invalid email")
	}

	return err
}

func addEmail(w http.ResponseWriter, r *http.Request) {
	existingEmail := &model.Email{}

	var m model.Email

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		fmt.Println("invalid input", err)
		return
	}
	validate := validator.New()
	if err := validate.Struct(m); err != nil {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	err := mr.GetOne(existingEmail, bson.M{"email": m.Email})
	if err == nil {
		http.Error(w, fmt.Sprintf("Email with name: %s already exist", m.Email), 400)
		return
	}

	_, err = mr.AddOne(&m)
	if err != nil {
		http.Error(w, fmt.Sprint(err), 400)
		return
	}

	w.Write([]byte("Email created successfully"))
	w.WriteHeader(201)
}

func deleteEmail(w http.ResponseWriter, r *http.Request) {
	existingEmail := &model.Email{}

	Email := chi.URLParam(r, "Email")
	if Email == "" {
		http.Error(w, "Email not found", http.StatusNotFound)
		return
	}
	err := mr.GetOne(existingEmail, bson.M{"Email": Email})
	if err != nil {
		http.Error(w, fmt.Sprintf("Email with name: %s does not exist", Email), 400)
		return
	}
	_, err = mr.RemoveOne(bson.M{"Email": Email})
	if err != nil {
		http.Error(w, fmt.Sprint(err), 400)
		return
	}
	w.Write([]byte("Email deleted"))
	w.WriteHeader(200)
}
