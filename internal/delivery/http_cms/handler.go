package httpcms

import (
	"html/template"
	"log"
	"net/http"

	"github.com/Ucell/client_manager/storage"
)

type Handler struct {
	storage storage.IStorage
	tmpl    *template.Template
}

func NewHandler(storage storage.IStorage) *Handler {
	// Template fayllarini to'g'ri load qilish
	tmpl, err := template.ParseGlob("internal/delivery/http_cms/templates/*.html")
	if err != nil {
		log.Printf("Template load xatoligi: %v", err)
		// Agar template topilmasa, oddiy template yaratamiz
		tmpl = template.New("default")
	}

	return &Handler{
		storage: storage,
		tmpl:    tmpl,
	}
}

func (h *Handler) Routes() *http.ServeMux {
	mux := http.NewServeMux()

	// Static files
	fs := http.FileServer(http.Dir("internal/delivery/http_cms/assets/"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// User routes
	mux.HandleFunc("/", h.listUsers)
	mux.HandleFunc("/users", h.listUsers)
	mux.HandleFunc("/user/view", h.viewUser)
	mux.HandleFunc("/user/new", h.newUserForm)
	mux.HandleFunc("/user/create", h.createUser)
	mux.HandleFunc("/user/edit", h.editUserForm)
	mux.HandleFunc("/user/update", h.updateUser)
	mux.HandleFunc("/user/delete", h.deleteUser)

	return mux
}
