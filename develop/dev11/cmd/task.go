package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mirhijinam/wb-l2/develop/dev11/internal/config"
	"github.com/mirhijinam/wb-l2/develop/dev11/internal/net"
	"github.com/mirhijinam/wb-l2/develop/dev11/internal/repository"
	"github.com/mirhijinam/wb-l2/develop/dev11/internal/service"
)

/* HTTP server.
 *
 * Реализовать HTTP сервер для работы с календарем.
 *
 * В рамках задания необходимо:
 * - Использовать стандартную библиотеку HTTP
 * - Реализовать вспомогательные функции для сериализации объектов
 * доменной области в JSON
 * - Реализовать вспомогательные функции для парсинга и валидации параметров
 * методов /create_event и /update_event
 * - Реализовать HTTP обработчики для каждого из методов API, используя
 * вспомогательные функции и объекты доменной области
 * - Реализовать middleware для логирования запросов
 *
 * Методы API:
 * - POST /create_event
 * - POST /update_event
 * - POST /delete_event
 * - GET /events_for_day
 * - GET /events_for_week
 * - GET /events_for_month
 *
 * Формат параметров:
 * - www-url-form-encoded (напр. user_id=3&date=2019-09-09)
 *
 * Способ передачи параметров:
 * - GET: через queryString
 * - POST: через тело запроса
 *
 * В результате каждого запроса должен возвращаться JSON документ, содержащий:
 * - {"result": "..."} в случае успешного выполнения метода
 * - {"error": "..."} в случае ошибки бизнес-логики
 *
 * Ошибки:
 * - В случае ошибки бизнес-логики сервер должен возвращать HTTP 503
 * - В случае ошибки входных данных сервер должен возвращать HTTP 400
 * - В случае остальных ошибок сервер должен возвращать HTTP 500
 *
 * Условия:
 * 1. Реализовать все методы.
 * 2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
 * 3. Web-сервер должен запускаться на порту указанном в конфиге.
 * 4. Web-сервер должен выводить в лог каждый обработанный запрос.
 * 5. Код должен проходить проверки go vet и golint.
 */

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	repo := repository.New()
	svc := service.New(repo)
	handler := net.New(svc)

	srv := http.Server{
		Handler: handler,
		Addr:    fmt.Sprintf(":%s", cfg.Port),
	}

	log.Printf("server started on port %s", cfg.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("error starting server: %v", err)
	}

}
