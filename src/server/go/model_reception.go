/*
 * backend service
 *
 * Сервис для управления ПВЗ и приемкой товаров
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

type Reception struct {

	Id string `json:"id,omitempty"`

	DateTime time.Time `json:"dateTime"`

	PvzId string `json:"pvzId"`

	Status string `json:"status"`
}
