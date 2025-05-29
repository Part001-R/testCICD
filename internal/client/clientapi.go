package clientapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type (
	// Для передачи даты и времени
	DateTimeT struct {
		Date string `json:"date"`
		Time string `json:"time"`
	}
)

// Запрос к серверу на получение текущих даты и времени.
func ReqDateTime(httpMethod string, u string, client *http.Client) (d, t string, err error) {

	// Проверка аргументов
	if httpMethod == "" {
		return "", "", errors.New("нет содержимого в параметре httpMethod")
	}
	if u == "" {
		return "", "", errors.New("нет содержимого в параметре u")
	}
	if client == nil {
		return "", "", errors.New("нет содержимого в параметре client")
	}

	// Создание запроса
	req, err := http.NewRequest(httpMethod, u, nil)
	if err != nil {
		return "", "", fmt.Errorf("ошибка при создании запроса:{%v}", err)
	}

	// Запрос
	res, err := client.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("ошибка запроса:{%v}", err)
	}

	if res.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("ожидался код 200, а принят:{%d}", res.StatusCode)
	}

	// Чтение тела ответа
	bodyRx, err := io.ReadAll(res.Body)
	if err != nil {
		return "", "", fmt.Errorf("ошибка при чтении тела ответа:{%v}", err)
	}
	defer func() {
		_ = res.Body.Close()
	}()

	var rxDateTime DateTimeT
	err = json.Unmarshal(bodyRx, &rxDateTime)
	if err != nil {
		return "", "", fmt.Errorf("ошибка при десиреализации ответа сервера:{%v}", err)
	}

	// Возврат результата запроса
	d = rxDateTime.Date
	t = rxDateTime.Time
	return d, t, nil
}
