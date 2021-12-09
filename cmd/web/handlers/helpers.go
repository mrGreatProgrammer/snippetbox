package handlers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	// "github.com/mrGreatProgrammer/snippetbox/cmd/web/handlers"
)



// Помощник serverError записывает сообщение об ошибке в errorLog и
// затем отправляет пользователю ответ 500 "Внутренняя ошибка сервера".
func (app *Application) ServerError(w http.ResponseWriter, err error)  {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// Помощник clientError отправляет определенный код состояния и соответвующее описание
// пользоваелю. Мы будем использовать это в следующих уроках, чтобы отправлять ответывроде 400
// "Bad Request", когда есть проблема с пользовательским запросом.
func (app *Application) ClientError(w http.ResponseWriter, status int)  {
	http.Error(w, http.StatusText(status), status)
}

// Мы также реализуем помощник notFound. Это просто
// удобная оболочка вокруг clientError, которая отправляет пользоваетелю ответ "404 Страница не найдена".
func (app *Application) NotFound(w http.ResponseWriter)  {
	app.ClientError(w, http.StatusNotFound)
}