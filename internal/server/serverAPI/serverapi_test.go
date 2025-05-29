package serverapi

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Тест обработчика запроса даты и времени сервера (Успешность)
func Test_HndlCurrentDateTime_Success(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/datetime", nil)
	res := httptest.NewRecorder()

	var curDT DateTimeT
	tn := time.Now()
	curDT.Date = tn.Format("2006-01-02")
	curDT.Time = tn.Format("15:04:05")

	curDT.HndlCurrentDateTime(res, req)
	require.Equalf(t, http.StatusOK, res.Result().StatusCode, "ожидался код:{%d}, а принято:{%d}", http.StatusOK, res.Result().StatusCode)

	// Обработка тела ответа
	body, err := io.ReadAll(res.Body)
	require.NoErrorf(t, err, "ошибка чтения тела ответа:{%s}", err)

	var rxDT DateTimeT
	err = json.Unmarshal(body, &rxDT)
	require.NoErrorf(t, err, "ошибка десериализации тела ответа:{%s}", err)

	assert.Equalf(t, curDT.Date, rxDT.Date, "ожидалась дата:{%s}, а принято:{%s}", curDT.Date, rxDT.Date)
	assert.Equalf(t, curDT.Time, rxDT.Time, "ожидалась время:{%s}, а принято:{%s}", curDT.Time, rxDT.Time)
}

// Тест обработчика запроса даты и времени сервера (Ошибки)
func Test_HndlCurrentDateTime_Error(t *testing.T) {

	var testData = []struct {
		testName string
		method   string
		wantCode int
	}{
		{
			testName: "Неподдерживаемый метод POST",
			method:   http.MethodPost,
			wantCode: http.StatusBadRequest,
		},
	}

	for _, tt := range testData {
		t.Run(tt.testName, func(t *testing.T) {

			req := httptest.NewRequest(tt.method, "/datetime", nil)
			res := httptest.NewRecorder()

			var curDT DateTimeT
			tn := time.Now()
			curDT.Date = tn.Format("2006-01-02")
			curDT.Time = tn.Format("15:04:05")

			curDT.HndlCurrentDateTime(res, req)
			assert.Equalf(t, tt.wantCode, res.Result().StatusCode, "ожидался код:{%d}, а принято:{%d}", tt.wantCode, res.Result().StatusCode)
		})
	}
}
