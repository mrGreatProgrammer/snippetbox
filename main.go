package main

import (
	"log"
	"net/http"
)

// Создается функция-обработчик "home", которая записывает байтовый слайс, содержащий
// Обработик главной страницы.
func home(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Привет из Snippetbox"))
}

// Обработчик для отображения сожержимого заметки.
func showSnippet(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Отображение заметки..."))
}

// Обработчик для создания новой заметки.
func createSnippet(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Форма для создания новой заметки..."))
}

func main()  {
	// Регистрируем два новых обработчика и соответствующие URL-шаблоны в
	// маршрутизаторе servemux
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}