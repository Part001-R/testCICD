package serverapi

import (
	"encoding/json"
	"net/http"
	"time"
)

type (
	// Для передачи даты и времени
	DateTimeT struct {
		Date string `json:"date"`
		Time string `json:"time"`
	}
)

// Обработчик запроса на предоставление текущей даты и времени
func (dt *DateTimeT) HndlCurrentDateTime(w http.ResponseWriter, r *http.Request) {

	// Проверка данных
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if _, err := time.Parse("2006-01-02", dt.Date); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if _, err := time.Parse("15:04:05", dt.Time); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Подготовка ответа
	byteTx, err := json.Marshal(dt)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Ответ
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteTx)
}
