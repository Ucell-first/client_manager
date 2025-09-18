package httpcms

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Ucell/client_manager/middleware"
	"github.com/Ucell/client_manager/storage"
)

type Handler struct {
	storage storage.IStorage
	tmpl    *template.Template
	auth    *middleware.AuthMiddleware
}

func NewHandler(storage storage.IStorage) *Handler {
	templatesPath := "internal/delivery/http_cms/templates/*.html"
	files, err := filepath.Glob(templatesPath)
	if err != nil {
		log.Fatalf("Template fayllarini izlashda xatolik: %v", err)
	}

	if len(files) == 0 {
		log.Fatal("Template fayllar topilmadi!")
	}

	tmpl, err := template.ParseGlob(templatesPath)
	if err != nil {
		log.Fatalf("Template fayllarini yuklashda xatolik: %v", err)
	}

	auth := middleware.NewAuthMiddleware()

	log.Println("Template fayllar va middleware muvaffaqiyatli yuklandi")

	return &Handler{
		storage: storage,
		tmpl:    tmpl,
		auth:    auth,
	}
}

func (h *Handler) Routes() *http.ServeMux {
	mux := http.NewServeMux()

	// Static files
	fs := http.FileServer(http.Dir("internal/delivery/http_cms/assets/"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Auth routes (middlewaresiz)
	mux.HandleFunc("/login", h.loginForm)
	mux.HandleFunc("/auth", h.loginAuth)

	// Barcha boshqa route'lar middleware orqali
	protectedMux := http.NewServeMux()

	// User routes
	protectedMux.HandleFunc("/", h.listUsers)
	protectedMux.HandleFunc("/users", h.listUsers)
	protectedMux.HandleFunc("/user/view", h.viewUser)
	protectedMux.HandleFunc("/user/new", h.newUserForm)
	protectedMux.HandleFunc("/user/create", h.createUser)
	protectedMux.HandleFunc("/user/edit", h.editUserForm)
	protectedMux.HandleFunc("/user/update", h.updateUser)
	protectedMux.HandleFunc("/user/delete", h.deleteUser)
	protectedMux.HandleFunc("/logout", h.logout)

	// Protected routes ni auth middleware bilan o'rash
	mux.Handle("/users", h.auth.RequireAuth(protectedMux))
	mux.Handle("/user/", h.auth.RequireAuth(protectedMux))
	mux.Handle("/logout", h.auth.RequireAuth(protectedMux))
	mux.Handle("/", h.auth.RequireAuth(protectedMux))

	log.Println("Routes muvaffaqiyatli sozlandi")
	return mux
}
