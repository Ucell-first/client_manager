package httpcms

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Ucell/client_manager/storage"
)

type Handler struct {
	storage storage.IStorage
	tmpl    *template.Template
}

func NewHandler(storage storage.IStorage) *Handler {
	// Template fayl yo'lini tekshirish
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

	log.Println("Template fayllar muvaffaqiyatli yuklandi")

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
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Root route called: %s", r.URL.Path)
		h.listUsers(w, r)
	})
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Users route called: %s", r.URL.Path)
		h.listUsers(w, r)
	})
	mux.HandleFunc("/user/view", h.viewUser)
	mux.HandleFunc("/user/new", h.newUserForm)
	mux.HandleFunc("/user/create", h.createUser)
	mux.HandleFunc("/user/edit", h.editUserForm)
	mux.HandleFunc("/user/update", h.updateUser)
	mux.HandleFunc("/user/delete", h.deleteUser)

	log.Println("Routes configured successfully")
	return mux
}
