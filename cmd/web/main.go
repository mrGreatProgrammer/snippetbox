package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/mrGreatProgrammer/snippetbox/cmd/web/handlers"
	"github.com/mrGreatProgrammer/snippetbox/pkg/models/mysql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	addr := flag.String("addr", ":4000", "Сетевой адрес HTTP")
	dsn := flag.String("dsn", "web:password@/snippetbox?parseTime=true", "Название MySQL источника данных")
	flag.Parse()

	// Логирование сообщений в файл
	// f, err := os.OpenFile("info.log", os.O_RDWR|os.O_CREATE, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()

	// infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)

	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	// Инициализируем новый кэш шаблона...
	templateCache, err := handlers.NewTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	// И добавляем его в зависимостях нашего
	// веб-приложения.
	app := &handlers.Application{
		ErrorLog: errorLog,
		InfoLog: infolog,
		Snippets: &mysql.SnippetModel{DB: db},
		TemplateCache: templateCache,
	}

	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.Routes(),
	}

	infolog.Printf("Запуск сервера на localhost%s\n", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

// Функция openDB() обертывает sql.Open() и возвращает пул соединений sql.DB
// для заданной строки подключения (DSN).
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}