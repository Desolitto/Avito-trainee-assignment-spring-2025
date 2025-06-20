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



type Pvz struct {

	Id string `json:"id,omitempty"`

	RegistrationDate time.Time `json:"registrationDate,omitempty"`

	City string `json:"city"`
}

// AssertPvzRequired checks if the required fields are not zero-ed
func AssertPvzRequired(obj Pvz) error {
	elements := map[string]interface{}{
		"city": obj.City,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertPvzConstraints checks if the values respects the defined constraints
func AssertPvzConstraints(obj Pvz) error {
	return nil
}
