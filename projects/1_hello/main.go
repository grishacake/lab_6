package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Регистрируем обработчики для разных путей
	http.HandleFunc("/get", handleRequest)

	// Запускаем веб-сервер на порту 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}
	w.Write([]byte("Hello, web!!!"))
}
