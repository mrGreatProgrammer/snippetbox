package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mrGreatProgrammer/snippetbox/cmd/web/handlers"
)

func main() {
	addr := flag.String("addr", ":4000", "Сетевой адрес HTTP")

	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/snippet", handlers.ShowSnippet)
	mux.HandleFunc("snippet/create", handlers.CreateSnippet)

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("ui/static/")})
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Инициализируем новую структуру http.Server. Мы устанавливаем поля Addr и Handler, так
	// что сервер использует тот же сетевой адрес и маршуруты, что и раньше, и назначаем
	// поле ErrorLog, чтобы сервер использовал наш логгер
	// при возникновении проблем.
	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: mux,
	}

	infolog.Printf("Запуск сервера на %s\n", *addr)
	// Вызываем метод ListenAndServe() от нашей новой структуры http.Server
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
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
