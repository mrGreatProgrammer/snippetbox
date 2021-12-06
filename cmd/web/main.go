package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"

	"github.com/mrGreatProgrammer/snippetbox/cmd/web/handlers"
)

func main() {
	// Создаем новый флаг командной строки, значение по умолчанию: ":4000".
	// Добавляем небольшую справку, объясняющая, что содержит данный флаг.
	// Значение флага будет сохранено в переменной addr.
	addr := flag.String("addr", ":4000", "Сетевой адрес HTTP")

	// МЫ вызываем функцию flag.Parse() для извлечение флага из командной строки.
	// Она считывает значение флага из командной строки и присваивает его содержимое
	// переменной. Вам нужно вызвать ее *до* использования переменной addr
	// иначе она всегда будет содержать значение по умолчанию ":4000".
	// Если есть ошибки во время извлечения данных - приложение будет остановлено.
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/snippet", handlers.ShowSnippet)
	mux.HandleFunc("snippet/create", handlers.CreateSnippet)

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("ui/static/")})
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Значение, возвращает функцией flag.String(), является указателем на значение
	// из флага, а не самим значением. Нам нужно убрать ссылку на указатель
	// то есть перед использованием добавьте к нему префикс *. Обратите внимание, что мы используем
	// функцию log.Printf() для записи логов в журнал работы нашего приложения.
	log.Printf("Запуск сервера на %s\n", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			cloeErr := f.Close()
			if cloeErr != nil {
				return nil, cloeErr
			}

			return nil, err
		}
	}

	return f, nil
}
