module github.com/Desolitto/Avito-trainee-assignment-spring-2025/server

go 1.22.3

require (
	github.com/Desolitto/Avito-trainee-assignment-spring-2025/internal/auth v0.0.0
	github.com/gorilla/mux v1.8.1
)

require github.com/golang-jwt/jwt/v5 v5.2.2 // indirect

replace github.com/Desolitto/Avito-trainee-assignment-spring-2025/internal/auth => ../internal/auth
