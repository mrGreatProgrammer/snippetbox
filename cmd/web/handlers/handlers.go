package handlers

import (
	"errors"
	"fmt"
	"html/template"

	// "html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/mrGreatProgrammer/snippetbox/pkg/models"
	"github.com/mrGreatProgrammer/snippetbox/pkg/models/mysql"
)

// Добавляемм поле snippets в структуру application. Это позволит
// сделать объект SnippetModel доступным для наших обработчиков.
type Application struct {
	ErrorLog *log.Logger
	InfoLog *log.Logger
	Snippets *mysql.SnippetModel
}

// Создается функция-обработчик "home", которая записывает байтовый слайс, содержащий
// Обработик главной страницы.
func (app *Application) Home(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path != "/" {
		app.NotFound(w) // Используем помощника notFound()
		return
	}

	s, err := app.Snippets.Latest()
	if err != nil {
		app.ServerError(w, err)
		return
	}

	for _, snippet := range s {
		fmt.Fprintf(w, "%v\n", snippet)
	}

	// files := []string {
	// 	"ui/html/home-page-tmpl.html",
	// 	"ui/html/base-layout-tmpl.html",
	// 	"ui/html/footer-partial-tmpl.html",
	// }

	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.ServerError(w, err) // Используем помощника serverError()
	// 	return
	// }

	// err = ts.Execute(w, nil)
	// if err != nil {
	// 	app.ServerError(w, err) // Используем помощника serverError()
	// }
}

// Обработчик для отображения сожержимого заметки.
func (app *Application) ShowSnippet(w http.ResponseWriter, r *http.Request)  {
	id, err := strconv.Atoi(r.URL.Query().Get("id")) 
	if err != nil || id < 1{
		app.NotFound(w)
		return
	}


	s, err := app.Snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.NotFound(w)
		} else {
			app.ServerError(w, err)
		}
		return
	}
	
	// Инициализируем срез, содержащий путь к файлу show.page.tmpl.html
	// Добавив еще базовый шаблон и часть футера, который мы сделали ранее.
	files := []string{
		"ui/html/show-page-tmpl.html",
		"ui/html/base-layout-tmpl.html",
		"ui/html/footer-partial-tmpl.html",
	}

	// Парсинг файлов шаблона...
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	// А затем выполняем их. Обратите внимание на передачу заметки с данными
	// (структура models.Snippet) в качестве последнего параметра.
	err = ts.Execute(w, s)
	if err != nil {
		app.ServerError(w, err)
	}
}

// Обработчик для создания новой заметки.
func (app *Application) CreateSnippet(w http.ResponseWriter, r *http.Request)  {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.ClientError(w, http.StatusMethodNotAllowed) // Используем помощника clientError()

		return
	}

	// Создаем несколько переменных, содержащих тестовые данные. Мы удалим их позже.
	title := "История про улитку"
	content := "Улитка выползла из раковины, \nвытянула рожки, \nи опять подобрала их."
	expires := "7"

	// Передаем данные в метод SnippetModel.Insert(), получая обратно
	// ID только что созданной записи в базу данных.
	id, err := app.Snippets.Insert(title, content, expires)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	// Перенаправляем пользователя на соответствующую страницу заметки

	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}