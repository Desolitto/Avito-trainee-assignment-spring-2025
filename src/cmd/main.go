package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	openapi  "github.com/Desolitto/Avito-trainee-assignment-spring-2025/server/go"

)

func main() {
	// Создаем новый роутер
	router := mux.NewRouter()

	// Создаем сервис
	apiService := openapi.NewDefaultAPIService()

	// Создаем контроллер
	apiController := openapi.NewDefaultAPIController(apiService)

	// Получаем все маршруты
	routes := apiController.Routes()

	// Регистрируем каждый маршрут
	for _, route := range routes {
		handler := http.HandlerFunc(route.HandlerFunc)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Handler(handler)
	
	// Middleware для логирования (опционально)
	router.Use(loggingMiddleware)

	}
    // Распечатаем зарегистрированные маршруты
    router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
        path, _ := route.GetPathTemplate()
        methods, _ := route.GetMethods()
        log.Printf("Route: %v, Methods: %v", path, methods)
        return nil
    })
	// Настройка сервера
	server := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	log.Println("Starting server on :8081")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

// Middleware для логирования запросов
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
