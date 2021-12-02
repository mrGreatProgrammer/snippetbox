package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Создается функция-обработчик "home", которая записывает байтовый слайс, содержащий
// Обработик главной страницы.
func Home(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Инициализируем срез содержащий пути к двум файлам. Обратите внимание, что
	// файл home-page-tmpl.html должен быть *первим* файлом в срезе.
	files := []string {
		"pkg/ui/html/home-page-tmpl.html",
		"pkg/ui/html/base-layout-tmpl.html",
		"pkg/ui/html/footer-partial-tmpl.html",
	}

	// Используем функцию tamplate.ParseFiles() для чтения файла шаблона.
	// Если возникла ошибка, мы запишем детальное сообщение ошибки и
	// используя функцию http.Error() мы отправим пользователю
	// ответ: 500 Internal Server Error (Внутреняя ошибка на сервере)
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() представляет
	// возможность отправки динамических данных в шаблон
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

// Обработчик для отображения сожержимого заметки.
func ShowSnippet(w http.ResponseWriter, r *http.Request)  {
	// Извлекаем значене параметра id из URL и попытаемся
	// конвертировать строку в integer используя функцию strconv.Atoi(). Если его нельзя
	// кновертировать в integer, или значение меньше 1, возвращаем ответ
	// 404 - страница не найдена!
	id, err := strconv.Atoi(r.URL.Query().Get("id")) 
	if err != nil || id < 1{
		http.NotFound(w, r)
		return
	}

	// Используем функцию fmt.Fprintf() для вставкизначения из id в строку ответа
	// и записываем его в http.ResponseWriter
	fmt.Fprintf(w, "Отображение выбранной заметки с ID %d ...", id)

	// w.Write([]byte("Отображение заметки..."))
}

// Обработчик для создания новой заметки.
func CreateSnippet(w http.ResponseWriter, r *http.Request)  {
	if r.Method != http.MethodPost {
		// Используем метод Header().Set() для добавления заголовка 'Allow: POST' в
        // карту HTTP-заголовков. Первый параметр - название заголовка, а
        // второй параметр - значение заголовка.
		w.Header().Set("Allow", http.MethodPost)

        // Используем функцию http.Error() для отправки кода состояния 405 с соответствующим сообщением
		http.Error(w, "метод запрещен!", http.StatusMethodNotAllowed)

		// Затем мы завершаем работу функции вызвав "return", чтобы
        // последующий код не выполнялся.
		return
	}
	w.Write([]byte("Форма для создания новой заметки..."))
}