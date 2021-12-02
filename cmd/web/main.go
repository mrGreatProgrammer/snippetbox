package main

import (
	"log"
	"net/http"
	"path/filepath"

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
	fileServer := http.FileServer(neuteredFileSystem{http.Dir("ui/static/")})

	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Запуск сервера на http//loaclhost:4000")
	err := http.ListenAndServe(":4000", mux)
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
