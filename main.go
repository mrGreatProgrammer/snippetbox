package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Создается функция-обработчик "home", которая записывает байтовый слайс, содержащий
// Обработик главной страницы.
func home(w http.ResponseWriter, r *http.Request)  {
	// Проверяется, если текущий путь URL запроса точно совпадает с шаблоном "/". Если нет, вызывается
	// функция http.NotFound() для возвращения клиенту ошибку 404.
	// Выжно, чтобы мы завершили работу обработчика через return. Если мы забудем про "return", то обработчик
	// продолжит работу и выведет сообщение "Привет из SnippetBox" как ни в чем не бывало.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Привет из Snippetbox"))
}

// Обработчик для отображения сожержимого заметки.
func showSnippet(w http.ResponseWriter, r *http.Request)  {
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
func createSnippet(w http.ResponseWriter, r *http.Request)  {
	if r.Method != http.MethodPost {
		// Используем метод Header().Set() для добавления заголовка 'Allow: POST' в
        // карту HTTP-заголовков. Первый параметр - название заголовка, а
        // второй параметр - значение заголовка.
		w.Header().Set("Allow", http.MethodPost)

        // Используем функцию http.Error() для отправки кода состояния 405 с соответствующим сообщением
		http.Error(w, "метод запрещен!", 405)

		// Затем мы завершаем работу функции вызвав "return", чтобы
        // последующий код не выполнялся.
		return
	}
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