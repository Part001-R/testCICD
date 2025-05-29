package main

import (
	"fmt"
	"log"
	"net/http"
	serverapi "testCICD/internal/server/serverAPI"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {

	cR := chi.NewRouter()

	// Ручка предоставления текущих даты и времени
	cR.Get("/datetime", func(w http.ResponseWriter, r *http.Request) {
		var curDT serverapi.DateTimeT

		tn := time.Now()
		curDT.Date = tn.Format("2006-01-02")
		curDT.Time = tn.Format("15:04:05")
		curDT.HndlCurrentDateTime(w, r)
	})

	// Запуск HTTP сервера
	fmt.Println("Запуск HTTP сервера")
	err := http.ListenAndServe(":50555", cR)
	if err != nil {
		log.Fatal("ошибка запуска сервера")
	}
}
