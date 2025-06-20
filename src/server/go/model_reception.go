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
	"time"
)



type Reception struct {

	Id string `json:"id,omitempty"`

	DateTime time.Time `json:"dateTime"`

	PvzId string `json:"pvzId"`

	Status string `json:"status"`
}

// AssertReceptionRequired checks if the required fields are not zero-ed
func AssertReceptionRequired(obj Reception) error {
	elements := map[string]interface{}{
		"dateTime": obj.DateTime,
		"pvzId": obj.PvzId,
		"status": obj.Status,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertReceptionConstraints checks if the values respects the defined constraints
func AssertReceptionConstraints(obj Reception) error {
	return nil
}
