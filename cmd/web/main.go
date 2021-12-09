package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/mrGreatProgrammer/snippetbox/cmd/web/handlers"
	
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	addr := flag.String("addr", ":4000", "Сетевой адрес HTTP")
	// Определение нового флага из командной строки для настройки MySQL подключения
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

	// Чтобы функция main() была более компактной, мы поместили код для создания
	// пула соединений в отдельную функцию openDB(). Мы передаем в нее полученный
	// источник данных (DSN) из флага командной строки.
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	// Мы также откладываем вызов db.Close(), чтобы пул соединений был закрыт
	// до выхода из функции main().
	// Подробнее про defer: https://golangs.org/errors#defer
	defer db.Close()

	app := &handlers.Application{
		ErrorLog: errorLog,
		InfoLog: infolog,
	}



	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.Routes(),
	}

	infolog.Printf("Запуск сервера на %s\n", *addr)
	// Поскольку переменная `err` уже объявлена в приведенном выше коде, нужно
	// использовать оператор присваивания =
	// вместо оператора := (объявить и присвоить)
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