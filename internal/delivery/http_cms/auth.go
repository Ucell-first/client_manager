package httpcms

import (
	"log"
	"net/http"
	"time"

	"github.com/Ucell/client_manager/auth"
)

func (h *Handler) loginForm(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "Admin Login",
	}

	if err := h.tmpl.ExecuteTemplate(w, "login.html", data); err != nil {
		log.Printf("Template xatolik: %v", err)
		http.Error(w, "Sahifani ko'rsatishda xatolik", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) loginAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Faqat POST so'rov", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Forma xatoligi", http.StatusBadRequest)
		return
	}

	login := r.FormValue("login")
	password := r.FormValue("password")

	admin, err := h.storage.Admin().Login(r.Context(), login, password)
	if err != nil {
		log.Printf("Login xatolik: %v", err)
		data := PageData{
			Title: "Admin Login",
			Error: "Login yoki parol noto'g'ri",
		}
		if err := h.tmpl.ExecuteTemplate(w, "login.html", data); err != nil {
			http.Error(w, "Template xatolik", http.StatusInternalServerError)
		}
		return
	}

	// Token yaratish
	token, err := auth.GenerateJWTToken(admin.ID, "admin")
	if err != nil {
		log.Printf("Token yaratish xatolik: %v", err)
		http.Error(w, "Token yaratishda xatolik", http.StatusInternalServerError)
		return
	}

	// Cookie orqali token saqlash
	cookie := &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(6 * 30 * 24 * time.Hour), // 6 oy
	}
	http.SetCookie(w, cookie)

	log.Printf("Admin %s muvaffaqiyatli login qildi", admin.Login)
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func (h *Handler) logout(w http.ResponseWriter, r *http.Request) {
	// Cookie o'chirish
	cookie := &http.Cookie{
		Name:     "auth_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1, // MaxAge -1 bilan darhol o'chiradi
	}
	http.SetCookie(w, cookie)

	log.Println("User logout qildi")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
