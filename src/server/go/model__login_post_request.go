/*
 * backend service
 *
 * Сервис для управления ПВЗ и приемкой товаров
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type LoginPostRequest struct {

	Email string `json:"email"`

	Password string `json:"password"`
}
