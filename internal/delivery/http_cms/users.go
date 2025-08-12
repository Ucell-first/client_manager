package httpcms

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Ucell/client_manager/storage/repo"
)

type PageData struct {
	Title string
	Users []*repo.User
	User  *repo.User
	Error string
}

func (h *Handler) listUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.storage.User().GetAll(r.Context())
	if err != nil {
		log.Printf("Database xatoligi: %v", err)
		http.Error(w, "Foydalanuvchilarni olishda xatolik: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Oddiy template ishlatish
	h.renderSimpleListPage(w, users)
}

func (h *Handler) viewUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "User ID ko'rsatilmagan", http.StatusBadRequest)
		return
	}

	user, err := h.storage.User().GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Foydalanuvchini topishda xatolik: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.NotFound(w, r)
		return
	}

	h.renderSimpleViewPage(w, user)
}

func (h *Handler) newUserForm(w http.ResponseWriter, r *http.Request) {
	h.renderSimpleNewPage(w)
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Faqat POST so'rov qabul qilinadi", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Forma ma'lumotlarini o'qishda xatolik", http.StatusBadRequest)
		return
	}

	isActive, _ := strconv.ParseBool(r.FormValue("is_active"))

	user := &repo.User{
		MSISDN:   r.FormValue("msisdn"),
		Name:     r.FormValue("name"),
		IsActive: isActive,
	}

	if err := h.storage.User().Create(r.Context(), user); err != nil {
		http.Error(w, "Foydalanuvchini yaratishda xatolik: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func (h *Handler) editUserForm(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "User ID ko'rsatilmagan", http.StatusBadRequest)
		return
	}

	user, err := h.storage.User().GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Foydalanuvchini topishda xatolik: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.NotFound(w, r)
		return
	}

	data := PageData{
		Title: "Foydalanuvchini tahrirlash",
		User:  user,
	}

	if err := h.tmpl.ExecuteTemplate(w, "edit.html", data); err != nil {
		log.Printf("Template xatoligi: %v", err)
		http.Error(w, "Sahifani ko'rsatishda xatolik", http.StatusInternalServerError)
	}
}

func (h *Handler) updateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Faqat POST so'rov qabul qilinadi", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Forma ma'lumotlarini o'qishda xatolik", http.StatusBadRequest)
		return
	}

	isActive, _ := strconv.ParseBool(r.FormValue("is_active"))

	user := &repo.User{
		UserID:   r.FormValue("user_id"),
		MSISDN:   r.FormValue("msisdn"),
		Name:     r.FormValue("name"),
		IsActive: isActive,
	}

	if err := h.storage.User().Update(r.Context(), user); err != nil {
		http.Error(w, "Foydalanuvchini yangilashda xatolik: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "User ID ko'rsatilmagan", http.StatusBadRequest)
		return
	}

	if err := h.storage.User().Delete(r.Context(), id); err != nil {
		http.Error(w, "Foydalanuvchini o'chirishda xatolik: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}
