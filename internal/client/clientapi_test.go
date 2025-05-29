package clientapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Запрос даты и времени сервера (Успешность)
func Test_ReqDateTime_Success(t *testing.T) {

	var curDT DateTimeT
	tn := time.Now()
	curDT.Date = tn.Format("2006-01-02")
	curDT.Time = tn.Format("15:04:05")

	method := http.MethodGet

	// Сервер
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Подготовка ответа
		byteTx, err := json.Marshal(curDT)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		// Ответ
		w.Header().Set("Content-Type", "application-json")
		w.WriteHeader(http.StatusOK)
		w.Write(byteTx)
	}))
	defer srv.Close()

	rxDate, rxTime, err := ReqDateTime(method, srv.URL, srv.Client())
	require.NoErrorf(t, err, "ожидалсь отсутствие ошибки, а принято:{%v}", err)
	assert.Equalf(t, curDT.Date, rxDate, "ожидалась дата:{%s}, а принята:{%s}", curDT.Date, rxDate)
	assert.Equalf(t, curDT.Time, rxTime, "ожидалась время:{%s}, а принята:{%s}", curDT.Time, rxTime)
}

// Запрос даты и времени сервера (Ошибки)
func Test_ReqDateTime_Error(t *testing.T) {

	var testData = []struct {
		testName  string
		useMethod string
		useURL    string
		useClient string
		wantErr   string
	}{
		{
			testName:  "Нет мотода",
			useMethod: "false",
			useURL:    "true",
			useClient: "true",
			wantErr:   "нет содержимого в параметре httpMethod",
		},
		{
			testName:  "Нет URL",
			useMethod: "true",
			useURL:    "false",
			useClient: "true",
			wantErr:   "нет содержимого в параметре u",
		},
		{
			testName:  "Нет клиента",
			useMethod: "true",
			useURL:    "true",
			useClient: "false",
			wantErr:   "нет содержимого в параметре client",
		},
	}

	for _, tt := range testData {
		t.Run(tt.testName, func(t *testing.T) {

			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
			defer srv.Close()

			method := http.MethodGet
			client := srv.Client()
			u := srv.URL

			if tt.useMethod == "false" {
				method = ""
			}
			if tt.useURL == "false" {
				u = ""
			}
			if tt.useClient == "false" {
				client = nil
			}

			_, _, err := ReqDateTime(method, u, client)
			rxErr := fmt.Sprintf("%s", err)
			assert.Equalf(t, tt.wantErr, rxErr, "ожидалась ошибка:{%s}, а принято:{%s}", tt.wantErr, rxErr)
		})
	}
}
