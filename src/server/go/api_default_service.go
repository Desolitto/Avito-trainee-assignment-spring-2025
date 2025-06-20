// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * backend service
 *
 * Сервис для управления ПВЗ и приемкой товаров
 *
 * API version: 1.0.0
 */

package openapi

import (
	"context"
	"net/http"
	"errors"
	"time"
	"log"
	"fmt"
	"github.com/Desolitto/Avito-trainee-assignment-spring-2025/internal/auth"


)
var allowedRoles = map[string]bool{
    "employee":   true,
    "moderator":  true,
    "admin":      true,
}

func isValidRole(role string) bool {
    return allowedRoles[role]
}

// DefaultAPIService is a service that implements the logic for the DefaultAPIServicer
// This service should implement the business logic for every endpoint for the DefaultAPI API.
// Include any external packages or services that will be required by this service.
type DefaultAPIService struct {
}

// NewDefaultAPIService creates a default api service
func NewDefaultAPIService() *DefaultAPIService {
	return &DefaultAPIService{}
}

// DummyLoginPost - Получение тестового токена
func (s *DefaultAPIService) DummyLoginPost(ctx context.Context, req DummyLoginPostRequest) (ImplResponse, error) {
    // Проверка на пустую роль
    if req.Role == "" {
        return Response(http.StatusBadRequest, Error{
            Message: "Role is required. Allowed roles: employee, moderator, admin",
        }), nil
    }

    // Проверка роли с использованием новой функции
    if !isValidRole(req.Role) {
        return Response(http.StatusBadRequest, Error{
            Message: fmt.Sprintf("Invalid role. Allowed roles: %v", 
                getValidRoles()),
        }), nil
    }

    // Генерация токена
    token, err := auth.GenerateJWT(req.Role)
    if err != nil {
        log.Printf("Token generation error for role %s: %v", req.Role, err)
        return Response(http.StatusInternalServerError, Error{
            Message: "Failed to generate token",
        }), nil
    }

    // Возвращаем JSON с токеном
    return Response(http.StatusOK, map[string]string{
        "token": token,
    }), nil
}


// Вспомогательная функция для получения списка допустимых ролей
func getValidRoles() []string {
    roles := make([]string, 0, len(allowedRoles))
    for role := range allowedRoles {
        roles = append(roles, role)
    }
    return roles
}

// RegisterPost - Регистрация пользователя
func (s *DefaultAPIService) RegisterPost(ctx context.Context, registerPostRequest RegisterPostRequest) (ImplResponse, error) {
	// TODO - update RegisterPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, User{}) or use other options such as http.Ok ...
	// return Response(201, User{}), nil

	// TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	// return Response(400, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("RegisterPost method not implemented")
}

// LoginPost - Авторизация пользователя
func (s *DefaultAPIService) LoginPost(ctx context.Context, loginPostRequest LoginPostRequest) (ImplResponse, error) {
	// TODO - update LoginPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, string{}) or use other options such as http.Ok ...
	// return Response(200, string{}), nil

	// TODO: Uncomment the next line to return response Response(401, Error{}) or use other options such as http.Ok ...
	// return Response(401, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("LoginPost method not implemented")
}

// PvzGet - Получение списка ПВЗ с фильтрацией по дате приемки и пагинацией
func (s *DefaultAPIService) PvzGet(ctx context.Context, startDate time.Time, endDate time.Time, page int32, limit int32) (ImplResponse, error) {
	// TODO - update PvzGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []PvzGet200ResponseInner{}) or use other options such as http.Ok ...
	// return Response(200, []PvzGet200ResponseInner{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PvzGet method not implemented")
}

// PvzPost - Создание ПВЗ (только для модераторов)
func (s *DefaultAPIService) PvzPost(ctx context.Context, pvz Pvz) (ImplResponse, error) {
	// TODO - update PvzPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, Pvz{}) or use other options such as http.Ok ...
	// return Response(201, Pvz{}), nil

	// TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	// return Response(400, Error{}), nil

	// TODO: Uncomment the next line to return response Response(403, Error{}) or use other options such as http.Ok ...
	// return Response(403, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PvzPost method not implemented")
}

// PvzPvzIdCloseLastReceptionPost - Закрытие последней открытой приемки товаров в рамках ПВЗ
func (s *DefaultAPIService) PvzPvzIdCloseLastReceptionPost(ctx context.Context, pvzId string) (ImplResponse, error) {
	// TODO - update PvzPvzIdCloseLastReceptionPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Reception{}) or use other options such as http.Ok ...
	// return Response(200, Reception{}), nil

	// TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	// return Response(400, Error{}), nil

	// TODO: Uncomment the next line to return response Response(403, Error{}) or use other options such as http.Ok ...
	// return Response(403, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PvzPvzIdCloseLastReceptionPost method not implemented")
}

// PvzPvzIdDeleteLastProductPost - Удаление последнего добавленного товара из текущей приемки (LIFO, только для сотрудников ПВЗ)
func (s *DefaultAPIService) PvzPvzIdDeleteLastProductPost(ctx context.Context, pvzId string) (ImplResponse, error) {
	// TODO - update PvzPvzIdDeleteLastProductPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	// return Response(200, nil),nil

	// TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	// return Response(400, Error{}), nil

	// TODO: Uncomment the next line to return response Response(403, Error{}) or use other options such as http.Ok ...
	// return Response(403, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PvzPvzIdDeleteLastProductPost method not implemented")
}

// ReceptionsPost - Создание новой приемки товаров (только для сотрудников ПВЗ)
func (s *DefaultAPIService) ReceptionsPost(ctx context.Context, receptionsPostRequest ReceptionsPostRequest) (ImplResponse, error) {
	// TODO - update ReceptionsPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, Reception{}) or use other options such as http.Ok ...
	// return Response(201, Reception{}), nil

	// TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	// return Response(400, Error{}), nil

	// TODO: Uncomment the next line to return response Response(403, Error{}) or use other options such as http.Ok ...
	// return Response(403, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ReceptionsPost method not implemented")
}

// ProductsPost - Добавление товара в текущую приемку (только для сотрудников ПВЗ)
func (s *DefaultAPIService) ProductsPost(ctx context.Context, productsPostRequest ProductsPostRequest) (ImplResponse, error) {
	// TODO - update ProductsPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, Product{}) or use other options such as http.Ok ...
	// return Response(201, Product{}), nil

	// TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	// return Response(400, Error{}), nil

	// TODO: Uncomment the next line to return response Response(403, Error{}) or use other options such as http.Ok ...
	// return Response(403, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ProductsPost method not implemented")
}
