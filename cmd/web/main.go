package main

import (
	"log"
	"net/http"
	// "path/filepath"

	"github.com/mrGreatProgrammer/snippetbox/cmd/web/handlers"
)

func main() {
mux := http.NewServeMux()
mux.HandleFunc("/", handlers.Home)
mux.HandleFunc("/snippet", handlers.ShowSnippet)
mux.HandleFunc("snippet/create", handlers.CreateSnippet)

// Инициализируем FileServer, он будет обрабатывать
// HTTP-запросы к статическим файлам из папки "./ui/static".
// Обратите внимание, что переданный в финукцию http.Dir путь
// является относительным корневой папке проекта
fileServer := http.FileServer(http.Dir("ui/static/"))

// Используем функцию mux.Handle() для регистрации обработчика для
// всех запросов, которые начинаются с "/static/". Мы Убираем
// префикс "/static" перед тем как запрос достигнет http.FileServer
mux.Handle("/static/", http.StripPrefix("/static", fileServer))

log.Println("Запуск сервера на http//loaclhost:4000")
err := http.ListenAndServe(":4000", mux)
log.Fatal(err)
}