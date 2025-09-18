package middleware

import (
	"context"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Ucell/client_manager/auth"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
)

type AuthMiddleware struct {
	enforcer *casbin.Enforcer
}

func NewAuthMiddleware() *AuthMiddleware {
	m, err := model.NewModelFromFile("casbin/conf.conf")
	if err != nil {
		log.Fatalf("Casbin modelni yuklashda xatolik: %v", err)
	}

	adapter := fileadapter.NewAdapter("casbin/policy.csv")
	enforcer, err := casbin.NewEnforcer(m, adapter)
	if err != nil {
		log.Fatalf("Casbin enforcerni yaratishda xatolik: %v", err)
	}

	return &AuthMiddleware{enforcer: enforcer}
}

func (a *AuthMiddleware) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Login va assets yo'llarini o'tkazib yuborish
		if r.URL.Path == "/login" || r.URL.Path == "/auth" {
			next.ServeHTTP(w, r)
			return
		}
		if matched, _ := filepath.Match("/assets/*", r.URL.Path); matched {
			next.ServeHTTP(w, r)
			return
		}

		// Token olish
		token := r.Header.Get("Authorization")
		if token == "" {
			cookie, err := r.Cookie("auth_token")
			if err != nil {
				http.Redirect(w, r, "/login", http.StatusFound)
				return
			}
			token = cookie.Value
		}

		// Token tekshirish
		valid, err := auth.ValidateToken(token)
		if !valid || err != nil {
			log.Printf("Token yaroqsiz: %v", err)
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		// User ID va role olish
		userID, role, err := auth.GetUserIdFromToken(token)
		if err != nil {
			log.Printf("Tokendan ma'lumot olishda xatolik: %v", err)
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		// Casbin bilan ruxsatni tekshirish
		allowed, err := a.enforcer.Enforce(role, r.URL.Path, r.Method)
		if err != nil || !allowed {
			log.Printf("Ruxsat rad etildi: user=%s, path=%s, method=%s", userID, r.URL.Path, r.Method)
			// Cookie ni tozalash
			cookie := &http.Cookie{
				Name:     "auth_token",
				Value:    "",
				Path:     "/",
				HttpOnly: true,
				MaxAge:   -1,
			}
			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		// Context ga user ma'lumotlarini qo'shish
		ctx := context.WithValue(r.Context(), "user_id", userID)
		ctx = context.WithValue(ctx, "role", role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
