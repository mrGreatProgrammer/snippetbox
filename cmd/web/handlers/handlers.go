package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Создаем структуру `application` для хранения зависимостей всего веб-приложения.
// Пока, что мы добавим поля только для двух логгеров, но
// мы будем расширать данную структуру по мере усложнения приложения.
type Application struct {
	ErrorLog *log.Logger
	InfoLog *log.Logger
}

// Создается функция-обработчик "home", которая записывает байтовый слайс, содержащий
// Обработик главной страницы.
func (app *Application) Home(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path != "/" {
		app.NotFound(w) // Используем помощника notFound()
		return
	}

	files := []string {
		"ui/html/home-page-tmpl.html",
		"ui/html/base-layout-tmpl.html",
		"ui/html/footer-partial-tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.ServerError(w, err) // Используем помощника serverError()
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.ServerError(w, err) // Используем помощника serverError()
	}
}

// Обработчик для отображения сожержимого заметки.
func (app *Application) ShowSnippet(w http.ResponseWriter, r *http.Request)  {
	id, err := strconv.Atoi(r.URL.Query().Get("id")) 
	if err != nil || id < 1{
		app.NotFound(w) // Используем помощника notFound()
		return
	}

	fmt.Fprintf(w, "Отображение выбранной заметки с ID %d ...", id)

	// w.Write([]byte("Отображение заметки..."))
}

// Обработчик для создания новой заметки.
func (app *Application) CreateSnippet(w http.ResponseWriter, r *http.Request)  {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.ClientError(w, http.StatusMethodNotAllowed) // Используем помощника clientError()

		return
	}
	w.Write([]byte("Форма для создания новой заметки..."))
}