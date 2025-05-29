package main

import (
	"fmt"
	"net/http"
	clientapi "testCICD/internal/client"
	"time"
)

func main() {

	for {
		time.Sleep(3 * time.Second)

		date, time, err := clientapi.ReqDateTime(http.MethodGet, "http://server:50555/datetime", http.DefaultClient)
		if err != nil {
			fmt.Printf("ошибка при выполнении запроса:{%v}\n", err)
			continue
		}
		fmt.Printf("на сервере - дата:{%s}, время:{%s}\n", date, time)
	}
}
