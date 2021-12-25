package handlers

import (
	"html/template"
	"path/filepath"

	"github.com/mrGreatProgrammer/snippetbox/pkg/models"
)

// Создаем тип templateData, который будет действовать как хранилище для
// любых динамических данных, которые нужно передать HTML-шаблонам.
// На данный момент он содержит только одно поле, но мы добавим в него другие
// по мере развития нашего приложения.
type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}

func NewTemplateCache(dir string) (map[string]*template.Template, error) {
	// Инициализируем новую карту, которая будет хранить кэш.
	cache := map[string]*template.Template{}

	// Используем функцию filepath.Glob, чтобы получить срез всех файловых путей с
	// расширением '-page-tmpl.html'. По сути, мы получим список всех файлов шаблона для страниц
	// нашего веб-приложения.
	pages, err := filepath.Glob(filepath.Join(dir, "*-page-tmpl.html"))
	if err != nil {
		return nil, err
	}

	// Перебираем файл шаблона от каждой страницы.
	for _, page := range pages {
		// Извлечение конечное названия файла (например, 'home-page-tmpl.html') из полного пути к файлу
		// и присваивание его переменной name.
		name := filepath.Base(page)

		// Обрабатываем итерируемый файл шаблона.
		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Используем  метод ParseGlob для добавления всех каркасных шаблонов.
		// В нашем случае это только файл base-layout-tmpl.html (основная структура шаблона)
		ts, err = ts.ParseGlob(filepath.Join(dir, "*-layout-tmpl.html"))
		if err != nil {
			return nil, err
		}

		// Используем метод ParseGlob для добавления всех вспомогательных шаблонов.
		// В нашем случае это footer-partial-tmpl.html "подвал" нашего шаблона.
		ts, err = ts.ParseGlob(filepath.Join(dir, "*-partial-tmpl.html"))
		if err != nil {
			return nil, err
		}

		// Добавляем полученный набор шаблонов в кэш, используя название страницы
		// (например, home-page-tmpl.html) в качестве ключа для нашей карты.
		cache[name] = ts
	}
	
	// Возвращаем полученную карту.
	return cache, nil
}