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
	log.Println("listUsers function called")

	users, err := h.storage.User().GetAll(r.Context())
	if err != nil {
		log.Printf("Database error: %v", err)
		http.Error(w, "Foydalanuvchilarni olishda xatolik: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Found %d users", len(users))

	data := PageData{
		Title: "Foydalanuvchilar ro'yxati",
		Users: users,
	}

	log.Printf("Executing template 'list.html' with data: %+v", data)

	// Template mavjudligini tekshirish
	if h.tmpl.Lookup("list.html") == nil {
		log.Println("ERROR: 'list.html' template not found!")
		http.Error(w, "Template topilmadi", http.StatusInternalServerError)
		return
	}

	// Base template bilan birga render qilish
	if err := h.tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Sahifani ko'rsatishda xatolik: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Template executed successfully")
}

func (h *Handler) viewUser(w http.ResponseWriter, r *http.Request) {
	log.Println("viewUser function called")

	id := r.URL.Query().Get("id")
	if id == "" {
		log.Println("ERROR: User ID not provided")
		http.Error(w, "User ID ko'rsatilmagan", http.StatusBadRequest)
		return
	}

	log.Printf("Looking for user with ID: %s", id)

	user, err := h.storage.User().GetByID(r.Context(), id)
	if err != nil {
		log.Printf("Database error: %v", err)
		http.Error(w, "Foydalanuvchini topishda xatolik: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if user == nil {
		log.Printf("User with ID %s not found", id)
		http.NotFound(w, r)
		return
	}

	log.Printf("Found user: %+v", user)

	data := PageData{
		Title: "Foydalanuvchi ma'lumotlari",
		User:  user,
	}

	if err := h.tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Sahifani ko'rsatishda xatolik: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("View template executed successfully")
}

func (h *Handler) newUserForm(w http.ResponseWriter, r *http.Request) {
	log.Println("newUserForm function called")

	data := PageData{
		Title: "Yangi foydalanuvchi qo'shish",
	}

	if err := h.tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Sahifani ko'rsatishda xatolik: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("New user form template executed successfully")
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	log.Println("createUser function called")

	if r.Method != http.MethodPost {
		log.Println("ERROR: Non-POST request to createUser")
		http.Error(w, "Faqat POST so'rov qabul qilinadi", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		log.Printf("Form parsing error: %v", err)
		http.Error(w, "Forma ma'lumotlarini o'qishda xatolik", http.StatusBadRequest)
		return
	}

	isActive, _ := strconv.ParseBool(r.FormValue("is_active"))

	user := &repo.User{
		MSISDN:   r.FormValue("msisdn"),
		Name:     r.FormValue("name"),
		IsActive: isActive,
	}

	log.Printf("Creating user: %+v", user)

	if err := h.storage.User().Create(r.Context(), user); err != nil {
		log.Printf("Database error creating user: %v", err)
		http.Error(w, "Foydalanuvchini yaratishda xatolik: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("User created successfully with ID: %s", user.UserID)
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func (h *Handler) editUserForm(w http.ResponseWriter, r *http.Request) {
	log.Println("editUserForm function called")

	id := r.URL.Query().Get("id")
	if id == "" {
		log.Println("ERROR: User ID not provided for edit")
		http.Error(w, "User ID ko'rsatilmagan", http.StatusBadRequest)
		return
	}

	log.Printf("Loading user for edit with ID: %s", id)

	user, err := h.storage.User().GetByID(r.Context(), id)
	if err != nil {
		log.Printf("Database error: %v", err)
		http.Error(w, "Foydalanuvchini topishda xatolik: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if user == nil {
		log.Printf("User with ID %s not found for edit", id)
		http.NotFound(w, r)
		return
	}

	data := PageData{
		Title: "Foydalanuvchini tahrirlash",
		User:  user,
	}

	if err := h.tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Sahifani ko'rsatishda xatolik: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Edit form template executed successfully")
}

func (h *Handler) updateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("updateUser function called")

	if r.Method != http.MethodPost {
		log.Println("ERROR: Non-POST request to updateUser")
		http.Error(w, "Faqat POST so'rov qabul qilinadi", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		log.Printf("Form parsing error: %v", err)
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

	log.Printf("Updating user: %+v", user)

	if err := h.storage.User().Update(r.Context(), user); err != nil {
		log.Printf("Database error updating user: %v", err)
		http.Error(w, "Foydalanuvchini yangilashda xatolik: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("User updated successfully: %s", user.UserID)
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("deleteUser function called")

	id := r.URL.Query().Get("id")
	if id == "" {
		log.Println("ERROR: User ID not provided for delete")
		http.Error(w, "User ID ko'rsatilmagan", http.StatusBadRequest)
		return
	}

	log.Printf("Deleting user with ID: %s", id)

	if err := h.storage.User().Delete(r.Context(), id); err != nil {
		log.Printf("Database error deleting user: %v", err)
		http.Error(w, "Foydalanuvchini o'chirishda xatolik: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("User deleted successfully: %s", id)
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}
