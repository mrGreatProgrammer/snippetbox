package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/mrGreatProgrammer/snippetbox/cmd/web/handlers"
)

func main() {
	addr := flag.String("addr", ":4000", "Сетевой адрес HTTP")

	// Логирование сообщений в файл
	// f, err := os.OpenFile("info.log", os.O_RDWR|os.O_CREATE, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()

	// infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)

	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Инициализируем новую структуру с зависимотями приложения.
	app := &handlers.Application{
		ErrorLog: errorLog,
		InfoLog: infolog,
	}

	flag.Parse()

	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.Routes(), // Вызов нового метода app.routes()
	}

	infolog.Printf("Запуск сервера на %s\n", *addr)
	// Вызываем метод ListenAndServe() от нашей новой структуры http.Server
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}