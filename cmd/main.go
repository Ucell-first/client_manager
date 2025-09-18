package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Ucell/client_manager/configuration"
	httpcms "github.com/Ucell/client_manager/internal/delivery/http_cms"
	"github.com/Ucell/client_manager/storage"
	"github.com/Ucell/client_manager/storage/postgres"
)

func main() {
	cfg, err := configuration.Load()
	if err != nil {
		log.Fatalf("Konfiguratsiya yuklashda xatolik: %v", err)
	}

	db, err := postgres.ConnectPdb(&cfg.Postgres)
	if err != nil {
		log.Fatalf("Bazaga ulashda xatolik: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Bazani yopishda xatolik: %v", err)
		}
	}()

	store := storage.NewStorage(db)
	handler := httpcms.NewHandler(store)

	server := &http.Server{
		Addr:    cfg.Server.GetAddress(),
		Handler: handler.Routes(),
	}

	go func() {
		log.Printf("Server %s portida ishlamoqda", cfg.Server.GetAddress())
		log.Println("Login uchun: admin/admin yoki test/1234")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server ishga tushirishda xatolik: %v", err)
		}
	}()

	// Signal kutish
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Serverni to'xtatishga tayyorlanish...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server to'xtatishda xatolik:", err)
	}

	log.Println("Server muvaffaqiyatli to'xtatildi")
}
